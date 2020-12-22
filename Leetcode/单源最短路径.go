package main

import "fmt"

func main() {
	graph := make([][]int, 6, 6)
	for a := range graph {
		graph[a] = make([]int, 6, 6)
	}

	graph[0][1] = 100
	graph[1][0] = 100

	graph[0][2] = 1200
	graph[2][0] = 1200

	graph[1][2] = 900
	graph[2][1] = 900

	graph[1][3] = 300
	graph[3][1] = 300

	graph[2][3] = 400
	graph[3][2] = 400

	graph[2][4] = 500
	graph[4][2] = 500

	graph[3][4] = 1300
	graph[4][3] = 1300

	graph[3][5] = 1400
	graph[5][3] = 1400

	graph[4][5] = 1500
	graph[5][4] = 1500

	dist := Dijkstra(graph, 0)
	for i := range dist {
		fmt.Printf("dist[%d]=%d\n", i, dist[i])
	}

}
func Dijkstra(graph [][]int, start int) []int {
	n := len(graph)         // 图中顶点的个数
	visit := make([]int, n) // 标记已经作为中间节点访问的顶点
	dist := make([]int, n)  // 存储从起点start到其他顶点的最短路径

	for i := 0; i < n; i++ {
		dist[i] = graph[start][i] // 将 dist 数组初始化为最初图中的路径长度
	}

	visit[start] = 1 // 标记为起始顶点

	for i := 0; i < n; i++ { // 更新最短路径的循环  循环 N 次
		minDist := int(^uint(0) >> 1) // 将最短路径标记为 最大
		middle := 0

		for j := 0; j < n; j++ {
			if visit[j] == 0 && minDist > dist[j] {
				minDist = dist[j]
				middle = j
			}
		}

		for j := 0; j < n; j++ {
			if visit[j] == 0 && dist[j] > (dist[middle]+graph[middle][j]) {
				dist[j] = dist[middle] + graph[middle][j]
			}
		}
		visit[middle] = 1
	}
	return dist
}
