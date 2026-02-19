package greeting

import "fmt"

// Hello は挨拶メッセージを返す。
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", formatName(name))
}

// formatName はパッケージ内部専用のヘルパー。
func formatName(name string) string {
	return "[" + name + "]"
}
