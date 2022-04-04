<h1 align="center">Nilo</h1> <br>

<p align="center">
  <a href="#--usage--explanation">Usage</a> •
  <a href="#--installation--requirements">Installation</a>
</p>

<h3 align="center">Nilo makes requests to urls returning only if status code is 200</h3>

## - Installation & Requirements:
```
> go install github.com/ferreiraklet/nilo@latest

OR

> git clone https://github.com/ferreiraklet/nilo.git

> cd nilo

> go build nilo.go

> chmod +x nilo

> ./nilo -h
```
<br>


## - Usage & Explanation:
* Inspired by hakcheckurl from hakluke, nilo reads a list of URL's and return only HTTP 200 response code.
* You can use to check if endpoints are alive, Ex: `echo http://xxxxx.com/app.js | nilo`
* IMPORTANT! Urls must contain protocol, http, https.
* You can specify headers: `echo http://xxxx.com/app.js | nilo -H "Header1: value1;Header2: Value2;Header3: Value3"` Be careful with syntax.

