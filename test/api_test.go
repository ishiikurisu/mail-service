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
    to_email := "crisjr@pm.me"
    msg := "おはよう！"

    response, err := http.PostForm(rootUrl + "/", url.Values{
        "to": {to_email},
        "msg": {msg},
    })

    if err != nil {
        t.Log("Error should be nil")
        t.Log(err)
        t.Fail()
    }

    rawData, err := ioutil.ReadAll(response.Body)
    data := string(rawData[:])
    if !strings.Contains(data, "sent") {
        t.Log("Failed to send email, response was: ", data)
        t.Fail()
    }
}
