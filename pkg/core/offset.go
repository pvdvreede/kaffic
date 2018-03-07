package core

import "fmt"

type Offset int64

func (this Offset) String() string {
	return fmt.Sprintf("%030d", int64(this))
}

func (this Offset) ToInt64() int64 {
	return int64(this)
}

func (this Offset) Next() Offset {
	return Offset(int64(this) + 1)
}

func InitialOffset() Offset {
	return Offset(1)
}

func FromOffset(frm int64) Offset {
	return Offset(frm)
}
