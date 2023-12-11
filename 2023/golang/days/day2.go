package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_game_id(line string) (int, error) {
	game_string := strings.Split(line, ":")
	game_id_string := strings.Split(game_string[0], " ")[1]

	return strconv.Atoi(game_id_string)
}

func get_rolls(line string) []string {
	game_string := strings.Split(line, ":")
	rolls := strings.Split(game_string[1], ";")

	for i, curr := range rolls {
		rolls[i] = strings.TrimPrefix(curr, " ")
	}

	return rolls
}

func check_dice(rolls []string) {
	for _, curr := range rolls {
		split := strings.Split(curr, ",")
		for _, x := range split {
			trimmed := strings.TrimPrefix(x, " ")
			value, err := strconv.Atoi(strings.Split(trimmed, " ")[0])
			if err != nil {
				panic(err)
			}
			color := strings.Split(trimmed, " ")[1]

			fmt.Println("Value: ", value, "\nColor: ", color)
		}

	}
}

func DayTwo() {
	file, err := os.Open("../data/day2.test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// sum_of_ids := 0
	// var red, green, blue = []int{}, []int{}, []int{}
	// var green = []int{}
	// var blue = []int{}
	for scanner.Scan() {
		line := scanner.Text()
		// game_id, err := get_game_id(line)
		if err != nil {
			panic(err)
		}
		rolls := get_rolls(line)
		check_dice(rolls)
		// fmt.Println(game_id)

		// sum_of_ids += line_id
	}
}
