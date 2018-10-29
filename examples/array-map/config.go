// This file is auto-generated by taxi v1.2.3, DO NOT EDIT!

package Config

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

const (
	KeyNewbieGuideDefineName = "newbieguidedefine"
)

// 新手引导配置
type NewbieGuideDefine struct {
	Name           string            // ID
	Type           string            // 任务类型
	Target         string            // 目标
	Accomplishment []int             // 完成步骤
	Goods          map[string]uint32 // 物品
	Description    string            // 描述
}

func (p *NewbieGuideDefine) ParseFromRow(row []string) error {
	if len(row) < 6 {
		log.Panicf("NewbieGuideDefine: row length out of index %d", len(row))
	}
	if row[0] != "" {
		p.Name = row[0]
	}
	if row[1] != "" {
		p.Type = row[1]
	}
	if row[2] != "" {
		p.Target = row[2]
	}
	if row[3] != "" {
		for _, item := range strings.Split(row[3], "\\") {
			var value = MustParseTextValue("int", item, row[3])
			p.Accomplishment = append(p.Accomplishment, value.(int))
		}
	}
	if row[4] != "" {
		p.Goods = map[string]uint32{}
		for _, text := range strings.Split(row[4], "=") {
			if text == "" {
				continue
			}
			var item = strings.Split(text, ";")
			var value = MustParseTextValue("string", item[0], row[4])
			var key = value.(string)
			value = MustParseTextValue("uint32", item[1], row[4])
			var val = value.(uint32)
			p.Goods[key] = val
		}
	}
	if row[5] != "" {
		p.Description = row[5]
	}
	return nil
}

func LoadNewbieGuideDefineList(loader DataSourceLoader) ([]*NewbieGuideDefine, error) {
	buf, err := loader.LoadDataByKey(KeyNewbieGuideDefineName)
	if err != nil {
		return nil, err
	}
	var list []*NewbieGuideDefine
	var r = csv.NewReader(buf)
	for i := 0; ; i++ {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("NewbieGuideDefine: read csv %v", err)
			return nil, err
		}
		var item NewbieGuideDefine
		if err := item.ParseFromRow(row); err != nil {
			log.Errorf("NewbieGuideDefine: parse row %d, %s, %v", i+1, row, err)
			return nil, err
		}
		list = append(list, &item)
	}
	return list, nil
}

//
func MustParseTextValue(typename, valueText string, msgtips interface{}) interface{} {
	switch typename {
	case "bool":
		b, err := strconv.ParseBool(valueText)
		if err != nil {
			log.Panicf("%s %s, %v, %v", typename, valueText, err, msgtips)
		}
		return b

	case "float32", "float64":
		f, err := strconv.ParseFloat(valueText, 64)
		if err != nil {
			log.Panicf("%s %s, %v, %v", typename, valueText, err, msgtips)
		}
		if typename == "float32" {
			return float32(f)
		}
		return f // float64

	case "uint", "uint8", "uint16", "uint32", "uint64":
		n, err := strconv.ParseUint(valueText, 10, 64)
		if err != nil {
			log.Panicf("%s %s, %v, %v", typename, valueText, err, msgtips)
		}
		if typename == "uint" {
			return uint(n)
		} else if typename == "uint8" {
			return uint8(n)
		} else if typename == "uint16" {
			return uint16(n)
		} else if typename == "uint32" {
			return uint32(n)
		}
		return n // uint64

	case "int", "int8", "int16", "int32", "int64":
		n, err := strconv.ParseInt(valueText, 10, 64)
		if err != nil {
			log.Panicf("%s %s, %v, %v", typename, valueText, err, msgtips)
		}
		if typename == "int" {
			return int(n)
		} else if typename == "int8" {
			return int8(n)
		} else if typename == "int16" {
			return int16(n)
		} else if typename == "int32" {
			return int32(n)
		}
		return n // int64

	default:
		return valueText
	}
}
