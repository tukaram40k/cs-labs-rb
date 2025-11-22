package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

var pStr = "32317006071311007300153513477825163362488057133489075174588434139269806834136210002792056362640164685458556357935330816928829023080573472625273554742461245741026202527916572972862706300325263428213145766931414223654220941111348629991657478268034230553086349050635557712219187890332729569696129743856241741236237225197346402691855797767976823014625397933058015226858730761197532436467475855460715043896844940366130497697812854295958659597567051283852132784468522925504568272879113720098931873959143374175837826000278034973198552060607533234122603254684088120031105907484281003994966956119696956248629032338072839127039"

func main() {
	// message
	msg := []byte("Nume Prenume")
	fmt.Printf("message %q\n\n", string(msg))

	// convert m to int
	hexStr := ""
	for _, ch := range msg {
		hexStr += fmt.Sprintf("%02x", ch)
	}
	fmt.Println("m (hex) =", hexStr)

	mInt, _ := new(big.Int).SetString(hexStr, 16)
	fmt.Println("m (decimal) =", mInt)
	fmt.Println()

	// set p and g
	p, _ := new(big.Int).SetString(pStr, 10)
	g := big.NewInt(2)

	fmt.Println("p bits =", p.BitLen())
	fmt.Println("g =", g)
	fmt.Println()

	// private key x in [1, p-2]
	x, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2)))
	if err != nil {
		log.Fatal(err)
	}
	x.Add(x, big.NewInt(1))

	// public key y = g^x mod p
	y := new(big.Int).Exp(g, x, p)

	fmt.Println("private key x =", x)
	fmt.Println("\npublic key y =", y)
	fmt.Println()

	// choose key k
	k, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2)))
	if err != nil {
		log.Fatal(err)
	}
	k.Add(k, big.NewInt(1))

	// c1 = g^k mod p
	c1 := new(big.Int).Exp(g, k, p)

	// s = y^k mod p
	s := new(big.Int).Exp(y, k, p)

	// c2 = m * y^k mod p
	c2 := new(big.Int).Mod(new(big.Int).Mul(mInt, s), p)

	fmt.Println("c1 = g^k mod p")
	fmt.Println("c1 =", c1)
	fmt.Println("\nc2 = m * s mod p")
	fmt.Println("c2 =", c2)

	// s = c1^x mod p
	sPrime := new(big.Int).Exp(c1, x, p)
	fmt.Println("\ns = c1^x mod p")
	fmt.Println("s =", sPrime)

	sInv := new(big.Int).ModInverse(sPrime, p)
	if sInv == nil {
		log.Fatal("no inverse for sPrime")
	}

	// m' = c2 * s^-1 mod p
	mPrime := new(big.Int).Mod(new(big.Int).Mul(c2, sInv), p)
	fmt.Println("\nm' = c2 * s^-1 mod p")
	fmt.Println("m' =", mPrime)

	decrypted := mPrime.Bytes()
	fmt.Printf("\ndecrypted message: %q\n", string(decrypted))
}
