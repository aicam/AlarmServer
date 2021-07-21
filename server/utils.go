package server

import (
	"bytes"
	"crypto/des"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func convertMonthToInt(month string) int {
	var m map[string]int
	m = make(map[string]int)
	m["January"] = 1
	m["February"] = 2
	m["March"] = 3
	m["April"] = 4
	m["May"] = 5
	m["June"] = 6
	m["July"] = 7
	m["August"] = 8
	m["September"] = 9
	m["October"] = 10
	m["November"] = 11
	m["December"] = 12
	return m[month]
}

type NotifReq struct {
	Token   string `json:"token"`
	User    string `json:"user"`
	Message string `json:"message"`
	Title   string `json:"title"`
}

func sendNotificationByPushOver(message string, title string) {
	url := "https://api.pushover.net/1/messages.json"

	var jsonBytes []byte
	jsonBytes, err := json.Marshal(&NotifReq{
		Token:   "atvfudwzqaiapnynb436d3bsji625s",
		User:    "uj19y8eotoue2gemw4aerdpkir9imq",
		Message: message,
		Title:   title,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonBytes))
	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println(err)
		return
	}
	//client := &http.Client{}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

func DesEncrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	// src = PKCS5Padding(src, bs)
	if len(src)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func DesDecrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	// out = PKCS5UnPadding(out)
	return out, nil
}
