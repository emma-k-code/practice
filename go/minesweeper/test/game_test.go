package test

import (
	"testing"

	Game "../game"
)

// 測試 16*30 地圖建立
func BenchmarkCreateMapBig(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Game.CreateMap(16, 30, 99)
	}
}

// 測試 9*9 地圖建立
func BenchmarkCreateMapSmall(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Game.CreateMap(9, 9, 10)
	}
}

// 測試 亂數取地雷位置
func BenchmarkGetMineIndex(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Game.GetMineIndex(99, 16, 30)
	}
}

func TestGetMineIndex(t *testing.T) {
	m := 10
	result := Game.GetMineIndex(m, 9, 9)
	if len(result) != m {
		t.Error("mine count error")
	}
}

func TestCreateMapSmall(t *testing.T) {
	result := Game.CreateMap(9, 9, 10)
	if len(result) != 9 {
		t.Error("map lenght error")
	}
}
