package services

import (
	"IdentityCardReader/backend/model"
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/otiai10/gosseract"
)

func CardImageDecrypt(imageBytes string, identityCard *model.IdentityCard, isBackData bool) string {
	client := gosseract.NewClient()
	defer client.Close()

	data, err := base64.StdEncoding.DecodeString(imageBytes)
	//fmt.Println(data)

	if err != nil {
		return ""
	}

	client.SetImageFromBytes(data)

	text, _ := client.Text()
	//fmt.Println(text)
	err = ioutil.WriteFile("imageText.txt", []byte(text), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	if isBackData {

		file, err := os.Open("imageText.txt")

		if err != nil {
			fmt.Printf("failed to open file: %v", err)

		}

		scanner := bufio.NewScanner(file)

		// The bufio.ScanLines is used as an
		// input to the method bufio.Scanner.Split()
		// and then the scanning forwards to each
		// new line using the bufio.Scanner.Scan()
		// method.
		scanner.Split(bufio.ScanLines)
		var line string
		var j = 0
		for scanner.Scan() {
			fmt.Println(line)
			line = scanner.Text()
			if j > 0 && j < 4 {
				if len(line) <= 1 {
					continue
				}
				_, err := strconv.Atoi(line[0:1])
				fmt.Println(j)
				fmt.Println(line[0:1])
				if err == nil {
					identityCard.Nif = line[0:9]
				}

			}
			j++

		}
		file.Close()
	}
	return text
}
