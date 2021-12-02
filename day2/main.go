package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e);
	}
}

func parse_command(command_s string) (string, int) {
	tmp := strings.Split(command_s, " ")

	command := tmp[0]
	distance, _ := strconv.Atoi(tmp[1])

	return command, distance
}

func exec_command(command_s string) (int, int, int) {
	command, distance := parse_command(command_s)

	x := 0
	y := 0
	z := 0

	switch command {
	case "forward":
		x = distance
		y = distance
	case "down":
		z = distance
	case "up":
		z = -distance
	}

	return x, y, z
}

func main() {
	dat, err := os.ReadFile("./data.txt");
	check(err)

	s := string(dat)
	commands := strings.Split(s, "\n")

	x, y, aim := 0, 0, 0

	for _, command := range commands {
		tmp_x, tmp_y, tmp_z := exec_command(command)

		x += tmp_x
		aim += tmp_z
		y += tmp_y * aim
	}

	fmt.Println("Final location ", x, y, aim)
	fmt.Println("Result", x * y)
}
