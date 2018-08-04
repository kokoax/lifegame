package cell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCell(t *testing.T) {
	result := NewCell()
	expect := &cellImpl{
		aroundCell: nil,
		isLive:     false,
		willLive:   false,
	}

	assert.Equal(t, result, expect)
}

func TestNewCellWithAround(t *testing.T) {
	result := NewCellWithAround(nil)
	expect := &cellImpl{
		aroundCell: nil,
		isLive:     false,
		willLive:   false,
	}

	assert.Equal(t, result, expect)
}

func TestSetAround(t *testing.T) {
	result := NewCell().SetAroundCell(nil)
	expect := NewCellWithAround(nil)

	assert.Equal(t, result, expect)
}

func TestNextBirth(t *testing.T) {
	liveCell := NewIsLiveCell()
	deadCell := NewIsDeadCell()
	around := [][]Cell{
		[]Cell{liveCell, liveCell, deadCell},
		[]Cell{liveCell, nil, deadCell},
		[]Cell{deadCell, deadCell, deadCell},
	}
	result, err := NewCellWithAround(around).Next()
	expect := true

	assert.NoError(t, err)
	assert.Equal(t, result.WillLive(), expect)
}

func TestNextLive(t *testing.T) {
	liveCell := NewIsLiveCell()
	deadCell := NewIsDeadCell()
	around := [][]Cell{
		[]Cell{deadCell, deadCell, deadCell},
		[]Cell{deadCell, nil, liveCell},
		[]Cell{deadCell, liveCell, liveCell},
	}
	result, err := (&cellImpl{
		isLive:     true,
		willLive:   false,
		aroundCell: around,
	}).Next()
	expect := true

	assert.NoError(t, err)
	assert.Equal(t, result.WillLive(), expect)
}

func TestNextDepopulation(t *testing.T) {
	liveCell := NewIsLiveCell()
	deadCell := NewIsDeadCell()
	around := [][]Cell{
		[]Cell{deadCell, deadCell, deadCell},
		[]Cell{deadCell, nil, liveCell},
		[]Cell{deadCell, deadCell, deadCell},
	}
	result, err := (&cellImpl{
		isLive:     true,
		willLive:   false,
		aroundCell: around,
	}).Next()

	expect := false

	assert.NoError(t, err)
	assert.Equal(t, result.WillLive(), expect)
}

func TestNextOvercrowding(t *testing.T) {
	liveCell := NewIsLiveCell()
	deadCell := NewIsDeadCell()
	around := [][]Cell{
		[]Cell{liveCell, liveCell, liveCell},
		[]Cell{liveCell, nil, deadCell},
		[]Cell{deadCell, deadCell, deadCell},
	}
	result, err := (&cellImpl{
		isLive:     true,
		willLive:   false,
		aroundCell: around,
	}).Next()

	expect := false

	assert.NoError(t, err)
	assert.Equal(t, result.WillLive(), expect)
}

func TestGeneration(t *testing.T) {
	liveCell := NewWillLiveCell()
	deadCell := NewWillDeadCell()
	liveGene, err := liveCell.Generation()
	assert.NoError(t, err)

	deadGene, err := deadCell.Generation()
	assert.NoError(t, err)

	assert.Equal(t, liveGene.IsLive(), true)
	assert.Equal(t, deadGene.IsLive(), false)
}
