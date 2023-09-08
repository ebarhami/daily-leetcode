func generate(numRows int) [][]int {
    ans := make([][]int,numRows)
    for i:=0;i<numRows;i++{
        ans[i] = make([]int,i+1)
        for j:=0;j<=i;j++{
            if i<=1 || j==i || j==0 {
                ans[i][j] = 1
            } else {
                ans[i][j] = ans[i-1][j]+ans[i-1][j-1]
            }
        }
    }
    
    return ans
}