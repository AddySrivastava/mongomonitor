package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type HTTPDigest struct {
	URI      string
	METHOD   string
	USERNAME string
	PASSWORD string
	BODY     []byte
	HEADERS  map[string]string
}

func sendRequest(uri string, method string, username string, password string, postBody []byte, headers map[string]string) (respBody []byte, statusCode int, err error) {
	req, err := http.NewRequest(method, uri, nil)

	// Set headers in the request headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	//create the http client and send a request to fetch client the response in 401 and prepare the next
	//req body for the digest api call
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	// If the response during the first request is not 401 then endpoint doesn't need auth and
	// the response can be sent to the caller
	if resp.StatusCode != http.StatusUnauthorized {
		log.Printf("Recieved status code '%v' auth skipped, reason - %s", resp.StatusCode, string(bodyBytes))
		respBody, err = ioutil.ReadAll(resp.Body)
		return respBody, resp.StatusCode, err
	}

	// Prepare the digest path for authorization header
	digestParts := digestParts(resp)
	digestParts["uri"] = uri
	digestParts["method"] = method
	digestParts["username"] = username
	digestParts["password"] = password

	//Prepare the digest request with the given postBody and set the authorization header
	req, err = http.NewRequest(method, uri, bytes.NewBuffer(postBody))
	req.Header.Set("Authorization", getDigestAuthorization(digestParts))

	//Send the HTTP request
	resp, err = httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body([]bytes)
	body, err := ioutil.ReadAll(resp.Body)

	return body, resp.StatusCode, err
}

func (httpDigest *HTTPDigest) MakeRequest() ([]byte, int, error) {
	return sendRequest(httpDigest.URI, httpDigest.METHOD, httpDigest.USERNAME, httpDigest.PASSWORD, httpDigest.BODY, httpDigest.HEADERS)
}

func digestParts(resp *http.Response) map[string]string {
	result := map[string]string{}

	if len(resp.Header["Www-Authenticate"]) > 0 {
		relevantHeaders := []string{"nonce", "realm", "qop"}
		responseHeaders := strings.Split(resp.Header["Www-Authenticate"][0], ",")
		for _, r := range responseHeaders {
			for _, w := range relevantHeaders {
				if strings.Contains(r, w) {
					result[w] = strings.Split(r, `"`)[1]
				}
			}
		}
	}
	return result
}

func getMD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func getCnonce() string {
	b := make([]byte, 8)
	io.ReadFull(rand.Reader, b)
	return fmt.Sprintf("%x", b)[:16]
}

func getDigestAuthorization(digestParts map[string]string) string {
	d := digestParts
	ha1 := getMD5(d["username"] + ":" + d["realm"] + ":" + d["password"])
	ha2 := getMD5(d["method"] + ":" + d["uri"])
	nonceCount := 00000001
	cnonce := getCnonce()
	response := getMD5(fmt.Sprintf("%s:%s:%v:%s:%s:%s", ha1, d["nonce"], nonceCount, cnonce, d["qop"], ha2))
	authorization := fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc="%v", qop="%s", response="%s"`,
		d["username"], d["realm"], d["nonce"], d["uri"], cnonce, nonceCount, d["qop"], response)
	return authorization
}
