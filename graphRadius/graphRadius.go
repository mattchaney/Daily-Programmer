package main

import (
	"fmt"
	"math"
)

func findRadius(dist [][]int) (int) {
	radius := make([]int, len(dist))
	var max int
	for i:=0;i<len(dist);i++ {
		if max < dist[i][0] {
			max = dist[i][0]	
		}
		for j:=0;j<len(dist[j])-1;j++ {
			if max < dist[i][j+1] {
				max = dist[i][j+1]
			}
		}
		radius[i] = max
	}
	res := radius[0]
	for i:=0;i<len(radius)-1;i++ {
		if res < radius[i+1] {
			res = radius[i+1]
		}
	}
	return res
}


func printMatrix(m [][]int) {
	for i:=0;i<len(m);i++ {
		fmt.Println(m[i])
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	m := make([][]int, n)
	dist := make([][]int, n)
	for i:=0;i<n;i++ {
		m[i] = make([]int, n)
		dist[i] = make([]int, n)
	}
	for i:=0;i<n;i++ {
		for j:=0;j<n;j++ {
			fmt.Scan(&m[i][j])
			if m[i][j] == 1 {
				dist[i][j] = 1
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	for k:=0;k<n;k++ {
		for i:=0;i<n;i++ {
			for j:=0;j<n;j++ {
				if dist[i][j] > dist[i][k] + dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	
	fmt.Println("Radius:",findRadius(dist))
}