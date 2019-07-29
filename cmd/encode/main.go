package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	// First element in os.Args is always the program name,
	// So we need at least 2 arguments to have a file name argument.
	//if len(os.Args) < 2 {
	//	fmt.Println("Missing parameter, provide file name!")
	//	return
	//}
	//data, err := ioutil.ReadFile(os.Args[1])
	//if err != nil {
	//	fmt.Println("Can't read file:", os.Args[1])
	//	panic(err)
	//}
	// data is the file content, you can use it

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	encoded, err := stringDataToBase64(data)
	if err != nil {
		log.Fatalln(err)
	}
	println(encoded)
}

func stringDataToBase64(secretContent []byte) (string, error) {
	secret := yaml.MapSlice{}
	err := yaml.Unmarshal(secretContent, &secret)

	if err != nil {
		return "", err
	}

	var dataIndex *int
	var stringDataIndex *int

	for i, item := range secret {
		if item.Key == "data" {
			newIndex := i
			dataIndex = &newIndex
		}

		if item.Key == "stringData" {
			newIndex := i
			stringDataIndex = &newIndex

			stringData, ok := item.Value.(yaml.MapSlice)
			if !ok {
				return "", fmt.Errorf("stringData is invalid format")
			}

			for k, dataItem := range stringData {
				valueBytes := []byte(fmt.Sprintf("%v", dataItem.Value))
				stringData[k].Value = base64.StdEncoding.EncodeToString(valueBytes)
			}
		}
	}

	if stringDataIndex != nil {
		if dataIndex != nil {
			data, ok := secret[*dataIndex].Value.(yaml.MapSlice)
			if !ok {
				return "", fmt.Errorf("data is invalid format")
			}
			stringData, ok := secret[*stringDataIndex].Value.(yaml.MapSlice)
			if !ok {
				return "", fmt.Errorf("stringData is invalid format")
			}

			data = append(data, stringData...)
			secret[*dataIndex].Value = data

			secret = removeMapItemByIndex(secret, *stringDataIndex)
		} else {
			secret[*stringDataIndex].Key = "data"
		}
	}

	out, err := yaml.Marshal(secret)

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func removeMapItemByIndex(slice yaml.MapSlice, s int) yaml.MapSlice {
	return append(slice[:s], slice[s+1:]...)
}
