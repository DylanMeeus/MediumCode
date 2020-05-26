package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	// paste your api key and user token here
	key  = ""
	user = ""
	url  = "https://3345423f-8bbf-4215-8004-215a3da67b60.users.dploy.ai/"
)

type Payload struct {
	Image string `json:"image"`
	Type  string `json:"type"`
}

func loadImage() (Payload, error) {
	file := os.Args[1]
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return Payload{}, err
	}
	return Payload{
		Image: base64.StdEncoding.EncodeToString(bs),
		Type:  strings.Split(file, ".")[1],
	}, nil
}

func main() {
	payload, err := loadImage()
	if err != nil {
		panic(err)
	}
	c := http.Client{}

	data, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(data)))
	// set our API keys
	req.Header.Set("x-api-key", key)
	req.Header.Set("x-api-user", user)
	req.Header.Set("content-type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ResponseBody := struct {
		// We only want the image, but we could also get the confidence and the outcome
		// (maksed or not) as a string array.
		AnnotatedImage string `json:"annotated_image"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&ResponseBody); err != nil {
		panic(err)
	}
	WriteResponse(ResponseBody.AnnotatedImage, fmt.Sprintf("output.%s", payload.Type))
	fmt.Println("done!")
}

// save the response as an image
func WriteResponse(image, outputfile string) error {
	bs, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		return err
	}
	ioutil.WriteFile(outputfile, bs, 0644)
	return nil
}
