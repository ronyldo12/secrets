package secret

import (
	"encoding/json"
	"log"

	"github.com/ronyldo12/secrets/crypt"
	"github.com/ronyldo12/secrets/masterkey"
	"github.com/ronyldo12/secrets/secret_file"
)

type SecretsData struct {
	Secrets map[string]string `json:"secrets"`
}

func (s *SecretsData) Save() {
	data, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	dataEncrypted, err := crypt.Encrypt(string(data), masterkey.Masterkey())
	if err != nil {
		log.Fatal(err)
	}
	secret_file.SaveFileContent(dataEncrypted)
}

func GetSecretsData() *SecretsData {
	masterkey.Check()
	data := secret_file.FileGetContent()
	secretsData := SecretsData{
		Secrets: map[string]string{},
	}
	if data == "" {
		return &secretsData
	}
	data, err := crypt.Decrypt(data, masterkey.Masterkey())
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(data), &secretsData)
	if err != nil {
		log.Fatal(err)
	}
	return &secretsData
}
