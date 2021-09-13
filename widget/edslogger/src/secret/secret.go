package secret

import (
	"encoding/json"
	"os"
)

const secretFile = "widget/secretdata/eds.json"

type Secret struct {
	UserId  string
	UserPsd string
	Cookie  string
}

func RetrieveSecret() (*Secret, error) {
	file, err := os.Open(secretFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var secret *Secret
	err = json.NewDecoder(file).Decode(&secret)

	return secret, err
}
