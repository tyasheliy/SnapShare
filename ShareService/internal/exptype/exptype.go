package exptype

import "time"

type ExpireType string

const (
	Snap   ExpireType = "Snap"
	Short  ExpireType = "Short"
	Medium ExpireType = "Medium"
	Long   ExpireType = "Long"
)

func (t ExpireType) GetDuration() time.Duration {
	switch t {
	case Snap:
		return time.Minute * 2
	case Short:
		return time.Minute * 5
	case Medium:
		return time.Minute * 10
	case Long:
		return time.Minute * 15
	default:
		return time.Minute
	}
}
