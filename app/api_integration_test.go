// +build integration

package app

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/axelspringer/swerve/config"
	"github.com/axelspringer/swerve/log"
	"github.com/axelspringer/swerve/model"
	"github.com/stretchr/testify/assert"
)

const (
	testDomainFrom       = "ilk.io"
	testDomainTo         = "https://bild.de"
	updateNewDescription = "this is a new description"
	updateNewStatuscode  = 302
	updateNewPromotable  = false
)

var (
	token             string
	cfg               *config.Configuration
	httpClient        http.Client
	baseUrlApi        string
	baseUrlRedirecter string
	updateNewPathMap  []model.PathMap = []model.PathMap{
		{From: "/", To: "/home"},
	}
)

func TestMain(m *testing.M) {
	os.Setenv("SWERVE_USE_PEBBLE", "true")
	os.Setenv("SWERVE_PEBBLE_CA_URL", "http://localhost:15000/roots/0")
	os.Setenv("SWERVE_LETSENCRYPT_URL", "https://localhost:14000/dir")
	os.Setenv("SWERVE_DB_ENDPOINT", "http://localhost:8000")
	os.Setenv("SWERVE_DB_REGION", "eu-west-1")
	os.Setenv("SWERVE_DB_KEY", "0")
	os.Setenv("SWERVE_DB_SECRET", "0")
	os.Setenv("SWERVE_USERS", "Users")
	os.Setenv("SWERVE_DOMAINS", "Domains")
	os.Setenv("SWERVE_DOMAINS_TLS_CACHE", "DomainsTLSCache")
	os.Setenv("SWERVE_API_UI_URL", "*")
	os.Setenv("SWERVE_API_VERSION", "v1")
	os.Setenv("SWERVE_BOOTSTRAP", "1")
	os.Setenv("SWERVE_LOG_LEVEL", "error")
	a := NewApplication()
	cfg = a.Config
	httpClient = http.Client{
		Timeout: time.Second * 5,
	}
	baseUrlApi = fmt.Sprintf("http://127.0.0.1:%d", cfg.API.Listener)
	if err := a.Setup(); err != nil {
		log.Fatal(err)
	}
	go a.Run()
	err := waitUntilServerIsUpAndReady(a.Config.API.Listener)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

type testcase struct {
	name               string
	user               string
	pass               string
	expectedStatuscode int
}

func Test_APILogin(t *testing.T) {
	testCases := []testcase{
		{"valid login", "admin", "mytestpw", http.StatusOK},
		{"invalid login", "admin", "noValidPW", http.StatusUnauthorized},
	}

	for _, te := range testCases {
		url := fmt.Sprintf("%s/login", baseUrlApi)
		t.Logf("run -> %s, user: %s wanted statuscode: %d, url: %s", te.name, te.user, te.expectedStatuscode, url)
		payload := []byte(fmt.Sprintf(`{"username":"%s", "pwd":"%s"}`, te.user, te.pass))
		resp, err := httpClient.Post(url, "content-type: application/json", bytes.NewReader(payload))
		if err != nil {
			t.Fatal(err)
		}
		if te.name == "valid login" && resp.StatusCode == http.StatusOK {
			for _, cookie := range resp.Cookies() {
				if cookie.Name == "token" {
					token = cookie.Value
					break
				}
			}
		}
		assert.Equal(t, te.expectedStatuscode, resp.StatusCode, fmt.Sprintf("Test %s", te.name))
	}
}

func Test_PostRedirects(t *testing.T) {
	if checkEmptyToken(t) {
		return
	}
	testCases := []struct {
		name               string
		payload            model.Redirect
		expectedStatuscode int
	}{
		{
			"post new redirect for domain " + testDomainFrom,
			model.Redirect{
				RedirectFrom: testDomainFrom,
				Description:  "",
				RedirectTo:   testDomainTo,
				Promotable:   true,
				Code:         301,
				PathMaps: []model.PathMap{
					{
						From: "/",
						To:   "/digital",
					},
				},
			},
			200,
		},
	}

	url := fmt.Sprintf(baseUrlApi + "/" + cfg.API.Version + "/redirects/")
	for _, te := range testCases {
		payload, err := json.Marshal(te.payload)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Cookie", fmt.Sprintf("token=%s", token))
		resp, err := httpClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, resp.StatusCode, te.expectedStatuscode, te.name)
	}

	if t.Failed() {
		t.Fatal(fmt.Errorf("POST redirect faild, skipp all other tests"))
	}

}

func Test_GetRedirects(t *testing.T) {
	if checkEmptyToken(t) {
		return
	}
	testCases := []struct {
		name               string
		domain             string
		expectedStatuscode int
	}{
		{"get redirect for domain " + testDomainFrom,
			testDomainFrom,
			200,
		},
	}

	url := fmt.Sprintf(baseUrlApi + "/" + cfg.API.Version + "/redirects/" + testDomainFrom)
	for _, te := range testCases {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Cookie", fmt.Sprintf("token=%s", token))
		resp, err := httpClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		r := model.Redirect{}
		assert.Equal(t, nil, json.NewDecoder(resp.Body).Decode(&r), "unmarshal response against model.Redirect")
		assert.Equal(t, testDomainFrom, r.RedirectFrom, "unmarshal response against model.Redirect")
		assert.Equal(t, te.expectedStatuscode, resp.StatusCode, te.name)
	}
}

