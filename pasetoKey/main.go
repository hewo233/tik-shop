package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/hertz-contrib/paseto"
)

func main() {
	// 生成 Ed25519 公钥和私钥
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating Ed25519 key:", err)
		return
	}

	// 将公钥和私钥转换为 Base64 格式便于保存和传输
	publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKey)
	privateKeyBase64 := base64.StdEncoding.EncodeToString(privateKey)

	// 打印公钥和私钥
	fmt.Println("Public Key (Base64):", publicKeyBase64)
	fmt.Println("Private Key (Base64):", privateKeyBase64)

	// 生成一个示例 Implicit 值
	implicit := base64.StdEncoding.EncodeToString([]byte("tik-shopIMPLICIThiho"))
	fmt.Println("Implicit (Base64):", implicit)

	fmt.Println("Default Paseto Public:", paseto.DefaultPublicKey)
	fmt.Println("Default Paseto Private:", paseto.DefaultPrivateKey)
	fmt.Println("Default Paseto Implicit:", paseto.DefaultImplicit)
}
