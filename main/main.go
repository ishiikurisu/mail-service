package main

import (
    "fmt"
    "log"
    "net/http"
    "net/smtp"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    http.HandleFunc("/", index)
    log.Print("Serving at ", port)
    log.Fatal(http.ListenAndServe(port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	hostname := os.Getenv("HOSTNAME")
    port := os.Getenv("HOSTPORT")
    address := fmt.Sprintf("%s%s", hostname, port)
    from := os.Getenv("HOSTMAIL")
    secret := os.Getenv("HOSTPWD")
	auth := smtp.PlainAuth("", from, secret, hostname)

    r.ParseForm()
    to := r.Form["to"]
    msg := []byte(r.Form["msg"][0])

	err := smtp.SendMail(address, auth, from, to, msg)

    if err == nil {
        fmt.Fprintf(w, "sent")
    } else {
        fmt.Fprintf(w, "failed: %#v", err)
    }
}
