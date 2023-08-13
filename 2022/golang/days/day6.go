package days

import (
	"bufio"
	"fmt"
	"os"
)

func DaySix(packetsize int) {
	file, err := os.Open("../data/day6.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	marker := make([]byte, 0, packetsize)
	answer := 1
	for {
		char, err := reader.ReadByte()

		if err != nil {
			fmt.Println(err)
			break
		}

		// push to queue
		marker = append(marker, char)

		// once size is four, check for uniques
		// if success, break with answer
		// if failure, dequeue
		if len(marker) == packetsize {
			visited := make(map[byte]bool)
			for i, curr := range marker {
				if visited[curr] == true {
					// we have a failure
					marker = marker[1:]
					break
				} else if i == packetsize-1 {
					goto done
				} else {
					visited[curr] = true
				}
			}
		}
		answer++
	}

done:
	fmt.Printf("Solution: %v\n", answer)
}