package main

import (
	"fmt"
	"gowith/config"
	"gowith/logging"
	_ "gowith/models/mysql"
	"gowith/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.HTTPPort),
		Handler:           router,
		ReadTimeout:       config.ReadTimeout,
		WriteTimeout:      config.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	//开启服务
	go func() {
		//http服务
		err := s.ListenAndServe()
		if err != nil{
			logging.Error(err)
		}

		//https服务
		//err := s.ListenAndServeTLS("config/3875272_www.zhxf.yuhualab.com.pem", "config/3875272_www.zhxf.yuhualab.com.key")
		//if err != nil{
		//	logging.Error(err)
		//	fmt.Println(err)
		//}
	}()
	select {

	}
}
