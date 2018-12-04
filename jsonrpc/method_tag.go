package jsonrpc

import "strings"

const sep = "."

// MethodTag Тэг метода
type MethodTag struct {
	Service, Method string
}

// MethodTagToString Преобразует тэг в строку
func MethodTagToString(mt MethodTag) string {
	return strings.Join([]string{mt.Service, mt.Method}, sep)
}

// MethodTagFromString Преобразует строку в тэг
func MethodTagFromString(method string) MethodTag {
	splittedName := strings.Split(method, sep)
	if len(splittedName) == 2 {
		return MethodTag{Service: splittedName[0], Method: splittedName[1]}
	}
	return MethodTag{Method: method}
}
