package main

import (
	"fmt"
	"math/rand"
	"time"
)

const dice = 7

type info struct {
	numdice int
	usedice bool
}

func main() {

	var n, m, k, pq string

	fmt.Println("")
	fmt.Println("============  GAME -  SHUT -  THE -  BOX =========== ")
	fmt.Println("")
	fmt.Print("ENTER YOUR NAME : ")
	fmt.Scanln(&n, &m, &k)
	fmt.Println("")
	fmt.Println("WELCOME TO THE GAME", n, m, k, "AND HAVE FUN SOLDIER")
	fmt.Println("")
	fmt.Println("----- BEFORE WE PLAY, WE HAVE SOME RULES IN THIS GAME -----")
	fmt.Println("")
	fmt.Println("-----(' THE SUM OF YOUR INDEX AND DEALER INDEX, WE CALL IT AS PAIR ')-----")
	fmt.Println("")
	fmt.Println("-----(' YOU MUST CHOOSE 1 NUMBER DICE INDEX ')-----")
	fmt.Println("")
	fmt.Println("-----(' DON'T ENTER THE SAME INDEX ')-----")
	fmt.Println("")
	fmt.Println("-----(' TO WIN, YOUR PAIR MUST BIGGER THEN PREVIOUS PAIR ')-----")
	fmt.Println("")
	fmt.Println("-----(' FOR THE FIRST 3 PAIRS YOU DON'T GET ANY POINT	')-----")
	fmt.Println("")
	fmt.Println("-----(' TO FOUR PAIRS YOU GET 10 POINT ')-----")
	fmt.Println("-----(' TO FIVE PAIRS YOU GET 30 POINT ')-----")
	fmt.Println("-----(' TO SIX PAIRS  YOU GET 40 POINT ')-----")
	fmt.Println("-----(' TO ALL PAIRS  YOU GET 150 POINT')-----")
	fmt.Println("")
	fmt.Println("----- PLAY HONESTLY AND GOOD LUCK SOLDIER -----")
	fmt.Println("")
	fmt.Println("Type 'PLAY' to Play The Game OR Type 'QUIT' To Exit The Game")
	fmt.Println("")
	fmt.Print("And Your Choice is ")
	fmt.Scanln(&pq)

	for pq == "PLAY" || pq == "play" || pq == "Play" && pq != "Quit" && pq != "quit" && pq != "QUIT" {
		fmt.Println("")
		var playerole [dice]info
		var dealerole [dice]info
		diceplayer(&playerole)
		dicedealer(&dealerole)
		fmt.Print("")
		choose(&playerole, &dealerole)
		fmt.Print("----- PLAY AGAIN SOLDIER ----- ? ")
		fmt.Scanln(&pq)
	}

	fmt.Println("")
	fmt.Println("###### GOOD == CHOICE == AND == GOOD BYE == SOLDIER ######")
	fmt.Println("")

}

func diceplayer(playerole *[dice]info) {
	fmt.Print("Your Dice : ")
	var a1 rand.Source
	var r1 *rand.Rand
	var t1 int
	for i := 0; i < dice; i++ {
		a1 = rand.NewSource(time.Now().UnixNano()) //to random number dice
		r1 = rand.New(a1)
		t1 = r1.Intn(6) + 1
		playerole[i].numdice = t1
		fmt.Print(t1, " ")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("")
}

func dicedealer(dealerole *[dice]info) {
	fmt.Println("")
	var a1 rand.Source
	var r1 *rand.Rand
	var t1 int
	fmt.Print("Please Wait For The Dealer Dice ! ! ! ! ")
	for i := 0; i < dice; i++ {
		a1 = rand.NewSource(time.Now().UnixNano()) //to random number dice
		r1 = rand.New(a1)
		t1 = r1.Intn(6) + 1
		dealerole[i].numdice = t1
		// fmt.Print(t1, " ")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("")
}

func choose(playerole *[dice]info, dealerole *[dice]info) {
	var d int
	var sumdice int
	var score int

	fmt.Println("")
	fmt.Println("---------------------------------------")
	fmt.Println("")
	poin := 0

	for i := 0; i < dice; i++ {
		fmt.Print("Your Dice   : ")
		for y := 0; y < dice; y++ {
			fmt.Print(playerole[y].numdice, " ")
		}
		fmt.Println("")

		fmt.Print("Dealer Dice : ")
		for x := 0; x <= i; x++ {
			fmt.Print(dealerole[x].numdice, " ")
		}
		fmt.Println("")
		fmt.Println("")
		fmt.Print("Choose Your Dice Index Number You Want To Sum, Please : ")
		fmt.Scanln(&d)
		if playerole[d-1].usedice !true {
			playerole[d-1].usedice = true
		} else {
			fmt.Println("")
			for playerole[d-1].usedice == true {
				fmt.Println("You Already Used The Index Number Kid ! ! !")
				fmt.Println("Come On Play Honestly ! ! !")
				fmt.Println("")
				fmt.Print("Input Again : ")
				fmt.Scanln(&d)
			}
		}
		sumdice = playerole[d-1].numdice + dealerole[i].numdice
		fmt.Println("==================", dealerole[i].numdice, "-", playerole[d-1].numdice, "==================")
		fmt.Println("=============== Result Is", sumdice, "===============")
		fmt.Println("")
		fmt.Println("")
		if poin < sumdice {
			fmt.Println("THE SUM OF YOUR DICE AND DEALER DICE IS BIGGER")
			// fmt.Println("|-------------------------------|")
			// fmt.Println("|			THE SUM OF			|")
			// fmt.Println("|								|")
			// fmt.Println("|			YOUR DICE			|")
			// fmt.Println("|								|")
			// fmt.Println("|			IS BIGGER			|")
			// fmt.Println("|-------------------------------|")
			fmt.Println("")
			poin = sumdice
			score = score + 1
		} else {
			fmt.Println("THE SUM OF YOUR DICE AND DEALER DICE IS SMALLER")
			// fmt.Println("|------------------------------|")
			// fmt.Println("|			THE SUM OF			|")
			// fmt.Println("|			YOUR DICE			|")
			// fmt.Println("|			IS SMALLER			|")
			// fmt.Println("|------------------------------|")
			fmt.Println("")
			poin = sumdice
		}
	}
	fmt.Println("Your Correct Answer Is", score)
	fmt.Println("")
	if score == 4 {
		fmt.Println("===== Your Point is 10 =====")
		fmt.Println(" CONGRATULATION KID ! ! !")
	} else if score == 5 {
		fmt.Println("===== Your Point Is 30 =====")
		fmt.Println(" CONGRATULATION KID ! ! !")
	} else if score == 6 {
		fmt.Println("===== Your Point Is 70 =====")
		fmt.Println(" CONGRATULATION KID ! ! !")
	} else if score == 7 {
		fmt.Println("===== Your Point Is 150=====")
		fmt.Println(" CONGRATULATION KID ! ! !")
	} else {
		fmt.Println("  ~~~~~~~~~~  (- SORRY YOU DONT RECEIVED ANY SCORE -) ~~~~~~~~~~  ")
	}
	fmt.Println("")
	fmt.Println(">>>>>>>>>> A HONOR CAN PLAY WITH YOU, SOLDIER <<<<<<<<<<")
	fmt.Println("")
}

/*

	|==========================|
	| NAME : ANDI ACHMAD ADJIE |
	|	 CLASS : IF-43-INT	   |
	|     SID : 1301194264	   |
	|==========================|


*/
