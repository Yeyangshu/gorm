package context

import (
	"context"
	"fmt"
	"gorm/callback"
	"gorm/config"
	"gorm/model"
	"log"
	"os"
	"testing"
	"time"
)

// 单会话模式
func TestSingleSessionMode(t *testing.T) {
	db := config.DB
	ctx := context.WithValue(context.Background(), "context", "123")
	user := model.User{ID: 1}
	// SELECT * FROM `users` WHERE `users`.`id` = 1
	db.WithContext(ctx).Find(&user)
	log.Println(user)
}

// 持续会话模式
func TestContinuousSessionMode(t *testing.T) {
	db := config.DB
	ctx := context.WithValue(context.Background(), "context", "123")
	user := model.User{}
	tx := db.WithContext(ctx)
	// SELECT * FROM `users` WHERE `users`.`id` = 1
	tx.Find(&user, 1)
	// UPDATE `users` SET `email`='123456@163.com',`updated_at`='2023-04-18 22:21:26.601' WHERE `id` = 1
	tx.Model(&user).Update("email", "123456@163.com")
	log.Println(user)
}

// context超时，参考Chi 中间件示例：https://gorm.io/zh_CN/docs/context.html#Chi-%E4%B8%AD%E9%97%B4%E4%BB%B6%E7%A4%BA%E4%BE%8B
func TestContextTimeout(t *testing.T) {
	db := config.DB
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// SELECT * FROM `users` WHERE `users`.`id` = 1
	user := model.User{ID: 1}
	db.WithContext(ctx).Find(&user)
	log.Println(user)

	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			fmt.Println("waiting")
			time.Sleep(time.Second)
		}
	}

	//2023/04/18 22:31:20 D:/gorm/context/context_test.go:44
	//[28.427ms] [rows:1] SELECT * FROM `users` WHERE `users`.`id` = 1
	//2023/04/18 22:31:20 {1 1 0xc0001caed0 0 <nil> { false} {0001-01-01 00:00:00 +0000 UTC false} 0001-01-01 00:00:00 +0000 UTC 2023-04-18 22:21:26 +0800 CST}
	//waiting
	//timeout
	//--- PASS: TestContextTimeout (1.04s)
	//PASS
}

// Hooks/Callbacks 中的 Context
func TestContextInHookOrCallback(t *testing.T) {
	db := config.DB
	db.Use(&callback.CallbackContext{})

	// gorm设置值，可传递到钩子
	db = db.Set("setKey", true)

	ctx := context.WithValue(context.Background(), "context", "123")
	user := model.User{ID: 1}
	// SELECT * FROM `users` WHERE `users`.`id` = 1
	db.WithContext(ctx).Find(&user)

	//=== RUN   TestContextInHookOrCallback
	//2023/04/19 07:03:54 hook_context get value: 123

	//2023/04/18 22:44:24 callback_context get value： 123
	//
	//2023/04/18 22:44:24 D:/gorm/context/context_test.go:78
	//[45.910ms] [rows:1] SELECT * FROM `users` WHERE `users`.`id` = 1
	//--- PASS: TestContextInHookOrCallback (0.05s)
	//PASS
	//
	//Process finished with the exit code 0
}

// 如果测试文件中包含该函数，那么生成的测试将调用TestMain(m)，而不是直接运行测试。
// TestMain 运行在主 goroutine 中 , 可以在调用 m.Run 前后做任何设置和拆卸。
// https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter09/09.5.html
func TestMain(m *testing.M) {
	config.InitMysql()

	code := m.Run()

	// 退出
	os.Exit(code)
}
