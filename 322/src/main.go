package main

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func CoinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}
	state := make([]int, amount+1)

	for tmpAmount := 1; tmpAmount <= amount; tmpAmount++ {
		//Init state[tmpAmount] can make it been set to the first state[tmpAmount-v] + 1
		state[tmpAmount] = 0x7FFFFFFF
		for _, v := range coins {
			if tmpAmount-v < 0 {
				continue
			}
			state[tmpAmount] = Min(state[tmpAmount], state[tmpAmount-v]+1)
		}
	}

	if state[amount] == 0x7FFFFFFF {
		return -1
	} else {
		return state[amount]
	}
}
