// This file is auto-generated by Tabular v0.5.0, DO NOT EDIT!
package config

import (
	"log"
	"strings"
)

var (
	_ = strings.Split
	_ = log.Panicf
)

// 新手引导配置, 新手任务.xlsx
type NewbieGuideDefine struct {
	Name           string            `json:"name"`           // ID
	Type           string            `json:"type"`           // 任务类型
	Target         string            `json:"target"`         // 目标
	Accomplishment []int16           `json:"accomplishment"` // 完成步骤
	Goods          map[string]uint32 `json:"goods"`          // 物品
	Description    string            `json:"description"`    // 描述
}
