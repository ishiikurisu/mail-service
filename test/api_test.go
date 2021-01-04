package test

import (
    "testing"
    "net/http"
    "net/url"
    "io/ioutil"
    "strings"
)

var (
    rootUrl = "http://localhost:8000"
)

func TestRootApiCall(t *testing.T) {
    email := "liberdade-organizacao@gmail.com"

    response, err := http.PostForm(rootUrl + "/", url.Values{
        "from": {email},
    })
    if err != nil {
        t.Log("Error should be nil", err)
        t.Fail()
    }

    rawData, err := ioutil.ReadAll(response.Body)
    data := string(rawData[:])
    if !strings.Contains(data, email) {
        t.Log("Email didn't appear on data")
        t.Fail()
    }
}
