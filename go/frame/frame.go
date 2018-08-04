package frame

import (
	"github.com/rthornton128/goncurses"

	"github.com/kokoax/lifegame/go/cell"
)

// Frame is
type Frame interface {
	Next() (Frame, error)
	Visualize() error
	Generation() (Frame, error)
	Close() error
	Ncurse() *goncurses.Window
}

type frameImpl struct {
	W      int
	H      int
	Cells  [][]cell.Cell
	ncurse *goncurses.Window
}

func newFrame(w int, h int, cells [][]cell.Cell, ncurse *goncurses.Window) Frame {
	return &frameImpl{
		W:      w,
		H:      h,
		Cells:  cells,
		ncurse: ncurse,
	}
}

func newFrameWithNcurse(w int, h int, cells [][]cell.Cell) (Frame, error) {
	ncurse, err := goncurses.Init()
	if err != nil {
		return nil, err
	}
	settedCells := setAround(w, h, cells)

	return newFrame(w, h, settedCells, ncurse), nil
}

// NewFrame is
func NewFrame(w int, h int) (Frame, error) {
	return newFrameWithNcurse(w, h, initCell(w, h))
}

// NewFrameWithCells is
func NewFrameWithCells(w int, h int, cells [][]cell.Cell) (Frame, error) {
	return newFrameWithNcurse(w, h, cells)
}

func initCell(w int, h int) (cells [][]cell.Cell) {
	for i := 0; i < h; i++ {
		cells = append(cells, []cell.Cell{})
		for j := 0; j < w; j++ {
			cells[i] = append(cells[i], cell.NewCell())
		}
	}

	return
}

func checkIndex(i int, j int, w int, h int, cells [][]cell.Cell) cell.Cell {
	if i >= 0 && j >= 0 && i < h && j < w {
		return cells[i][j]
	}

	return nil
}

func getAround(i int, j int, w int, h int, cells [][]cell.Cell) [][]cell.Cell {
	return [][]cell.Cell{
		[]cell.Cell{checkIndex(i-1, j-1, w, h, cells), checkIndex(i-1, j, w, h, cells), checkIndex(i-1, j+1, w, h, cells)},
		[]cell.Cell{checkIndex(i, j-1, w, h, cells), nil, checkIndex(i, j+1, w, h, cells)},
		[]cell.Cell{checkIndex(i+1, j-1, w, h, cells), checkIndex(i+1, j, w, h, cells), checkIndex(i+1, j+1, w, h, cells)},
	}
}

func setAround(w int, h int, cells [][]cell.Cell) (ret [][]cell.Cell) {
	// f.Cellsをdeep copyしている
	for i := range cells {
		ret = append(ret, []cell.Cell{})
		for j := range cells[i] {
			ret[i] = append(ret[i], cells[i][j].Clone())
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ret[i][j] = ret[i][j].SetAroundCell(getAround(i, j, w, h, cells))
			// ret[i][j].SetAroundCell(getAround(i, j, w, h, cells))
		}
	}

	return
}

func (f *frameImpl) Ncurse() *goncurses.Window {
	return f.ncurse
}

func (f *frameImpl) Next() (Frame, error) {
	cells := [][]cell.Cell{}
	// f.Cellsをdeep copyしている
	for i := range f.Cells {
		cells = append(cells, []cell.Cell{})
		for j := range f.Cells[i] {
			cells[i] = append(cells[i], f.Cells[i][j].Clone())
		}
	}

	for i := range cells {
		for j := range cells[i] {
			cell, err := cells[i][j].Next()
			if err != nil {
				return nil, err
			}
			cells[i][j] = cell
		}
	}

	return newFrame(f.W, f.H, cells, f.ncurse), nil
}

func (f *frameImpl) Visualize() error {
	f.ncurse.Refresh()

	if err := f.ncurse.Clear(); err != nil {
		return err
	}
	for i := range f.Cells {
		for j := range f.Cells[i] {
			if f.Cells[i][j].IsLive() == true {
				f.ncurse.MovePrint(i, j, "#")
			} else {
				f.ncurse.MovePrint(i, j, ".")
			}
		}
	}

	return nil
}

func (f *frameImpl) Generation() (Frame, error) {
	cells := [][]cell.Cell{}
	// f.Cellsをdeep copyしている
	for i := range f.Cells {
		cells = append(cells, []cell.Cell{})
		for j := range f.Cells[i] {
			cells[i] = append(cells[i], f.Cells[i][j].Clone())
		}
	}

	for i := range cells {
		for j := range cells[i] {
			cell, err := cells[i][j].Generation()
			if err != nil {
				return nil, err
			}
			cells[i][j] = cell
		}
	}
	return newFrame(f.W, f.H, setAround(f.W, f.H, cells), f.ncurse), nil
}

func (f *frameImpl) Close() error {
	goncurses.End()

	return nil
}
