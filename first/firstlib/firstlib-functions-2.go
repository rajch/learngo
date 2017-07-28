package firstlib

var cnt int = 0


func Init() {
	cnt = 0;
}

func Next() (int) {
	result := cnt
	cnt++
	return result
}
 
