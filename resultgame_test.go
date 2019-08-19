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
