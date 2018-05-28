package main

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

// 測試原生 json encode
func BenchmarkJsonEncodeStd(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		JSONEncodeStd()
	}
}

// 測試 json-iterator 的 json encode
func BenchmarkJsonEncodeIter(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		JSONEncodeIter()
	}
}

// 測試原生 json decode
func BenchmarkJsonDecodeStd(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		JSONEncodeStd()
	}
}

// 測試 json-iterator 的 json decode
func BenchmarkJsonDecodeIter(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		JSONDecodeIter()
	}
}

// 以下為測試使用
type user struct {
	Name string `json:user`
	Text string `json:text`
	Note string `json:note`
}

func JSONEncodeStd() []byte {
	u := user{"Emma", "https://www.google.com.tw/", "abc123"}
	b, _ := json.MarshalIndent(u, "", "  ")

	return b
}

func JSONEncodeIter() []byte {
	u := user{"Emma", "https://www.google.com.tw/", "abc123"}
	b, _ := jsoniter.ConfigCompatibleWithStandardLibrary.MarshalIndent(u, "", "  ")

	return b
}

func JSONDecodeStd() error {
	u := user{}
	j := `{"user":"Emma","text":"https://www.google.com.tw/","note":"abc123"}`
	return json.Unmarshal([]byte(j), &u)
}

func JSONDecodeIter() error {
	u := user{}
	j := `{"user":"Emma","text":"https://www.google.com.tw/","note":"abc123"}`
	return jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(j), &u)
}
