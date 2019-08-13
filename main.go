package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	defaultNSFile = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
	defaultPort   = "8000"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	h := &PodInfoHandler{
		info: &PodInfo{
			Name:      Podname(),
			Namespace: Namespace(),
		},
	}

	s := http.Server{
		Addr:           ":" + port,
		Handler:        h,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// Namespace determines which namespace the pod is running in.
func Namespace() string {
	fileName := os.Getenv("NS_FILE")
	if fileName == "" {
		fileName = defaultNSFile
	}

	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("ERROR reading namespace file at '%s': %s\n", fileName, err)
	}

	return string(buf)
}

// Podname determines the name of the pod
func Podname() string {
	name, err := os.Hostname()
	if err != nil {
		log.Fatalf("ERROR determining hostname: %s\n", err)
	}
	return name
}

// PodInfo holds basic information about the pod the app is running in
type PodInfo struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

// PodInfoHandler handles HTTP requests
type PodInfoHandler struct {
	info *PodInfo
}

func (h *PodInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; encoding=utf-8")
	json.NewEncoder(w).Encode(h.info)
}
