/**
* save Json
* 将 JSON 数据存入文件
**/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "jan@newmarch.name"}}}

	saveJson("data/person.json", person)
}

func saveJson(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	// 编码 json 转 string
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}

/////////////////////////////////////////////////////////////////////////
/**
* load Json
* 将 JSON 加载到内存中
**/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func main() {
	var person Person
	loadJson("data/person.json", &person)
	fmt.Println("Person:", person.String())
}

func loadJson(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	// 解码  string 转 json
	dncoder := json.NewDecoder(inFile)
	err = dncoder.Decode(key)
	checkError(err)
	inFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}

