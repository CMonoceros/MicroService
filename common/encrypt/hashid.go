package encrypt

import (
	"SnowBrick-Backend/common/errcode"
	"SnowBrick-Backend/common/log"
	"fmt"
	"github.com/pkg/errors"
	"github.com/speps/go-hashids"
)

type HashIDConfig struct {
	Salt string
}

func EncodeID(salt string, id int) (data string, err error) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 8
	h, err := hashids.NewWithData(hd)
	if err != nil {
		log.Error("EncodeID NewWithData error(%v)", err)
		err = errors.WithMessage(errcode.EncryptEncodeError, fmt.Sprintf("加密错误"))
		return "", err
	}
	data, err = h.Encode([]int{id})
	if err != nil {
		log.Error("EncodeID Encode error(%v)", err)
		err = errors.WithMessage(errcode.EncryptEncodeError, fmt.Sprintf("加密错误"))
		return "", err
	}
	return
}

func DecodeID(salt string, data string) (id int, err error) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 8
	h, err := hashids.NewWithData(hd)
	if err != nil {
		log.Error("DecodeID NewWithData error(%v)", err)
		err = errors.WithMessage(errcode.EncryptDecodeError, fmt.Sprintf("解密错误"))
		return 0, err
	}
	e, err := h.DecodeWithError(data)
	if err != nil {
		log.Error("DecodeID DecodeWithError error(%v)", err)
		err = errors.WithMessage(errcode.EncryptDecodeError, fmt.Sprintf("解密错误"))
		return 0, err
	}
	return e[0], nil
}
