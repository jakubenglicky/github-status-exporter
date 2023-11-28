package main

import (
	"flag"
	"net/http"
	"log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ePort int

func init() {
	flag.IntVar(&ePort, "port", 8080, "Exporter port to listen on")
}

func main() {
	components := GetGithubStatusComponents()
	monitor := NewMonitor()

	for _, component := range components {
		if component.Status == "operational" {
			monitor.GithubComponentStatus.WithLabelValues(component.Name, component.Status).Set(1)
		} else {
			monitor.GithubComponentStatus.WithLabelValues(component.Name).Set(0)
		}
	}

	flag.Parse()

	http.Handle("/metrics", promhttp.HandlerFor(monitor.Registry, promhttp.HandlerOpts{Registry: monitor.Registry}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
