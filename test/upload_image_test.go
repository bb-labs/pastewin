package test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestUploadImage(t *testing.T) {
	URL := "https://jzd5pgu3q3axdvxbw6kupn3bg40wbnkz.lambda-url.us-west-2.on.aws/"

	raw, _ := os.ReadFile("/Users/trumanpurnell/Desktop/vanessa.pdf")
	response, err := http.Post(URL, "application/pdf", bytes.NewBuffer([]byte(raw)))

	fmt.Println(err)

	defer response.Body.Close()

	bytes, _ := io.ReadAll(response.Body)

	fmt.Println(string(bytes))

}
