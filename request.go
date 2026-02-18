/**
 * @package go-request (2026)
 * @author Emmanuel Analike <emmanuel@analike.dev>
 * @created Feb 18, 2026; 2:17 PM
 */

package request

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.analike.dev/request/data"
)

type Parameters map[string]string

type Request struct {
	Scheme        string
	Uri           string
	Method        string
	Headers       *http.Header
	RemoteAddress string
	Data          *data.Data
	Parameters    func(string) string
	Queries       *Values
	Form          *Values
	Client        *any
}

func (r *Request) GetHeader(key string) string {
	return r.Headers.Get(key)
}

func (r *Request) GetAuthHeader() string {
	return r.GetHeader("Authorization")
}

func (r *Request) HasAuthHeader() bool {
	return r.GetAuthHeader() != ""
}

func (r *Request) GetBearerToken() string {
	if r.HasAuthHeader() {
		split := strings.Split(r.GetHeader("Authorization"), "Bearer ")
		if len(split) == 2 {
			return strings.TrimSpace(split[1])
		}
	}
	return ""
}

func (r *Request) HasBearerToken() bool {
	return r.GetBearerToken() != ""
}

func (r *Request) GetIpAddresses(keys ...string) string {
	for _, key := range keys {
		val := r.GetHeader(key)
		if val != "" {
			return val
		}
	}
	return ""
}

func FromGin(c *gin.Context) *Request {
	r := *c.Request
	var scheme string
	if r.Header["X-Forwarded-Proto"][0] == "https" {
		scheme = "https"
	} else {
		scheme = "http"
	}
	req := Request{
		Scheme:        scheme,
		Uri:           r.RequestURI,
		Method:        r.Method,
		Headers:       &r.Header,
		RemoteAddress: r.RemoteAddr,
		Data:          data.FromGin((c)),
		Parameters:    c.Param,
		Queries:       nil,
		Form:          nil,
		Client:        nil,
	}

	return &req
}
