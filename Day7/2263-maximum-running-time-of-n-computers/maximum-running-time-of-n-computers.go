func maxRunTime(n int, batteries []int) int64 {
    sort.Ints(batteries)
    m := len(batteries)

    plugged := make([]int64, 0)

    for i:=m-n;i<m;i++{
        plugged = append(plugged, int64(batteries[i]))
    }

    extra := 0
    for i:=0;i<m-n;i++{
        extra += batteries[i]
    }

    for i:=0;i<n-1;i++{
        req := int((plugged[i+1] - plugged[i]) * int64(i+1))
        if req <= extra {
            extra -= req
        } else {
            return plugged[i] + int64(extra/(i+1))
        }
    }

    return plugged[n-1] + int64(extra/n)
}

/*
m

take largest n from m batteries

m-n -> combined -> a single extra supply -> T (total : i=0 -> m-n T+=batteries[i])

distribute T to n batteries with minimum 1 unit battery

*/