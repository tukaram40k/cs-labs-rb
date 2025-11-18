package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func main() {
	// message
	msg := []byte("Nume Prenume")
	fmt.Printf("message %q\n\n", string(msg))

	// convert m to int
	hexStr := ""
	for _, ch := range msg {
		hexStr += fmt.Sprintf("%02x", ch)
	}

	fmt.Println("m (hex) = ", hexStr)

	mInt, _ := new(big.Int).SetString(hexStr, 16)

	fmt.Println("m (decimal) = ", mInt)
	fmt.Println()

	// chose p and q
	p, _ := rand.Prime(rand.Reader, 1022)
	q, _ := rand.Prime(rand.Reader, 1036)

	fmt.Println("p = ", p)
	fmt.Println("q = ", q)
	fmt.Println()

	// n = p * q
	fmt.Println("n = p * q")
	n := new(big.Int).Mul(p, q)
	fmt.Printf("n is %d bits\n", n.BitLen())
	fmt.Println("n = ", n)
	fmt.Println()

	// phi = (p-1)*(q-1)
	fmt.Println("phi = (p-1)*(q-1)")
	one := big.NewInt(1)
	pMinus1 := new(big.Int).Sub(p, one)
	qMinus1 := new(big.Int).Sub(q, one)
	phi := new(big.Int).Mul(pMinus1, qMinus1)
	fmt.Println("phi = ", phi)
	fmt.Println()

	// choose public key e
	fmt.Println("public key e = 65537")
	e := big.NewInt(65537)
	//fmt.Println("e:", e)
	fmt.Println()

	// d = e^{-1} mod phi
	fmt.Println("d = e^{-1} mod phi")
	d := new(big.Int).ModInverse(e, phi)
	if d == nil {
		log.Fatalf("no modular inverse for e modulo phi")
	}
	fmt.Println("d = ", d)
	fmt.Println()

	// c = m^e mod n
	fmt.Println("c = m^e mod n")
	c := new(big.Int).Exp(mInt, e, n)
	fmt.Println("c = ", c)

	// m' = c^d mod n
	fmt.Println("m' = c^d mod n")
	mPrime := new(big.Int).Exp(c, d, n)
	fmt.Println("m' = ", mPrime)

	// convert to string
	decrypted := mPrime.Bytes()
	fmt.Println()
	fmt.Printf("decrypted message: %q\n", string(decrypted))
}
