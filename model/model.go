package model

import (
	"context"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"
)

//PeriodicalExec 周期性传入方法 周期为period
func PeriodicalExec(method func(), period time.Duration) {
	go func() {
		for {
			method()
			time.Sleep(period)
		}
	}()
}

// SnakeString
// 蛇形命名法转驼峰命名法
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// CamelString
// 驼峰命名法转蛇形命名法
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if !k && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || !k) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// ContextInMetadataConversionOut grpc单次调用传递metadata需要放到mdOutgoingKey中，
// 但是grpc被调用方会解析数据并放入到mdIncomingKey，为了方便将metadata在调用过程中传递下去，通过该方法，将mdIncomingKey内容拷贝至mdOutgoingKey
func ContextInMetadataConversionOut(ctx context.Context) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	return metadata.NewOutgoingContext(ctx, md)
}
