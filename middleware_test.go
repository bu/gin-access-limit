package accessLimit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter(CIDRs string) *gin.Engine {
	// no debug mode
	gin.SetMode(gin.ReleaseMode)

	// create a default
	r := gin.Default()

	// our middle-ware
	r.Use(CIDR(CIDRs))

	// routes
	r.GET("/", testGET)

	return r
}

func TestAllowAccessSource(t *testing.T) {
	r := setupRouter("127.0.0.1/32")

	// prepare
	ExpectedResponseStatus := 200

	// run
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "127.0.0.1:80"
	r.ServeHTTP(w, req)

	// check
	assert.Equal(t, ExpectedResponseStatus, w.Code)
}

func TestNotAllowAccessSource(t *testing.T) {
	r := setupRouter("172.18.0.0/16")

	// prepare
	ExpectedResponseStatus := 403

	// run
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "127.0.0.1:80"
	r.ServeHTTP(w, req)

	// check
	assert.Equal(t, ExpectedResponseStatus, w.Code)
}

func TestAllowAccessFromManySource(t *testing.T) {
	r := setupRouter("172.18.0.0/16, 127.0.0.1/32, ::1/128")

	// prepare
	ExpectedResponseStatus := 200

	// run
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "127.0.0.1:80"
	r.ServeHTTP(w, req)

	// check
	assert.Equal(t, ExpectedResponseStatus, w.Code)
}

func TestNotAllowAccessFromManySource(t *testing.T) {
	r := setupRouter("172.18.0.0/16, 127.0.0.1/32, ::1/128")

	// prepare
	ExpectedResponseStatus := 403

	// run
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.168.1.12:80"
	r.ServeHTTP(w, req)

	// check
	assert.Equal(t, ExpectedResponseStatus, w.Code)
}

func TestTrustedHeader(t *testing.T) {
	// Allow Trust Header
	TrustedHeaderField = "X-Real-Ip"

	r := setupRouter("172.18.0.0/16, 127.0.0.1/32, ::1/128")

	// prepare
	ExpectedResponseStatus := 200

	// run
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.168.1.12:80"
	req.Header.Add("X-Real-Ip", "127.0.0.1")

	r.ServeHTTP(w, req)

	// check
	assert.Equal(t, ExpectedResponseStatus, w.Code)
}

func TestNotInTrustedHeader(t *testing.T) {
	// Allow Trust Header
	TrustedHeaderField = "X-Real-Ip"

	r := setupRouter("172.18.0.0/16, 127.0.0.1/32, ::1/128")

	// prepare
	ExpectedResponseStatus := 403

	// run
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.168.1.12:80"
	req.Header.Add("X-Forwarded-For", "127.0.0.1")

	r.ServeHTTP(w, req)

	// check
	assert.Equal(t, ExpectedResponseStatus, w.Code)
}

func testGET(c *gin.Context) {
	c.String(200, "pong")
}
