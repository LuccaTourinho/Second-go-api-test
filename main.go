package main

import (
    "log"
)

func main() {
    store, err := NewPostgressStorage()
    if err != nil {
        log.Fatal(err)
    }

    server := NewAPIServer(":8080", store) // This should work if NewAPIServer is defined properly
    log.Fatal(server.Run())
}