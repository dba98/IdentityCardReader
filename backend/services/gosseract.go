package services

import (
	"fmt"
	"github.com/otiai10/gosseract"
)

func CardImageDecrypt(){
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("C:\\Users\\dgbor\\Documents\\image.png")
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!
}