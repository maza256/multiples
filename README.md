## Multiples Test API

This is an API implemented in the Go Language developed by Google, which reads in from a GET request a value and returns a JSON format file whether the number is a multiple of 7 and/or 9.


Due to time constraints error handling and code re-use have not been a priority.

# Build and Run

To run the server an Installation of Go lang must be present (v1.6.2+) and the path to the install location added to the path environment variable. 

Once installed, the following command must be issued to install a required package:
`go get github.com/gorilla/mux`

Once installed the code can be compiled with:
`go build multiple_API.go`

The compiled executable can now be run to start the server. This may need to be run as a root due to using port 80, which some systems may protect from use.
The code may be modified to utilise a different port number to solve this. If changed, the browser call will need to include the port number used in the following manner, example port number 8080:
`localhost:8080/GET/<number_to_be_evaluated>`