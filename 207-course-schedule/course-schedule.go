func canFinish(numCourses int, prerequisites [][]int) bool {
    indegree := make([]int, numCourses)
    nextCourses := make(map[int][]int)
    for _, preqs := range prerequisites {
        first, next := preqs[1], preqs[0]
        indegree[next]++
        if _, exist := nextCourses[first]; !exist {
            nextCourses[first] = make([]int, 0)
        }
        nextCourses[first] = append(nextCourses[first], next)
    }

    q := make([]int, 0)
    for i:=0;i<numCourses;i++{
        if indegree[i] == 0 {
            q = append(q, i)
        }
    }

    completed := 0
    for len(q) > 0 {
        front := q[0]
        q = q[1:]
        completed++
        for _, next := range nextCourses[front] {
            indegree[next]--
            if indegree[next] == 0 {
                q = append(q, next)
            }
        }
    }

    return completed == numCourses
}