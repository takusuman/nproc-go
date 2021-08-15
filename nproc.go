package main
import (
	"fmt";
	"runtime"
)

func main(){
	var n int
	n = runtime.NumCPU()
	fmt.Println(n)
}
