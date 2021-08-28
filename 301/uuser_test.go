package main

import (
	"testing"
)

// 成功的测试用例
func TestCheckUserOK(t *testing.T) {
	if pass := checkUser("francis", "francis"); false == pass {
		t.Errorf("checkUser(francis, francis) = %v; want true", pass)
	}
}

// 失败的测试用例
func TestCheckUserError(t *testing.T) {
	if pass := checkUser("francis", "a"); true == pass {
		t.Errorf("checkUser(francis, a) = %v; want false", pass)
	}
}

// 基准测试  性能测试
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkUser("francis", "a")
	}
}
