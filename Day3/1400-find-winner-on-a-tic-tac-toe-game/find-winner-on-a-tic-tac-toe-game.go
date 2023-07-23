func tictactoe(moves [][]int) string {
    board := make([][]int, 3)
    for i:=0;i<3;i++{
        board[i] = make([]int, 3)
    }

    A, B := make([]int, 8), make([]int, 8)

    for i:=0;i<len(moves);i++{
        x,y := moves[i][0], moves[i][1]
        if i % 2 == 0 {
            A[x]++
            A[y+3]++
            if x == y {A[6]++}
            if x + y == 2 {A[7]++}
        } else {
            B[x]++
            B[y+3]++
            if x == y {B[6]++}
            if x + y == 2 {B[7]++}
        }
    }

    for i:=0;i<8;i++{
        if A[i] == 3 {return "A"}
        if B[i] == 3 {return "B"}
    }

    if len(moves) == 9 {
        return "Draw"
    }
    return "Pending"
}

/*
  0 1 2 
0
1
2


*/