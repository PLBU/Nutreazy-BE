package utils

import (
	"fmt"
	"log"
)

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Log(text string) {
	fmt.Println(text)
}