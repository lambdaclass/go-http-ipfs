# Dependencies
- [Go](https://golang.org/) (1.11)
- [IPFS](https://docs.ipfs.io/) (0.4.17)
- [go-ipfs-api](https://github.com/ipfs/go-ipfs-api)

# Setup
1. [Install Go](https://golang.org/dl/)
1. [Install IPFS](https://docs.ipfs.io/introduction/install/)
1. Get get-ipfs-api by doing: `go get -u github.com/ipfs/go-ipfs-api`
1. Start the IPFS daemon: `ipfs daemon`
1. In a new shell clone and move into this directory: `git clone https://github.com/lambdaclass/go-http-ipfs.git && go-http-ipfs`
1. Build and run the Go server: `go run`

Now you will have the HTTP server running in port 8080.

# Endpoints

## /upload

This endpoint receives the file the user wants to upload to IPFS. The request must be a POST with a multipart form in which the field `file` has the file to upload.

Example:
```
$> curl localhost:8000/upload -F 'file=@example.txt'
QmbZGZZFbc9eFB1hfZj8PHsp8ZrszwXgf3nkrFx8z3v6ri
```

## /file/:hash

This endpoint fetches a file from IPFS using its hash and returns it to the client for download. The request must be a GET where `:hash` is the hash identifying the file in IPFS

Example:
```
curl localhost:8000/file/QmbZGZZFbc9eFB1hfZj8PHsp8ZrszwXgf3nkrFx8z3v6ri -O
```
