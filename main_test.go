package main

import (
	"encoding/json"
	"log"
	"testing"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func BenchmarkJSON(b *testing.B) {
	p := &Person{
		Name:  "John Doe",
		Id:    1234,
		Email: "johndoe@example.com",
	}

	b.Run("serialization", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := json.Marshal(p)
			if err != nil {
				log.Fatal("JSON Marshaling error: ", err)
			}
		}
	})

	b.Run("deserialization", func(b *testing.B) {
		data, _ := json.Marshal(p)
		for i := 0; i < b.N; i++ {
			newPerson := &Person{}
			if err := json.Unmarshal(data, newPerson); err != nil {
				log.Fatal("JSON Unmarshaling error: ", err)
			}
		}
	})
}

func BenchmarkProtobuf(b *testing.B) {
	p := &Person{
		Name:  "John Doe",
		Id:    1234,
		Email: "johndoe@example.com",
	}

	b.Run("serialization", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := proto.Marshal(p)
			if err != nil {
				log.Fatal("Protobuf Marshaling error: ", err)
			}
		}
	})

	b.Run("deserialization", func(b *testing.B) {
		data, _ := proto.Marshal(p)
		for i := 0; i < b.N; i++ {
			newPerson := &Person{}
			if err := proto.Unmarshal(data, newPerson); err != nil {
				log.Fatal("Protobuf Unmarshaling error: ", err)
			}
		}
	})
}

func BenchmarkComparison(b *testing.B) {
	p := &Person{
		Name:  "John Doe",
		Id:    1234,
		Email: "johndoe@example.com",
	}

	// JSON
	b.Run("JSON", func(b *testing.B) {
		b.Run("serialization", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := json.Marshal(p)
				if err != nil {
					log.Fatal("JSON Marshaling error: ", err)
				}
			}
		})

		b.Run("deserialization", func(b *testing.B) {
			data, _ := json.Marshal(p)
			for i := 0; i < b.N; i++ {
				newPerson := &Person{}
				if err := json.Unmarshal(data, newPerson); err != nil {
					log.Fatal("JSON Unmarshaling error: ", err)
				}
			}
		})
	})

	// Protobuf
	b.Run("Protobuf", func(b *testing.B) {
		b.Run("serialization", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := proto.Marshal(p)
				if err != nil {
					log.Fatal("Protobuf Marshaling error: ", err)
				}
			}
		})

		b.Run("deserialization", func(b *testing.B) {
			data, _ := proto.Marshal(p)
			for i := 0; i < b.N; i++ {
				newPerson := &Person{}
				if err := proto.Unmarshal(data, newPerson); err != nil {
					log.Fatal("Protobuf Unmarshaling error: ", err)
				}
			}
		})
	})
}

func BenchmarkComplex(b *testing.B) {
	friends := make([]*Person, 50)
	for i := range friends {
		friends[i] = &Person{
			Name:  fmt.Sprintf("Friend %d", i+1),
			Id:    int32(i + 1235),
			Email: fmt.Sprintf("friend%d@example.com", i+1),
		}
	}

	p := &Person{
		Name:    "John Doe",
		Id:      1234,
		Email:   "johndoe@example.com",
		Friends: friends,
	}

	// JSON
	b.Run("JSON", func(b *testing.B) {
		b.Run("serialization", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := json.Marshal(p)
				if err != nil {
					log.Fatal("JSON Marshaling error: ", err)
				}
			}
		})

		b.Run("deserialization", func(b *testing.B) {
			data, _ := json.Marshal(p)
			for i := 0; i < b.N; i++ {
				newPerson := &Person{}
				if err := json.Unmarshal(data, newPerson); err != nil {
					log.Fatal("JSON Unmarshaling error: ", err)
				}
			}
		})
	})

	// Protobuf
	b.Run("Protobuf", func(b *testing.B) {
		b.Run("serialization", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := proto.Marshal(p)
				if err != nil {
					log.Fatal("Protobuf Marshaling error: ", err)
				}
			}
		})

		b.Run("deserialization", func(b *testing.B) {
			data, _ := proto.Marshal(p)
			for i := 0; i < b.N; i++ {
				newPerson := &Person{}
				if err := proto.Unmarshal(data, newPerson); err != nil {
					log.Fatal("Protobuf Unmarshaling error: ", err)
				}
			}
		})
	})
}