package main

import(
	"fmt"
)


func insertionSort(v []int) {
	for insercao:=1; insercao < len(v); insercao++{
		//inserção E {1, 2, 3, 4,..., n-1}
		for i:=insercao; i >= 1 && v[i] < v[i-1]; i--{
			//para inserção = 1, i E {1}
			//para inserção = 2, i E {1, 2}
			//para inserção = 3, i E {1, 2, 3}
			//...
			//para inserção = 1, i E {1, 2, 3, 4,..., n-1}
			v[i], v[i-1] = v[i-1], v[i]
		}
	}
}

func main(){
	v := [] int {5, 9, 1, 2, 8, 10, 0}
	insertionSort(v)
	fmt.Println(v)
}

