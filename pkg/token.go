package pkg

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// CheckTokenValidityTime 检查token 有效期
// token数据格式是token_value:timestamp 秒时间
// 如果正确 就刷新token有效时间
func CheckTokenValidityTime(tokenT *string) bool {
	if tokenT == nil {
		return false
	}
	// 需要转换为int64类型
	timestamp, err := strconv.ParseInt(*tokenT, 10, 64)
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
		return false
	}
	// 检查有效期
	tokenTime := time.Unix(timestamp, 0)
	expirationTime := tokenTime.Add(7 * 24 * time.Hour)
	now := time.Now()

	if now.Before(expirationTime) {
		*tokenT = strconv.FormatInt(now.Unix(), 10)
		return true
	} else {
		return false
	}
}

// CheckSpecialCharacters 检查token格式是否正确
func CheckSpecialCharacters(token *string) bool {
	if token == nil {
		return false
	}
	matched, err := regexp.MatchString("[^a-zA-Z0-9]+", *token)
	if err != nil {
		return false
	}
	if matched {
		return false
	} else {
		//不包含特殊字符
		return true
	}
}
