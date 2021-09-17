package weather

import (
	"fmt"
)

func ExampleToS() {
	w := new(Weather)
	w.TargetArea = "テスト地方"
	w.HeadlineText = "一日良い天気です。"
	w.Text = "傘を持ち歩く必要はないでしょう。"

	fmt.Println(w.ToS())
	// Output: テスト地方の天気です。
	// 一日良い天気です。
	// 傘を持ち歩く必要はないでしょう。
}
