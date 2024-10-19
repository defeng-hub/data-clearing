package main

import (
	"context"
	"dataclearing/gkld-guokao-2025/model"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var client *mongo.Client

func init() {
	var err error
	// 设置 MongoDB 连接选项
	clientOptions := options.Client().ApplyURI("mongodb://bythjy:bythjy-mongo-123@43.248.100.60:27017/")

	// 连接到 MongoDB
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

}

func Run() {

	dirPath := "/Users/Apple/Desktop/byth/ai/抓包v2/10-19/"
	// 读取文件夹下的所有文件
	files, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			filePath := filepath.Join(dirPath, file.Name())

			// 读取文件内容
			content, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", filePath, err)
				continue
			}

			// 解析JSON到结构体
			var yourStruct model.Response
			err = json.Unmarshal(content, &yourStruct)
			if err != nil {
				fmt.Printf("Error parsing JSON in file %s: %v\n", filePath, err)
				continue
			}

			v1 := yourStruct.Data.JobDataList.Get("部门代码")
			v2 := yourStruct.Data.JobDataList.Get("职位代码")
			if v1 != "" {
				if v2 != "" {
					yourStruct.Data.JobCodeBase = v1
					yourStruct.Data.JobCode = v2
				} else {
					continue
				}
			} else {
				continue
			}

			client.Database("gkld")
			collection := client.Database("gkld").Collection("2024-10-19")
			_, err = collection.InsertOne(context.Background(), yourStruct.Data)
			if err != nil {
				continue
			}
			fmt.Println("插入成功")
		}
	}
}
func main() {
	Run()
}
