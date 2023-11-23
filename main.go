package main

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"profit-loss-calculator.com/utils"
	"strconv"
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
	Number               int `json:"number"`
	Principal            int `json:"principal"`
	Interest             int `json:"interest"`
	PrincipalAndInterest int `json:"principal and interest"`
	LoanBalance          int `json:"loan balance"`
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

	var queryParam = r.URL.Query()
	var calValue CalValue

	// 에러 핸들링 필요
	calValue.supplyPrice, _ = utils.InNumberCharRemove(queryParam["supplyPrice"][0])
	calValue.cash, _ = utils.InNumberCharRemove(queryParam["cash"][0])
	calValue.loanAmount, _ = utils.InNumberCharRemove(queryParam["loanAmount"][0])
	calValue.interestRate, _ = strconv.ParseFloat(queryParam["interestRate"][0], 64)
	calValue.principalAndInterest, _ = utils.InNumberCharRemove(queryParam["principalAndInterest"][0])

	list := make(installmentSlice, 0)
	var idx = 1
	for calValue.loanAmount > 0 {
		// 이자 계산
		interest := int(float64(calValue.loanAmount)*calValue.interestRate/100) / 12

		// 원금 계산
		principal := calValue.principalAndInterest - interest

		// 대출잔액
		loanBalance := calValue.loanAmount - principal
		calValue.loanAmount = loanBalance

		// 날짜 계산해서 사이즈 지정 필요

		data := Installment{
			Number:               idx,
			Principal:            principal,
			Interest:             interest,
			PrincipalAndInterest: calValue.principalAndInterest,
			LoanBalance:          loanBalance,
		}

		list = append(list, data)
		idx++
	}

	// 반환 값에 필드가 소문자면 반환되지 않음
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
