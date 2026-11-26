func canAliceWin(n int) bool {
    alice := false
    x := 10
    for n > 0 {
        n -= x
        x--
        alice = !alice
    }

    if n < 0 {
        alice = !alice
    }
    return alice
}