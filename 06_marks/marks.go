package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

func marks(input io.Reader, output io.Writer) error {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(input)
	out = bufio.NewWriter(output)
	defer out.Flush()

	var dataInputs int
	fmt.Fscan(in, &dataInputs)

	var i = 0
	for i < dataInputs {

		var n, m int
		fmt.Fscan(in, &n, &m)

		marks := make([][]int, n)

		worstMark := 5
		worstMarkCoords := MarkCoords{0, 0, worstMark}
		secondWorstMark := 5
		//secondWorstMarkCoords := MarkCoords{0, 0}
		worstMarkIndex := make(map[int][]MarkCoords)
		j := 0

		rowSums := make([]int, n)
		colSums := make([]int, m)
		for j < n {
			var line string
			fmt.Fscan(in, &line)

			marks[j] = make([]int, m)
			for k, v := range line {
				marks[j][k], _ = strconv.Atoi(string(v))
				currentMark := marks[j][k]

				if currentMark < worstMark {
					secondWorstMark = worstMark
					//secondWorstMarkCoords = worstMarkCoords
					worstMark = marks[j][k]
					worstMarkCoords = MarkCoords{j, k, worstMark}
				} else if worstMark < currentMark && currentMark < secondWorstMark {
					secondWorstMark = currentMark
				}
				rowSums[j] += currentMark
			}

			j++
		}

		for j = 0; j < m; j++ {
			for k := 0; k < n; k++ {
				colSums[j] += marks[k][j]
			}
		}

		fillWorskMarkIndex(marks, n, m, worstMark, worstMarkIndex)
		fillWorskMarkIndex(marks, n, m, secondWorstMark, worstMarkIndex)

		if n <= 30 && m <= 30 {
			result := BruteForce(marks, n, m, worstMark, worstMarkCoords)

			fmt.Fprintln(output, result.x+1, result.y+1)

			i++
			continue

		}

		if worstMark == 5 {
			fmt.Fprintln(output, worstMarkCoords.x+1, worstMarkCoords.y+1)
		} else {

			if len(worstMarkIndex[worstMark]) > 1 {

				adjacentX, sortedByX := FindAdjacentXWorstMark(worstMarkIndex[worstMark])
				adjacentY, sortedByY := FindAdjacentYWorstMark(worstMarkIndex[worstMark])

				if adjacentX != -1 && adjacentY != -1 {

					if worstMark == 4 {
						fmt.Fprintln(output, worstMarkCoords.x+1, worstMarkCoords.y+1)
					} else {
						fmt.Fprintln(output, adjacentX+1, adjacentY+1)
					}

				} else if adjacentX != -1 {

					existsNonAdjacent := false
					var nonAdjacent MarkCoords
					for _, v := range sortedByY {
						if v.x != adjacentX {
							existsNonAdjacent = true
							nonAdjacent = v
							break
						}
					}

					if existsNonAdjacent {
						fmt.Fprintln(output, adjacentX+1, nonAdjacent.y+1)
					} else {
						fmt.Fprintln(output, adjacentX+1, worstMarkCoords.y+1)
					}

				} else if adjacentY != -1 {

					existsNonAdjacent := false
					var nonAdjacent MarkCoords
					for _, v := range sortedByX {
						if v.y != adjacentY {
							existsNonAdjacent = true
							nonAdjacent = v
							break
						}
					}

					if existsNonAdjacent {
						fmt.Fprintln(output, nonAdjacent.x+1, adjacentY+1)
					} else {
						fmt.Fprintln(output, worstMarkCoords.x+1, adjacentY+1)
					}
				} else {

					if len(worstMarkIndex[secondWorstMark]) == 1 && worstMarkIndex[secondWorstMark][0].x == worstMarkCoords.x || worstMarkIndex[secondWorstMark][0].y == worstMarkCoords.y {

						if len(worstMarkIndex[worstMark]) == 2 && n != 2 && m != 2 {

							worstMarkFirst, _ := FindWorstMark(marks, n, m, worstMarkIndex[worstMark][0].x, worstMarkIndex[worstMark][1].y)
							worstMarkSecond, _ := FindWorstMark(marks, n, m, worstMarkIndex[worstMark][1].x, worstMarkIndex[worstMark][0].y)

							if worstMarkFirst > worstMarkSecond {
								fmt.Fprintln(output, worstMarkIndex[secondWorstMark][0].x+1, worstMarkCoords.y+1)
							} else if worstMarkFirst < worstMarkSecond {

								fmt.Fprintln(output, worstMarkCoords.x+1, worstMarkIndex[secondWorstMark][0].y+1)
							} else {
								if worstMarkIndex[worstMark][0].x < worstMarkIndex[worstMark][1].x {
									fmt.Fprintln(output, worstMarkIndex[worstMark][0].x+1, worstMarkIndex[worstMark][1].y+1)
								} else {
									fmt.Fprintln(output, worstMarkIndex[worstMark][1].x+1, worstMarkIndex[worstMark][0].y+1)
								}
							}

						} else {
							fmt.Fprintln(output, worstMarkIndex[secondWorstMark][0].x+1, worstMarkIndex[secondWorstMark][0].y+1)
						}

						//fmt.Fprintln(output, worstMarkIndex[secondWorstMark][0].x+1, worstMarkIndex[secondWorstMark][0].y+1)
					} else {

						if m == 2 && n == 2 && len(worstMarkIndex[worstMark]) == 2 {

							secondWorstXToEliminate := worstMarkIndex[worstMark][0].x + 1
							secondWorstYToEliminiate := worstMarkIndex[worstMark][0].y + 1

							worstXToEliminate := worstMarkIndex[worstMark][1].x + 1
							worstYToEliminiate := worstMarkIndex[worstMark][1].y + 1

							worstMarkFirst, worstMarkFirstCoords := FindWorstMark(marks, n, m, worstMarkIndex[worstMark][0].x, worstMarkIndex[worstMark][1].y)
							worstMarkSecond, worstMarkSecondCoords := FindWorstMark(marks, n, m, worstMarkIndex[worstMark][1].x, worstMarkIndex[worstMark][0].y)

							if worstMarkFirst > worstMarkSecond {
								fmt.Fprintln(output, worstXToEliminate, secondWorstYToEliminiate)
							} else if worstMarkFirst < worstMarkSecond {
								fmt.Fprintln(output, secondWorstXToEliminate, worstYToEliminiate)
							} else {
								if worstMarkFirstCoords.x < worstMarkSecondCoords.x {
									fmt.Fprintln(output, worstXToEliminate, secondWorstYToEliminiate)
								} else {
									fmt.Fprintln(output, secondWorstXToEliminate, worstYToEliminiate)
								}
							}

						} else {

							if len(worstMarkIndex[worstMark]) == 2 {

								if worstMarkIndex[worstMark][0].y < worstMarkIndex[worstMark][1].y {
									fmt.Fprintln(output, worstMarkIndex[worstMark][1].x+1, worstMarkIndex[worstMark][0].y+1)
								} else {
									fmt.Fprintln(output, worstMarkIndex[worstMark][0].x+1, worstMarkIndex[worstMark][1].y+1)
								}

							} else {

								//firstWorstMark, _ := FindWorstMark(marks, n, m, worstMarkIndex[worstMark][0].x, worstMarkIndex[worstMark][len(worstMarkIndex[worstMark])-1].y)
								//secondWorstMark, _ := FindWorstMark(marks, n, m, worstMarkIndex[worstMark][0].x, worstMarkIndex[worstMark][0].y)
								//
								//if firstWorstMark > secondWorstMark {
								//	fmt.Fprintln(output, worstMarkIndex[worstMark][0].x+1, worstMarkIndex[worstMark][len(worstMarkIndex[worstMark])-1].y+1)
								//} else {
								fmt.Fprintln(output, worstMarkIndex[worstMark][0].x+1, worstMarkIndex[worstMark][0].y+1)
								//}

							}
						}
						//fmt.Fprintln(output, worstMarkIndex[worstMark][len(worstMarkIndex[worstMark])-1].x+1, worstMarkIndex[worstMark][0].y+1)
					}

				}

			} else {

				if len(worstMarkIndex[secondWorstMark]) > 1 {

					adjacentX, _ := FindAdjacentXWorstMark(worstMarkIndex[secondWorstMark])
					adjacentY, _ := FindAdjacentYWorstMark(worstMarkIndex[secondWorstMark])

					if adjacentX != -1 && adjacentY != -1 {
						if m == 2 && n == 2 {
							fmt.Fprintln(output, adjacentX+1, worstMarkCoords.y+1)
						} else {

							firstWorstMark, _ := FindWorstMark(marks, n, m, adjacentX, worstMarkCoords.y)
							secondWorstMark, _ := FindWorstMark(marks, n, m, worstMarkCoords.x, adjacentY)

							if firstWorstMark > secondWorstMark {
								fmt.Fprintln(output, adjacentX+1, worstMarkCoords.y+1)
							} else if firstWorstMark < secondWorstMark {
								fmt.Fprintln(output, worstMarkCoords.x+1, adjacentY+1)
							} else {
								if worstMarkCoords.x <= adjacentX {
									fmt.Fprintln(output, worstMarkCoords.x+1, adjacentY+1)
								} else {
									fmt.Fprintln(output, adjacentX+1, worstMarkCoords.y+1)
								}
							}
						}
					} else if adjacentX != -1 {
						fmt.Fprintln(output, adjacentX+1, worstMarkCoords.y+1)
					} else if adjacentY != -1 {
						fmt.Fprintln(output, worstMarkCoords.x+1, adjacentY+1)
					} else {
						fmt.Fprintln(output, worstMarkIndex[worstMark][0].x+1, worstMarkIndex[worstMark][len(worstMarkIndex[worstMark])-1].y+1)
					}
				} else {
					if m == 2 && n == 2 {

						secondWorstXToEliminate := worstMarkIndex[secondWorstMark][0].x + 1
						secondWorstYToEliminiate := worstMarkIndex[secondWorstMark][0].y + 1

						worstXToEliminate := worstMarkCoords.x + 1
						worstYToEliminiate := worstMarkCoords.y + 1

						worstMarkFirst, worstMarkFirstCoords := FindWorstMark(marks, n, m, worstMarkCoords.x, worstMarkIndex[secondWorstMark][0].y)
						worstMarkSecond, worstMarkSecondCoords := FindWorstMark(marks, n, m, worstMarkIndex[secondWorstMark][0].x, worstMarkCoords.y)

						if worstMarkFirst > worstMarkSecond {
							fmt.Fprintln(output, n-worstMarkFirstCoords.x, m-worstMarkFirstCoords.y)
						} else if worstMarkFirst < worstMarkSecond {
							fmt.Fprintln(output, n-worstMarkSecondCoords.x, m-worstMarkSecondCoords.y)
						} else {
							if worstMarkFirstCoords.x < worstMarkSecondCoords.x {
								fmt.Fprintln(output, secondWorstXToEliminate, worstYToEliminiate)
							} else {
								fmt.Fprintln(output, worstXToEliminate, secondWorstYToEliminiate)
							}
						}

					} else {

						worstMarkFirst, _ := FindWorstMark(marks, n, m, worstMarkCoords.x, worstMarkIndex[secondWorstMark][0].y)
						worstMarkSecond, _ := FindWorstMark(marks, n, m, worstMarkIndex[secondWorstMark][0].x, worstMarkCoords.y)

						if worstMarkFirst > worstMarkSecond {
							fmt.Fprintln(output, worstMarkIndex[secondWorstMark][0].x+1, worstMarkCoords.y+1)
						} else if worstMarkFirst < worstMarkSecond {

							fmt.Fprintln(output, worstMarkCoords.x+1, worstMarkIndex[secondWorstMark][0].y+1)
						} else {
							if worstMarkCoords.x < worstMarkIndex[secondWorstMark][0].x {
								fmt.Fprintln(output, worstMarkCoords.x+1, worstMarkIndex[secondWorstMark][0].y+1)
							} else {
								fmt.Fprintln(output, worstMarkIndex[secondWorstMark][0].x+1, worstMarkCoords.y+1)
							}
						}
					}

				}

			}

		}

		i++
	}

	return nil
}

