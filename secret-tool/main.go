package main

import (
	aesenv ".."
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	filename = flag.String("filename", "/etc/secret", "reads the specified secret file")
	create   = flag.String("create", "", "creates a new secret and writes it to the specified file")
	encrypt  = flag.Bool("encrypt", false, "encrypt data from stdin")
	decrypt  = flag.Bool("decrypt", false, "decrypt data from stdin")
)

func main() {
	flag.Parse()

	if *create != "" {
		secret := aesenv.NewSecret()
		secret.WriteFile(*create)
		return
	}

	secret, err := aesenv.ReadFile(*filename)
	if err != nil {
		log.Fatalln(err)
	}

	if *encrypt || *decrypt {
		input := bufio.NewReader(os.Stdin)
		line, _, _ := input.ReadLine()
		text := strings.TrimSpace(string(line))

		if *encrypt {
			cipherText := secret.Seal([]byte(text))
			base64Text := base64.StdEncoding.EncodeToString(cipherText)
			fmt.Println(base64Text)

		} else {
			cipherText, err := base64.StdEncoding.DecodeString(text)
			if err != nil {
				log.Fatal(err)
			}
			plainText, ok := secret.Open(cipherText)
			if !ok {
				log.Fatal("unable to open secret")
			}
			fmt.Println(string(plainText))
		}
		return
	}

	command := flag.Args()[0]
	args := flag.Args()[1:]

	secret.Exec(command, args...)
}
