package main

import (
    "fmt"
    "log"

    "google.golang.org/protobuf/proto"
)

func main() {
    p := &Person{
        Name:  "John Doe",
        Id:    1234,
        Email: "johndoe@example.com",
    }

    // Serialize the person to a byte slice
    data, err := proto.Marshal(p)
    if err != nil {
        log.Fatal("Marshaling error: ", err)
    }

    // Deserialize the byte slice back to a person
    newPerson := &Person{}
    if err := proto.Unmarshal(data, newPerson); err != nil {
        log.Fatal("Unmarshaling error: ", err)
    }

    // Use the newPerson object
    fmt.Printf("New Person: %+v\n", newPerson)
}

