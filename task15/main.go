package main

import "strings"

var justString string

func createHugeString(num int) string {
	return strings.Repeat("q", num)
}

// GC не сможет освободить память, выделенную под строку, созданную в justString, т.к. для
// освобождения необходимо отсутсвие указателей не только на начало строки, но и на все ее элементы
// и т.к. justString - глобальная переменная, то после выхода из someFunc() justString будеь содержать
// элементы из v, что будет препятствовать сборке мусора для v
// Clone же копирует строку с новой аллокацией, т.е. выделяет новый участок в памяти для justString
// и тогда GC сможет освободить память, выделенную для v
/*func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}*/

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}