func BruteForce(marks [][]int, n int, m int, worstMark int, worstMarkCoords MarkCoords) MarkCoords {
	minWorstMark := worstMark
	minWorstMarkCoords := MarkCoords{0, 0, minWorstMark}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			worstMark, _ := FindWorstMark(marks, n, m, i, j)
			if worstMark > minWorstMark {
				minWorstMark = worstMark
				minWorstMarkCoords = MarkCoords{i, j, worstMark}
			}
		}
	}

	return minWorstMarkCoords
}

func FindWorstMark(marks [][]int, n int, m int, rowToExclude int, colToExclude int) (int, MarkCoords) {
	worstMark := 5
	worstMarkCoords := MarkCoords{-1, -1, worstMark}
	for i := 0; i < n; i++ {
		if i == rowToExclude {
			continue
		}
		for j := 0; j < m; j++ {
			if j == colToExclude {
				continue
			}
			currentMark := marks[i][j]
			if currentMark < worstMark || (currentMark == worstMark && (worstMarkCoords.x == -1 && worstMarkCoords.y == -1)) {
				worstMark = currentMark
				worstMarkCoords = MarkCoords{i, j, worstMark}
			}
		}
	}
	return worstMark, worstMarkCoords
}

func FindAdjacentXWorstMark(marks []MarkCoords) (int, []MarkCoords) {
	sort.Slice(marks, func(i, j int) bool {
		return marks[i].x < marks[j].x
	})

	sorted := make([]MarkCoords, 0, len(marks))

	sorted = append(sorted, marks...)

	for i := 0; i < len(marks)-1; i++ {
		if marks[i].x == marks[i+1].x {
			return marks[i].x, sorted
		}
	}

	return -1, sorted
}

func FindAdjacentYWorstMark(marks []MarkCoords) (int, []MarkCoords) {
	sort.Slice(marks, func(i, j int) bool {
		return marks[i].y < marks[j].y
	})

	sorted := make([]MarkCoords, 0, len(marks))

	sorted = append(sorted, marks...)

	for i := 0; i < len(marks)-1; i++ {
		if marks[i].y == marks[i+1].y {
			return marks[i].y, sorted
		}
	}

	return -1, sorted
}

type MarkCoords struct {
	x, y int
	mark int
}

func fillWorskMarkIndex(marks [][]int, n, m, worstMark int, worstMarkIndex map[int][]MarkCoords) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if marks[i][j] == worstMark {
				worstMarkIndex[worstMark] = append(worstMarkIndex[worstMark], MarkCoords{i, j, worstMark})
			}
		}
	}
}
