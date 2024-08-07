package gob_cache

import (
	"encoding/gob"
	"fmt"
	"os"
)

func Set(key string, value interface{}) error {
	path := "./cache/"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return err
	}
	file, err := os.Create(path + key + ".gob")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(value)
}
func Get(key string, value interface{}) error {
	path := "./cache/"
	file, err := os.Open(path + key + ".gob")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(value)
}
