package share

import (
	"time"

	"github.com/sony/sonyflake"
)

type IdGenerator struct {
	flake *sonyflake.Sonyflake
}

func NewIdGenerator(baseTime string) *IdGenerator {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	dt, err := time.ParseInLocation("2006-01-02", baseTime, loc)
	if err != nil {
		panic(err)
	}
	settings := sonyflake.Settings{
		StartTime:      dt,
		MachineID:      nil,
		CheckMachineID: nil,
	}
	flake := sonyflake.NewSonyflake(settings)
	return &IdGenerator{flake: flake}
}

func (this *IdGenerator) GetId() int64 {
	nextID, _ := this.flake.NextID()
	return int64(nextID)
}
