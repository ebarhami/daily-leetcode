func bulbSwitch(n int) int {
    return sqrt(n)
}

func sqrt(n int) int {
    i := 1
    for i*i < n {
        i++
    }
    if i*i > n {return i-1}
    return i
}


/*
1 = 1
2 = 1
3 = 1
4 = 2 -> 1, 4

a bulb toggle odd times -> it is on
              even times -> it is off

even -> we can reach that using even number of toggle = even number of factor


2 -> 1, 2
3 -> 1, 3

12 -> 1, 2, 3, 4, 6, 12 = 6 = off

odd -> odd number of factor
1 -> 1
4 -> 1, 2, 4 -> perfect square

*/