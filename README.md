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
* Read a list of URL's and return only HTTP 200 response code.
* YOu can use to check if endpoints are alive, Ex: `echo http://xxxxx.com/app.js | nilo` 
