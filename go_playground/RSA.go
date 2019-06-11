package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n, e, d, bb, ptn, etn, dtn big.Int
	plain := "Hello, World!"
	fmt.Println("Plain text:")
	fmt.Println(plain)

	n.SetString("9516311845790656153499716760847001433441357", 10)
	e.SetString("65537", 10)
	d.SetString("5617843187844953170308463622230283376298685", 10)

	for _, b := range []byte(plain) {
		ptn.Or(ptn.Lsh(&ptn, 8), bb.SetInt64(int64(b)))
	}
	if ptn.Cmp(&n) >= 0 {
		fmt.Println("Plain text too long.")
		return
	}
	fmt.Println("Plain text as number:")
	fmt.Println(&ptn)
	etn.Exp(&ptn, &e, &n)
	fmt.Println("Encoded:")
	fmt.Println(&etn)
	dtn.Exp(&etn, &d, &n)
	fmt.Println("Decoded::")
	fmt.Println(&dtn)

	var db [16]byte
	dx := 16
	bff := big.NewInt(0xff)
	for dtn.BitLen() > 0 {
		dx--
		db[dx] = byte(bb.And(&dtn, bff).Int64())
		dtn.Rsh(&dtn, 8)
	}
	fmt.Println("Decoded text:")
	fmt.Println(string(db[dx:]))
}

/*func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Printf("rsa.GenerateKey: %v\n", err)
	}

	message := "Hello World!"
	messageBytes := bytes.NewBufferString(message)
	sha1 := sha1.New()

	encrypted, err := rsa.EncryptOAEP(sha1, rand.Reader, &privateKey.PublicKey, messageBytes.Bytes(), nil)
	if err != nil {
		fmt.Printf("EncryptOAEP: %s\n", err)
	}

	decrypted, err := rsa.DecryptOAEP(sha1, rand.Reader, privateKey, encrypted, nil)
	if err != nil {
		fmt.Printf("decrypt: %s\n", err)
	}

	decryptedString := bytes.NewBuffer(decrypted).String()
	fmt.Printf("message: %v\n", message)
	fmt.Printf("encrypted: %v\n", encrypted)
	fmt.Printf("decryptedString: %v\n", decryptedString)
}*/
