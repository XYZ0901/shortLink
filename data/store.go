package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"shortLink/initialize"
	"shortLink/utils"
	"time"
)

var (
	redisCli     = initialize.RedisCli
	urlIdKey     = "next.url.id"
	shortlinkUrl = "shortlink:%s:url"
	UrlShortLink = "url:%s:shortlink"
)

func Shorten(s string, ttl int) (string, error) {
	// 检查url是否已转化过且未过期
	v, err := redisCli.Get(context.Background(), fmt.Sprintf(UrlShortLink, s)).Result()
	if err != nil {
		// 发生错误
		if err != redis.Nil {
			return "", err
		}
		// 如果不存在
		v, err := redisCli.Incr(context.Background(), urlIdKey).Result()
		if err != nil {
			return "", err
		}
		shortlink := utils.Base62Encode(int(v))
		// 防止分布式多机并行重复更新短网址
		isOk, err := redisCli.SetNX(context.Background(), fmt.Sprintf(UrlShortLink, s),
			shortlink, time.Duration(ttl)*time.Second).Result()
		if err != nil {
			return "", err
		}
		// 如果被别机先更新
		if !isOk {
			v, err := redisCli.Get(context.Background(), fmt.Sprintf(UrlShortLink, s)).Result()
			if err != nil {
				return "", err
			}
			return v, nil
		}
		// TODO: 可以利用hash存储 这里可改用hash存储
		_, err = redisCli.Set(context.Background(), fmt.Sprintf(shortlinkUrl, shortlink),
			s, time.Duration(ttl)*time.Second).Result()
		if err != nil {
			redisCli.Del(context.Background(), fmt.Sprintf(UrlShortLink, s))
			return "", err
		}
		return shortlink, err
	}
	// 短地址存在
	// TODO: 是否需要考虑重设短地址时间
	_, err = redisCli.Expire(context.Background(), fmt.Sprintf(UrlShortLink, s), time.Duration(ttl)*time.Second).Result()
	if err != nil {
		return "", err
	}
	_, err = redisCli.Expire(context.Background(), fmt.Sprintf(shortlinkUrl, v), time.Duration(ttl)*time.Second).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

func Redirect(s string) (string, error) {
	url, err := redisCli.Get(context.Background(), fmt.Sprintf(shortlinkUrl, s)).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return url, nil
}
