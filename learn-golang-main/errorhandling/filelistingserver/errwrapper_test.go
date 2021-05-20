package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

// todo: move user error out for both
// 		test and normal expectedCode
type testUserError string

func (e testUserError) Error() string {
	return e.Message()
}

func (e testUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errForbidden(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println(writer, "no error")
	return nil
}

var tests = []struct {
	handler appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errForbidden, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, ""},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.handler)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "https://www.google.com", nil)
		f(response, request)

		verifyResponse(t, response.Result(), tt.code, tt.message)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.handler)
		server := httptest.NewServer(http.HandlerFunc(f))

		response, _ := http.Get(server.URL)
		verifyResponse(t, response, tt.code, tt.message)
	}
}

func verifyResponse(t *testing.T, response *http.Response, expectedCode int, expectedMessage string) {
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(bodyBytes), "\n")
	if response.StatusCode != expectedCode || body != expectedMessage {
		t.Errorf("expect (%d, %s), got (%d, %s) \n",
			expectedCode,
			expectedMessage,
			response.StatusCode,
			body)
	}
}
