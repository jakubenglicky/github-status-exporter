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
			components := GetGithubStatusComponents()
			for _, component := range components {
				if component.Status == "operational" {
					monitor.GithubComponentStatus.WithLabelValues(component.Name, component.Status).Set(1)
				} else {
					monitor.GithubComponentStatus.WithLabelValues(component.Name).Set(0)
				}
			}
			handlerFor.ServeHTTP(w, r)
		})
	}

	http.Handle("/metrics", middleware(promhttp.HandlerFor(monitor.Registry, promhttp.HandlerOpts{Registry: monitor.Registry})))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", ePort), nil))
}
