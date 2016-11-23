package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"

	"gopkg.in/yaml.v2"
)

type Item struct {
	Input  string
	Sheet  int
	PkCols []string
	Cols   []string
	Output string
}

type Config struct {
	Template string
	Sheet    int
	PkCols   []string
	PkSep    string
	Ext      string
	List     []Item
}

// 字母表示的列转为列序号, 例如A:1
func toIntColumn(name string) (int, error) {
	col := 0
	isMinus := false
	if name[0] == '-' {
		isMinus = true
		name = name[1:]
	}
	n := len(name)
	base := int('Z'-'A') + 1
	for i, v := range name {
		if v < 'A' || v > 'Z' {
			return 0, fmt.Errorf("无效列名:%s", name)
		}
		col += ((int(v-'A') + 1) * int(math.Pow(float64(base), float64(n-i-1))))
	}

	if isMinus {
		col = -col
	}
	return col, nil
}

func toIntColumns(names []string) ([]int, error) {
	ret := []int{}
	for _, name := range names {
		v, err := toIntColumn(name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, v)
	}
	return ret, nil
}

// idx: 配置项索引
func (cfg *Config) GetPkCols(idx int) []int {
	if len(cfg.List[idx].PkCols) > 0 {
		cols, err := toIntColumns(cfg.List[idx].PkCols)
		if err != nil {
			panic(fmt.Errorf("主键列无效:%s %s", cfg.List[idx].Input, err))
		}
		return cols
	}
	cols, err := toIntColumns(cfg.PkCols)
	if err != nil {
		panic(fmt.Errorf("主键列无效: %s", err))
	}
	return cols
}

// idx: 配置项索引
func (cfg *Config) GetSheet(idx int) int {
	if cfg.List[idx].Sheet > 0 {
		return cfg.List[idx].Sheet
	}
	return cfg.Sheet
}

// idx: 配置项索引
// nil表示所有列(不可能处理没有数据的表)
func (cfg *Config) GetCols(idx int) []int {
	cols, err := toIntColumns(cfg.List[idx].Cols)
	if err != nil {
		panic(fmt.Errorf("包含列无效: %s, %s", cfg.List[idx].Input, err))
	}

	if len(cols) == 0 {
		return nil
	}

	// 校验
	if cols[0] >= 0 {
		// 包含
		for _, ci := range cols {
			if ci <= 0 {
				panic(fmt.Errorf("包含冲突: %s %v %d", cfg.Template, cols, ci))
			}
		}
	} else {
		// 不包含
		for _, ci := range cols {
			if ci >= 0 {
				panic(fmt.Errorf("包含冲突: %s %v %d", cfg.Template, cols, ci))
			}
		}
	}
	return cols
}

func readConfig(fp string) *Config {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
