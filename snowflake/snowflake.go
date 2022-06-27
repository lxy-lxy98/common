package snowflake

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Snowflake struct {
	sync.Mutex       //锁
	timestamp  int64 //时间戳
	workerid   int64 //机器id
	sequence   int64 //序列号
}

const (
	epoch         = int64(1577808000000)              // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	timestampBits = uint(41)                          // 时间戳占用位数
	workeridBits  = uint(10)                          // 机器id所占位数
	sequenceBits  = uint(12)                          // 序列所占的位数
	timestampMax  = int64(-1 ^ (-1 << timestampBits)) // 时间戳最大值
	//-1 二进制：11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111111
	//41bits 的 timestamp 最大值可以将 -1 向左位移 41 位，得到：
	//11111111 11111111 11111110 00000000 00000000 00000000 00000000 00000000
	//再和 -1 进行 ^异或运算：
	//00000000 00000000 00000001 11111111 11111111 11111111 11111111 11111111
	workeridMax    = int64(-1 ^ (-1 << workeridBits)) // 支持的最大机器id数量,同理
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits)) // 支持的最大序列id数量
	workeridShift  = sequenceBits                     // 机器id左移位数
	timestampShift = sequenceBits + workeridBits      // 时间戳左移位数
)

// 将十进制数字转化为二进制字符串
func convertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s
}
func (s *Snowflake) NextVal() int64 {
	s.Lock()
	now := time.Now().UnixMilli() //转毫秒
	fmt.Println(s.timestamp, now)
	if s.timestamp == now {
		// 当同一时间戳（精度：毫秒）下多次生成id会增加序列号
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			// 如果当前序列超出12bit长度，则需要等待下一毫秒
			// 下一毫秒将使用sequence:0
			for now <= s.timestamp {
				now = time.Now().UnixNano()
			}
		}
	} else {
		// 不同时间戳（精度：毫秒）下直接使用序列号：0
		s.sequence = 0
	}
	//t表示现在距离epoch的时间差
	t := now - epoch
	//时间戳超过69年限制，退出
	if t > timestampMax {
		s.Unlock()
		fmt.Sprintf("epoch must be between 0 and %d", timestampMax-1)
		return 0
	}
	s.timestamp = now
	//组装id
	r := int64((t)<<timestampShift | (s.workerid << workeridShift) | (s.sequence))

	s.Unlock()
	return r
}
