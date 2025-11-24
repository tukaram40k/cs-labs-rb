package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var pStr = "32317006071311007300153513477825163362488057133489075174588434139269806834136210002792056362640164685458556357935330816928829023080573472625273554742461245741026202527916572972862706300325263428213145766931414223654220941111348629991657478268034230553086349050635557712219187890332729569696129743856241741236237225197346402691855797767976823014625397933058015226858730761197532436467475855460715043896844940366130497697812854295958659597567051283852132784468522925504568272879113720098931873959143374175837826000278034973198552060607533234122603254684088120031105907484281003994966956119696956248629032338072839127039"

func main() {
	p := new(big.Int)
	p.SetString(pStr, 10)

	g := big.NewInt(2)

	// private keys a, b in [2, p-2]
	a, _ := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2)))
	a.Add(a, big.NewInt(1))

	b, _ := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2)))
	b.Add(b, big.NewInt(1))

	// public keys A = g^a mod p, B = g^b mod p
	A := new(big.Int).Exp(g, a, p)
	B := new(big.Int).Exp(g, b, p)

	// shared secret
	Ka := new(big.Int).Exp(B, a, p)

	// convert shared key to 32-byte AES key
	keyBytes := Ka.Bytes()
	keyBytes = keyBytes[len(keyBytes)-32:]

	fmt.Println("private a:", a)
	fmt.Println("private b:", b)
	fmt.Println("\npublic A:", A)
	fmt.Println("public B:", B)
	fmt.Println("\nshared key:", Ka)
	fmt.Println("\nAES key (32 bytes): ", keyBytes)
}
