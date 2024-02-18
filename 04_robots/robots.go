package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func robots(input io.Reader, output io.Writer) error {

	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(input)
	//scanner := bufio.NewScanner(input)
	out = bufio.NewWriter(output)
	defer out.Flush()

	var dataInputs int
	_, err := fmt.Fscan(in, &dataInputs)
	if err != nil {
		return fmt.Errorf("unable to read data inputs")
	}

	var i = 0
	for i < dataInputs {

		var n, m int

		fmt.Fscan(in, &n, &m)
		matrix := make([]string, n)

		var aRobot Robot
		var bRobot Robot

		for j := 0; j < n; j++ {
			fmt.Fscan(in, &matrix[j])

			aX := strings.Index(matrix[j], "A")
			if aX != -1 {
				aRobot.y = j
				aRobot.x = aX
				aRobot.printSymbol = "a"
			}

			bX := strings.Index(matrix[j], "B")
			if bX != -1 {
				bRobot.y = j
				bRobot.x = bX
				bRobot.printSymbol = "b"
			}
		}

		var topLeftRobot Robot
		var bottomRightRobot Robot
		if aRobot.y < bRobot.y {
			topLeftRobot = aRobot
			bottomRightRobot = bRobot
		} else if aRobot.y == bRobot.y && aRobot.x < bRobot.x {
			topLeftRobot = aRobot
			bottomRightRobot = bRobot
		} else {
			topLeftRobot = bRobot
			bottomRightRobot = aRobot
		}

		leftRobotMove(topLeftRobot, matrix)

		rightRobotMove(bottomRightRobot, m, n, matrix)

		print(matrix, output)

		i++
	}

	return nil
}

func rightRobotMove(b Robot, m int, n int, matrix []string) {
	robotBX := b.x
	robotBY := b.y

	prevRobotBX := robotBX
	prevRobotBY := robotBY

	for !(robotBX == m-1 && robotBY == n-1) {

		for robotBY < n {
			if robotBY+1 < n && matrix[robotBY+1][robotBX] == '.' {
				matrix[robotBY+1] = matrix[robotBY+1][:robotBX] + b.printSymbol + matrix[robotBY+1][robotBX+1:]
				robotBY++
			}
			if prevRobotBX == robotBX && prevRobotBY == robotBY {
				break
			}

			prevRobotBX = robotBX
			prevRobotBY = robotBY
		}

		for robotBX < m {
			if robotBX+1 < m && matrix[robotBY][robotBX+1] == '.' {
				matrix[robotBY] = matrix[robotBY][:robotBX+1] + b.printSymbol + matrix[robotBY][robotBX+2:]
				robotBX++
				break
			}
			if prevRobotBX == robotBX && prevRobotBY == robotBY {
				break
			}

			prevRobotBX = robotBX
			prevRobotBY = robotBY
		}

		//if prevRobotBX == robotBX && prevRobotBY == robotBY {
		//	break
		//}
	}
}

func leftRobotMove(a Robot, matrix []string) {
	robotAX := a.x
	robotAY := a.y

	prevRobotAX := robotAX
	prevRobotAY := robotAY

	for !(robotAY == 0 && robotAX == 0) {

		for robotAY >= 0 {
			if robotAY-1 >= 0 && matrix[robotAY-1][robotAX] == '.' {
				matrix[robotAY-1] = matrix[robotAY-1][:robotAX] + a.printSymbol + matrix[robotAY-1][robotAX+1:]
				robotAY--
			}
			if prevRobotAX == robotAX && prevRobotAY == robotAY {
				break
			}

			prevRobotAX = robotAX
			prevRobotAY = robotAY
		}

		for robotAX >= 0 {
			if robotAX-1 >= 0 && matrix[robotAY][robotAX-1] == '.' {
				matrix[robotAY] = matrix[robotAY][:robotAX-1] + a.printSymbol + matrix[robotAY][robotAX:]
				robotAX--
				break
			}
			if prevRobotAX == robotAX && prevRobotAY == robotAY {
				break
			}

			prevRobotAX = robotAX
			prevRobotAY = robotAY
		}

	}
}

type Robot struct {
	x, y        int
	printSymbol string
}

func print(matrix []string, output io.Writer) {
	for i := 0; i < len(matrix); i++ {
		fmt.Fprintln(output, matrix[i])
	}
}
