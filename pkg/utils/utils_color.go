package utils

import (
	"fmt"
)

// 前景 背景 颜色
// -------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色

// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

const (
	TextColorBlack = iota + 30
	TextColorRed
	TextColorGreen
	TextColorYellow
	TextColorBlue
	TextColorMagenta
	TextColorCyan
	TextColorWhite
)

func Black(msg string) string {
	return SetColor(msg, 0, 0, TextColorBlack)
}

func Red(msg string) string {
	return SetColor(msg, 0, 0, TextColorRed)
}

func Green(msg string) string {
	return SetColor(msg, 0, 0, TextColorGreen)
}

func Yellow(msg string) string {
	return SetColor(msg, 0, 0, TextColorYellow)
}

func Blue(msg string) string {
	return SetColor(msg, 0, 0, TextColorBlue)
}

func Magenta(msg string) string {
	return SetColor(msg, 0, 0, TextColorMagenta)
}

func Cyan(msg string) string {
	return SetColor(msg, 0, 0, TextColorCyan)
}

func White(msg string) string {
	return SetColor(msg, 0, 0, TextColorWhite)
}

func SetColor(msg string, conf, bg, text int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}
