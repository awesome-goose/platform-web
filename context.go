package web

import (
	"log"
	"net/http"

	"github.com/awesome-goose/contracts"
)

type Context struct {
	request  *Request
	response *Response

	// params
	// query
	// headers
	// url
	// domain
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	c := &Context{
		request:  NewRequest(r),
		response: NewResponse(w),
	}
	c.parse()

	return c
}

func (c *Context) parse() {

	// Extracting the path and query parameters from the rawArgs
	// domain, headers etc

}

func (c *Context) Paths() []string {
	return []string{}
}

func (c *Context) Segments() []string {
	return c.Paths()
}

func (c *Context) Request() contracts.Request {
	return c.request
}

func (c *Context) Response() contracts.Response {
	return c.response
}

func (c *Context) Html(data any) {
	err := c.response.Write([]byte("<html><body>" + data.(string) + "</body></html>"))
	if err != nil {
		log.Printf("Error writing HTML response: %v", err)
	}
}

func (c *Context) Json(data any) {
	err := c.response.Write([]byte("{\"message\": \"" + data.(string) + "\"}"))
	if err != nil {
		log.Printf("Error writing JSON response: %v", err)
	}
}

func (c *Context) Print(data any) {
	err := c.response.Write([]byte(data.(string)))
	if err != nil {
		log.Printf("Error writing Print response: %v", err)
	}
}
