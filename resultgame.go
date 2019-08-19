package katatennis

func ResultGame(player1score, player2score int) string {
	mapResult := map[int]string{
		0: "LOVE",
		1: "15",
		2: "30",
		3: "40",
	}

	return mapResult[player1score] + "-" + mapResult[player2score]
}
