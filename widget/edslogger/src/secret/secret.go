package secret

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

// go run main 的相对路径
const secretFile = "widget/secretdata/eds.json"

type Secret struct {
	UserId  string
	UserPsd string
	Cookie  string
}

var secret *Secret

func RetrieveSecret() (*Secret, error) {
	file, err := os.Open(secretFile)
	log.Println(os.Getwd())
	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&secret)

	return secret, err
}

func Parse(str *string) *Secret {
	r := strings.NewReader(*str)
	err := json.NewDecoder(r).Decode(&secret)
	if err != nil {
		log.Fatal(err)
	}
	return secret
}
