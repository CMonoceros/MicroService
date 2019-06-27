package oss

import (
	"SnowBrick-Backend/common/ctime"
	"SnowBrick-Backend/common/errcode"
	"SnowBrick-Backend/common/log"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
	"net/http"
	"regexp"
	"strings"
	"time"
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

type Bucket struct {
	*oss.Bucket
	Config *Config
}

func New(c *Config, bucketName string) (bucket *Bucket, err error) {
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
		ossBucket, err := client.Bucket(bucketName)
		if err != nil {
			log.Error("OSS New bucket error(%v)", err)
			return nil, err
		}
		return &Bucket{
			Bucket: ossBucket,
			Config: c,
		}, nil
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

func (bucket *Bucket) SignUrl(src string, processParam ...string) (url string, err error) {
	objectName, err := ParseObjectName(bucket.BucketName, src)
	if err != nil {
		log.Error("Bucket.SignUrl ParseObjectName error(%v)", err)
		return "", err
	}

	processString := strings.Join(processParam, "")
	signSrc, err := bucket.SignURL(objectName, http.MethodGet,
		int64(time.Duration(bucket.Config.SignExpireTime).Seconds()),
		oss.Process(processString))
	if err != nil {
		log.Error("Bucket.SignUrl SignURL error(%v)", err)
		return "", err
	}

	return signSrc, nil
}
