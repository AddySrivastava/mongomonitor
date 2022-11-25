package utils

type httpClient interface {
	makeRequest() ([]byte, int, error)
}
