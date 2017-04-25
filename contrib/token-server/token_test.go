package main

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	"github.com/docker/distribution/registry/auth"
	"github.com/docker/libtrust"
)

func TestTokenIssuerCreateJWT(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal(err)
	}
	pk, err := libtrust.FromCryptoPrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	tokenIssuer := TokenIssuer{
		Expiration: time.Duration(100),
		Issuer:     "localhost",
		SigningKey: pk,
	}

	grantedAccessList := make([]auth.Access, 0, 0)
	token, err := tokenIssuer.CreateJWT("test", "test", grantedAccessList)

	if len(token) == 0 {
		t.Fatal("token not generated.")
	}

}
