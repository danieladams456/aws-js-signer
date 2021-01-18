package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"testing"
)

func getTestPublicKey() (*rsa.PublicKey, error) {
	pubKeyStr := "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAo009XAA3N0XFKrrWda1gTc0ehoXmxWsL6/THCTMKWZt82uENiYXoSVoekjd7WYb5Er5oN59W36YEL65hH402I0Frw5YEDsM0epum9eJgyuZ/2bmH/DA7ncwpRXggkOfzoqjXzubCqQYVnsb9ErDzgBPlug3LjyL8ju/Oe/P9dFIXie/egdRYr3QXbRYMl3UuebtdMNbRdFaCkv8IMeYGPvrrDXE78w4o6dvsnnBDr7aHMTprCt4Yzttbigg69OpmjA4qBz4+33YHXKNtNcUb9CuVlBla2ZZkfU7BQ4TEi/87kAorXuUZDnV1atUBiwwlTxlQeZhcVrM5sFtjKDsJGij6rxvyTrsXTjSH/UhCHr4e2tbDutVdvuG+NVjAWhlMHGJGbjphxwAbB7KX5JqSMpYsKFfp5QTHA4Hbiu5xaBSi8jzDanDfzufHXQvuf1cpkShxjSc2gyiYaflodjxiQ2OJr3MtzRxm977aOyJmytEuY7Zpya9lmwzAv/6uIrmAu+iFq5cje46kwES16dsGYI8v1YSSc9vIGAzhl27qNTjEnjmYPcyLeL5MiuGk+3eh4asRAZMvZ+vvIAnuhM9N3cCmK4Tn0L6vAwmosQrOf3H11QxXD3KQGP8sJXAxxTJ7pb4GKYITZW2sJ3Jaifi/04MBtm0n4N8xCGUxO4tiXDkCAwEAAQ=="
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyStr)
	if err != nil {
		return nil, err
	}
	pubKey, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return nil, err
	}
	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("could not cast to *rsa.PublicKey")
	}
	return rsaPubKey, nil
}

func TestVerifySignature(t *testing.T) {
	message := "testmessage"
	signature := "WCw/zVvnwKv5WXFWY7/HGVatOpMsOakw53yi+9a32G+ogSM7JgaG046JiOVVI/F0svCcGEmJ38FI1+/7QsJpyShP1nGBaen/Ydov+ddG3ISGU53zhM1Gq5i6pBiXSjTYecTQpOLE27zoS967CglS1XMmg+mHnigis25rksMcJwbuvP9PZ0wMv7YUjEth4FAA49YPNBw9To7vyLPZpAwDndZ/PUFSJhA5CMruWN1HvnRXlBkOC7LN4FZfUJ0UpxisycKEcJf6LcvqTstzovKmo+84Zp6ArQ3ODBcIBaSpF5Rbwo1nDUGTTgHF8XYRVy9uCEMSS4z/jSjaiUtLgzEE5AcetTEMl5ZWPd71MReiI0lEfDBC1gKek+7oFEfcq54Fmf0jdzgeKCxiTSegIbd4gNqYvAks+JWJzEhTHTZWbE0HHZ5QWAjus9V/gRbrBegOuVdXRSY1O7XAlqhIJdSd/RG4147t3a9c8oWW71/CILIFbIT7ZuaBGCsfD1q71maaz5rhe3TH5OGh58WWfgas3lxnoxyQ2P1s7MjhSCanVM5+H4XDfsFbP3zPHQnsAhczFNnDctUfGkO4SRnNrh4Tj6L7jIin8ERXLPSMzHg6x+8gQXzRy0KnHJfaHQNc1W4zctXh7Bqr8M0h02ErBaD6WkCv6oPl7+aPNYs3ohOrg0Q="
	pubKey, _ := getTestPublicKey()
	err := verifySignature(message, signature, pubKey)
	if err != nil {
		t.Error(err)
	}
}

func Test_verifyJWT(t *testing.T) {
	jwt := "eyJ0eXAiOiAiSldUIiwgImFsZyI6ICJSUzI1NiIsICJraWQiOiAiZDE1ODNiMTEtZjdmNy00OWI2LWEzYzctMDFmYmFlNTY5MTVmIn0.eyJpYXQiOiAxNjExMDExMzE4LCAiZXhwIjogMTYxMTAxNDkxOCwgImFybiI6ICJhcm46YXdzOnN0czo6OTMyNzUyMzA2NTk3OmFzc3VtZWQtcm9sZS9zYW1sLUFkbWluL0RhbmllbC5BZGFtc0BnZW53b3J0aC5jb20ifQ.gWWCX1XCEwDyhghrm-_SSqc-0oJ7TinUxJAUxYbs_FglIEX5v89O1tLmVtivwUNDezkpBctDT6_8GznFv9Sj_nJIlXX3Au9dYJQ-9VTAYFJ3X7Vy_f-aZAORG4ICTX1VIBs8oFUkcvET5Et1lW4Ths185Qi3IfLpIFZAmTihNiPbmYgQu6gnBL8Pxfb8hK9j5kUV0DilAo7qN7eV6gpO3fOKrKunYqbzPWIVQ9vSZYxJcsOtlD8xmza0g-0hfPQx-LjvfENea3Z3_bb9wRkiVARoOpzJAVWHXdwTwgzqM2-9pGr829oeDg83Q3XB3o0pwLoD0J_M1WPOet7fT0TmTz-rpyxKZnwhc3B7bcfqNqmXICH9PJ_UEyP5HivE18wdMNsZLwCnhvFi7InUroIXf21H_2d9zlZUVN6u4X_bJRLBDKwibKPVTSeBvIp4OZZdy4s5pCIEETzKb9BPryXU-ANvI8XnxmoCoEs-3eKFP2ke4DYCGMUr8abQmD8A3GOWGFL2-ChbZlZKI__ZeMWmED3OV-OAh8lceX4s-izVpsbFU83IGiTDszHeLyAKYi6yvoDXJx4bS1CpUT2rdjqRNloQDrSElByjYxYdOAXiTvP-KIDk1Ikfcs1QHpvCWXDBIsdBvjmj0siEtJawW8OPYbADWAzNWUk7iDzrQfPfSzE"
	_, err := verifyJWT(jwt)
	if err != nil {
		t.Error(err)
	}
}
