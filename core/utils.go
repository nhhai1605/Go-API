package core

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"github.com/mitchellh/mapstructure"
)

func DatabaseMapping(rows *sql.Rows, cols []string) ([]map[string]interface{}, error) {
	var maps = []map[string]interface{}{}
	for rows.Next() {
		vals := make([]interface{}, len(cols))
		for i := range cols {
			var ii interface{}
			vals[i] = &ii
		}

		err := rows.Scan(vals...)
		if err != nil {
			return nil, err
		}

		m := map[string]interface{}{}
		for i, colName := range cols {
			val := vals[i].(*interface{})
			m[colName] = *val
		}
		maps = append(maps, m)
	}
	return maps, nil
}

func StructMapping(source interface{}, target interface{}) error {
	return mapstructure.Decode(source, &target)
}


func Sha256(data string) string {
	bv := []byte(data)
	hasher := sha256.New()
	hasher.Write(bv)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func Encrypt(str string) (string, error) {
    secretKey := os.Getenv("SECRET_KEY")
    key := []byte(secretKey)
    text := []byte(str)
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    b := base64.StdEncoding.EncodeToString(text)
    ciphertext := make([]byte, aes.BlockSize+len(b))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    cfb := cipher.NewCFBEncrypter(block, iv)
    cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
    return hex.EncodeToString(ciphertext), nil
}

func Decrypt(str string) (string, error) {
    secretKey := os.Getenv("SECRET_KEY")
    key := []byte(secretKey)
    text, err := hex.DecodeString(str)
    if err != nil {
        return "", err
    }
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    if len(text) < aes.BlockSize {
        return "", errors.New("ciphertext too short")
    }
    iv := text[:aes.BlockSize]
    text = text[aes.BlockSize:]
    cfb := cipher.NewCFBDecrypter(block, iv)
    cfb.XORKeyStream(text, text)
    data, err := base64.StdEncoding.DecodeString(string(text))
    if err != nil {
        return "", err
    }
    return string(data), nil
}