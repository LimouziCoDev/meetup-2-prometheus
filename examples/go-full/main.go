package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Errors channel
	errc := make(chan error)

	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		// index endpoint
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Welcome to the LimouziCoDev Prometheus meetup !")
		})
		http.Handle("/metrics", promhttp.Handler())
		http.HandleFunc("/demo/", prometheus.InstrumentHandler("demo-app", demo()))

		log.Println("The microservice prometheus-demo-app has started on port 8888")
		errc <- http.ListenAndServe(":8888", nil)
	}()

	log.Println("exit", <-errc)
}

func demo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := rand.Intn(10)
		if value == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else if value == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
