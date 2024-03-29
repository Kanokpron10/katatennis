package katatennis_test

import (
	"katatennis"
	"testing"
)

func Test_ResultGame_Input_Player1_Score_0_Player2_Score_0_Should_Be_LOVE_LOVE(t *testing.T) {
	expectResult := "LOVE-LOVE"
	score := katatennis.PlayerScore{
		Player1score: 0,
		Player2score: 0,
	}

	actualResult := score.ResultGame()

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_1_Player2_Score_0_Should_Be_15_LOVE(t *testing.T) {
	expectResult := "15-LOVE"
	score := katatennis.PlayerScore{
		Player1score: 0,
		Player2score: 0,
	}
	score.AddScore(1)
	actualResult := score.ResultGame()

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_2_Player2_Score_0_Should_Be_30_LOVE(t *testing.T) {
	expectResult := "30-LOVE"
	score := katatennis.PlayerScore{
		Player1score: 0,
		Player2score: 0,
	}
	score.AddScore(1)
	score.AddScore(1)
	actualResult := score.ResultGame()

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_2_Player2_Score_1_Should_Be_30_15(t *testing.T) {
	expectResult := "30-15"
	score := katatennis.PlayerScore{
		Player1score: 0,
		Player2score: 0,
	}
	score.AddScore(1)
	score.AddScore(1)
	score.AddScore(2)
	actualResult := score.ResultGame()

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_2_Player2_Score_2_Should_Be_30_30(t *testing.T) {
	expectResult := "30-30"
	score := katatennis.PlayerScore{
		Player1score: 0,
		Player2score: 0,
	}
	score.AddScore(1)
	score.AddScore(1)
	score.AddScore(2)
	score.AddScore(2)
	actualResult := score.ResultGame()

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_2_Player2_Score_3_Should_Be_30_40(t *testing.T) {
	expectResult := "30-40"
	score := katatennis.PlayerScore{
		Player1score: 0,
		Player2score: 0,
	}
	score.AddScore(1)
	score.AddScore(1)
	score.AddScore(2)
	score.AddScore(2)
	score.AddScore(2)
	actualResult := score.ResultGame()

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_2_Player2_Score_4_Should_Be_Player_2_Win(t *testing.T) {
	expectResult := "Player 2 Win"
	score := katatennis.PlayerScore{
		Player1score: 0,
		Player2score: 0,
	}
	score.AddScore(1)
	score.AddScore(1)
	score.AddScore(2)
	score.AddScore(2)
	score.AddScore(2)
	score.AddScore(2)

	actualResult := score.ResultGame()

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_4_Player2_Score_0_Should_Be_Player_1_Win(t *testing.T) {
	expectResult := "Player 1 Win"
	score := katatennis.PlayerScore{
		Player1score: 0,
		Player2score: 0,
	}
	score.AddScore(1)
	score.AddScore(1)
	score.AddScore(1)
	score.AddScore(1)

	actualResult := score.ResultGame()

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_4_Player2_Score_1_Should_Be_Player_1_Win(t *testing.T) {
	expectResult := "Player 1 Win"
	score := katatennis.PlayerScore{
		Player1score: 0,
		Player2score: 0,
	}
	score.AddScore(1)
	score.AddScore(1)
	score.AddScore(1)
	score.AddScore(1)
	score.AddScore(2)

	actualResult := score.ResultGame()

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}
