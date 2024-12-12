package main

import (
	"fmt"
	"strings"
	"unsafe"

	"golang.org/x/exp/rand"
)

const minLength = 100

func someFunc() {

	// если переменная используется только внутри одной функции, то лучше объявлять ее внутри нее,
	// чтобы не возникало проблем с многопоточностью
	var justString string
	v := createHugeString(1 << 10)

	// для предотвращения утечки памяти необходимо создать копию большой строки
	// варианты создания копии:
	// 1. Создание среза строки через []byte()
	// 2. Использование метода string.Clone(v)
	// 3. Использование copy([]byte, string)
	// 4. Использование temp := bytes.Buffer -> temp.WriteString(v[]) -> temp.String()

	// проверка если
	if len(v) > minLength {
		justString = strings.Clone(v[:minLength])
	} else {
		justString = v
	}

	fmt.Println(unsafe.StringData(justString) == unsafe.StringData(v))

}
func main() {
	someFunc()
}

func createHugeString(n int) string {
	tempStr := "qwertyuiopasdfhklzxcvbnm"
	runes := make([]rune, n)
	for i := 0; i < len(runes); i++ {
		runes[i] = rune(tempStr[rand.Intn(len(tempStr))])
	}
	return string(runes)
}
