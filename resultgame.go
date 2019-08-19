package katatennis

const RESULTWIN = 4

type PlayerScore struct {
	Player1score int
	Player2score int
}

func (playerscore *PlayerScore) AddScore(player int) {
	if player == 1 {
		playerscore.Player1score++
	} else {
		playerscore.Player2score++
	}
}

func (score PlayerScore) ResultGame() string {
	mapResult := map[int]string{
		0: "LOVE",
		1: "15",
		2: "30",
		3: "40",
	}

	if score.Player1score == RESULTWIN || score.Player2score == RESULTWIN {
		return checkWin(score.Player1score)
	}
	return mapResult[score.Player1score] + "-" + mapResult[score.Player2score]
}

func checkWin(Player1score int) string {
	if Player1score == RESULTWIN {
		return "Player 1 Win"
	}
	return "Player 2 Win"
}
