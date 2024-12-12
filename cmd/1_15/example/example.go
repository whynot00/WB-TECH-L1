package main

var justString string

func someFunc() {
	// основная проблема - утечка памяти

	v := createHugeString(1 << 10)

	// при создании среза строки по индексу,
	// новая строка содержащая [:100] ссылается на первые 100 элементов строки v
	// и большая строка будет существовать до тех пор, пока строка поменьше существует
	justString = v[:100]
}
func main() {
	someFunc()
}

func createHugeString(n int) string {
	return "s"
}
