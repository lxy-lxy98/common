package message

import (
	"context"
	"fmt"
	"sort"
	"strings"
)

type messageHeadKey struct{}

// MessageHead
// 统一的消息前缀打印，该结构并不保证线程安全,
type MessageHead struct {
	RequestID string
	values    map[string]interface{}
	Key       interface{} //按照需要定义
}

func NewMessageHeadContext(ctx context.Context, mh *MessageHead) context.Context {
	return context.WithValue(ctx, messageHeadKey{}, mh)
}

// MessageHeadContext
// 保证会返回一个消息头，代码中使用这个更方便
func MessageHeadContext(ctx context.Context) *MessageHead {
	raw, ok := FromMessageHeadContext(ctx)
	if ok {
		return raw
	}
	return &MessageHead{}
}

// FromMessageHeadContext
// 不保证一定会有值返回，所以需要判断值是否存在
func FromMessageHeadContext(ctx context.Context) (*MessageHead, bool) {
	raw, ok := ctx.Value(messageHeadKey{}).(*MessageHead)
	return raw, ok
}

func MessageHeadExist(ctx context.Context) bool {
	_, ok := FromMessageHeadContext(ctx)
	return ok
}

// Set
// 如果key存在，会返回错误
func (mh *MessageHead) Set(key string, value interface{}) error {
	if mh.values == nil {
		mh.values = map[string]interface{}{}
	}
	_, exist := mh.values[key]
	if exist {
		return fmt.Errorf("messageHead key %s duplicate", key)
	}
	mh.values[key] = value
	return nil
}

// Append
// 重复append相同的key会直接覆盖
func (mh *MessageHead) Append(key string, value interface{}) *MessageHead {
	if mh.values == nil {
		mh.values = map[string]interface{}{}
	}
	mh.values[key] = value
	return mh
}

func (mh MessageHead) String() string {
	messageHead := fmt.Sprintf("requestID = %s", mh.RequestID)
	if len(mh.values) > 0 {
		var sortKey []string
		var values []interface{}

		for key := range mh.values {
			sortKey = append(sortKey, key)
		}
		sort.Strings(sortKey)
		for _, key := range sortKey {
			values = append(values, mh.values[key])
		}
		messageHead += fmt.Sprintf(", "+strings.Join(sortKey, " = %v, ")+" = %v", values...)
	}
	return messageHead
}

func NewMessageHead(requestID string) *MessageHead {
	return &MessageHead{
		RequestID: requestID,
	}
}
