package katatennis_test

import (
	"katatennis"
	"testing"
)

func Test_ResultGame_Input_Player1_Score_0_Player2_Score_0_Should_Be_LOVE_LOVE(t *testing.T) {
	expectResult := "LOVE-LOVE"

	actualResult := katatennis.ResultGame(0, 0)

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_1_Player2_Score_0_Should_Be_15_LOVE(t *testing.T) {
	expectResult := "15-LOVE"

	actualResult := katatennis.ResultGame(1, 0)

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_2_Player2_Score_0_Should_Be_30_LOVE(t *testing.T) {
	expectResult := "30-LOVE"

	actualResult := katatennis.ResultGame(2, 0)

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_2_Player2_Score_1_Should_Be_30_15(t *testing.T) {
	expectResult := "30-15"

	actualResult := katatennis.ResultGame(2, 1)

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_2_Player2_Score_2_Should_Be_30_30(t *testing.T) {
	expectResult := "30-30"

	actualResult := katatennis.ResultGame(2, 2)

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}

func Test_ResultGame_Input_Player1_Score_2_Player2_Score_3_Should_Be_30_40(t *testing.T) {
	expectResult := "30-40"

	actualResult := katatennis.ResultGame(2, 3)

	if expectResult != actualResult {
		t.Errorf("Expect %s but got %s", expectResult, actualResult)
	}
}
