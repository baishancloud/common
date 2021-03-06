package model

import (
	"fmt"

	"github.com/open-falcon/common/utils"
)

type Expression struct {
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
	ActionId   int               `json:"actionId"`
}

func (this *Expression) String() string {
	return fmt.Sprintf(
		"<Id:%d, Metric:%s, Field:%s, Tags:%v, %s%s%s MaxStep:%d, P%d %s ActionId:%d>",
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
		this.ActionId,
	)
}

type ExpressionResponse struct {
	Expressions []*Expression `json:"expressions"`
}
