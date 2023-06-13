package masterkey

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/term"
)

const MasterKeyIndex string = "SCTRS_MK_INDEX_2023123"

func IsDefined() bool {
	masterKey := os.Getenv(MasterKeyIndex)
	return len(masterKey) != 0
}
func Masterkey() string {
	Check()
	return os.Getenv(MasterKeyIndex)
}

func Check() {
	if !IsDefined() {
		log.Fatal("You should login!")
	}
}

func Read() (string, error) {
	fmt.Print("type you Master Key: ")
	byteValue, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	return string(byteValue), nil
}

func Prepare(masterkey string) string {
	hsha1 := sha256.Sum256([]byte(masterkey))
	return fmt.Sprintf("%x", hsha1)[:8] + fmt.Sprintf("%x", hsha1)[56:]
}

func Define(masterkey string) {
	masterkey = Prepare(masterkey)
	os.Setenv(MasterKeyIndex, masterkey)
	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
}
