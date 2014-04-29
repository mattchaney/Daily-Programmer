package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
)

func printCombo(res map[int]int) {
	for k,v := range res {
		fmt.Printf("%d:%d ",k,v)
	}
	fmt.Print("\n")
}

func reduce(coins []int, res map[int]int, currentCoin int, target int, currentTotal *int) {
	if *currentTotal == target  {
		printCombo(res)
		return
	}
	if currentCoin >= len(coins) {
		return	
	}
	if *currentTotal + coins[currentCoin] <= target { 
		newRes := make(map[int]int)
		for k,v := range res {
			newRes[k] = v
		}
		newRes[coins[currentCoin]] += 1
		newTotal := new(int)
		*newTotal = *currentTotal
		*newTotal += coins[currentCoin]
		reduce(coins, newRes, currentCoin, target, newTotal)
	}
	reduce(coins, res, currentCoin + 1, target, currentTotal)
}

func main() {
	var n, target int
	fmt.Scan(&n, &target)
	in := bufio.NewReader(os.Stdin)
	coins := make([][]int, n)
	for i:=0;i<n;i++ {
		str,_ := in.ReadString('\n')
		for _,val := range strings.Fields(str) {
			num,_ := strconv.Atoi(val)
			coins[i] = append(coins[i], num)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(coins[i])))
	}
	currentCoin := 0
	currentTotal := new(int)
	res := make(map[int]int)
	for i := range coins {
		fmt.Println("Currency", i+1, "Combinations")
		reduce(coins[i], res, currentCoin, target, currentTotal)
	}
}