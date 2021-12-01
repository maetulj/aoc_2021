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

func main() {
	dat, err := os.ReadFile("./data.txt");
	check(err)

	s := string(dat)
	depths := strings.Split(s, "\n")

	increases := -1
	tmp_depth := 0

	var depths_mean []int

	for i := 0; i < (len(depths) - 2); i++ {
		tmp_0, _ := strconv.Atoi(depths[i])
		tmp_1, _ := strconv.Atoi(depths[i+1])
		tmp_2, _ := strconv.Atoi(depths[i+2])
		
		tmp := tmp_0 + tmp_1 + tmp_2

		depths_mean = append(depths_mean, tmp)
	}

	for _, depth := range depths_mean {
		// for _, depth := range depths_mean {
		// depth_m, err := strconv.Atoi(depth)
		// check(err)
		depth_m := depth

		if depth_m > tmp_depth {
			increases += 1
		}

		tmp_depth = depth_m
	}

	fmt.Println(increases)
}
