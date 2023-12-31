package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var ePort int

func init() {
	flag.IntVar(&ePort, "port", 8080, "Exporter port to listen on")
}

func main() {
	flag.Parse()

	monitor := NewMonitor()
	middleware := func(handlerFor http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			components, err := GetGithubStatusComponents()

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("could not scrape GitHub status"))
				return
			}

			for _, component := range components {
				val := 0.0
				if component.IsOperational() {
					val = 1
				}
				monitor.GithubComponentStatus.WithLabelValues(component.Name).Set(val)
			}

			handlerFor.ServeHTTP(w, r)
		})
	}

	http.Handle("/metrics", middleware(promhttp.HandlerFor(monitor.Registry, promhttp.HandlerOpts{Registry: monitor.Registry})))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`
		<html>
		<head>
		<title>Github Status Exporter</title>
		</head>
		<body>
		<h1>Github Status Exporter</h1>
		<p><a href="/metrics">Metrics</a></p>
		</body>
		</html>
		`))
	}))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", ePort), nil))
}
