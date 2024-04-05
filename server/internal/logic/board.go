package logic

type Field struct {
	Value     string
	crossedBy []string
}

type Board struct {
	Fields [][]Field
	Size   uint8
}

func NewBoard(size uint8) *Board {
	fields := make([][]Field, size)
	for i := range fields {
		fields[i] = make([]Field, size)
	}
	return &Board{
		Fields: fields,
		Size:   size,
	}
}
