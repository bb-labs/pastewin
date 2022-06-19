package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestPingLambda(t *testing.T) {
	URL := "https://jzd5pgu3q3axdvxbw6kupn3bg40wbnkz.lambda-url.us-west-2.on.aws/"

	body, _ := json.Marshal(map[string]string{
		"name":  "Toby",
		"email": "Toby@example.com",
	})

	response, err := http.Post(URL, "application/json", bytes.NewBuffer(body))

	fmt.Println(err)

	defer response.Body.Close()

	bytes, _ := io.ReadAll(response.Body)

	fmt.Println(string(bytes))

}
