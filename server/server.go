package server

import (
	"go.uber.org/zap"
	"net/http"
)

var Log *zap.Logger

func Start(l *zap.Logger, port string) {
	Log = l
	//http.HandleFunc("/echoPayload", echoPayload)

	http.HandleFunc("/LoadDoc", loadDoc)
	http.HandleFunc("/ReserveFunds", reserveFunds)
	http.HandleFunc("/MarkJournal", markJournal)

	http.ListenAndServe(port, nil)
}

//func echoPayload(w http.ResponseWriter, req *http.Request ) {
//	Log.Info("MarkJournals start")
//	Log.Info("Request connection", zap.String("Протокол:",req.Proto), zap.String("URL.Path:",req.URL.Path[1:]))
//	defer req.Body.Close()
//	contents, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		Log.Panic("Oops! Failed reading body of the request", zap.Error(err))
//		http.Error(w, err.Error(), 500)
//	}
//
//	req.ParseForm()
//	params := req.Form
//	for s := range params {
//		Log.Info("В запросе", zap.String("key:",s),zap.String("val:",params.Get(s)))
//	}
//	Log.Info("В запросе", zap.String("Body:",string(contents)))
//}

func loadDoc(w http.ResponseWriter, req *http.Request) {
	Log.Info("loadDoc start")
}

func reserveFunds(w http.ResponseWriter, req *http.Request) {
	Log.Info("reserveFunds start")
}

func markJournal(w http.ResponseWriter, req *http.Request) {
	Log.Info("markJournals start")
}
