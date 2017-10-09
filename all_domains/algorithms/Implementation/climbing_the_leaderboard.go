package implementation

import (
	"fmt"
	"os"
)

//ClimbingTheLeaderboard task solution
func ClimbingTheLeaderboard() {
	//Leaderboard scores
	var total int

	fmt.Fscanf(os.Stdin, "%d", &total)
	ranks := make([]int, total)
	scores := make([]int, total)

	rank := 1

	for i := 0; i < total; i++ {
		fmt.Fscanf(os.Stdin, "%d", &scores[i])
		if i > 0 && scores[i] < scores[i-1] {
			rank++
		}
		ranks[i] = rank
	}

	//todo reduce size of scores, split by chunks of 250k
	//Total Alice games played
	var totalGames int
	fmt.Fscanf(os.Stdin, "%d", &totalGames)
	var currentScore int
	currentRank := ranks[total-1 ] + 1
	for i := 0; i < totalGames; i++ {
		fmt.Fscanf(os.Stdin, "%d", &currentScore)
		for j := total - 1; j >= 0; j-- {
			if scores[j] > currentScore {
				break
			}
			if scores[j] == currentScore {
				currentRank = ranks[j]
				break
			}
			if scores[j] < currentScore {
				currentRank = ranks[j]
				continue
			}
			currentRank = 1
		}
		fmt.Println(currentRank)
	}
}
