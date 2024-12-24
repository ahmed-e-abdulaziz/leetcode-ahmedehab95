func validPath(n int, edges [][]int, source int, destination int) bool {
    if source == destination {
        return true
    }
    if len(edges) == 0 {
        return false
    }
    graph:= make([][]int, n)
    for _, e:=range edges {
        graph[e[0]] = append(graph[e[0]], e[1])
        graph[e[1]] = append(graph[e[1]], e[0])
    }
    return dfs(graph, make([]bool, n), source, destination)
}

func dfs(graph [][]int, visited []bool, source, dest int) bool {
    if source == dest {
        return true
    }
    for _, neighbour := range graph[source] {
        if visited[neighbour] {
            continue
        }
        visited[neighbour]=true
        found:=dfs(graph, visited, neighbour, dest)
        if found {
            return true
        }
    }
    return false
}