package frame

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kokoax/lifegame/go/cell"
)

func TestNewFrame(t *testing.T) {
	w, h := 1, 1

	_, err := NewFrame(w, h)
	assert.NoError(t, err)
}

func TestInitCell(t *testing.T) {
	w, h := 1, 1
	result := initCell(w, h)
	var expect [][]cell.Cell

	expect = append(expect, []cell.Cell{})
	expect[0] = append(expect[0], cell.NewCell())

	assert.Equal(t, result, expect)
}

func TestSetAround(t *testing.T) {
	w, h := 2, 2

	setAround(w, h, initCell(w, h))
	// assert.NoError(t, err)
}

func TestNext(t *testing.T) {
	w, h := 3, 3
	cells := initCell(w, h)
	frame, _ := NewFrameWithCells(w, h, cells)
	_, err := frame.Next()

	assert.NoError(t, err)
}

func TestGeneration(t *testing.T) {
	w, h := 3, 3
	cells := initCell(w, h)
	frame, _ := NewFrameWithCells(w, h, cells)
	next, _ := frame.Next()
	_, err := next.Generation()

	assert.NoError(t, err)
}

func TestVisualize(t *testing.T) {
	w, h := 1, 1
	frame, _ := NewFrame(w, h)
	err := frame.Visualize()

	assert.NoError(t, err)
}
