package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	message := []byte("testmessage")
	pubKeyStr := "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAo009XAA3N0XFKrrWda1gTc0ehoXmxWsL6/THCTMKWZt82uENiYXoSVoekjd7WYb5Er5oN59W36YEL65hH402I0Frw5YEDsM0epum9eJgyuZ/2bmH/DA7ncwpRXggkOfzoqjXzubCqQYVnsb9ErDzgBPlug3LjyL8ju/Oe/P9dFIXie/egdRYr3QXbRYMl3UuebtdMNbRdFaCkv8IMeYGPvrrDXE78w4o6dvsnnBDr7aHMTprCt4Yzttbigg69OpmjA4qBz4+33YHXKNtNcUb9CuVlBla2ZZkfU7BQ4TEi/87kAorXuUZDnV1atUBiwwlTxlQeZhcVrM5sFtjKDsJGij6rxvyTrsXTjSH/UhCHr4e2tbDutVdvuG+NVjAWhlMHGJGbjphxwAbB7KX5JqSMpYsKFfp5QTHA4Hbiu5xaBSi8jzDanDfzufHXQvuf1cpkShxjSc2gyiYaflodjxiQ2OJr3MtzRxm977aOyJmytEuY7Zpya9lmwzAv/6uIrmAu+iFq5cje46kwES16dsGYI8v1YSSc9vIGAzhl27qNTjEnjmYPcyLeL5MiuGk+3eh4asRAZMvZ+vvIAnuhM9N3cCmK4Tn0L6vAwmosQrOf3H11QxXD3KQGP8sJXAxxTJ7pb4GKYITZW2sJ3Jaifi/04MBtm0n4N8xCGUxO4tiXDkCAwEAAQ=="
	pubKey, err := parsePublicKey(pubKeyStr)
	if err != nil{
		fmt.Fprintf(os.Stderr, "Error parsing public key: %s\n", err)
	}

	signature, err := hex.DecodeString("WCw/zVvnwKv5WXFWY7/HGVatOpMsOakw53yi+9a32G+ogSM7JgaG046JiOVVI/F0svCcGEmJ38FI1+/7QsJpyShP1nGBaen/Ydov+ddG3ISGU53zhM1Gq5i6pBiXSjTYecTQpOLE27zoS967CglS1XMmg+mHnigis25rksMcJwbuvP9PZ0wMv7YUjEth4FAA49YPNBw9To7vyLPZpAwDndZ/PUFSJhA5CMruWN1HvnRXlBkOC7LN4FZfUJ0UpxisycKEcJf6LcvqTstzovKmo+84Zp6ArQ3ODBcIBaSpF5Rbwo1nDUGTTgHF8XYRVy9uCEMSS4z/jSjaiUtLgzEE5AcetTEMl5ZWPd71MReiI0lEfDBC1gKek+7oFEfcq54Fmf0jdzgeKCxiTSegIbd4gNqYvAks+JWJzEhTHTZWbE0HHZ5QWAjus9V/gRbrBegOuVdXRSY1O7XAlqhIJdSd/RG4147t3a9c8oWW71/CILIFbIT7ZuaBGCsfD1q71maaz5rhe3TH5OGh58WWfgas3lxnoxyQ2P1s7MjhSCanVM5+H4XDfsFbP3zPHQnsAhczFNnDctUfGkO4SRnNrh4Tj6L7jIin8ERXLPSMzHg6x+8gQXzRy0KnHJfaHQNc1W4zctXh7Bqr8M0h02ErBaD6WkCv6oPl7+aPNYs3ohOrg0Q")
	if err != nil{
		fmt.Fprintf(os.Stderr, "Error decoding signature: %s\n", err)
	}

	// Only small messages can be signed directly; thus the hash of a
	// message, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should be used for this at the time
	// of writing (2016).
	hashed := sha512.Sum512(message)

	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA512, hashed[:], signature)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return
	}
}
