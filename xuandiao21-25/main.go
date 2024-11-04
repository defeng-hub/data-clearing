package main

import (
	"dataclearing/conf"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type TbDingxiangxuandiao struct {
	ID                   int64      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                        // 主键编码
	ZhiweiDaima          string     `gorm:"size:20;column:zhiwei_daima" json:"zhiwei_daima"`                     // 职位代码
	DanweiMingcheng      string     `gorm:"size:255;column:danwei_mingcheng" json:"danwei_mingcheng"`            // 单位名称
	ZhiweiMingcheng      string     `gorm:"size:255;column:zhiwei_mingcheng" json:"zhiwei_mingcheng"`            // 职位名称
	ZhiweiJianjie        string     `gorm:"size:1023;column:zhiwei_jianjie" json:"zhiwei_jianjie"`               // 职位简介
	ZhaokaoRenshu        string     `gorm:"size:20;column:zhaokao_renshu" json:"zhaokao_renshu"`                 // 招考人数
	MianshiBili          string     `gorm:"size:20;column:mianshi_bili" json:"mianshi_bili"`                     // 面试比例
	XueliYaoqiu          string     `gorm:"size:50;column:xueli_yaoqiu" json:"xueli_yaoqiu"`                     // 学历要求
	XueweiYaoqiu         string     `gorm:"size:50;column:xuewei_yaoqiu" json:"xuewei_yaoqiu"`                   // 学位要求
	ZhuanyeYaoqiu        string     `gorm:"size:1023;column:zhuanye_yaoqiu" json:"zhuanye_yaoqiu"`               // 专业要求
	ZhengzhiMianmao      string     `gorm:"size:50;column:zhengzhi_mianmao" json:"zhengzhi_mianmao"`             // 政治面貌
	QitaTiaojian         string     `gorm:"size:1023;column:qita_tiaojian" json:"qita_tiaojian"`                 // 其它条件
	KaoshengZixunDianhua string     `gorm:"size:20;column:kaosheng_zixun_dianhua" json:"kaosheng_zixun_dianhua"` // 考生咨询电话
	Beizhu               string     `gorm:"size:1023;column:beizhu" json:"beizhu"`                               // 备注
	Year                 string     `gorm:"size:10;column:year" json:"year"`                                     // 年份
	CreatedAt            time.Time  `gorm:"column:created_at" json:"created_at"`                                 // 创建时间
	UpdatedAt            time.Time  `gorm:"column:updated_at" json:"updated_at"`                                 // 最后更新时间
	DeletedAt            *time.Time `gorm:"column:deleted_at" json:"deleted_at"`                                 // 删除时间
	CreateBy             int64      `gorm:"column:create_by" json:"create_by"`                                   // 创建者
	UpdateBy             int64      `gorm:"column:update_by" json:"update_by"`                                   // 更新者
}

// TableName specifies the table name for GORM
func (TbDingxiangxuandiao) TableName() string {
	return "tb_dingxiangxuandiao"
}
func dataclean() {
	db, err := gorm.Open(mysql.Open(conf.DSNL), &gorm.Config{Logger: logger.Discard})
	if err != nil {
		panic("failed to connect database")
	}

	var year = "2025"
	f, err := excelize.OpenFile("./xuandiao21-25/dxxd" + year + "-最低分.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := f.GetRows("定向选调")
	var datas []TbDingxiangxuandiao
	for i, row := range rows {
		if i <= 2 {
			continue
		}
		fmt.Println(i, row)

		// 假设每行的数据顺序与表结构一致
		data := TbDingxiangxuandiao{
			Year:                 year,
			ZhiweiDaima:          row[1],
			DanweiMingcheng:      row[2],
			ZhiweiMingcheng:      row[3],
			ZhiweiJianjie:        row[4],
			ZhaokaoRenshu:        row[5],
			MianshiBili:          row[6],
			XueliYaoqiu:          row[7],
			XueweiYaoqiu:         row[8],
			ZhuanyeYaoqiu:        row[9],
			ZhengzhiMianmao:      row[10],
			QitaTiaojian:         row[11],
			KaoshengZixunDianhua: row[12],
			Beizhu:               row[13],
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		}
		datas = append(datas, data)
	}

	fmt.Println(datas)

	err = db.CreateInBatches(datas, 100).Error
	if err != nil {
		panic(err)
	}
}

func zuidifen() {
	db, err := gorm.Open(mysql.Open(conf.DSNL), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到标准输出
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 设定日志级别为 Info（也可以是 Debug）
				IgnoreRecordNotFoundError: true,
				Colorful:                  true, // 启用彩色打印
			},
		),
	})
	if err != nil {
		panic("failed to connect database")
	}

	var year = "2025"
	f, err := excelize.OpenFile("./xuandiao21-25/dxxd" + year + "-最低分.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := f.GetRows("定向选调")
	for i, row := range rows {
		if i <= 2 {
			continue
		}
		//fmt.Println(i, row)
		fmt.Println(row[2], row[1])
		db.Table("tb_dingxiangxuandiao").Where("year = ? && zhiwei_daima = ?", year, row[2]).Update("v2_min_score", row[1])
	}

}

func main() {
	zuidifen()
}
