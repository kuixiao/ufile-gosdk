package main

import (
	ufsdk "github.com/kuixiao/ufile-gosdk"
	"github.com/kuixiao/ufile-gosdk/example/helper"
	"log"
	"os"
)

const (
	ConfigFile = "./config.json"
	FilePath = "IOMutipartAsyncUpload.txt"
	KeyName = "IOMutipartAsyncUpload.txt"
	MimeType = ""
)

func main() {
	if _, err := os.Stat(FilePath); os.IsNotExist(err) {
		helper.GenerateFakefile(FilePath, helper.FakeBigFileSize)
	}

	// 加载配置，创建请求
	config, err := ufsdk.LoadConfig(ConfigFile)
	if err != nil {
		panic(err.Error())
	}
	req, err := ufsdk.NewFileRequest(config, nil)
	if err != nil {
		panic(err.Error())
	}

	// 异步分片上传本地文件
	f, err := os.Open(FilePath)
	if err != nil {
		panic(err.Error())
	}
	err = req.IOMutipartAsyncUpload(f, KeyName, MimeType)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	log.Println("文件上传成功!!")

	err = req.HeadFile(KeyName)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	log.Printf(" %s", req.LastResponseHeader)
}
