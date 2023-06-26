package main

import (
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "gopkg.in/natefinch/lumberjack.v2"

    "blog/global"
    "blog/internal/model"
    "blog/internal/routers"
    "blog/pkg/logger"
    "blog/pkg/setting"
)

// 2.1.0 初始化配置读取
func init() {
    err := setupSetting()
    if err != nil {
        log.Fatalf("init.setupSetting err: %v", err)
    }
    err = setupDBEngine()
    if err != nil {
        log.Fatalf("init.setupDBEngine err: %v", err)
    }
    err = setupLogger()
    if err != nil {
        log.Fatalf("init.setupLogger err: %v", err)
    }
}

// 2.1.1 初始化配置读取
func setupSetting() error {
    setting, err := setting.NewSetting()
    if err != nil {
        return err
    }
    err = setting.ReadSection("Server", &global.ServerSetting)
    if err != nil {
        return err
    }
    err = setting.ReadSection("App", &global.AppSetting)
    if err != nil {
        return err
    }
    err = setting.ReadSection("Database", &global.DatabaseSetting)
    if err != nil {
        return err
    }

    global.ServerSetting.ReadTimeout *= time.Second
    global.ServerSetting.WriteTimeout *= time.Second
    return nil
}

// 2.1.2 初始化数据库配置读取
func setupDBEngine() error {
    var err error
    global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
    if err != nil {
        return err
    }
    return nil
}

// 2.1.3 初始化日志
func setupLogger() error {
    global.Logger = logger.NewLogger(&lumberjack.Logger{
        Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
        MaxSize:   600,
        MaxAge:    10,
        LocalTime: true,
    }, "", log.LstdFlags).WithCaller(2)

    return nil
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/hd2yao/blog
func main() {
    // 1.0.0 初始
    //r := gin.Default()
    //r.GET("/ping", func(context *gin.Context) {
    //    context.JSON(http.StatusOK, gin.H{"message": "pong"})
    //})
    //r.Run()

    // 2.0.0 自定义路由
    //router := routers.NewPouter()
    //s := &http.Server{
    //    Addr:           ":8080",
    //    Handler:        router,
    //    ReadTimeout:    10 * time.Second,
    //    WriteTimeout:   10 * time.Second,
    //    MaxHeaderBytes: 1 << 20,
    //}
    //s.ListenAndServe()

    // 2.1.0 配置管理
    gin.SetMode(global.ServerSetting.RunMode)
    router := routers.NewPouter()
    s := &http.Server{
        Addr:           ":" + global.ServerSetting.HttpPort,
        Handler:        router,
        ReadTimeout:    global.ServerSetting.ReadTimeout,
        WriteTimeout:   global.ServerSetting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}
