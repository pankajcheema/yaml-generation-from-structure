package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// creating an struture
type Test struct {
	Name    string `yaml:"name"`
	Address []Add  `yaml:"address"`
}

//creating another structure
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
	//creating an array of type Add
	addarray := make([]Add, 0, 0)
	//pushing a value to array
	addarray = append(addarray, address)
	//Assigning new value to the array

	address = Add{Village: "anyari urf alinagar", Street: "188", Post: "chajlet", Dist: "moradabad", State: "UP"}

	//again pushing the new value to array

	addarray = append(addarray, address)
	//Now assigning the value to parent structure

	data := Test{Name: "pankaj", Address: addarray}

	//Creating bytes array from the struture
	marshalbytes, err := yaml.Marshal(&data)
	if err != nil {
		fmt.Println(err.Error)
	}
	//writing the file here with created dbytes
	writeErr := ioutil.WriteFile("test.yaml", marshalbytes, 0644)
	if writeErr != nil {
		fmt.Println(writeErr.Error)
	}
}
