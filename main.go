package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var verticeGroups map[int][]int

//if only a count of edge groups is required, then there isn't a need to store group members; this can be replaced with a simple counter instead
//this ds represents a group of neighbours or in other words, connected vertices
var verticeGroupAddr map[int]int

func numGroups(edgeGraph []string) int32 {
	//make space for verticeGroups
	verticeGroups = make(map[int][]int)
	verticeGroupAddr = make(map[int]int)

	//loop through each string
	for i := 0; i < len(edgeGraph); i++ {
		currentVerticeProfile := edgeGraph[i]

		groupIndex := verticeNeighbourGroup(i, currentVerticeProfile)
		verticeGroups[groupIndex] = addVerticeToGroup(i, verticeGroups[groupIndex])
	}
	return int32(len(verticeGroups))
}

func addVerticeToGroup(newVerticeIndex int, verticeGroup []int) []int {
	verticeGroup = append(verticeGroup, newVerticeIndex)
	return verticeGroup
}

//loop through all values, if 1 is found, then add that index to group if the mapaddre has it, or else create a new verticegroup and attach to groupAddr
func verticeNeighbourGroup(verticeIndex int, currentverticeProfile string) int {
	for i := 0; i < len(currentverticeProfile); i++ {
		c := currentverticeProfile[i]

		if verticeIndex == i {
			continue
		}

		if c == '1' {
			//is Neighbour, add i(neighbour) to whatever is in addr, if no addr then create a new group
			if verticeGroupAddr[verticeIndex] == 0 {
				verticeGroupAddr[verticeIndex] = verticeIndex
				verticeGroupAddr[i] = verticeIndex
			} else {
				verticeGroupIndex := verticeGroupAddr[verticeIndex]
				verticeGroupAddr[i] = verticeGroupIndex
			}
		}
	}
	return verticeGroupAddr[verticeIndex]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	edgeGraphCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var edgeGraph []string

	for i := 0; i < int(edgeGraphCount); i++ {
		edgeGraphItem := readLine(reader)
		edgeGraph = append(edgeGraph, edgeGraphItem)
	}

	res := numGroups(edgeGraph)
	fmt.Printf("\n Number of groups: %d\n", res)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