func Test_UpdateRedirects(t *testing.T) {
	if checkEmptyToken(t) {
		return
	}
	testCases := []struct {
		name               string
		redirect           model.Redirect
		expectedStatuscode int
	}{
		{"update redirect for domain " + testDomainFrom,
			model.Redirect{
				RedirectFrom: testDomainFrom,
				Description:  updateNewDescription,
				RedirectTo:   testDomainTo,
				Promotable:   updateNewPromotable,
				Code:         updateNewStatuscode,
				PathMaps:     updateNewPathMap,
			},
			http.StatusOK,
		},
	}

	url := fmt.Sprintf(baseUrlApi + "/" + cfg.API.Version + "/redirects/" + testDomainFrom)
	for _, te := range testCases {
		payload, err := json.Marshal(te.redirect)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Cookie", fmt.Sprintf("token=%s", token))
		resp, err := httpClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, te.expectedStatuscode, resp.StatusCode, te.name)
	}
}

func Test_UpdatedRedirects(t *testing.T) {
	if checkEmptyToken(t) {
		return
	}
	testCases := []struct {
		name               string
		domain             string
		expectedStatuscode int
	}{
		{"get redirect for domain " + testDomainFrom,
			testDomainFrom,
			http.StatusOK,
		},
	}

	url := fmt.Sprintf(baseUrlApi + "/" + cfg.API.Version + "/redirects/" + testDomainFrom)
	for _, te := range testCases {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Cookie", fmt.Sprintf("token=%s", token))
		resp, err := httpClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		r := model.Redirect{}
		assert.Equal(t, nil, json.NewDecoder(resp.Body).Decode(&r), "unmarshal response against model.Redirect")
		assert.Equal(t, testDomainTo, r.RedirectTo, "check to domain are equal")
		assert.Equal(t, testDomainFrom, r.RedirectFrom, "check previous updated redirect.RedirectFrom")
		assert.Equal(t, updateNewDescription, r.Description, "check previous updated redirect.description")
		assert.Equal(t, updateNewPromotable, r.Promotable, "check previous updated redirect.Promotable")
		assert.Equal(t, updateNewPathMap, r.PathMaps, "check previous updated redirect.PathMap")
		assert.Equal(t, te.expectedStatuscode, resp.StatusCode, te.name)
	}
}

func Test_HTTPSRedirect(t *testing.T) {
	httpClient := getHttpClientWithPebbleIntermediateCert(t)
	resp, err := httpClient.Get(fmt.Sprintf("https://%s:8081/", testDomainFrom))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, updateNewStatuscode, resp.StatusCode, "test redirect for https://ilk.io")
}

func Test_DeleteRedirect(t *testing.T) {
	return
	if checkEmptyToken(t) {
		return
	}
	testCases := []struct {
		name               string
		domain             string
		expectedStatuscode int
	}{
		{"delete redirect for domain " + testDomainFrom,
			testDomainFrom,
			http.StatusOK,
		},
	}

	url := fmt.Sprintf(baseUrlApi + "/" + cfg.API.Version + "/redirects/" + testDomainFrom)
	for _, te := range testCases {
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Cookie", fmt.Sprintf("token=%s", token))
		resp, err := httpClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, te.expectedStatuscode, resp.StatusCode, te.name)
	}
}

func Test_RedirectExistAfterDelete(t *testing.T) {
	return
	if checkEmptyToken(t) {
		return
	}
	testCases := []struct {
		name               string
		domain             string
		expectedStatuscode int
	}{
		{"test if redirect for domain " + testDomainFrom + " still exists",
			testDomainFrom,
			http.StatusNotFound,
		},
	}

	url := fmt.Sprintf(baseUrlApi + "/" + cfg.API.Version + "/redirects/" + testDomainFrom)
	for _, te := range testCases {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Cookie", fmt.Sprintf("token=%s", token))
		resp, err := httpClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, te.expectedStatuscode, resp.StatusCode, te.name)
	}
}

// Tests end here
func getHttpClientWithPebbleIntermediateCert(t *testing.T) *http.Client {
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM([]byte(cfg.PebbleCA))

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	tlsconfig := &tls.Config{
		RootCAs:   caPool,
		ClientCAs: caPool,
	}
	tr := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			if !strings.Contains(addr, ":15000") {
				addr = "127.0.0.1:8081"
			}
			return dialer.DialContext(ctx, network, addr)
		},
		TLSClientConfig: tlsconfig,
	}

	pebbleClient := &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := pebbleClient.Get("https://127.0.0.1:15000/roots/0")
	if err != nil {
		t.Fatal(err)
	}
	intermediateCert, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	caPool.AppendCertsFromPEM(intermediateCert)

	return pebbleClient
}

func waitUntilServerIsUpAndReady(apiport int) error {
	for i := 0; i < 30; i++ {
		resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/health", apiport))
		if err != nil {
		}
		if resp != nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		time.Sleep(time.Second * 1)
	}
	return fmt.Errorf("Can't reach api server on http://127.0.0.1:%d/health", apiport)
}

func checkEmptyToken(t *testing.T) bool {
	if token == "" {
		t.Fail()
		t.Log("test skipped because of missing token")
		return true
	}
	return false
}
