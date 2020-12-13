package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"gowith/config"
	"hash"
)

type Sha1Stream struct{
	_sha1 hash.Hash
}

//更新哈希
func (obj *Sha1Stream) Update(data []byte) {
	if obj._sha1 == nil {
		obj._sha1 = sha1.New()
	}
	obj._sha1.Write(data)
}

func (obj *Sha1Stream) Sum() string {
	return hex.EncodeToString(obj._sha1.Sum([]byte("")))
}

func Sha1(data string) string {
	dataWithSalt := []byte(data + config.JwtSecret)

	_sha1 := sha1.New()
	_sha1.Write(dataWithSalt)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

func MD5(data string) string {
	dataWithSalt := []byte(data + config.JwtSecret)
	_md5 := md5.New()
	_md5.Write(dataWithSalt)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}