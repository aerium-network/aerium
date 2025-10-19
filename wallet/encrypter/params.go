package encrypter

import (
	"encoding/base64"
	"math"
	"strconv"
)

type params map[string]string

func newParams() params {
	return make(map[string]string)
}

func (p params) SetUint8(key string, val uint8) {
	p.SetUint64(key, uint64(val))
}

func (p params) SetUint32(key string, val uint32) {
	p.SetUint64(key, uint64(val))
}

func (p params) SetUint64(key string, val uint64) {
	p[key] = strconv.FormatUint(val, 10)
}

func (p params) SetBytes(key string, val []byte) {
	p[key] = base64.StdEncoding.EncodeToString(val)
}

func (p params) SetString(key, val string) {
	p[key] = val
}

func (p params) GetUint8(key string) uint8 {
	v := p.GetUint64(key)
	if v > math.MaxUint8 {
		return 0
	}
	return uint8(v)
}

func (p params) GetUint32(key string) uint32 {
	v := p.GetUint64(key)
	if v > math.MaxUint32 {
		return 0
	}
	return uint32(v)
}

func (p params) GetUint64(key string) uint64 {
	val, err := strconv.ParseUint(p[key], 10, 64)
	if err != nil {
		return 0
	}

	return val
}

func (p params) GetBytes(key string) []byte {
	val, err := base64.StdEncoding.DecodeString(p[key])
	if err != nil {
		return []byte{}
	}

	return val
}

func (p params) GetString(key string) string {
	return p[key]
}
