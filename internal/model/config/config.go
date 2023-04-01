/**
 * @Author: lenovo
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2023/03/20 14:39
 */

package config

import "time"

type Settings struct {
	Serve    Serve    `json:"Serve" mapstructure:"Serve"`
	Mysql    *Mysql   `json:"Mysql" mapstructure:"Mysql"`
	Log      Log      `json:"Log" mapstructure:"Log"`
	SMTPInfo SMTPInfo `json:"SMTPInfo" mapstructure:"SMTPInfo"`
	Redis    Redis    `json:"Redis" mapstructure:"Redis"`
	Work     Work     `json:"Work" mapstructure:"Work"`
	Rule     Rule     `json:"Rule" mapstructure:"Rule"`
	Auto     Auto     `json:"Auto" mapstructure:"Auto"`
	Token    Token    `json:"Token" mapstructure:"Token"`
}

type Mysql struct {
	User     string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	DbName   string `json:"dbName" mapstructure:"dbName"`
}

type Serve struct {
	Addr           string        `json:"addr" mapstructure:"addr"`
	ReadTimeout    time.Duration `json:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout   time.Duration `json:"write_timeout" mapstructure:"write_timeout"`
	DefaultTimeout time.Duration `json:"default_timeout" mapstructure:"default_timeout"`
}

type Log struct {
	Level         string `json:"Level" mapstructure:"Level"`
	LogSavePath   string `json:"LogSavePath" mapstructure:"LogSavePath"`     // 保存路径
	LogFileExt    string `json:"LogFileExt" mapstructure:"LogFileExt"`       // 日志文件后缀
	MaxSize       int    `json:"MaxSize" mapstructure:"MaxSize"`             // 备份的大小(M)
	MaxBackups    int    `json:"MaxBackups" mapstructure:"MaxBackups"`       // 最大备份数
	MaxAge        int    `json:"MaxAge" mapstructure:"MaxAge"`               // 最大备份天数
	Compress      bool   `json:"Compress" mapstructure:"Compress"`           // 是否压缩过期日志
	LowLevelFile  string `json:"LowLevelFile" mapstructure:"LowLevelFile"`   // 低级别文件名
	HighLevelFile string `json:"HighLevelFile" mapstructure:"HighLevelFile"` // 高级别文件名
}

type SMTPInfo struct {
	Host     string   `json:"host" mapstructure:"host"`
	Port     int      `json:"port" mapstructure:"port"`
	IsSSL    bool     `json:"isSSL" mapstructure:"isSSL"`
	UserName string   `json:"userName" mapstructure:"userName"`
	Password string   `json:"password" mapstructure:"password"`
	From     string   `json:"from" mapstructure:"from"`
	To       []string `json:"to" mapstructure:"to"`
}

type Redis struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Password string `json:"password" mapstructure:"password"`
	PoolSize int    `json:"poolSize" mapstructure:"poolSize"`
}

type Work struct {
	TaskChanCapacity   int `json:"taskChanCapacity" mapstructure:"taskChanCapacity"`
	WorkerChanCapacity int `json:"workerChanCapacity" mapstructure:"workerChanCapacity"`
	WorkerNum          int `json:"workerNum" mapstructure:"workerNum"`
}

type Rule struct {
	DelUserTime    time.Duration `json:"delUserTime" mapstructure:"delUserTime"`  //延时删除用户的时间
	DelCodeTime    time.Duration `json:"delCodeTime" mapstructure:"delCodeTime"`  //延时删除验证码的时间
	AccountMaxNums int           `json:"accountMaxNum" mapstructure:"accountMax"` //账户可以创建的最大的数目

}

type Auto struct {
	Retry Retry `json:"retry" mapstructure:"retry"`
}

type Retry struct {
	TimeDuration time.Duration `json:"timeDuration" mapstructure:"timeDuration" ` //重试的时间间隔
	TimeCount    int           `json:"timeCount" mapstructure:"timeCount"`        //重试的次数
}

type Token struct {
	Key                string        `mapstructure:"Key"`
	AccessTokenExpire  time.Duration `mapstructure:"AccessTokenExpire"`
	RefreshTokenExpire time.Duration `mapstructure:"RefreshTokenExpire"`
	AccountTokenExpire time.Duration `mapstructure:"AccountTokenExpire"`
	AuthType           string        `mapstructure:"AuthType"`
	AuthKey            string        `mapstructure:"AuthKey"`
}
