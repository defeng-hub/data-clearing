package main

import (
	"dataclearing/conf"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
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
			if title == "20241118a" {
				title = "11.18-9:00"
				t1 = 202411181
			}
			if title == "20241118b" {
				title = "11.18-18:00"
				t1 = 202411182
			}
			if title == "20241119a" {
				title = "11.19-9:00"
				t1 = 202411191
			}
			if title == "20241119b" {
				title = "11.19-18:00"
				t1 = 202411192
			}
			if title == "20241120a" {
				title = "11.20-9:00"
				t1 = 202411201
			}
			if title == "20241120b" {
				title = "11.20-18:00"
				t1 = 202411202
			}
			if title == "20241121a" {
				title = "11.21-9:00"
				t1 = 202411211
			}
			if title == "20241121b" {
				title = "11.21-18:00"
				t1 = 202411212
			}
			if title == "20241122a" {
				title = "11.22-9:00"
				t1 = 202411221
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
	var year = "2025"
	type2 := "最终清洗结果" //国考更改（sheet名）
	f, err := excelize.OpenFile("./jingkao-2025quxian/2025quxian.xlsx")
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
