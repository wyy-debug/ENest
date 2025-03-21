package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
)

const (
	AESKeySize = 32 // AES-256
	NonceSize  = 12 // GCM模式的Nonce大小
)

// CryptoManager 加密管理器
type CryptoManager struct {
	key   []byte
	block cipher.Block
	aead  cipher.AEAD
}

// NewCryptoManager 创建新的加密管理器
func NewCryptoManager() (*CryptoManager, error) {
	// 生成随机密钥
	key := make([]byte, AESKeySize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}

	// 创建AES加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 创建GCM模式的AEAD
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &CryptoManager{
		key:   key,
		block: block,
		aead:  aead,
	}, nil
}

// Encrypt 加密数据
func (cm *CryptoManager) Encrypt(data []byte) ([]byte, error) {
	// 生成随机Nonce
	nonce := make([]byte, NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// 加密数据
	ciphertext := cm.aead.Seal(nil, nonce, data, nil)

	// 组合Nonce和密文
	result := make([]byte, NonceSize+len(ciphertext))
	copy(result, nonce)
	copy(result[NonceSize:], ciphertext)

	return result, nil
}

// Decrypt 解密数据
func (cm *CryptoManager) Decrypt(data []byte) ([]byte, error) {
	if len(data) < NonceSize {
		return nil, errors.New("invalid ciphertext length")
	}

	// 提取Nonce和密文
	nonce := data[:NonceSize]
	ciphertext := data[NonceSize:]

	// 解密数据
	plaintext, err := cm.aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// EncodeInt32LE 将int32编码为小端字节序
func EncodeInt32LE(value int32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(value))
	return buf
}

// DecodeInt32LE 从小端字节序解码int32
func DecodeInt32LE(data []byte) int32 {
	if len(data) < 4 {
		return 0
	}
	return int32(binary.LittleEndian.Uint32(data))
}