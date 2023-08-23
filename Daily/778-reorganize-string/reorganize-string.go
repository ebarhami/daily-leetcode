type letter struct {
    x rune
    freq int 
}

func reorganizeString(s string) string {
    pocket := make([]letter, 26)

    for _, c := range s {
        if pocket[int(rune(c) - 'a')].freq == 0 {
            pocket[int(rune(c) - 'a')] = letter{x: rune(c)}
        }
        pocket[int(rune(c) - 'a')].freq++
    }

    sort.Slice(pocket, func(i,j int) bool {
        return pocket[i].freq >= pocket[j].freq
    })

    answer := make([]rune, len(s))
    p := 0
    for i:=0;i<len(s);i+=2 {
        obj := pocket[p]
        if obj.freq > 0 {
            answer[i] = obj.x
            pocket[p].freq--
        } else {
            p++
            i-=2
        }
    }

    for i:=1;i<len(s);i+=2 {
        obj := pocket[p]
        if obj.freq > 0 {
            if p == 0 {return ""}
            answer[i] = obj.x
            pocket[p].freq--
        } else {
            p++
            i-=2
        }
    }


    return string(answer)
}