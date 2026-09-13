type point struct{ r, c int }

func largestOverlap(img1 [][]int, img2 [][]int) int {
    n := len(img1)
    
    // Collect positions of all 1s in each image
    type point struct{ r, c int }
    var ones1, ones2 []point
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if img1[i][j] == 1 {
                ones1 = append(ones1, point{i, j})
            }
            if img2[i][j] == 1 {
                ones2 = append(ones2, point{i, j})
            }
        }
    }
    
    // For each pair of 1s, compute the offset that aligns them.
    // Count how many pairs share the same offset.
    counts := make(map[point]int)
    answer := 0
    for _, p1 := range ones1 {
        for _, p2 := range ones2 {
            offset := point{p2.r - p1.r, p2.c - p1.c}
            counts[offset]++
            if counts[offset] > answer {
                answer = counts[offset]
            }
        }
    }
    return answer
}