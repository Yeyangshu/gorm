package callback

import (
	"gorm.io/gorm"
	"log"
)

// CallbackContext callback上下文测试
type CallbackContext struct {
}

func (e *CallbackContext) Name() string {
	return "callbackContext"
}

func (e *CallbackContext) Initialize(db *gorm.DB) error {
	return db.Callback().Query().After("gorm:query").Register("callback_context", getContext)
}

func getContext(db *gorm.DB) {
	ctx := db.Statement.Context
	value := ctx.Value("context")
	log.Printf("callback_context get context value: %s", value)

	// 获取gorm设置值
	setKey, ok := db.Get("setKey")
	if ok {
		log.Printf("callback_context get set value: %t", setKey)
	}
}
