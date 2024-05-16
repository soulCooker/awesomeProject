package backtracking

import (
	"fmt"
)

func findLargestTeamNum(scores []int, minTeamScore int) int {
	remainScore := 0
	teamScoreGroup := make([]int, 0)

	group := make([][]int, 0)

	for _, v := range scores {
		remainScore += v
	}

	var backtracking func(curSum int, start int, curGroup []int)
	maxTeamCount := 0

	backtracking = func(curSum, start int, curGroup []int) {
		defer func() {
			if curSum == 0 && len(teamScoreGroup) > 0 {
				group = group[0 : len(group)-1]
				teamScoreGroup = teamScoreGroup[0 : len(teamScoreGroup)-1]
			}
		}()

		if curSum >= minTeamScore {
			teamScoreGroup = append(teamScoreGroup, curSum)
			group = append(group, curGroup)
			fmt.Println("teamScoreGroup:", teamScoreGroup)
			fmt.Println("group:", group)
			curSum = 0
			curGroup = make([]int, 0)
			maxTeamCount = max(maxTeamCount, len(teamScoreGroup))
		}

		if remainScore < minTeamScore {
			if remainScore+curSum > minTeamScore {
				maxTeamCount = max(maxTeamCount, len(teamScoreGroup)+1)
			}
			return
		}

		for i := start; i < len(scores); i++ {
			remainScore -= scores[i]
			curSum += scores[i]
			curGroup = append(curGroup, scores[i])
			scores[start], scores[i] = scores[i], scores[start]
			backtracking(curSum, i+1, curGroup)
			scores[start], scores[i] = scores[i], scores[start]
			remainScore += scores[i]
			curGroup = curGroup[0 : len(curGroup)-1]
			curSum -= scores[i]
		}
	}

	backtracking(0, 0, []int{})

	return maxTeamCount
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	} else {
		return val2
	}
}
