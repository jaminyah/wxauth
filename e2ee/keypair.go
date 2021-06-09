package e2ee

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
)

func SendPublicKey(w http.ResponseWriter, r *http.Request) {

	publicKey := fetchPubKey()

	body := map[string]interface{}{"pubkey": publicKey, "msg": "success"}

	//set json response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}

func fetchPubKey() []byte {

	pubkey, err := os.ReadFile("pubkey.pem")
	if err != nil {
		fmt.Println("Cannot read pubkey.pem")
	}

	block, _ := pem.Decode([]byte(pubkey))
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("Failed: " + err.Error())
	}
	rsaPubKey, _ := pub.(*rsa.PublicKey)
	if err != nil {
		fmt.Println("Unable to convert to rsa public key")
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(rsaPubKey)
	if err != nil {
		fmt.Println("Cannot Marshal pubkey")
	}
	return publicKeyBytes
}

func FetchPriKey() *rsa.PrivateKey {

	prikey, err := os.ReadFile("prikey.pem")
	if err != nil {
		fmt.Println("Cannot read prikey.pem")
	}

	block, _ := pem.Decode([]byte(prikey))
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Failed: " + err.Error())
	}

	if err != nil {
		fmt.Println("Unable to convert to rsa private key")
	}

	//fmt.Print("Private key")
	//fmt.Print(pri)

	return pri
}
