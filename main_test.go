package aesenv

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSaveAndLoad(t *testing.T) {
	filename := fmt.Sprintf("data-%d.secret", time.Now().Unix())
	secret := NewSecret()

	secret.WriteFile(filename)
	actual, err := ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	if actual.Nonce != secret.Nonce {
		t.Fatal("expected Nonce to match")
	}
	if actual.Key != secret.Key {
		t.Fatal("expected Key to match")
	}

	os.Remove(filename)
}

func TestEncryptAndDecrypt(t *testing.T) {
	secret := NewSecret()
	plainText := make([]byte, 128)
	rand.Reader.Read(plainText)
	cipherText := secret.Seal(plainText)
	recovertedText, ok := secret.Open(cipherText)

	// Then
	if !ok {
		t.Fatal("expected Open to succeed")
	}
	if bytes.Compare(recovertedText, plainText) != 0 {
		t.Fatalf("expected recoverted data to match the original")
	}
}

func TestExec(t *testing.T) {
	secret := NewSecret()
	cipherText := secret.Seal([]byte("hello world"))
	text := base64.StdEncoding.EncodeToString(cipherText)
	os.Setenv("HELLO"+suffix, text)

	err := secret.Exec("/bin/echo", "hello world")
	if err != nil {
		t.Fatal(err)
	}
}
