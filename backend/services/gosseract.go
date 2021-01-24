package services

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/otiai10/gosseract"
)

func CardImageDecrypt(imageBytes string) string {
	client := gosseract.NewClient()
	defer client.Close()

	data, err := base64.StdEncoding.DecodeString(imageBytes)
	fmt.Println(data)

	if err != nil {
		return ""
	}

	client.SetImageFromBytes(data)

	text, _ := client.Text()
	fmt.Println(text)
	err = ioutil.WriteFile("imageText.txt", []byte(text), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	return text
}
