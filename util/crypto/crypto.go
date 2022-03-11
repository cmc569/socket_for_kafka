package crypto

import (
	"embed"
	"github.com/wenzhenxi/gorsa"
	"net/url"
)

//go:embed *
var f embed.FS

func KeyEncrypt(cryptoText string) (string, error) {
	publicKey, err := f.ReadFile("public.pem")
	pubEncryptStr, err := gorsa.PublicEncrypt(cryptoText, string(publicKey))
	if err != nil {
		return "", err
	}
	return pubEncryptStr, err
}

func KeyDecrypt(cryptoText string) (string, error) {
	privateKey, err := f.ReadFile("private.pem")
	priDecrypt, err := gorsa.PriKeyDecrypt(cryptoText, string(privateKey))
	if err != nil {
		return "", err
	}
	text, err := url.QueryUnescape(priDecrypt)
	return text, err
}
