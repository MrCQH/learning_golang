package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	Derivative()
}

func serialization() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	// 编码成json
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("data_format/json_related/vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	// 写到file中
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}

func Derivative() {
	var v VCard
	file, err := os.Open("data_format/json_related/vcard.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "文件打开出错")
	}
	reader := bufio.NewReader(file)
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "文件读取出错")
	}
	err = json.Unmarshal(bytes, &v)
	if err != nil {
		fmt.Fprintln(os.Stderr, "byte反序列化错误")
	}
	fmt.Println(string(bytes))
}
