package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
)

var pStr = "32317006071311007300153513477825163362488057133489075174588434139269806834136210002792056362640164685458556357935330816928829023080573472625273554742461245741026202527916572972862706300325263428213145766931414223654220941111348629991657478268034230553086349050635557712219187890332729569696129743856241741236237225197346402691855797767976823014625397933058015226858730761197532436467475855460715043896844940366130497697812854295958659597567051283852132784468522925504568272879113720098931873959143374175837826000278034973198552060607533234122603254684088120031105907484281003994966956119696956248629032338072839127039"

func main() {
	// hash
	hashStr := "R7CMxilhk+QZWdk6h3OWlQO0fgnw2y5G9bY/45bsJxE="
	hashBytes, _ := base64.StdEncoding.DecodeString(hashStr)
	h := new(big.Int).SetBytes(hashBytes)

	p, _ := new(big.Int).SetString(pStr, 10)
	g := big.NewInt(2)

	h.Mod(h, p)

	// private key a in [2, p-2]
	a, _ := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2)))
	a.Add(a, big.NewInt(1)) // ensure >= 1

	// public key A = g^a mod p
	A := new(big.Int).Exp(g, a, p)

	// choose k such that gcd(k, p-1)=1
	var k, kInv *big.Int
	pp1 := new(big.Int).Sub(p, big.NewInt(1))
	for {
		k, _ = rand.Int(rand.Reader, pp1)
		k.Add(k, big.NewInt(1))
		if new(big.Int).GCD(nil, nil, k, pp1).Cmp(big.NewInt(1)) == 0 {
			kInv = new(big.Int).ModInverse(k, pp1)
			if kInv != nil {
				break
			}
		}
	}

	// r = g^k mod p
	r := new(big.Int).Exp(g, k, p)

	// s = (h - a*r) * k^{-1} mod (p-1)
	s := new(big.Int).Mul(a, r)
	s.Mod(s, pp1)
	s.Sub(h, s)
	s.Mod(s, pp1)
	s.Mul(s, kInv)
	s.Mod(s, pp1)

	fmt.Println("signature r =", r)
	fmt.Println("signature s =", s)
	fmt.Println()

	// check A^r * r^s mod p == g^h mod p
	left1 := new(big.Int).Exp(A, r, p)
	left2 := new(big.Int).Exp(r, s, p)
	left := new(big.Int).Mul(left1, left2)
	left.Mod(left, p)

	right := new(big.Int).Exp(g, h, p)

	fmt.Println("v1 =", left)
	fmt.Println("v2 =", right)
	fmt.Println()

	if left.Cmp(right) == 0 {
		fmt.Println("signature is valid")
	} else {
		fmt.Println("signature is invalid")
	}
}
