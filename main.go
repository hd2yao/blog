package main

import (
    "net/http"
    "time"

    "blog/internal/routers"
)

func main() {
    // 1.0.0
    //r := gin.Default()
    //r.GET("/ping", func(context *gin.Context) {
    //    context.JSON(http.StatusOK, gin.H{"message": "pong"})
    //})
    //r.Run()

    // 2.0.0
    router := routers.NewPouter()
    s := &http.Server{
        Addr:           ":8080",
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}
