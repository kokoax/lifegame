package main

import (
	"time"

	"github.com/rthornton128/goncurses"

	"github.com/kokoax/lifegame/go/cell"
	"github.com/kokoax/lifegame/go/frame"
)

func generator(mapStr []string) (cells [][]cell.Cell) {
	for i := range mapStr {
		cells = append(cells, []cell.Cell{})
		for j := range mapStr[i] {
			if mapStr[i][j] == '#' {
				cells[i] = append(cells[i], cell.NewIsLiveCell())
			} else {
				cells[i] = append(cells[i], cell.NewIsDeadCell())
			}
		}
	}
	return
}

func main() {
	cells := generator([]string{
		"......................................",
		".........................#............",
		".......................#.#............",
		".............##......##............##.",
		"............#...#....##............##.",
		".##........#.....#...##...............",
		".##........#...#.##....#.#............",
		"...........#.....#.......#............",
		"............#...#.....................",
		".............##.......................",
		"......................................",
		"......................................",
		"......................................",
		"......................................",
		"......................................",
		"......................................",
		"......................................",
		"......................................",
	})

	frame, err := frame.NewFrameWithCells(len(cells[0]), len(cells), cells)
	if err != nil {
		panic("Fatal: cannot generate frame")
	}
	defer frame.Close()
	goncurses.Cursor(0)

	for i := 0; i < 200; i++ {
		frame.Visualize()
		time.Sleep(10 * time.Millisecond)

		frame, err = frame.Next()
		if err != nil {
			panic(err)
		}
		frame, err = frame.Generation()
		if err != nil {
			panic(err)
		}
	}
}
