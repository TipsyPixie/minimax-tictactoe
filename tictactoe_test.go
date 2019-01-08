package minimax

import (
    "fmt"
    "testing"
)

var testMarks = [][]int{
    {-1, 0, 0},
    {0, 0, 0},
    {0, 0, 1},
}

var testBoardState = NewBoardState(testMarks)

func TestBoardState_OptimalMove(t *testing.T) {
    for testBoardState != nil {
        printMarks(testBoardState)
        fmt.Println()
        testBoardState = testBoardState.OptimalMove()
    }
}

func printMarks(boardState *BoardState) {
    for _, rowMarks := range boardState.marks {
        fmt.Println(rowMarks)
    }
}
