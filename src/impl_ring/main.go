package main

import (
	"fmt"
	"strconv"
)

const RingMaxCapacity = 50

type Ring struct {
	rIndex      int
	curSize     int
	maxCapacity int
	ring        []string
}

func NewRing() *Ring {
	return &Ring{
		rIndex:      0,
		curSize:     0,
		maxCapacity: RingMaxCapacity,
		ring:        make([]string, RingMaxCapacity),
	}
}

func (r *Ring) isFull() bool {
	return r.curSize == r.maxCapacity
}

func (r *Ring) isEmpty() bool {
	return r.curSize == 0
}

// 先插入，后增加Index,
func (r *Ring) AddCoverMsg(msg string) {
	r.ring[r.rIndex] = msg
	r.rIndex = (r.rIndex + 1) % r.maxCapacity
	if !r.isFull() {
		r.curSize++
	}
}

// 从uIndex读取的消息，从r.index里返回消息
func (r *Ring) GetSeqMsg(uIndex int) []string {
	if uIndex <= r.rIndex {
		return r.ring[uIndex:r.rIndex]
	} else {
		if r.isFull() {
			return append(r.ring[uIndex:r.maxCapacity], r.ring[0:r.rIndex]...)
		}
		panic("不应该出现环形数组没满，但user index 大于 ringIndex 的情况")
	}
}

func main() {
	r := NewRing()
	for i := 0; i < 70; i++ {
		r.AddCoverMsg(strconv.Itoa(i))
	}
	fmt.Println(r.GetSeqMsg(0))
}
