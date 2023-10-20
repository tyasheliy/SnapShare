package exptype

import (
	"LinkService/internal/models/dto"
	"time"
)

type ExpireType string

const (
	Snap   ExpireType = "Snap"
	Short  ExpireType = "Short"
	Medium ExpireType = "Medium"
	Long   ExpireType = "Long"
)

func (t ExpireType) getDTO() dto.ExpireTypeDTO {
	return dto.ExpireTypeDTO{
		Name:     string(t),
		Duration: int(t.GetDuration().Minutes()),
	}
}

func GetTypes() []dto.ExpireTypeDTO {
	types := make([]dto.ExpireTypeDTO, 0, 4)

	types = append(types, Snap.getDTO())
	types = append(types, Short.getDTO())
	types = append(types, Medium.getDTO())
	types = append(types, Long.getDTO())

	return types
}

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
