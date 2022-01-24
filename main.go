package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/mpawlowski/echo-http/handler"
)

type options struct {
	httpServerAddress      string
	httpServerPort         int
	httpServerWriteTimeout time.Duration
	httpServerReadTimeout  time.Duration
	httpResponseCode       int
	httpResponseBody       string
}

var flags options

func init() {

	flag.Usage = func() {
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("Usage: %s <flags>\n", os.Args[0]))
		sb.WriteString("\nSimple one-liner http server with a static response.\n")
		sb.WriteString("\nFlags:\n\n")
		fmt.Fprint(os.Stderr, sb.String())
		flag.PrintDefaults()
	}

	flag.StringVar(&flags.httpServerAddress, "http-server-address", "127.0.0.1", "Address for the http server to bind to.")
	flag.IntVar(&flags.httpServerPort, "http-server-port", 8000, "Port for the http server to bind to.")
	flag.DurationVar(&flags.httpServerReadTimeout, "http-server-read-timeout", 5*time.Second, "Read timeout for the http server.")
	flag.DurationVar(&flags.httpServerWriteTimeout, "http-server-write-timeout", 5*time.Second, "Write timeout for the http server.")
	flag.IntVar(&flags.httpResponseCode, "http-response-code", 200, "Http response code returned.")
	flag.StringVar(&flags.httpResponseBody, "http-response-body", "OK", "Http response body returned.")
	flag.Parse()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.NewSimpleHandler(flags.httpResponseCode, []byte(flags.httpResponseBody)))
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", flags.httpServerAddress, flags.httpServerPort),
		WriteTimeout: flags.httpServerWriteTimeout,
		ReadTimeout:  flags.httpServerReadTimeout,
	}

	log.Fatal(srv.ListenAndServe())
}
