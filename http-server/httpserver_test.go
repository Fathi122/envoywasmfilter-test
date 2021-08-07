package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// expected values
type expectedValues struct {
	httpCode    int
	httpHeader  string
	httpMessage string
}

// test vector
type testInput struct {
	testName       string
	testEndpoint   string
	expectedResult expectedValues
}

// test expected values
var expectedValuesTable = []expectedValues{
	{httpCode: http.StatusOK, httpHeader: "application/json", httpMessage: `{"message": "HTTP Endpoint OK!"}`},
	{httpCode: http.StatusInternalServerError, httpHeader: "application/json", httpMessage: `{"message": "HTTP Endpoint Internal Error!"}`},
	{httpCode: http.StatusNotFound, httpHeader: "text/plain; charset=utf-8", httpMessage: "404 page not found\n"},
}

// test table
var tests = []testInput{
	{testName: "TestTwoHundredEndpoint", testEndpoint: "/test200", expectedResult: expectedValuesTable[0]},
	{testName: "TestFiveHundredEndpoint", testEndpoint: "/test500", expectedResult: expectedValuesTable[1]},
	{testName: "TestMissingEndpoint", testEndpoint: "/testnonexisting", expectedResult: expectedValuesTable[2]},
}

// mock server
var server = httptest.NewServer(router())

// testFunction actual test
func testFunction(URL string, testcase testInput, t *testing.T) {
	assert := assert.New(t)
	resp, err := http.Get(fmt.Sprintf("%s%s", URL, testcase.testEndpoint))
	assert.NoError(err)
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	assert.Equal(testcase.expectedResult.httpCode, resp.StatusCode)
	assert.Equal(testcase.expectedResult.httpHeader, resp.Header.Get("Content-Type"))
	assert.NoError(err)
	assert.Equal(testcase.expectedResult.httpMessage, string(bodyBytes))
}

func TestEndpoints(t *testing.T) {
	// loop through test table
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			testFunction(server.URL, tc, t)
		})
	}
}

func TestParallelEndpoints(t *testing.T) {
	// loop through test table
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			testFunction(server.URL, tc, t)
		})
	}
}
