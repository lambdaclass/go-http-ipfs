package main

import (
	shell "github.com/ipfs/go-ipfs-api"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	ipfsShell := shell.NewShell("localhost:5001")

	uploadHandler := func(w http.ResponseWriter, req *http.Request) {
		file, _, error := req.FormFile("file")
		if error != nil {
			panic(error)
		}
		defer file.Close()

		hash, error := ipfsShell.Add(file)
		if error != nil {
			panic(error)
		}

		io.WriteString(w, hash)
	}

	fileHandler := func(w http.ResponseWriter, req *http.Request) {
		url := req.URL.Path
		hash := strings.Split(url, "/")[2]

		error := ipfsShell.Get(hash, "/tmp/")
		if error != nil {
			panic(error)
		}

		w.Header().Add("Content-Disposition", "Attachment")
		http.ServeFile(w, req, "/tmp/"+hash)
	}

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/file/", fileHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
