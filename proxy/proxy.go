package proxy

import (
	"io"
	"net/http"
	"web-server/config"
)

var customTransport = http.DefaultTransport

func ForwardRequest(w http.ResponseWriter, r *http.Request) {
	target := config.AllConfig.TargetUrl

	proxyReq, err := http.NewRequest(r.Method, target, r.Body)
	if err != nil {
		http.Error(w, "error createing proxy request", http.StatusInternalServerError)
		return
	}

	for name, values := range r.Header {
		for _, v := range values {
			proxyReq.Header.Add(name, v)
		}
	}

	resp, err := customTransport.RoundTrip(proxyReq)
	if err != nil {
		http.Error(w, "Error sending request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}
