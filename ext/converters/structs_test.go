package converters

import (
	"fmt"
	"github.com/fluent-qa/gobase/pkg/utils"
	"testing"
)

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

type Address struct {
	Location string `json:"location"`
	ZipCode  string `json:"zipCode"`
}

func TestToStruct(t *testing.T) {
	addr := Address{Location: "test", ZipCode: "109032"}
	fmt.Println(addr.Location)
	person := &Person{
		Name:    "Person",
		Age:     10,
		Address: addr,
	}
	result := utils.JsonConverter.ToStructuredString(&addr)
	personResult := utils.JsonConverter.ToStructuredString(&person)
	fmt.Println(result, personResult)
	yamlPersonResult := ConverterIt(utils.YamlConverter, person)
	fmt.Println(yamlPersonResult)
	converter := NewFormatConverter(utils.JsonConverter)
	converter.Converter.ToStructuredString(person)
}

func ConverterIt(converter StringFormatConverter, any2 any) string {
	return converter.ToStructuredString(any2)
}
