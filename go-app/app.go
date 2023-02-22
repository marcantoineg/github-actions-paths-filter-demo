package app

import "os"

const filePath = "../super-important-data.txt"

func app() string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data) + ", in golang."
}
