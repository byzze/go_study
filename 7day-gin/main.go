package main

/*
(1)
$ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello Gee</h1>

(2)
$ curl "http://localhost:9999/hello?name=geektutu"
hello geektutu, you're at /hello

(3)
$ curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
{"password":"1234","username":"geektutu"}

(4)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/

import (
	"log"
	"net/http"
	"time"

	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	// r.GET("/hello/name", func(c *gee.Context) {
	// 	// expect /hello?name=geektutu
	// 	c.String(http.StatusOK, "name %s, you're at %s\n", c.Query("name"), c.Path)
	// })
	r.GET("/hello/:hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello1 %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:hello/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello2 %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:hello1", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello3 %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}