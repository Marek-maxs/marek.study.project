package main

import (
	"fmt"
	"github.com/Marek-maxs/govalidator"
)

// 这里可以适当调整

const (
	bucket = "sbos-dev-graphql"

	key = "company_logo/Dulwich+College+Singapore_78c44730-77d.jpg"

	region = "ap-east-1"
)

func main() {
	fmt.Println(int(-100))
	fmt.Println(govalidator.IsURL("wfeoefqwf"))
	//creds := credentials.NewEnvCredentials() // 确保环境变量AWS_ACCESS_KEY_ID和AWS_SECRET_ACCESS_KEY存在
	//
	//sess := session.Must(session.NewSessionWithOptions(
	//
	//	session.Options{
	//
	//		Config: aws.Config{
	//
	//			Credentials: creds,
	//
	//			Region: aws.String(region),
	//		},
	//	},
	//))
	//
	//sClient := s3.New(sess)
	//
	//inputConf := &s3.GetObjectInput{
	//
	//	Bucket: aws.String(bucket),
	//
	//	Key: aws.String(key),
	//
	//	ResponseContentType: aws.String(mime.TypeByExtension(filepath.Ext(key))),
	//}
	//
	//presignReq, _ := sClient.GetObjectRequest(inputConf)
	//
	//u, _, err := presignReq.PresignRequest(10 * time.Minute) // 时间可以随意 但不建议太长
	//savepic(u, key)
	//fmt.Println(u, err)

}

// 保存图片
//func savepic(url string, name string) {
//	resp, err := http.Get(url)
//	defer resp.Body.Close()
//	if err != nil {
//		fmt.Println(err)
//	}
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println(err)
//	}
//	err = ioutil.WriteFile("pic/"+name, body, 0755)
//	fmt.Println("erro:", err)
//}

//func GetImageNameFromUrl(url string) {
//	arr := strings.Split(url, "/")
//	return arr[len(arr)-1]
//}
