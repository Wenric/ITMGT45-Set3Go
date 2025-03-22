package main

func relationshipStatus(fromMember string, toMember string, socialGraph map[string]map[string][]string) string {
	fromFollows := contains(socialGraph[fromMember]["following"], toMember)
	toFollows := contains(socialGraph[toMember]["following"], fromMember)

	if fromFollows && toFollows {
		return "friends"
	} else if fromFollows {
		return "follower"
	} else if toFollows {
		return "followed by"
	} else {
		return "no relationship"
	}
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func ticTacToe(board [][]string) string {
	size := len(board)

	for i := 0; i < size; i++ {
		symbol := board[i][0]
		win := symbol != ""
		for j := 1; j < size; j++ {
			if board[i][j] != symbol {
				win = false
				break
			}
		}
		if win {
			return symbol
		}
	}

	for i := 0; i < size; i++ {
		symbol := board[0][i]
		win := symbol != ""
		for j := 1; j < size; j++ {
			if board[j][i] != symbol {
				win = false
				break
			}
		}
		if win {
			return symbol
		}
	}

	symbol := board[0][0]
	win := symbol != ""
	for i := 1; i < size; i++ {
		if board[i][i] != symbol {
			win = false
			break
		}
	}
	if win {
		return symbol
	}

	symbol = board[0][size-1]
	win = symbol != ""
	for i := 1; i < size; i++ {
		if board[i][size-1-i] != symbol {
			win = false
			break
		}
	}
	if win {
		return symbol
	}

	return "NO WINNER"
}

type Leg struct {
	travelTimeMins int
}

func eta(firstStop string, secondStop string, routeMap map[[2]string]Leg) int {
	totalTime := 0
	current := firstStop

	for current != secondStop {
		for key, val := range routeMap {
			if key[0] == current {
				totalTime += val.travelTimeMins
				current = key[1]
				break
			}
		}
	}

	return totalTime
}

func main() {

}
