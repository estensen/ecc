package main

import (
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
	c := crypto.S256()
	params := c.Params()
	fmt.Printf("P=%d\n", params.P)   // Order of curve (prime number)
	fmt.Printf("N=%d\n", params.N)   // Order of base point
	fmt.Printf("Gx=%d\n", params.Gx) // x of base point
	fmt.Printf("Gy=%d\n", params.Gy) // y of base point

	r, err := rand.Int(rand.Reader, params.P)
	if err != nil {
		log.Fatal("Could not generate random number")
	}
	p1x, p1y := c.ScalarMult(params.Gx, params.Gy, r.Bytes())
	fmt.Println("Private key")
	fmt.Printf("n=%s\n", r)
	fmt.Println("Public key point")
	fmt.Printf("Px=%d\nPy=%d\n", p1x, p1y)
	fmt.Printf("Is on curve: %t\n", c.IsOnCurve(p1x, p1y))

	r1, err := rand.Int(rand.Reader, params.P)
	if err != nil {
		log.Fatal("Could not generate random number")
	}
	p2x, p2y := c.ScalarBaseMult(r1.Bytes())
	// or
	// p2x, p2y := c.ScalarMult(params.Gx, params.Gy, r1.Bytes())
	fmt.Println("New point")
	fmt.Printf("Px=%d\nPy=%d\n", p2x, p2y)
	fmt.Printf("Is on curve: %t\n", c.IsOnCurve(p2x, p2y))

	p3x, p3y := c.Add(p1x, p1y, p2x, p2y)
	fmt.Println("Result of adding p1 and p2")
	fmt.Printf("Px=%d\nPy=%d\n", p3x, p3y)
	fmt.Printf("Is on curve: %t\n", c.IsOnCurve(p3x, p3y))
}

