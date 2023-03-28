package main
import "time"
type Item struct {
	A string
	B string
	C string
	D string
	E string
	F string
	G G
}
type G struct {
	H int
	I int
	K int
	L int
	M int
	N int
}
func main() {
	m := make(map[int]*Item, 10*1024*1024)
	for i := 0; i < 1024*1024; i++ {
		m[i] = &Item{}
	}
	for i := 0; ; i++ {
		delete(m, i)
		m[1024*1024+i] = &Item{}
		time.Sleep(10 * time.Millisecond)
	}
}