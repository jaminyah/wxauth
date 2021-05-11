
var captchaForm = {
    ShowLineOptions: [],
    CaptchaType: "string",
    Id: '',
    VerifyValue: '',
    DriverAudio: {
        Length: 6,
        Language: 'en'
    },
    DriverString: {
        Height: 60,
        Width: 240,
        ShowLineOptions: 0,
        NoiseCount: 0,
        Source: "1234567890qwertyuioplkjhgfdsazxcvbnm",
        Length: 6,
        Fonts: ["wqy-microhei.ttc"],
        BgColor: {R: 0, G: 0, B: 0, A: 0},
    },
    DriverMath: {
        Height: 60,
        Width: 240,
        ShowLineOptions: 0,
        NoiseCount: 0,
        Length: 6,
        Fonts: ["wqy-microhei.ttc"],
        BgColor: {R: 0, G: 0, B: 0, A: 0},
    },
    DriverChinese: {
        Height: 60,
        Width: 320,
        ShowLineOptions: 0,
        NoiseCount: 0,
        Source: "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,,不想要,的值",
        Length: 2,
        Fonts: ["wqy-microhei.ttc"],
        BgColor: {R: 125, G: 125, B: 0, A: 118},
    },
    DriverDigit: {
        Height: 80,
        Width: 240,
        Length: 5,
        MaxSkew: 0.7,
        DotCount: 80
    },
blob: "",
}


$(document).ready(function() {
    $("#captcha-solution").val("");
    generateCaptcha()
});


function displayCaptcha(captcha) {
    let captchaImage = "<img src='" + captcha.data + "'/>";
    $("#captcha-img").html(captchaImage);
}

function generateCaptcha() {
    console.log('Generating captcha');
    const url = '/api/getcaptcha';
    var blob = "";

    let fetchData = {
        method: 'post',
        body: JSON.stringify(captchaForm),
        headers: new Headers()
    }

    fetch(url, fetchData)
    .then(function(response){
        console.log(" fetch .then");
        return response.json();
    })
    .then(function(data){
        console.log(data.captchaId);
        captchaForm.Id = data.captchaId;
        blob = data.data;
        displayCaptcha(data);
    })
    .catch(function(error){
        console.log("fetch error: ")
        console.log(error);
    });
}

$('#comment-form').submit(function(e) {

    e.preventDefault();         // avoid executing the actual submit form

    var form = $(this);
    $.ajax({
        type: form.attr('method'),
        url: "/submit",
        contentType: 'application/x-www-form-urlencoded',
        data: form.serialize(),
        success: function(data) {
            $("#username").val("");
            $("#message").val("");
            $("#captcha-solution").val("");
            $(".captcha-row").show(2000);
        },
        error: function(data) {
            console.log('There is an error');
            console.log(data);
        }
    });
});

function verifyCaptcha() {
    console.log("verifyCaptcha");
    captchaForm.VerifyValue = document.getElementById("captcha-solution").value;
    console.log(captchaForm.VerifyValue)

    const url = '/api/verifycaptcha';

    let fetchData = {
        method: 'post',
        body: JSON.stringify(captchaForm),
        headers: new Headers()
    }

    fetch(url, fetchData)
    .then(function(response){
        return response.json();
    })
    .then(function(data){
        console.log("verify server response: ", data.msg);
        if (data.msg == "ok") {
            console.log("verify captcha success.")
            $("#captcha-solution").val("");
            $("#captcha-blk").hide("slow", function(){
                console.log(data.msg);
            });
        } else {
            console.log("verify captcha fail.")
            console.log(data.code);
            showMessage(data.msg);
            $("#captcha-solution").val("");
            generateCaptcha();
        }
    })
    .catch(function(error){
        console.log("fetch error: ")
        console.log(error);
    });
}

function showMessage(msgText) {
    console.log("Show message.")
    /*let msg = msgText;
    $(".alert").find('.message').text(msg);
    $(".alert").fadeIn("slow", function() {
        setTimeout(function(){
            $(".alert").fadeOut("slow");
        }, 2000);
    });*/
}