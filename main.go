package main

import "fmt"

func main() {
	// problema #1 :kak hranit mnogo dannih
	score1 :=10
	score2 := 10
	score3 :=10

	//nps = 100
	promotersLowerBound :=0
	defactorsUpperBound :=0
	//neutrals?- practice vs theory
	//if'i - usloviya
	// boolean - tip dannih
	//problema dannih
	//problem2: dublirovanie koda
	//problem 3: magic values(9,6) - durnoy ton
	// 3 magic values
// refactoring  ulichshenie strukturi koda bez modifikacii povedenya
//ctrl +alt+ v s videleniem pozvolyaet sozdat lokalnuyu
//shift + f6 - pereimenovanie vseh imyon

	if score1>=9 {
		promotersLowerBound= promotersLowerBound +1
	}

	if score1<=6 {
		defactorsUpperBound = defactorsUpperBound +1

	}

	if score2>=9 {
		promotersLowerBound= promotersLowerBound +1
	}

	if score2<=6 {
		defactorsUpperBound = defactorsUpperBound +1

	}

	if score3>=9 {
		promotersLowerBound= promotersLowerBound +1
	}

	if score3<=6 {
		defactorsUpperBound = defactorsUpperBound +1

	}


	nps:=(promotersLowerBound - defactorsUpperBound)/3*100
	//2/3*100->0*100-> 0             2*100/3->200/3->66
	fmt.Println(nps)
	//ctrl+alt+shift+leviy klik mishi(mnogo kursorov)
}
