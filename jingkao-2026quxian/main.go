package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"dataclearing/conf"
	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type RegistrationData struct {
	jobCode string  `json:"-"`
	Title   string  `json:"title"`
	List    []Entry `json:"list"`
}

type Entry struct {
	Time          int64  `json:"time"`
	Title         string `json:"title"`
	Baomingrenshu int64  `json:"baomingrenshu"`
	Guoshenrenshu int64  `json:"guoshenrenshu"`
	Daishenrenshu int64  `json:"daishenrenshu"`
}

func NewRegistrationData(row []string) RegistrationData {
	if len(row) <= 1 {
		panic("xxx")
	}
	var rdata = RegistrationData{
		jobCode: row[0],
		Title:   "报名数据",
		List:    []Entry{},
	}
	for i, _ := range row {
		if i == 0 {
			continue
		}
		if i%3 == 1 {

			//t1, _ := strconv.ParseInt(Ddate[i/3], 10, 64)
			row1, _ := strconv.ParseInt(row[i], 10, 64)
			row2, _ := strconv.ParseInt(row[i+1], 10, 64)
			row3, _ := strconv.ParseInt(row[i+2], 10, 64)

			var t1 int64
			title := Ddate[i/3]
			if title == "20251117b" {
				title = "11.17-18:00"
				t1 = 202511172
			}
			if title == "20251118a" {
				title = "11.18-09:00"
				t1 = 202511181
			}
			if title == "20251118b" {
				title = "11.18-18:00"
				t1 = 202511182
			}
			if title == "20251119a" {
				title = "11.19-09:00"
				t1 = 202511191
			}
			if title == "20251119b" {
				title = "11.19-18:00"
				t1 = 202511192
			}
			if title == "20251120a" {
				title = "11.20-09:00"
				t1 = 202511201
			}
			if title == "20251120b" {
				title = "11.20-18:00"
				t1 = 202511202
			}
			if title == "20251121a" {
				title = "11.21-09:00"
				t1 = 202511211
			}
			if title == "20251121b" {
				title = "11.21-18:00"
				t1 = 202511212
			}
			e := Entry{
				Time:          t1,
				Title:         title,
				Baomingrenshu: row1,
				Guoshenrenshu: row2,
				Daishenrenshu: row3,
			}
			rdata.List = append(rdata.List, e)
		} else {
			continue
		}
	}

	return rdata
}

var Ddate []string

func main() {
	db, err := gorm.Open(mysql.Open(conf.DSNL), &gorm.Config{Logger: logger.Discard})
	if err != nil {
		panic("failed to connect database")
	}
	var year = "2026"
	type2 := "最终清洗结果" //国考更改（sheet名）
	f, err := excelize.OpenFile("./jingkao-2026quxian/2026quxian.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := f.GetRows(type2)

	for i, row := range rows {
		if i == 0 {
			func1(row)
		}

		if i <= 1 {
			continue
		}
		dd := NewRegistrationData(row)
		fmt.Println(dd)

		v, errJ := json.Marshal(dd)
		if errJ != nil {
			fmt.Println("出错1个", errJ)
		}
		db.Table("tb_jingkao").Where("year = ? and job_code = ?", year, row[0]).
			Update("v3_submit_info", string(v))
	}
}

func func1(row []string) {
	for i, s := range row {
		if i == 0 {
			continue
		}
		if s != "" {
			Ddate = append(Ddate, s)
		}
	}
}
