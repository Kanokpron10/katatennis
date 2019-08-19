package katatennis

func ResultGame(player1score, player2score int) string {
	mapResult := map[int]string{
		0: "LOVE",
		1: "15",
		2: "30",
		3: "40",
	}

	if player1score == 4 {
		return "Player 1 Win"
	}

	if player2score == 4 {
		return "Player 2 Win"
	}

	return mapResult[player1score] + "-" + mapResult[player2score]
}
