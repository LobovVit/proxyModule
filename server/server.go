package server

import (
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"proxymodule/utils/config"
	"syscall"
	"time"
)

var Log *zap.Logger
var Srv *http.Server

func Start(l *zap.Logger, Config *config.Cfg) error {
	Log = l
	router := mux.NewRouter()
	router.HandleFunc("/LoadDoc", loadDoc).Methods("GET")
	router.HandleFunc("/ReserveFunds", reserveFunds)
	router.HandleFunc("/MarkJournal", markJournal)
	router.HandleFunc("/shutdown", shutdown)

	Srv = &http.Server{
		Addr:    Config.PORT,
		Handler: router,
	}

	go func() {
		if err := Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Log.Fatal("listen: ", zap.Error(err))
		}
	}()
	Log.Info("Server Started")

	return nil
}

func Shutdown() {
	Log.Info("Server Stopped")
	time.Sleep(10 * time.Second)
	Log.Info("graceful shutdown 0 here") //еще принимаем новые подключения
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //можно установить предельное ожидание
	Log.Info("graceful shutdown 1 here")                                    //уже не принимаем новые подключения
	defer func() {
		Log.Info("graceful shutdown 2 here") //после остановки всего
		cancel()
	}()

	if err := Srv.Shutdown(ctx); err != nil {
		Log.Fatal("Server Shutdown Failed", zap.Error(err))
	}
	Log.Info("Server stopped ОК")
}

func loadDoc(w http.ResponseWriter, req *http.Request) {
	Log.Info("loadDoc start")
	time.Sleep(30 * time.Second)
	Log.Info("loadDoc 30 s")
	w.WriteHeader(200)
	w.Write([]byte("loadDoc Test is what we usually do"))
}

func reserveFunds(w http.ResponseWriter, req *http.Request) {
	Log.Info("reserveFunds start")
	time.Sleep(30 * time.Second)
	Log.Info("reserveFunds 30 s")
	w.WriteHeader(200)
	w.Write([]byte("reserveFunds Test is what we usually do"))
}

func markJournal(w http.ResponseWriter, req *http.Request) {
	Log.Info("markJournals start")
	time.Sleep(30 * time.Second)
	Log.Info("markJournals 30 s")
	w.WriteHeader(200)
	w.Write([]byte("markJournal Test is what we usually do"))
}

func shutdown(w http.ResponseWriter, req *http.Request) {
	Log.Info("shutdown start")
	w.WriteHeader(200)
	w.Write([]byte("shutdown Test is what we usually do"))
	//Shutdown()
	syscall.Kill(syscall.Getpid(), syscall.SIGINT) //посылаем ctl+c
}
