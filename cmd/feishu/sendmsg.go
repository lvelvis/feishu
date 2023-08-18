package feishu

import (
	"encoding/json"
	"fmt"
	"os"
)

type TextTag struct {
	Tag  string `json:"tag"`
	Text string `json:"text"`
}

type ContentItem struct {
	Tags []TextTag `json:"tags"`
}

type ZhCn struct {
	Title   string         `json:"title"`
	Content [][]ContentItem `json:"content"`
}

type PostJson struct {
	ZhCn ZhCn `json:"zh_cn"`
}

func sendMsg() (string, error){
	title := os.Getenv("TITLE_ENV")              // 从环境变量中获取 title 的值
	jobName := os.Getenv("JOB_NAME_ENV")         // 从环境变量中获取 JOB_NAME 的值
	buildNumber := os.Getenv("BUILD_NUMBER_ENV") // 从环境变量中获取 BUILD_NUMBER 的值
	durationString := os.Getenv("DURATION_ENV")  // 从环境变量中获取 currentBuild.durationString 的值
	branchTag := os.Getenv("BRANCH_TAG_ENV")     // 从环境变量中获取 params.BRANCH_TAG 的值
	currentResult := os.Getenv("RESULT_ENV")     // 从环境变量中获取 currentBuild.currentResult 的值

	// 构建结构体
	textTags := []TextTag{
		{Tag: "text", Text: "发布应用名称：" + jobName},
		{Tag: "text", Text: "构建编号：" + buildNumber},
		{Tag: "text", Text: "持续时间：" + durationString},
		{Tag: "text", Text: "发布分支：" + branchTag},
		{Tag: "text", Text: "构建结果：" + currentResult},
	}

	contentItems := []ContentItem{
		{Tags: textTags},
	}

	content := [][]ContentItem{
		contentItems,
	}

	zhCn := ZhCn{
		Title:   title,
		Content: content,
	}

	post := PostJson{
		ZhCn: zhCn,
	}

	// 将结构体转换为 JSON 字符串
	jsonData, err := json.Marshal(post)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return "error", err
	}
	fmt.Println(jsonData)
	return string(jsonData), nil
}