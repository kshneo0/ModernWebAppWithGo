package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, e *http.Request) {

}

// go test -v
// go test -cover
// go test -coverprofile=cove rage.out && go tool cover -html=coverage.out
// go test -v ./...
