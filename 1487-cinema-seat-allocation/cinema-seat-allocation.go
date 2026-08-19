func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
    seatsPerRow := make(map[int][]bool)

    for _, seat := range reservedSeats {
        x, y := seat[0]-1, seat[1]-1
        if _, ok := seatsPerRow[x]; !ok {
            seatsPerRow[x] = make([]bool, 11)
        }
        seatsPerRow[x][y] = true
    }

    answer := n * 2
    for key, seats := range seatsPerRow {
        b := getBlocks(seats)
        if b == 0 {
            answer -=2
        }
        if b == 1 {
            answer--
        }
        fmt.Println(key, answer)
    }

    return answer
}

func getBlocks(s []bool) int {
    one, two, three, four := true, true, true, true

    if s[1] || s[2] {
        one = false
    }
    if s[3] || s[4] {
        two = false
    }
    if s[5] || s[6] {
        three = false
    }
    if s[7] || s[8] {
        four = false
    }

    if one && two && three && four {
        return 2
    }
    if (one && two) || (three && four) || (two && three) {
        return 1
    }

    return 0
}