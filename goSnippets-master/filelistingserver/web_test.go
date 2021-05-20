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

var tests = []struct {
	h    appHandler
	code int
	msg  string
}{
	{mockPanic, 500, "Internal Server Error"},
	{mockUserError, 400, "user error"},
	{mockNotFound, 404, "Not Found"},
	{mockForbidden, 403, "Forbidden"},
	{mockUnknown, 500, "Internal Server Error"},
	{mockNoError, 200, "no error"},
}

func TestErrorWrapper(t *testing.T) {
	for _, test := range tests {
		recorder := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		handler := errorWrapper(test.h)
		handler(recorder, req)

		verifyResponse(recorder.Result(), test.code, test.msg, t)
	}
}

func TestErrorWrapperInServer(t *testing.T) {
	for _, test := range tests {
		server := httptest.NewServer(errorWrapper(test.h))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, test.code, test.msg, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	bytes, _ := ioutil.ReadAll(resp.Body)
	body := strings.TrimSpace(string(bytes))
	if expectedCode != resp.StatusCode || expectedMsg != body {
		t.Errorf("expected (%d, %s) got (%d, %s)",
			expectedCode, expectedMsg, resp.StatusCode, body)
	}
}

func mockNoError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintf(writer, "no error")
	return nil
}

func mockUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown")
}

func mockForbidden(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func mockNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

type MockUserErr string

func (e MockUserErr) Message() string {
	return string(e)
}

func (e MockUserErr) Error() string {
	return e.Message()
}
func mockUserError(writer http.ResponseWriter, request *http.Request) error {
	return MockUserErr("user error")
}

func mockPanic(writer http.ResponseWriter, request *http.Request) error {
	panic("PANIC")
}
