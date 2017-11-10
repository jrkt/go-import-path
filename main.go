package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

const (
	baseDomain = "exameple.com"
)

var (
	router = map[string]string{
		"/package1": "https://github.com/jrkt/package1",
		"/package2": "https://github.com/jrkt/package2",
		"/package3": "https://gitlab.com/jrkt/package3",
	}
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("serving on port :9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.RequestURI, "?go-get=1") {
		r.RequestURI = strings.Replace(r.RequestURI, "?go-get=1", "", 1)
	}

	parts := strings.Split(r.RequestURI, "/")
	root := parts[1]
	url, ok := router["/"+root]
	if !ok {
		http.Error(w, "import not found", http.StatusNotFound)
		return
	}

	if len(parts) > 2 {
		for _, part := range parts[2:] {
			url += "/" + part
		}
	}

	opts := struct {
		RouterKey, RouterValue, BaseURL string
	}{
		RouterKey:   r.RequestURI,
		RouterValue: url,
		BaseURL:     baseDomain,
	}

	t, err := template.New("").Parse(html)
	if err != nil {
		http.Error(w, "error building template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, opts)
	if err != nil {
		http.Error(w, "error executing template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(buf.Bytes())
}

var html = `<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <meta name="go-import" content="{{.BaseURL}}{{.RouterKey}} git {{.RouterValue}}">
    <meta http-equiv="refresh" content="0; url={{.RouterValue}}">
</head>
<body>
    Nothing to see here; <a href="{{.RouterValue}}">move along</a>.
</body>`
