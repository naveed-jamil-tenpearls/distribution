package main

import (
	"testing"
	"github.com/docker/libtrust"
	"time"
	"crypto/rsa"
	"crypto/rand"
	"github.com/docker/distribution/registry/auth"
)

func TestTokenIssuer_CreateJWT(t *testing.T) {
	d := time.Duration(100)
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	pk, err := libtrust.FromCryptoPrivateKey(key)
	if err != nil {
		panic(err)
	}
	tokenIssuer := TokenIssuer{
		Expiration: d,
		Issuer: "localhost",
		SigningKey: pk,

	}

	grantedAccessList := make([]auth.Access, 0,0)
	tokenIssuer.CreateJWT("test", "test", grantedAccessList)

}
