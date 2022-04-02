package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"flag"
	"time"
)

func init() {
	flag.Usage = func() {
	help := []string{
			"",
			"",
			"Usage:",
			"+=======================================================+",
			"       -H,            Headers",
			"	-h             Show This Help Message",
			"",
			"+=======================================================+",
			"",
		}

		fmt.Fprintf(os.Stderr, strings.Join(help, "\n"))
	}

}

func main() {


	var headers string
	flag.StringVar(&headers, "H","","")
	flag.Parse()


	
	var urls []string
	std := bufio.NewScanner(os.Stdin)
	for std.Scan() {
		var line string = std.Text()
		//hline := strings.Replace(line, "%2F", "/", -1)
		//line = hline
		
		urls = append(urls, line)

	}
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(url string) {

			defer wg.Done()
			if headers != ""{
				x := getParams(url, headers)
				 if x != "ERROR" {
                         	 	fmt.Println(x)
                        }
			}else{
				x := getParams(url,"0")
				 if x != "ERROR" {
                                 	fmt.Println(x)
                        }
		}

		}(u)
	}

	wg.Wait()

}

func getParams(url string, headers string) string{
	
	var trans = &http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			KeepAlive: time.Second,
			Timeout: 5 * time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: trans,
		Timeout: 5 * time.Second,
	}

	res, err := http.NewRequest("GET", url, nil)
	
	if headers != "0"{
		if strings.Contains(headers, ";"){
			parts := strings.Split(headers, ";")
			for _, q := range parts{
				separatedHeader := strings.Split(q,":")
				res.Header.Set(separatedHeader[0],separatedHeader[1])
	}	
}else{
	sHeader := strings.Split(headers,":")
	res.Header.Set(sHeader[0], sHeader[1])
}
}
	res.Header.Set("Connection", "close")
	resp, err := client.Do(res)
	// res.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36")
	
	if err != nil {
		return "ERROR"
	}
	defer resp.Body.Close()

	if err != nil {
		return "ERROR"
	}
	
	if resp.StatusCode >= 200 && resp.StatusCode <= 299{
		return url
	}else{
		return "ERROR"
	}


}

