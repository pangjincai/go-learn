package test

import (
	"fmt"
	"testing"
	"time"
)

var isPrint bool = true

func init()  {
	fmt.Println("hi test")
}
func TestSum(t *testing.T) {
	val := Sum(1, 2)
	if isPrint {
		fmt.Println("sum=", val)
	}
}

// -bench 基准测试
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(i, i+1)
	}
}

// -timeout 测试
func TestSumLongTime(t *testing.T) {
	time.Sleep(time.Second * 2)
	Sum(1, 2)
}

// 子测试
func TestSumSubTest(t *testing.T) {
	t.Run("1+2", func(t *testing.T) {
		val := Sum(1, 2)
		fmt.Println("1+2=", val)
	})
	t.Run("2+3", func(t *testing.T) {
		val := Sum(2, 3)
		fmt.Println("2+3=", val)
	})
}

// 子测试，无具体子测试
func TestSumSubTest2(t *testing.T) {
	val := Sum(2, 3)
	t.Log("no subtest=", val)
}

// 并发测试
func TestSumParallel(t *testing.T) {
	t.Parallel()
	Sum(1, 2)
}

func TestSumParallel2(t *testing.T) {
	t.Parallel()
	Sum(1, 2)
}
