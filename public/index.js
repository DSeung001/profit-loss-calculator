// use strict 엄격한 체크 모드
(function ($) {
    'use strict';
    $(function () {
        $('.money-number').on('keyup', function () {
            updateTextView($(this));
        });

        $('#calculatorForm').submit(function (event) {
            event.preventDefault();
            getVar()
            $.ajax({
                url: "/profitLossCalculator",
                type: "GET",
                data: {
                    supplyPrice: supplyPrice,
                    cash: cash,
                    loanAmount: loanAmount,
                    interestRate: interestRate,
                    principalAndInterest: principalAndInterest
                },
                success: function(data) {
                    console.log(data)
                },
                fail: function(jqXHR, textStatus, errorThrown) {
                    console.log(jqXHR)
                    console.log("API 요청 실패:", textStatus, errorThrown);
                }
            })
        })

        setVar()
    });
})(jQuery);

let supplyPrice = 156830000
let cash = 62830000
let loanAmount = 94000000
let interestRate = 5.9
let principalAndInterest = 557548


function setVar(){
    $("#supplyPrice").val(supplyPrice);
    $("#cash").val(cash);
    $("#loanAmount").val(loanAmount);
    $("#interestRate").val(interestRate);
    $("#principalAndInterest").val(principalAndInterest);

    updateTextView($("#supplyPrice"));
    updateTextView($("#cash"));
    updateTextView($("#loanAmount"));
    updateTextView($("#principalAndInterest"));
}

function getVar() {
    supplyPrice = $("#supplyPrice").val();
    cash = $("#cash").val();
    loanAmount = $("#loanAmount").val();
    interestRate = $("#interestRate").val();
    principalAndInterest = $("#principalAndInterest").val();
}

// Text를 toLocaleString() 함수를 이용하여 숫자에 콤마를 추가
function updateTextView(_obj) {
    var num = getNumber(_obj.val());
    if (num == 0) {
        _obj.val('');
    } else {
        _obj.val(num.toLocaleString());
    }
}

// _str 에서 숫자만 추출
function getNumber(_str) {
    var arr = _str.split('');
    var out = new Array();
    for (var cnt = 0; cnt < arr.length; cnt++) {
        if (isNaN(arr[cnt]) == false) {
            out.push(arr[cnt]);
        }
    }
    return Number(out.join(''));
}