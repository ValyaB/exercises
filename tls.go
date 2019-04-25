package router

import (
	"encoding/pem"
	"io/ioutil"

	"github.com/youmark/pkcs8"
)

func getPrivateKey(path string, passphrase []byte) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(b)
	der, err := pkcs8.ParsePKCS8PrivateKey(block.Bytes, passphrase)
	if err != nil {
		return nil, err
	}

	p, err := pkcs8.ConvertPrivateKeyToPKCS8(der)
	if err != nil {
		return nil, err
	}

	t := &pem.Block{Type: "PRIVATE KEY", Bytes: p}

	return pem.EncodeToMemory(t), nil
}

func getCert(path string) ([]byte, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return f, nil
}
