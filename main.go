package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

// 자동 랜더링
var rd *render.Render

type CalValue struct {
	supplyPrice          int
	cash                 int
	loanAmount           int
	interestRate         float64
	principalAndInterest int
}

type Installment struct {
	number int
	//date                 string
	principal            int
	interest             int
	principalAndInterest int
	loanBalance          int
}

// 나중에 기간에 맞게 make로 바꿔서 slice 안쓰게 수정 필요
type installmentSlice []Installment

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/profitLossCalculator", getCalculatorHandler).Methods(http.MethodGet)
	return mux
}

func getCalculatorHandler(w http.ResponseWriter, r *http.Request) {

	// r.Body 빔
	fmt.Println(r.Body)
	var calValue CalValue
	// err : EOF 발생
	err := json.NewDecoder(r.Body).Decode(&calValue)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	list := make(installmentSlice, 400)
	var idx = 1
	for calValue.loanAmount > 0 {
		// 이자 계산
		interest := int(float64(calValue.loanAmount) * calValue.interestRate / 100)

		// 원금 계산
		principal := calValue.principalAndInterest - interest

		// 대출잔액
		loanBalance := calValue.loanAmount - principal
		calValue.loanAmount -= loanBalance

		// 날짜 계산해서 사이즈 지정 필요

		list = append(list,
			Installment{
				number:               idx,
				principal:            principal,
				interest:             interest,
				principalAndInterest: calValue.principalAndInterest,
				loanBalance:          loanBalance,
			})
	}

	rd.JSON(w, http.StatusOK, list)
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
