func minSpeedOnTime(dist []int, hour float64) int {
    target := ceil(hour)
    if target < len(dist) {return -1}
    l, r := 1, 10000000

    for l < r {
        mid := (l+r)/2
        x := getHourByVelocity(mid, dist)
        fmt.Printf("%d %v\n", mid, x)
        if x > hour {
            l = mid+1
        } else if x < hour {
            r = mid
        } else {
            return mid
        }
    }

    return l
}

func getHourByVelocity(v int, dist []int) float64 {
    var x float64

    for i, d := range dist {
        if i < len(dist)-1 {
            x += float64(d/v)
            if d%v != 0 {x++}
        } else {
            x += float64(d)/float64(v)
        }
    }

    return x
}

func ceil(x float64) int {
    intX := int(x)
    if float64(intX) < x {
        intX++
    }

    return intX
}