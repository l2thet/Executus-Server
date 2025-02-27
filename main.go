package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	var addr = os.Getenv("ServerPort")
	if addr == "" {
		log.Fatal("ServerPort environment variable not set")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/listavailablemusic", listAvailableMusic)

	mux.HandleFunc("/api/servesong", serveSong)

	log.Printf("Server listening on %s", addr)

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Starting server on %s", server.Addr)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func listAvailableMusic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		files, err := readDir("./assets/music")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(files)
	}

}

func readDir(dirname string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})
	return files, err
}

func serveSong(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		fileName := r.URL.Query().Get("name")
		if fileName == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		http.ServeFile(w, r, "./assets/music/"+fileName)
	}
}
