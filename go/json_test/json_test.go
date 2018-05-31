package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/francoispqt/gojay"

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

// // 測試 gojay 的 json encode
// func BenchmarkJSONEncodeJay(b *testing.B) {
// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		JSONEncodeJay()
// 	}
// }

// // 測試 gojay 的 json decode
// func BenchmarkJsonDecodeJay(b *testing.B) {
// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		JSONDecodeJay()
// 	}
// }

// func TestJsonDecodeJay(t *testing.T) {
// 	err := JSONDecodeJay()
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// 以下為測試使用
type user struct {
	Name string `json:"user"`
	Text string `json:"text"`
	Note string `json:"note"`
}

func JSONEncodeStd() []byte {
	u := user{"Emma", "https://www.google.com.tw/", "abc123"}
	b, _ := json.Marshal(u)

	return b
}

func JSONEncodeIter() []byte {
	u := user{"Emma", "https://www.google.com.tw/", "abc123"}
	b, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(u)

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

func JSONEncodeJay() []byte {
	u := user{"Emma", "https://www.google.com.tw/", "abc123"}
	b, _ := gojay.Marshal(u)

	return b
}

func JSONDecodeJay() error {
	u := &user{}
	j := `{"user":"Emma","text":"https://www.google.com.tw/","note":"abc123"}`

	dec := gojay.NewDecoder(strings.NewReader(j))
	return dec.Decode(u)
}
