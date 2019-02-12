package authenticate

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/pomerium/pomerium/internal/templates"
)

func testAuthenticate() *Authenticate {
	var auth Authenticate
	auth.RedirectURL, _ = url.Parse("https://auth.example.com/oauth/callback")
	auth.SharedKey = "IzY7MOZwzfOkmELXgozHDKTxoT3nOYhwkcmUVINsRww="
	auth.AllowedDomains = []string{"*"}
	auth.ProxyRootDomains = []string{"example.com"}
	auth.templates = templates.New()
	return &auth
}

func TestAuthenticate_RobotsTxt(t *testing.T) {
	auth := testAuthenticate()
	req, err := http.NewRequest("GET", "/robots.txt", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(auth.RobotsTxt)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := fmt.Sprintf("User-agent: *\nDisallow: /")
	if rr.Body.String() != expected {
		t.Errorf("handler returned wrong body: got %v want %v", rr.Body.String(), expected)
	}
}