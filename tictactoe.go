package minimax

const (
    MaxRow           = 3
    MaxCol           = MaxRow
    EmptyMark        = 0
    Omark            = 1
    Xmark            = -1
    winnersGain      = 100
    scoreErosionRate = 9
)

type BoardState struct {
    marks      [][]int
    marker     int
    prevState  *BoardState
    nextStates []*BoardState
    score      int
}

func NewBoardState(newMarks [][]int) *BoardState {
    markCount := 0
    for _, rowMarks := range newMarks {
        for _, mark := range rowMarks {
            if mark != EmptyMark {
                markCount++
            }
        }
    }
    marker := -1
    if markCount % 2 == 1 {
        marker = 1
    }

    boardState := &BoardState{
        nextStates:[]*BoardState{},
        marker: marker,
        prevState: nil,
        score: 0,
        marks: newMarks,
    }
    boardState.genNextStates()
    boardState.sumScore()
    return boardState
}

func (boardState *BoardState) OptimalMove() *BoardState {
    nextStates := boardState.nextStates
    var optimalNextState *BoardState = nil
    for _, nextState := range nextStates {
        if optimalNextState == nil || boardState.marker * nextState.score > boardState.marker * optimalNextState.score {
            optimalNextState = nextState
        }
    }
    return optimalNextState
}

func (boardState *BoardState) sumScore() int {
    if winnerFlag := boardState.isFinished(); winnerFlag != 0 {
        boardState.score = winnerFlag * winnersGain
    } else if nextStatesLength := len(boardState.nextStates); nextStatesLength != 0 {
        scoreSum := 0
        for _, nextState := range boardState.nextStates {
            scoreSum += nextState.sumScore() * scoreErosionRate / 10
        }
        boardState.score = scoreSum / nextStatesLength
    } else {
        boardState.score = 0
    }

    return boardState.score
}

func (boardState *BoardState) genNextStates() []*BoardState {
    if boardState.isFinished() != 0 {
        return []*BoardState{}
    }

    marks := boardState.marks

    for row := 0; row < MaxRow; row++ {
        for col := 0; col < MaxCol; col++ {
            if marks[row][col] == EmptyMark {
                nextState := boardState.markAndGenNextState(row, col)
                nextState.genNextStates()
                boardState.nextStates = append(boardState.nextStates, nextState)
            }
        }
    }
    return boardState.nextStates
}

func (boardState *BoardState) markAndGenNextState(row int, col int) *BoardState {
    marks := boardState.marks
    newMarks := make([][]int, 0, MaxCol)
    for _, rowMarks := range marks {
        newMarksRow := append(make([]int, 0, len(rowMarks)), rowMarks...)
        newMarks = append(newMarks, newMarksRow)
    }
    newMarks[row][col] = boardState.marker

    return &BoardState{
        marker: - boardState.marker,
        marks: newMarks,
        prevState: boardState,
        nextStates: []*BoardState{},
        score: 0,
    }
}

func (boardState *BoardState) isFinished() int {
    marks := boardState.marks

    //horizontal
    for row := 0; row < MaxRow; row++ {
        marksSum := 0
        for col := 0; col < MaxCol; col++ {
            marksSum += marks[row][col]
        }
        if marksSum >= MaxCol || marksSum <= -MaxCol {
            return marksSum / MaxCol
        }
    }

    //vertical
    for col := 0; col < MaxRow; col++ {
        marksSum := 0
        for row := 0; row < MaxRow; row++ {
            marksSum += marks[row][col]
        }
        if marksSum >= MaxRow || marksSum <= -MaxRow {
            return marksSum / MaxRow
        }
    }

    //diagonal
    marksSum := 0
    for diagonalIndex := 0; diagonalIndex < MaxRow; diagonalIndex++ {
        marksSum += marks[diagonalIndex][diagonalIndex]
    }
    if marksSum >= MaxRow || marksSum <= -MaxRow {
        return marksSum / MaxRow
    }

    //anti-diagonal
    marksSum = 0
    for diagonalIndex := 0; diagonalIndex < MaxRow; diagonalIndex++ {
        marksSum += marks[diagonalIndex][MaxCol-1-diagonalIndex]
    }
    if marksSum >= MaxRow || marksSum <= -MaxRow {
        return marksSum / MaxRow
    }

    return 0
}
