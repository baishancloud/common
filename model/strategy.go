package model

import (
	"fmt"

	"github.com/open-falcon/common/utils"
)

//策略结构
//TODO 先添加单一判断的field。下一步变好为列表，判断多个field。
type Strategy struct {
	Id         int               `json:"id"`
	Metric     string            `json:"metric"`
	Field      string            `json:"field"` // default is "" , eq value.
	Tags       map[string]string `json:"tags"`
	Func       string            `json:"func"`       // e.g. max(#3) all(#3)
	Operator   string            `json:"operator"`   // e.g. < !=
	RightValue float64           `json:"rightValue"` // critical value
	MaxStep    int               `json:"maxStep"`
	Priority   int               `json:"priority"`
	Note       string            `json:"note"`
	Tpl        *Template         `json:"tpl"`
}

func (this *Strategy) String() string {
	return fmt.Sprintf(
		"<Id:%d, Metric:%s, Field:%s Tags:%v, %s%s%s MaxStep:%d, P%d, %s, %v>",
		this.Id,
		this.Metric,
		this.Field,
		this.Tags,
		this.Func,
		this.Operator,
		utils.ReadableFloat(this.RightValue),
		this.MaxStep,
		this.Priority,
		this.Note,
		this.Tpl,
	)
}

type HostStrategy struct {
	Hostname   string     `json:"hostname"`
	Strategies []Strategy `json:"strategies"`
}

type StrategiesResponse struct {
	HostStrategies []*HostStrategy `json:"hostStrategies"`
}
