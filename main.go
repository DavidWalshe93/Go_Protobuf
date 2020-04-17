// David Walshe
// 17/04/2020

package main

import (
	complexpb "GoProtoBuffer/src/complex"
	enumpb "GoProtoBuffer/src/enum"
	simplepb "GoProtoBuffer/src/simple"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func main() {

	// Simple Demo
	// ===========
	sm := createProtoBuffer()
	readAndWriteDemo(sm)
	jsonDemo(sm)

	// Enum Demo
	// =========
	ep := createEnum()
	fmt.Println(ep)

	// Complex Demo
	// ============
	cp := createComplexProtoBuf()
	fmt.Println(cp)
}

// *********************************************************************************************************************
// Complex
// *********************************************************************************************************************

// Create an Hierarchical Protobuf construct.
func createComplexProtoBuf() *complexpb.ComplexMessage {
	cp := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}

	return &cp
}

// *********************************************************************************************************************
// Enum
// *********************************************************************************************************************

// Create a ProtoBuf with an Enum construct.
func createEnum() *enumpb.EnumMessage {
	ep := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}

	return &ep
}

// *********************************************************************************************************************
// Simple
// *********************************************************************************************************************

// Helper method to show to/from JSON functionality.
func jsonDemo(pb proto.Message) {
	smAsString := toJSON(pb)

	fmt.Println(smAsString)

	sm2 := createProtoBuffer()
	fromJSON(smAsString, sm2)
	fmt.Println(sm2)
}

// Converts a protobuf message to a json string.
func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{
		Indent: "    ",
	}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return out
}

// Converts a json string to a protobuf message.
func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Helper method to show read and write functionality
// between a binary file and a protobuf.
func readAndWriteDemo(sm proto.Message) {
	fmt.Println(sm)
	_ = writeToFile("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}
	_ = readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

// Write a protobuf into a binary file.
func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	err = ioutil.WriteFile(fname, out, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}

// Read data from a binary file and load it into
// a protobuf.
func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}

// Creates a new ProtoBuffer with some initial values
// and returns it to caller.
func createProtoBuffer() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 5, 6, 8},
	}

	return &sm
}
