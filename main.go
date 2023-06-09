/**
 * @Author: lyc
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/03/20 14:09
 */

package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go-chat/internal/global"
	"go-chat/internal/model/common"
	"go-chat/internal/routing/router"
	"go-chat/internal/setting"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
func init() {
	log.SetFlags(log.Ltime | log.Llongfile)
}

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

// @title        chat
// @version      1.0
// @description  在线聊天系统

// @license.name  lyc,why
// @license.url

// @host      127.0.0.1:8084
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

func main() {
	setting.InitAll()

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", common.ValidateMobile) //这里的mobile和from表单里的是一样的
		_ = v.RegisterValidation("email", common.ValidateEmail)   //这里的email和from表单里的是一样的
		//翻译错误
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "非法的手机号码", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
		_ = v.RegisterTranslation("email", global.Trans, func(ut ut.Translator) error {
			return ut.Add("email", "非法的邮箱", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("email", fe.Field())
			return t
		})
	}
	r, ws := router.NewRouter() // 注册路由
	s := &http.Server{
		Addr:           global.Settings.Serve.Addr,
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Info("Server started!")
	global.Logger.Info(global.Settings.Rocketmq.Addr)
	//global.Logger.Info( global.Settings.Rocketmq.Port)
	errChan := make(chan error, 1)
	defer close(errChan)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		defer ws.Close()
		if err := ws.Serve(); err != nil {
			errChan <- err
		}
	}()
	gracefulExit(s, errChan)
}

// 优雅退出
func gracefulExit(s *http.Server, errChan chan error) {
	// 优雅退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errChan:
		global.Logger.Error(err.Error())
	case <-quit:
		global.Logger.Info("ShutDown Server...")
		// 给几秒完成剩余任务
		ctx, cancel := context.WithTimeout(context.Background(), global.Settings.Serve.DefaultTimeout)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil { // 优雅退出
			if !errors.Is(err, context.DeadlineExceeded) {
				global.Logger.Error("Server forced to ShutDown,Err:" + err.Error())
			}
		}
	}
}
