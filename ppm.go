package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func beginFile(path string, dim [2]int) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	header := "P3\n" + strconv.Itoa(dim[0]) + " " + strconv.Itoa(dim[1]) + "\n" + "255\n"
	_, err = file.WriteString(header)
	return err
}

func writePPMLine(path string, line string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}
	_, err = fmt.Fprintln(file, line)
	return err
}

func writePPMLines(path string, arr [][]int) error {
	for _, row := range arr {
		for _, val := range row {
			l := "000 000 000"
			switch val {
			case 0:
				l = "255 255 255"
			case 1:
				l = "000 000 000"
			case 2:
				l = "000 128 000"
			case 3:
				l = "128 128 000"
			case 4:
				l = "255 128 000"
			case 5:
				l = "255 255 000"
			case 6:
				l = "255 0 255"
			}
			l = strings.ReplaceAll(l, " ", "\n")
			err := writePPMLine(path, l)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	mapa := [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 3, 3, 1, 0, 0, 1, 4, 1, 0, 2, 2, 0, 4, 0, 1, 1, 3, 3, 1, 0, 0, 1, 4, 1, 0, 2, 2, 0, 4, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 1, 0, 1},
		{1, 2, 1, 6, 1, 0, 1, 2, 1, 1, 1, 1, 2, 1, 0, 1, 1, 2, 1, 6, 1, 0, 1, 2, 1, 1, 1, 1, 2, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 1, 0, 2, 0, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0, 2, 0, 0, 0, 0, 1, 0, 1},
		{1, 2, 1, 1, 1, 2, 1, 0, 1, 1, 1, 2, 1, 1, 0, 1, 1, 2, 1, 1, 1, 2, 1, 0, 1, 1, 1, 2, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 2, 0, 1, 0, 2, 2, 1, 5, 0, 1, 1, 0, 0, 0, 0, 0, 2, 0, 1, 0, 2, 2, 1, 5, 0, 1},
		{1, 0, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 3, 3, 1, 0, 0, 1, 4, 1, 0, 2, 2, 0, 4, 0, 1, 1, 3, 3, 1, 0, 0, 1, 4, 1, 0, 2, 2, 0, 4, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 1, 0, 1},
		{1, 2, 1, 6, 1, 0, 1, 2, 1, 1, 1, 1, 2, 1, 0, 1, 1, 2, 1, 6, 1, 0, 1, 2, 1, 1, 1, 1, 2, 1, 0, 1},
		{1, 0, 1, 0, 0, 0, 1, 0, 2, 0, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0, 2, 0, 0, 0, 0, 1, 0, 1},
		{1, 2, 1, 1, 1, 2, 1, 0, 1, 1, 1, 2, 1, 1, 0, 1, 1, 2, 1, 1, 1, 2, 1, 0, 1, 1, 1, 2, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 2, 0, 1, 0, 2, 2, 1, 5, 0, 1, 1, 0, 0, 0, 0, 0, 2, 0, 1, 0, 2, 2, 1, 5, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}
	// fileContent, err := ioutil.ReadFile("image.ppm")
	err := beginFile("image.ppm", [2]int{32, 16})
	if err != nil {
		panic(err)
	}

	err = writePPMLines("image.ppm", mapa)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(fileContent))
}
