package web

import (
	"encoding/json"
	"go-IM/pkg/errs"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
)

type Middleware struct {
}

func (m Middleware) LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %q %v", r.Method, r.URL.String(), time.Now().Sub(startTime))
	}
	return http.HandlerFunc(fn)
}

func (m Middleware) CORSHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "*")
		header.Set("Access-Control-Allow-Credentials", "true")
		header.Set("Access-Control-Allow-Methods", "POST,GET,DELETE,PUT,OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (m Middleware) RecoverPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				var httpErr errs.HttpErr
				log.Printf("Recover from panic:%+v", err)
				printStack()
				switch err.(type) {
				case errs.HttpErr:
					httpErr = err.(errs.HttpErr)
				default:
					httpErr = errs.ServerInternalError
				}

				w.Header().Add("Content-Type", "application/json;charset=UTF-8")
				w.WriteHeader(httpErr.HttpSC)
				resStr, _ := json.Marshal(httpErr.Err)
				_, _ = io.WriteString(w, string(resStr))
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	log.Print(string(buf[:n]))
}
