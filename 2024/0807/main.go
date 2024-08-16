package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	// 创建一个 AWS session
	sess, err := session.NewSession(&aws.Config{Region: aws.String("ap-east-1")})
	if err != nil {
		fmt.Println("创建 session 失败:", err)
		return
	}

	// 创建上传器
	uploader := s3manager.NewUploader(sess)

	// 打开要上传的文件
	file, err := os.Open("D:\\Tools\\data.csv")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}

	folder := "meters_data" // bucket 中的文件夹
	fileName := "test-meter-data"
	date := time.Now().Format(time.DateTime)
	ext := "csv"
	key := fmt.Sprintf("%s/%s_%s.%s", folder, fileName, date, ext)

	// 上传文件
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("influxdb-meters-data"),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		fmt.Println("上传文件失败:", err)
		return
	}

	fmt.Println("文件上传成功")
}

func goRunPythonCode() {
	newExpress := "print('{}'.format(eval('%s')))"
	pycodeGo := fmt.Sprintf(newExpress, "10+10")
	cmd := exec.Command("python", "-c", pycodeGo) // 要执行的 Python 代码
	output, err := cmd.Output()                   // 获取输出结果
	if err != nil {
		fmt.Println(err)
		return
	}
	result := string(output)
	fmt.Println(result)
}
