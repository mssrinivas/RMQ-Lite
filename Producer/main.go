package main


import (
    "fmt"
    "log"
    "net/http"
)

type Producer struct {
  produceRecords ProduceRecordsToBroker
}


func main() {
   // Initialize all the configuration parameters
   port := flag.String("port", "9090", "a string")

    fmt.Printf("Starting producer at port \n")
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}