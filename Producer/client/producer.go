package producer

import (
    "log"
    "net/http"
)

type Producer struct {
 // produceRecords ProduceRecordsToBroker
}


func HttpKeepAlive(port string) {
	errChannel := make(chan error)
	go func() {
		errChannel <- http.ListenAndServe(port, nil)
	}()

	log.Fatal("exit", <-errChannel)
}