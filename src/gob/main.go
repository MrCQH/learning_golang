package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	Encoding()
	Decoding()
}

func Decoding() {
	file, err := os.OpenFile("gob/vcard.gob", os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal("Error in open file: ", err)
	}
	dec := gob.NewDecoder(file)
	vc := new(VCard)
	err = dec.Decode(vc)
	if err != nil {
		log.Fatal("Error in decoding gob: ", err)
	}
	log.Println(*vc)
}

func Encoding() {
	pa := &Address{"private", "chengdu", "china"}
	wa := &Address{"work", "chongqing", "china"}
	vc := VCard{"qh", "c", []*Address{pa, wa}, "none"}
	file, _ := os.OpenFile("gob/vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(&vc)
	if err != nil {
		log.Fatal("Error in encoding gob")
	}
}
