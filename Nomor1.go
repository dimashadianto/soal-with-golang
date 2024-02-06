package main

import "fmt"

func main() {
    fmt.Println(minNum(3, 5, 1))
}

func minNum(samDaily, kellyDaily, difference int) int {
    samSolved := difference + samDaily
    kellySolved := kellyDaily
    day := 1

    if kellyDaily <= samDaily {
        return -1
    }

    for samSolved-kellySolved >= 0 {
        samSolved += samDaily
        kellySolved += kellyDaily
        day++
    }

    return day
}
