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

func errPanic(writer http.ResponseWriter,
	request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter,
	request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter,
	request *http.Request) error {
	return errors.New("unknown error")
}

// test no error
func noError(writer http.ResponseWriter,
	request *http.Request) error {
	fmt.Fprint(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNotPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	//httptest.NewRecorder()

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.immoc.cn", nil)
		f(response, request)

		// b is []byte
		//b, _ := ioutil.ReadAll(response.Body)
		//body := strings.Trim(string(b), "\n")
		//if response.Code != tt.code || body != tt.message {
		//
		//	t.Errorf("expect (%d, %s);"+
		//		"got (%d,%s)", tt.code, tt.message,
		//		response.Code, body)
		//
		//}

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		// convert f to http.HandlerFunc convert
		server := httptest.NewServer(http.HandlerFunc(f))

		// Client call
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
	}

}

func verifyResponse(
	resp *http.Response,
	expectedCode int,
	expectedMsg string,
	t *testing.T) {
	// b is []byte
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {

		t.Errorf("expect (%d, %s);"+
			"got (%d,%s)", expectedCode, expectedMsg,
			resp.StatusCode, body)

	}
}
