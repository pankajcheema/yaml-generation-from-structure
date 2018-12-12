package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Test struct {
	Name    string `yaml:"name"`
	Address []Add  `yaml:"address"`
}
type Add struct {
	Village string `yaml:"village"`
	Street  string `yaml:"street"`
	Post    string `yaml:"post"`
	Dist    string `yaml:"dist"`
	State   string `yaml:"state"`
}

func main() {
	var address Add
	address = Add{Village: "anyari", Street: "188", Post: "chajlet", Dist: "moradabad", State: "UP"}
	addarray := make([]Add, 0, 0)
	addarray = append(addarray, address)
	address = Add{Village: "anyari urf alinagar", Street: "188", Post: "chajlet", Dist: "moradabad", State: "UP"}

	addarray = append(addarray, address)
	data := Test{Name: "pankaj", Address: addarray}

	marshalbytes, err := yaml.Marshal(&data)
	if err != nil {
		fmt.Println(err.Error)
	}
	writeErr := ioutil.WriteFile("test.yaml", marshalbytes, 0644)
	if writeErr != nil {
		fmt.Println(writeErr.Error)
	}
}
