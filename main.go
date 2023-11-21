package main

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

// 자동 랜더링
var rd *render.Render

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	return mux
}

func main() {
	rd = render.New()
	m := MakeWebHandler()
	/*
		negroni.Classic()은 기본적으로 다음과 같은 미들웨어를 포함합니다.
		negroni.Recovery- 패닉 복구 미들웨어.
		negroni.Logger- 요청/응답 로거 미들웨어.
		negroni.Static- "public" 디렉토리 아래에 제공되는 정적 파일입니다.
	*/
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
