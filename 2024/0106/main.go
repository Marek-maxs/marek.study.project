package main

/**
*
* Author: Marek
* Date: 2024-01-06 23:20
* Email: 364021318@qq.com
*
* 手势控制的脚本
 */

func main() {
	//runtime.LockOSThread()
	//defer runtime.UnlockOSThread()
	//
	//// 创建视频捕获对象
	//videoGapture, err := gocv.VideoCaptureDevice(0)
	//if err != nil {
	//	log.Error().Err(err).Msg("打开摄像头设备失败")
	//	return
	//}
	////videoGapture := opencv.NewCameraCapture(0)
	//if !videoGapture.IsOpened() {
	//	log.Error().Msg("无法打开摄像头")
	//	os.Exit(1)
	//}
	//
	//// 创建窗口用于显示视频流
	//window := opencv.NewWindow("Gesture Control")
	//defer window.Destroy()
	//
	//// 创建手势识别器对象
	//gestureDetector := gocv.NewSimpleBlobDetector()
	//
	//for {
	//	m := gocv.NewMat()
	//	// 读取一个帧视频图像
	//	ret := videoGapture.Read(&m)
	//	if !ret {
	//		log.Error().Err(err).Msg("无法读取视频帧")
	//		os.Exit(1)
	//	}
	//
	//	// 将图像转换为灰度图像以进行手势识别
	//	gray := gocv.NewMat()
	//	gocv.CvtColor(m, &gray, gocv.ColorBGRToGray)
	//}
}