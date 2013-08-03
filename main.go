package aesenv

import (
	"code.google.com/p/go.crypto/nacl/secretbox"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	suffix = "_CIPHER"
)

type Secret struct {
	Nonce [24]byte `json:"nonce"`
	Key   [32]byte `json:"key"`
}

func (s *Secret) Seal(plainText []byte) []byte {
	var cipherText []byte
	return secretbox.Seal(cipherText[:0], plainText, &s.Nonce, &s.Key)
}

func (s *Secret) Open(cipherText []byte) ([]byte, bool) {
	var plainText []byte
	return secretbox.Open(plainText[:0], cipherText, &s.Nonce, &s.Key)
}

func (s *Secret) Exec(name string, args ...string) error {
	for _, entries := range os.Environ() {
		pair := strings.SplitN(entries, "=", 2)
		cipherKey := pair[0]

		if strings.HasSuffix(cipherKey, suffix) {
			base64Text := pair[1]
			cipherText, err := base64.StdEncoding.DecodeString(base64Text)
			if err != nil {
				log.Printf("unable to decode key, %s => %s\n", cipherKey, base64Text)
				return err
			}

			key := cipherKey[0 : len(cipherKey)-len(suffix)]
			value, ok := s.Open(cipherText)
			if !ok {
				return errors.New(fmt.Sprintf("unable to decrypt %s, %s", cipherKey, base64Text))
			}
			os.Setenv(key, string(value))
		}
	}

	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)
	cmd.Wait()
	return nil
}

func NewSecret() *Secret {
	var key [32]byte
	var nonce [24]byte

	rand.Reader.Read(key[:])
	rand.Reader.Read(nonce[:])

	return &Secret{Nonce: nonce, Key: key}
}

func (s *Secret) WriteFile(filename string) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	text := base64.StdEncoding.EncodeToString(data)
	return ioutil.WriteFile(filename, []byte(text), 0644)
}

func ReadFile(filename string) (*Secret, error) {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}

	secret := new(Secret)
	err = json.Unmarshal(data, secret)
	return secret, err
}
