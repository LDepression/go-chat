/**
 * @Author: lyc
 * @Description:
 * @File:  global
 * @Version: 1.0.0
 * @Date: 2023/03/20 14:39
 */

package global

import (
	ut "github.com/go-playground/universal-translator"
	"go-chat/internal/model/config"
	"go-chat/internal/pkg/goroutine/work"
	"go-chat/internal/pkg/logger"
	"go-chat/internal/pkg/snowflake"
	"go-chat/internal/pkg/token"
)

var (
	Settings  = &config.Settings{}
	RootDir   string
	Logger    *logger.Log
	Trans     ut.Translator
	Worker    *work.Worker
	Maker     token.Maker
	SnowFlake *snowflake.Worker
)
