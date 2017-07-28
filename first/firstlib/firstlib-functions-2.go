package firstlib

var cnt int = 0

func increment() {
	cnt++
}

func Init() {
	cnt = 0;
}

func Next() (int) {
	defer increment()
	return cnt
}
 
