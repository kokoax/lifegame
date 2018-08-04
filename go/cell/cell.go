package cell

// Cell はlifegameの盤面の一マス
type Cell interface {
	SetAroundCell(cells [][]Cell) Cell
	Next() (Cell, error)
	WillLive() bool
	IsLive() bool
	AroundCell() [][]Cell
	CheckBirth() (bool, error)
	CheckLive() (bool, error)
	CheckDepopulation() (bool, error)
	CheckOvercrowding() (bool, error)
	Generation() (Cell, error)
	Clone() Cell
}

type cellImpl struct {
	aroundCell [][]Cell
	isLive     bool
	willLive   bool
}

func newCell(aroundCell [][]Cell, isLive bool, willLive bool) Cell {
	return &cellImpl{
		aroundCell: aroundCell,
		isLive:     isLive,
		willLive:   willLive,
	}
}

// NewCell is
func NewCell() Cell {
	return &cellImpl{
		aroundCell: nil,
		isLive:     false,
		willLive:   false,
	}
}

// NewLiveCell is
func NewLiveCell() Cell {
	return &cellImpl{
		aroundCell: nil,
		isLive:     true,
		willLive:   true,
	}
}

// NewIsLiveCell is
func NewIsLiveCell() Cell {
	return &cellImpl{
		aroundCell: nil,
		isLive:     true,
		willLive:   false,
	}
}

// NewDeadCell is
func NewDeadCell() Cell {
	return &cellImpl{
		aroundCell: nil,
		isLive:     false,
		willLive:   false,
	}
}

// NewIsDeadCell is
func NewIsDeadCell() Cell {
	return &cellImpl{
		aroundCell: nil,
		isLive:     false,
		willLive:   false,
	}
}

// NewWillLiveCell is
func NewWillLiveCell() Cell {
	return &cellImpl{
		aroundCell: nil,
		isLive:     false,
		willLive:   true,
	}
}

// NewWillDeadCell is
func NewWillDeadCell() Cell {
	return &cellImpl{
		aroundCell: nil,
		isLive:     false,
		willLive:   false,
	}
}

// NewCellWithAround is
func NewCellWithAround(cells [][]Cell) Cell {
	return &cellImpl{
		aroundCell: cells,
		isLive:     false,
		willLive:   false,
	}
}

func (c *cellImpl) SetAroundCell(cells [][]Cell) Cell {
	return newCell(cells, c.isLive, c.willLive)
}

func (c *cellImpl) IsLive() bool {
	return c.isLive
}

func (c *cellImpl) WillLive() bool {
	return c.willLive
}

func (c *cellImpl) AroundCell() [][]Cell {
	return c.aroundCell
}

func (c *cellImpl) CountLive() (int, error) {
	live := 0
	for i := range c.aroundCell {
		for j := range c.aroundCell[i] {
			if c.aroundCell[i][j] != nil && c.aroundCell[i][j].IsLive() {
				live++
			}
		}
	}
	return live, nil
}

// CheckBirth is return live: true, dead: false
func (c *cellImpl) CheckBirth() (bool, error) {
	live, err := c.CountLive()
	if err != nil {
		return false, err
	}
	if !c.IsLive() && live == 3 {
		return true, nil
	}
	return false, nil
}

// CheckLive is return live: true, dead: false
func (c *cellImpl) CheckLive() (bool, error) {
	live, err := c.CountLive()
	if err != nil {
		return false, err
	}
	if c.IsLive() && (live == 2 || live == 3) {
		return true, nil
	}
	return false, nil
}

// CheckDepopulation is return live: true, dead: false
func (c *cellImpl) CheckDepopulation() (bool, error) {
	live, err := c.CountLive()
	if err != nil {
		return false, err
	}
	if live <= 1 {
		return false, nil
	}
	return true, nil
}

// CheckOvercrowding is return live: true, dead: false
func (c *cellImpl) CheckOvercrowding() (bool, error) {
	live, err := c.CountLive()
	if err != nil {
		return false, err
	}
	if live >= 4 {
		return false, nil
	}
	return true, nil
}

func (c *cellImpl) Next() (Cell, error) {
	birth, err := c.CheckBirth()
	if err != nil {
		return nil, err
	}
	live, err := c.CheckLive()
	if err != nil {
		return nil, err
	}
	depopulation, err := c.CheckDepopulation()
	if err != nil {
		return nil, err
	}
	overcrowding, err := c.CheckOvercrowding()
	if err != nil {
		return nil, err
	}
	willLive := birth || live || (!depopulation && !overcrowding)

	return &cellImpl{
		aroundCell: c.aroundCell,
		isLive:     c.isLive,
		willLive:   willLive,
	}, nil
}

func (c *cellImpl) Generation() (Cell, error) {
	// cp := c.clone().(*cellImpl)
	c.isLive = c.willLive
	return c, nil
}

func (c *cellImpl) clone() Cell {
	copy := *c
	return &copy
}

func (c *cellImpl) Clone() Cell {
	clone := c.clone()
	return clone
}
