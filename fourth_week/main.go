package main

import (
	"sync"
	"time"
)

type Number struct {
	Buckets map[int64]*numberBucket // 以 unixtimestamp 为 key
	Mutex   *sync.RWMutex
}

type numberBucket struct {
	Value float64
}


// 每一次对熔断器的状态进行修改（更新）时，Number 都要先得到当前的 timestamp（秒级），如果 Bucket 不存在则创建。Rolling 包像个时间序列数据库，Buckets 的 Key 是 Unix 时间戳，Number 只保存 10s 内的数据。
func (r *Number) getCurrentBucket() *numberBucket {
	// 先得到当前的 timestamp
	now := time.Now().Unix()
	var bucket *numberBucket
	var ok bool
	// 判断是否存在，不存在则创建
	if bucket, ok = r.Buckets[now]; !ok {
		bucket = &numberBucket{}
		r.Buckets[now] = bucket
	}

	return bucket
}

// 修改完成后去掉 10s 以外的数据
func (r *Number) removeOldBuckets() {
	now := time.Now().Unix() - 10

	for timestamp := range r.Buckets {
		// TODO: configurable rolling window
		if timestamp <= now {
			delete(r.Buckets, timestamp)
		}
	}
}

// Increment Push
func (r *Number) Increment(i float64) {
	if i == 0 {
		return
	}

	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	// 先得到 bucket（当前 timestamp）
	b := r.getCurrentBucket() //代码在上面，先以timestamp检查，不存在则新建
	b.Value += i
	// 删除掉旧的
	r.removeOldBuckets()
}

func (r *Number) Sum(now time.Time) float64 {
	sum := float64(0)

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for timestamp, bucket := range r.Buckets {
		// TODO: configurable rolling window --- 目前滑动窗口的范围还是不可配置的
		if timestamp >= now.Unix()-10 {
			// 确定 timestamp 在区间内
			sum += bucket.Value
		}
	}

	return sum
}

func main() {

}
