package tmaic

import (
	"fmt"
	"github.com/farmerx/gorsa"
	c "github.com/pangxianfei/framework/config"
	"io/ioutil"
	"log"
)

/*
func init() {
	var Pokey string
	var Private string
	//获取公钥
	pucker, err := ioutil.ReadFile(c.GetString("app.public_key"))
	if err != nil {
		fmt.Print(err)
	}
	Pokey = string(pucker)
	if err := gorsa.RSA.SetPublicKey(Pokey); err != nil {
		log.Fatalln(`set public key :`, err)
	}
	//获取私钥
	PrivateStr, err := ioutil.ReadFile(c.GetString("app.private_key"))
	if err != nil {
		fmt.Print(err)
	}
	Private = string(PrivateStr)
	if err := gorsa.RSA.SetPrivateKey(Private); err != nil {
		log.Fatalln(`set private key :`, err)
	}
}
*/

/********************公钥加密 && 私钥解密***************************/
// 公钥加密
func Encryption(str string) (string, bool) {
	if len(str) <= 0 {
		return string(`Encryption string cannot be empty`), false
	}
	//获取公钥
	var Pokey string
	pucker, err := ioutil.ReadFile(c.GetString("app.public_key"))
	if err != nil {
		panic("get private key failed")
	}
	Pokey = string(pucker)
	if err := gorsa.RSA.SetPublicKey(Pokey); err != nil {
		log.Fatalln(`set public key :`, err)
		panic("set public key failed")

	}
	//end 获取公钥
	//开始加密
	puberty, err := gorsa.RSA.PubKeyENCTYPT([]byte(str))
	if err == nil {
		fmt.Println(string(puberty))
		return string(puberty), true
	}
	return string(`Encryption failed`), false
}

// 私钥解密
func Decrypt(str string) (string, bool) {
	if len(str) <= 0 {
		return string(`Decryption string cannot be empty`), false
	}
	//获取私钥
	var Private string
	PrivateStr, err := ioutil.ReadFile(c.GetString("app.private_key"))
	if err != nil {
		panic("get private key failed")
	}
	Private = string(PrivateStr)
	if err := gorsa.RSA.SetPrivateKey(Private); err != nil {
		panic("set private key failed")
	}
	//开始解密
	decrypt, err := gorsa.RSA.PriKeyDECRYPT([]byte(str))
	if err == nil {
		return string(decrypt), true
	}
	return string(`Decryption failed`), false
}

/***********************************************************/

/******************私钥加密 && 公钥解密**********************/
// 私钥加密
func PrivateEncryption(str string) (string, bool) {
	if len(str) <= 0 {
		return string(`Decryption string cannot be empty`), false
	}
	prienctypt, err := gorsa.RSA.PriKeyENCTYPT([]byte(str))
	if err == nil {
		return string(prienctypt), true
	}
	return string(`Encryption failed`), false
}

// 公钥解密
func PublicDecrypt(str string) (string, bool) {

	if len(str) <= 0 {
		return string(`The Decrypt string cannot be empty`), false
	}

	//获取公钥
	var Pokey string
	pucker, err := ioutil.ReadFile(c.GetString("app.public_key"))
	if err != nil {
		panic("get private key failed")
		return string(`get private key failed`), false
	}
	Pokey = string(pucker)
	if err := gorsa.RSA.SetPublicKey(Pokey); err != nil {
		log.Fatalln(`set public key :`, err)
		panic("set public key failed")
		return string(`set public key failed`), false
	}
	//end 获取公钥

	pubdecrypt, err := gorsa.RSA.PubKeyDECRYPT([]byte(str))
	if err != nil {
		return string(pubdecrypt), true
	}

	return string(`Decryption failed`), false
}

/***********************************************************/
