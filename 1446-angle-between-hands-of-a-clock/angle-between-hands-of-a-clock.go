func angleClock(hour int, minutes int) float64 {
    h, m := getHourPos(hour, minutes), getMinutePos(minutes)
    fmt.Println(h, m)

    return 360 * getDiff(h,m)
}

func getDiff(h, m float64) float64{
    diff := (h-m)
    fmt.Println(diff)
    if diff < 0 {
        diff *= float64(-1)
    }

    return minimum(diff)
}

func minimum(x float64) float64{
    if float64(12) - x < x {
        return (float64(12) - x)/float64(12)
    }
    return x/float64(12)
}

func getHourPos(hour int, minutes int) float64 {
    return float64(hour%12) + float64(float64(minutes)/60)
}

func getMinutePos(minutes int) float64 {
    return float64(minutes) / 5
}


/*
pos hour
y + x/60


pos minute
x/60 * 12 =>



*/


