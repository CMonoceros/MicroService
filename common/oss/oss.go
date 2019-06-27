package oss

import (
	"SnowBrick-Backend/common/ctime"
	"SnowBrick-Backend/common/errcode"
	"SnowBrick-Backend/common/log"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
	"regexp"
)

const (
	BUCKET_BRICK_SET = "snowbrick"
)

type Config struct {
	Key            string
	Secret         string
	Endpoint       map[string]string
	SignExpireTime ctime.Duration
}

func New(c *Config, bucketName string) (bucket *oss.Bucket, err error) {
	if endpoint, ok := c.Endpoint[bucketName]; !ok {
		log.Error("OSS endpoint don't have %s", bucketName)
		err = errors.WithMessage(errcode.OssConfigError, fmt.Sprintf("端点不存在(%s)", bucketName))
		return nil, err
	} else {
		client, err := oss.New(endpoint, c.Key, c.Secret)
		if err != nil {
			log.Error("OSS New client error(%v)", err)
			return nil, err
		}
		bucket, err = client.Bucket(bucketName)
		if err != nil {
			log.Error("OSS New bucket error(%v)", err)
			return nil, err
		}
		return bucket, nil
	}
}

func ParseObjectName(bucketName, src string) (name string, err error) {
	res := regexp.MustCompile(bucketName + `.+?/([^^]*)$`).FindStringSubmatch(src)
	if len(res) < 2 {
		err = errors.WithMessage(
			errcode.OssResourceError,
			fmt.Sprintf("正则匹配错误(%s,%s)", bucketName, src),
		)
		return
	}
	return res[1], nil
}
