package models

type DisplayMode string

const (
	LeftRight DisplayMode = "LR"
	TopBottom DisplayMode = "TB"
)

func (m DisplayMode) ToString() string {
	if m == LeftRight {
		return string(LeftRight)
	}
	return string(TopBottom)
}
