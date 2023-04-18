package callback

import (
	"gorm.io/gorm"
	"log"
)

// CallbackContext
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
	log.Println("callback_context get value：", value)
}
