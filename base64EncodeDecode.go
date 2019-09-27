package main

import (
	"encoding/base64"
	"fmt"
)

/**
 * base64编解码
 */
func main() {
	message := "hello, this is a mail"
	encodeMsg := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encodeMsg)

	data, err := base64.StdEncoding.DecodeString(encodeMsg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
