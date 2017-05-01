package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
)

func TestFoo(t *testing.T) {
    router := getEngine()

    req, _ := http.NewRequest("POST", "/resistance", strings.NewReader("{\"bands\":[\"Orange\",\"Blue\",\"Black\",\"Silver\"]}"))
    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    actual := strings.TrimRight(resp.Body.String(), "\n")
    expected := "{\"value\":3.6,\"prefix\":\"da\",\"tolerance\":10}"
    if strings.Compare(actual, expected) != 0  {
      t.Fatalf("\"%v\" != \"%v\"", actual, expected)
    }
}
