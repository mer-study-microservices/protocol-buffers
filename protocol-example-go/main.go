package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	simplepb "github.com/PROTOCOL-EXAMPLE-GO/src/simple"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()
	readAndWriteDemo(sm)
}

func readAndWriteDemo(pb proto.Message) {
	// writeToFile()
	writeToFile("simple.bin", pb)
	sm := &simplepb.SimpleMessage{}
	// readFromFile()
	readFromFile("simple.bin", sm)
	fmt.Println(sm)
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	sm.Name = "I renamed you"

	fmt.Println(sm)
	fmt.Println("The ID is:" + strconv.FormatInt(int64(sm.GetId()), 10))

	return &sm
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err)
		return err2
	}

	return nil
}
