package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func (t inventoryConfig) toString() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}

func getInventoryConfig() []inventoryConfig {
	myconfig := make([]inventoryConfig, 3)
	raw, err := ioutil.ReadFile("inventoryConfig.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &myconfig)
	return myconfig
}

func main() {
	myConfigs := getInventoryConfig()
	fmt.Println(myConfigs)
	for _, myconfig := range myConfigs {
		fmt.Println(myconfig.toString())
		var files []string
		path := myconfig.Path
		err := filepath.Walk(path, visit(&files))
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			fmt.Println(file)
		}
	}

}
