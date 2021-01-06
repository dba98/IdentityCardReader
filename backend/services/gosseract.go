package services

import (
	"fmt"
	"io/ioutil"

	"github.com/otiai10/gosseract"
)

func CardImageDecrypt() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("test.jpg")
	text, _ := client.Text()
	fmt.Println(text)
	err := ioutil.WriteFile("imageText.txt", []byte(text), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	// Hello, World!
}
