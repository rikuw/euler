package main

import (
    "math"
    "fmt"
)

func main() {
    var s [4000000]int

    for k := 0; k <= 55; k++ {
        s[k] = (100003 - 200003 * k + 300007 * e(k, 3)) % 1000000 - 500000
    }

    for k := 56; k <= 3999999; k++ {
        s[k] = (s[k - 24] + s[k - 55] + 1000000) % 1000000 - 500000
    }

    fmt.Printf("s[10] : %d\n", s[10])
    fmt.Printf("s[100] : %d\n", s[100])

    var a [2000][2000]int
    p := 0

    for i := 0; i < 2000; i++ {
        for j := 0; j < 2000; j++ {
            a[i][j] = s[p]
            p++
        }
    }

    fmt.Printf("a[1000][1000] : %d\n", a[1000][1000])

    var b [4][4]int
    b[0][0] = -2
    b[0][1] = 5
    b[0][2] = 3
    b[0][3] = 2
    b[1][0] = 9
    b[1][1] = -6
    b[1][2] = 5
    b[1][3] = 1
    b[2][0] = 3
    b[2][1] = 2
    b[2][2] = 7
    b[2][3] = 3
    b[3][0] = -1
    b[3][1] = 8
    b[3][2] = -4
    b[3][3] = 8

    /*
    def max_subarray(A):
    max_ending_here = max_so_far = A[0]
    for x in A[1:]:
        max_ending_here = max(x, max_ending_here + x)
        max_so_far = max(max_so_far, max_ending_here)
    return max_so_far
     */
    var maxEndingHereHorizontal, maxSoFarHorizontal int
    /*for i := 0; i < 2000; i++ {
        for j := 1; j < 2000; j++ {
            maxEndingHereHorizontal, maxSoFarHorizontal = a[i][0], a[i][0]
            maxEndingHereHorizontal = max(a[i][j], maxEndingHereHorizontal + a[i][j])
            maxSoFarHorizontal = max(maxSoFarHorizontal, maxEndingHereHorizontal)
        }
    }*/

    var maxEndingHereVertical, maxSoFarVertical int
    /*for i := 0; i < 2000; i++ {
        for j := 1; j < 2000; j++ {
            maxEndingHereHorizontal, maxSoFarVertical = a[i][0], a[i][0]
            maxEndingHereVertical = max(a[j][i], maxEndingHereVertical + a[j][i])
            maxSoFarVertical = max(maxSoFarVertical, maxEndingHereVertical)
        }
    }*/

    var maxEndingHereDiagonal, maxSoFarDiagonal int
    /*for i := 1; i < 2000; i++ {
        maxEndingHereDiagonal, maxSoFarDiagonal = a[0][0], a[0][0]
        maxEndingHereDiagonal = max(a[i][i], maxEndingHereDiagonal + a[i][i])
        maxSoFarDiagonal = max(maxSoFarDiagonal, maxEndingHereDiagonal)
    }*/

    var maxEndingHereAntiDiagonal, maxSoFarAntiDiagonal int
    /*for i := 1998; i >= 0; i-- {
        maxEndingHereAntiDiagonal, maxSoFarAntiDiagonal = a[1999][1999], a[1999][1999]
        maxEndingHereAntiDiagonal = max(a[i][i], maxEndingHereAntiDiagonal + a[i][i])
        maxSoFarAntiDiagonal = max(maxSoFarAntiDiagonal, maxEndingHereAntiDiagonal)
    }

    verticalOrHorizontal := max(maxEndingHereVertical, maxEndingHereHorizontal)
    diagonalOrAntiDiagonal := max(maxEndingHereDiagonal, maxEndingHereAntiDiagonal)

    realMax := max(verticalOrHorizontal, diagonalOrAntiDiagonal)*/

    // --------

    for i := 0; i < 4; i++ {
        for j := 1; j < 4; j++ {
            maxEndingHereHorizontal, maxSoFarHorizontal = b[i][0], b[i][0]
            maxEndingHereHorizontal = max(b[i][j], maxEndingHereHorizontal + b[i][j])
            maxSoFarHorizontal = max(maxSoFarHorizontal, maxEndingHereHorizontal)
        }
    }

    for i := 0; i < 4; i++ {
        for j := 1; j < 4; j++ {
            maxEndingHereVertical, maxSoFarVertical = b[i][0], b[i][0]
            maxEndingHereVertical = max(b[j][i], maxEndingHereVertical + b[j][i])
            maxSoFarVertical = max(maxSoFarVertical, maxEndingHereVertical)
        }
    }

    for i := 1; i < 4; i++ {
        maxEndingHereDiagonal, maxSoFarDiagonal = b[0][0], b[0][0]
        maxEndingHereDiagonal = max(b[i][i], maxEndingHereDiagonal + b[i][i])
        maxSoFarDiagonal = max(maxSoFarDiagonal, maxEndingHereDiagonal)
    }

    for i := 2; i >= 0; i-- {
        maxEndingHereAntiDiagonal, maxSoFarAntiDiagonal = b[3][3], b[3][3]
        maxEndingHereAntiDiagonal = max(b[i][i], maxEndingHereAntiDiagonal + b[i][i])
        maxSoFarAntiDiagonal = max(maxSoFarAntiDiagonal, maxEndingHereAntiDiagonal)
    }

    verticalOrHorizontal := max(maxSoFarVertical, maxSoFarHorizontal)
    diagonalOrAntiDiagonal := max(maxSoFarDiagonal, maxSoFarAntiDiagonal)

    realMax := max(verticalOrHorizontal, diagonalOrAntiDiagonal)

    fmt.Printf("Hori : %d\n", maxSoFarHorizontal)
    fmt.Printf("Vert : %d\n", maxSoFarVertical)
    fmt.Printf("Diag : %d\n", maxSoFarDiagonal)
    fmt.Printf("Anti-Diag : %d\n", maxSoFarAntiDiagonal)
    fmt.Printf("Biggest in example : %d\n", realMax)
}

func e(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func max(x, y int) int {
    return int(math.Max(float64(x), float64(y)))
}