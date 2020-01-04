package main

import "fmt"

func main() {
	// problema #1 :kak hranit mnogo dannih
	score1 :=10
	score2 := 10
	score3 :=10

	//nps = 100
	promoters:=0
	detractorov:=0
	//neutrals?- practice vs theory
	//if'i - usloviya
	// boolean - tip dannih
	//problema dannih
	//problem2: dublirovanie koda
	//problem 3: magic values(9,6) - durnoy ton
	// 3 magic values


	if score1>=9 {
		promoters:=promoters+1
	}

	if score1<=6 {
		detractorov= detractorov+1

	}

	nps:=(promoters-detractorov)/3*100
	//2/3*100->0*100-> 0             2*100/3->200/3->66
	fmt.Println(nps)
	//ctrl+alt+shift+leviy klik mishi(mnogo kursorov)
}
