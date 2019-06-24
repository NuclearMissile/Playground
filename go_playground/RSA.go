package main

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/NuclearMissile/Playground/go_playground/Mont"
	"math/big"
)

func String2BigInt(s *string) *big.Int {
	var number, temp big.Int
	for _, b := range []byte(*s) {
		number.Or(number.Lsh(&number, 8), temp.SetInt64(int64(b)))
	}
	return &number
}

func BigInt2String(n *big.Int) *string {
	var db [16]byte
	var temp big.Int
	dx := 16
	bff := big.NewInt(0xff)
	for n.BitLen() > 0 {
		dx--
		db[dx] = byte(temp.And(n, bff).Int64())
		n.Rsh(n, 8)
	}
	s := string(db[dx:])
	return &s
}

func RSAEncrypt(plain *string, e, n *big.Int) (*big.Int, error) {
	pAsBigInt := String2BigInt(plain)
	if pAsBigInt.Cmp(n) >= 0 {
		return nil, errors.New("Plain text too long. ")
	}
	return Mont.Exp(pAsBigInt, e, n), nil
}

func RSADecrypt(encrypted, d, n *big.Int) *string {
	temp := Mont.Exp(encrypted, d, n)
	return BigInt2String(temp)
}

func main() {
	pkey, _ := rsa.GenerateKey(rand.Reader, 1024)
	n := pkey.N
	e := big.NewInt(int64(pkey.E))
	d := pkey.D

	plain := "Hello, World!"
	fmt.Println("Plain text:")
	fmt.Println(plain)

	plainAsBigInt := String2BigInt(&plain)
	fmt.Println("Plain text as number:")
	fmt.Println(plainAsBigInt.String())

	fmt.Println("Encrypted:")
	encrypted, err := RSAEncrypt(&plain, e, n)
	if err != nil {
		panic(err)
	}
	fmt.Println(encrypted.String())

	fmt.Println("Decrypted:")
	fmt.Println(*RSADecrypt(encrypted, d, n))
}
