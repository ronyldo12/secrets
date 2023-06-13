package secret_file

import (
	"fmt"
	"log"
	"os"
)

func FilePath() string {
	if os.Getenv("HOME") == "" {
		log.Fatal("HOME env not found.")
	}
	return os.Getenv("HOME") + "/.secrets"
}

func FileExist() bool {
	if _, err := os.Stat(FilePath()); err != nil {
		return false
	}
	return true
}

func SaveFileContent(data string) error {
	if FileExist() == false {
		if _, err := os.Create(FilePath()); err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}
	dataByte := []byte(data)
	err := os.WriteFile(FilePath(), dataByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func FileGetContent() string {
	if FileExist() == false {
		return ""
	}

	content, err := os.ReadFile(FilePath())
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
