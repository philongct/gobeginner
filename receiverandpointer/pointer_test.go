package receiverandpointer

import (
    "testing"
    "fmt"
)

func turnOn(bulbs *uint64, pos int) {
    binPos := uint64(1) << uint(pos)
    *bulbs = *bulbs | binPos
}

func isMatched(bulbs uint64, m *uint64) bool {
    match := false
    for {
        result := uint64(0)
        if (*m <= bulbs) {
            result = *m & bulbs
        }

        if (result == *m) {
            match = true
            *m = (*m << 1) | 1
        } else {
            break
        }
    }

    return match
}

func solve(A []int) int {
    bulbs := uint64(0)
    m := uint64(1)
    count := 0

    for i, v := range A {
        turnOn(&bulbs, v - 1)
        if isMatched(bulbs, &m) {
            count++
            fmt.Printf("%d => %b\n", i, bulbs)
        }
    }

    return count
}

func Test1(t *testing.T) {
    A := []int{3, 2, 5, 1, 4, 10, 6, 9, 7, 8}
    solve(A)
}

func Test2(t *testing.T) {
    A := []int{3, 9, 5, 1, 4, 10, 6, 2, 7, 8}
    solve(A)
}

func Test3(t *testing.T) {
    A := []int{3, 9, 5, 1, 4, 2, 6, 10, 8, 7}
    solve(A)
}
