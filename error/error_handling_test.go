package error

import (
	"errors"
	"gorm.io/gorm"
	"gorm/config"
	"gorm/model"
	"log"
	"os"
	"testing"
)

// 错误检查
// 鼓励在调用任何 Finisher 方法后，都进行错误检查
// 官方错误类：errors.go

// TestErrorHandling 测试处理错误
func TestErrorHandling(t *testing.T) {
	db := config.DB

	if err := db.Where("email = 'baidu.com'").First(&model.User{}).Error; err != nil {
		// 处理错误
		log.Println("Error, 记录不存在", err)
	}

	if result := db.Where("email = 'baidu.com'").First(&model.User{}); result.Error != nil {
		// 处理错误
		log.Println("result, 记录不存在", result.Error)
	}

	//API server listening at: 127.0.0.1:61185
	//=== RUN   TestErrorHandling
	//
	//2023/04/19 07:21:50 D:/gorm/error/error_handling_test.go:18 record not found
	//[1.590ms] [rows:0] SELECT * FROM `users` WHERE email = 'baidu.com' ORDER BY `users`.`id` LIMIT 1
	//2023/04/19 07:21:51 Error, 记录不存在 record not found
	//
	//2023/04/19 07:21:51 D:/gorm/error/error_handling_test.go:23 record not found
	//[0.000ms] [rows:0] SELECT * FROM `users` WHERE email = 'baidu.com' ORDER BY `users`.`id` LIMIT 1
	//2023/04/19 07:22:14 result, 记录不存在 record not found
	//--- PASS: TestErrorHandling (43.05s)
	//PASS
}

// TestErrRecordNotFound
// 当 First、Last、Take 方法找不到记录时，GORM 会返回 ErrRecordNotFound 错误。
// 如果发生了多个错误，你可以通过 errors.Is 判断错误是否为 ErrRecordNotFound
func TestErrRecordNotFound(t *testing.T) {
	db := config.DB

	if err := db.Where("email = 'baidu.com'").First(&model.User{}).Error; err != nil {
		// 处理错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Error, 记录不存在, 错误类型：ErrRecordNotFound", err)
		}
	}

	//API server listening at: 127.0.0.1:61442
	//=== RUN   TestErrRecordNotFound
	//
	//2023/04/19 07:25:42 D:/gorm/error/error_handling_test.go:50 record not found
	//[2.120ms] [rows:0] SELECT * FROM `users` WHERE email = 'baidu.com' ORDER BY `users`.`id` LIMIT 1
	//2023/04/19 07:26:04 Error, 记录不存在, 错误类型：ErrRecordNotFound record not found
	//--- PASS: TestErrRecordNotFound (27.20s)
	//PASS
}

func TestMain(m *testing.M) {
	config.InitMysql()

	code := m.Run()

	// 退出
	os.Exit(code)
}
