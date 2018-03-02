/*
/   Author: Marek Stefanowski
/   Description: A Golang server implementation to identify if a number is a multiple of 7 and/or 9
/   Licence: Free
*/
package main

// Imports required to carry out the functionality
import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "time"
    "github.com/gorilla/mux"
)

// Struct for the output JSON file
type jsonMessage struct {
    Program     string
    Version     string
    Timestamp   time.Time
    Input       int 
    Output      string
}

// Function that takes in an empty interface
// Allowing for future updates of the API to reuse this code
func respondJSON (output interface{}, w http.ResponseWriter, r *http.Request) {
    // Convert Message to JSON
    jsonResponse, _ := json.MarshalIndent(output, "", "\t")
    
    // Indicate to browser file is to be downloaded and serve
    w.Header().Add("Content-Disposition", "Attachment;filename=response")
    w.WriteHeader(http.StatusCreated)
    w.Write(jsonResponse) 
}

// Home page to act as a help page
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Multiples API\nPlease provide an input value in the form:\n" +
                   "127.0.0.1/GET/<x> or localhost/GET/<x>\n" +
                   "Where \"<x>\" is the number you wish to evaluate")
}

// Error Handling for no input eval value provided
func noInputValue(w http.ResponseWriter, r *http.Request) {
    // Generate JSON Structure Message and send via respondJSON
    createResponse := jsonMessage{Program: "multiplesAPI", 
                                  Version: "1.0", 
                                  Timestamp: time.Now().UTC(), 
                                  Input: 0, 
                                  Output: "No Input Value Provided"}
    respondJSON(createResponse, w, r)
}

// Calculate response to input provided
func evaluateValue(w http.ResponseWriter, r *http.Request) {
    // Receive input value to test
    vars := mux.Vars(r)
    inputVal, err := strconv.Atoi(vars["id"])
    if err != nil {
        fmt.Fprintf(w, "Invalid input, an integer must be provided\n\n" +
                     "Please provide an input value in the form:\n" +
                     "127.0.0.1/GET/<x> or localhost/GET/<x>\n" +
                     "Where \"<x>\" is the number you wish to evaluate")
        return
    }
    
    // Calculate Response
    var response string = ""  
    if (inputVal % 7) == 0 { 
        response += "C"                   
    }    
    if (inputVal % 9) == 0 { 
        response += "N"                   
    }
    if response == ""      {
        response = strconv.Itoa(inputVal) 
    } 
    
    // Generate JSON Structure Message and send via respondJSON
    createResponse := jsonMessage{Program: "multiplesAPI", 
                                  Version: "1.0", 
                                  Timestamp: time.Now().UTC(), 
                                  Input: inputVal, 
                                  Output: response}
    respondJSON(createResponse, w, r)
}

// Initialise request router and handler for each expected input path
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/GET/", noInputValue)
    myRouter.HandleFunc("/GET/{id}", evaluateValue)
    // Create server to listen on port 80
    http.ListenAndServe(":80", myRouter)
}

// Main function to call handle requests
func main() {
    handleRequests()
}
