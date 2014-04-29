package main

import (
	"fmt"
)

func printMatrix(m [][]byte) {
	for i:=0;i<len(m);i++ {
		fmt.Printf("%q\n", m[i])
	}
}

func fall(m [][]byte) {
	n:=len(m)
	for i:=n-1;i>0;i-- {
		for j:=0;j<n;j++ {
			if m[i-1][j] == '.' && m[i][j] == ' ' {
				for k:=i;k<n && m[k][j] == ' ';k++ {
					m[k-1][j] = ' '
					m[k][j] = '.'
				}
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	m := make([][]byte, n)
	for i:=0; i<n; i++ {
        m[i] = make([]byte, n)
        for j:=0; j<n; j++ {
            fmt.Scanf("%c", &m[i][j])
        }
        fmt.Scanln()
    }
    fmt.Println("\nBefore:")
    printMatrix(m)
	fall(m)
	fmt.Println("\nAfter:")
	printMatrix(m)
}