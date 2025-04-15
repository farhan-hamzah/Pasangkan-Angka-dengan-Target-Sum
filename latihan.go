package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var n, i, target, jum, j int
	fmt.Scan(&n, &target)

	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}

	for i =0; i < n; i++{
		for j = i+1; j < n; j++{
			if A[i]+A[j] == target{
				jum = A[i]+A[j]
			}
		}
	}
	fmt.Print(jum)
}