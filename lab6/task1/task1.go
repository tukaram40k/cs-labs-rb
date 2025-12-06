package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
)

func main() {
	// set hash
	hashStr := "R7CMxilhk+QZWdk6h3OWlQO0fgnw2y5G9bY/45bsJxE="
	hashBytes, _ := base64.StdEncoding.DecodeString(hashStr)
	fmt.Println("hash string:", hashStr)

	// create big.Int from bytes
	hInt := new(big.Int).SetBytes(hashBytes)
	fmt.Println()

	// choose p and q
	p, _ := rand.Prime(rand.Reader, 1536)
	q, _ := rand.Prime(rand.Reader, 1538)

	// n = p * q
	n := new(big.Int).Mul(p, q)
	fmt.Printf("n is %d bits\n", n.BitLen())
	fmt.Println()

	// phi = (p-1)*(q-1)
	one := big.NewInt(1)
	phi := new(big.Int).Mul(new(big.Int).Sub(p, one), new(big.Int).Sub(q, one))

	// public key e
	e := big.NewInt(65537)

	// compute private key d
	d := new(big.Int).ModInverse(e, phi)

	// signature s = h^d mod n
	s := new(big.Int).Exp(hInt, d, n)
	fmt.Println("signature =", s)
	fmt.Println()

	// verify v = s^e mod n
	v := new(big.Int).Exp(s, e, n)
	fmt.Println("original hash as integer =", hInt)
	fmt.Println("verified integer =", v)
	fmt.Println()

	// compare v with original hash int
	if v.Cmp(hInt) == 0 {
		fmt.Println("signature is valid")
	} else {
		fmt.Println("signature is invalid")
	}
}
