package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
	"withChat/internal/config"
)

func LoggerToFile() gin.HandlerFunc {
	logFileToPath := config.Config.Log.Log_Path
	logFileName := config.Config.Log.Log_Name
	fileName := path.Join(logFileToPath, logFileName)

	exists, err := PathExists(fileName)
	if err != nil { // 获取文件信息失败
		panic(err)
	}
	if !exists { // 如果文件不存在就创建
		os.Create(fileName)
	}

	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	// 实例化
	logger := logrus.New()
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"}) // 设置时间格式
	// 设置输出
	writers := []io.Writer{
		src,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logger.SetOutput(fileAndStdoutWriter)
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		//日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode, latencyTime, clientIP, reqMethod, reqUri)
	}

}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
