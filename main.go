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

	components := GetGithubStatusComponents()
	monitor := NewMonitor()

	for _, component := range components {
		if component.Status == "operational" {
			monitor.GithubComponentStatus.WithLabelValues(component.Name, component.Status).Set(1)
		} else {
			monitor.GithubComponentStatus.WithLabelValues(component.Name).Set(0)
		}
	}

	http.Handle("/metrics", promhttp.HandlerFor(monitor.Registry, promhttp.HandlerOpts{Registry: monitor.Registry}))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", ePort), nil))
}
