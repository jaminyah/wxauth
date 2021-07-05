package e2ee

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
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

func fetchPriKey() *rsa.PrivateKey {

	prikey, err := os.ReadFile("prikey.pem")
	if err != nil {
		fmt.Println("Cannot read prikey.pem")
	}

	block, _ := pem.Decode([]byte(prikey))
	rsaPrivate, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Failed: " + err.Error())
	}

	if err != nil {
		fmt.Println("Unable to convert to rsa private key")
	}

	return rsaPrivate
}

func DecodeRSA(encoded string) string {

	cipherText, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Println(err)
	}

	privateKey := fetchPriKey()
	data, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		log.Println(err)
	}

	return string(data)
}

func PerformHash(passwd string) (string, error) {

	var passwdBytes = []byte(passwd)
	hashedBytes, err := bcrypt.GenerateFromPassword(passwdBytes, bcrypt.MinCost)

	return string(hashedBytes), err
}

func ComparePass(dbHash, loginPass string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(dbHash), []byte(loginPass))
	return err == nil
}
