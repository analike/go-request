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

func (r *Request) GetIpAddresses(trustedHeaders ...string) string {
	for _, key := range trustedHeaders {
		val := r.GetHeader(key)
		if val != "" {
			return val
		}
	}
	return ""
}

func (r *Request) GetIpAddress(trustedKeys ...string) string {
	ip := r.GetIpAddresses(trustedKeys...)
	if ip != "" {
		ip = r.RemoteAddress
		pieces := strings.Split(ip, ",")
		if len(pieces) > 0 {
			found := pieces[0]
			return strings.TrimSpace(found)
		}
	}
	return ""
}

func FromGin(c *gin.Context) *Request {
	r := *c.Request
	scheme := "http"
	if fwProto := r.Header.Get("X-Forwarded-Proto"); fwProto == "https" {
		scheme = "https"
	}
	req := Request{
		Scheme:        scheme,
		Uri:           r.RequestURI,
		Method:        r.Method,
		Headers:       &r.Header,
		RemoteAddress: r.RemoteAddr,
		Data:          data.FromGin(c),
		Parameters:    c.Param,
		Queries:       nil,
		Form:          nil,
		Client:        nil,
	}

	return &req
}

func FromGoHttp(r *http.Request) *Request {
	scheme := "http"
	if fwProto := r.Header.Get("X-Forwarded-Proto"); fwProto == "https" {
		scheme = "https"
	}
	req := Request{
		Scheme:        scheme,
		Uri:           r.RequestURI,
		Method:        r.Method,
		Headers:       &r.Header,
		RemoteAddress: r.RemoteAddr,
		Data:          data.FromGoHttp(r),
		Parameters:    r.URL.Query().Get,
		Queries:       nil,
		Form:          nil,
		Client:        nil,
	}

	return &req
}
