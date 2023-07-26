func isValid(s string) bool {
    q := make([]rune, 0)
    
    q = append(q, rune(s[0]))
    
    per := make(map[rune]rune)
    per['('] = ')'
    per['{'] = '}'
    per['['] = ']'
    
    
    for i:=1;i<len(s);i++{
        c := rune(s[i])
        
        if len(q) == 0 {
            q = append(q, c)
            continue
        }
        top := q[len(q)-1]
        if per[top] == c { // pop
            q = q[:len(q)-1]
        } else { // push
            q = append(q, c)
        }
    }
    
    return len(q) == 0
}


/*
())() // invalid
)))(((

-> we have more close bracket than the open one
nOpen++, nOpen--

nOpen == 0 -> valid
nOpen < 0

0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0
0 0 1 0 1 0 0 0
0 1 1 1 1 1 0 0
1 1 1 1 1 1 1 0
( ( ( ) ( ) ) )

2*n -> n open + n close
C(2*n, n)/

*/