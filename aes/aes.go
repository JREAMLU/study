package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

func main() {
	data := "Hello World"
	key := "12345678901234561234567890123456"
	iv := []byte("0123456789abcdef")
	s, iv, err := AESEncrypter(data, key, iv)
	if err != nil {
		fmt.Println("++++++++++++: err:", err)
		return
	}

	r, err := AESDecrypter(s, key, iv)
	if err != nil {
		fmt.Println("++++++++++++: err:", err)
		return
	}

	fmt.Println("++++++++++++: ", s, string(r), string(iv))

	/*
				<?php
		    		$encryptMethod = 'AES-256-CBC';
		    		// 明文数据
		    		$data = 'Hello World';

		    		$key = '12345678901234561234567890123456';

		    		// 生成IV
		    		$ivLength = openssl_cipher_iv_length($encryptMethod);
		    		// $iv = openssl_random_pseudo_bytes($ivLength, $isStrong);
		    		$iv = '0123456789abcdef';
		    		if (false === $iv && false === $isStrong) {
		    		    die('IV generate failed');
		    		}

		    		// 加密
		    		$encrypted = openssl_encrypt($data, $encryptMethod, $key, 0, $iv);
		    		// 解密
		    		$decrypted = openssl_decrypt($encrypted, $encryptMethod, $key, 0, $iv);

		    		var_dump($encrypted, $decrypted, $iv);
				?>
	*/
}

// AESEncrypter aes encrypt
func AESEncrypter(src string, encrypteKey string, iv []byte) (string, []byte, error) {
	key := []byte(encrypteKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	content := []byte(src)
	content = PKCS7Padding(content, block.BlockSize())
	dst := make([]byte, len(content))
	mode.CryptBlocks(dst, content)

	ciphertext := base64.StdEncoding.EncodeToString(dst)

	return ciphertext, iv, nil
}

// AESDecrypter aes decrypt
func AESDecrypter(src string, encrypteKey string, iv []byte) ([]byte, error) {
	key := []byte(encrypteKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	content, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	dst := make([]byte, len(content))
	mode.CryptBlocks(dst, content)

	return PKCS7UnPadding(dst, block.BlockSize())
}

// PKCS7Padding pkcs7 padding
func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - (len(src) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(src, padText...)
}

// PKCS7UnPadding pkcs7 unpadding
func PKCS7UnPadding(src []byte, blockSize int) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding >= length || unpadding > blockSize {
		return nil, errors.New("padding size error")
	}

	return src[:(length - unpadding)], nil
}
