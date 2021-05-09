
var captchaForm = {
    ShowLineOptions: [],
    CaptchaType: "math",
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
        Length: 6,
        Fonts: ["wqy-microhei.ttc"],
        BgColor: {R: 0, G: 0, B: 0, A: 0},
    },
    DriverMath: {
        Height: 34,
        Width: 240,
        ShowLineOptions: 0,
        NoiseCount: 0,
        Length: 6,
        Fonts: ["wqy-microhei.ttc"],
        BgColor: {R: 0, G: 0, B: 0, A: 0},
    },
    DriverDigit: {
        Height: 80,
        Width: 240,
        Length: 5,
        MaxSkew: 0.7,
        DotCount: 80
    },
    blob: "",
    isLoading: false   
}

$(document).ready(function() {
    generateCaptcha();
});

function generateCaptcha() {
    console.log('Generating captcha');
    const url = 'api/getCaptcha';
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

function displayCaptcha(captcha) {
let captchaImage = "<h2>Human Verification</h2>" + "<h3>Solve the math problem:</h3>" + "<img src='" + captcha.data + "'/>";
$("#captcha-img").html(captchaImage);
}

function verifyCaptcha() {
    console.log("verifyCaptcha");
    captchaForm.VerifyValue = document.getElementById("captcha-solution").value;

    const url = '/api/verifyCaptcha';

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
        console.log(data.msg);
        if (data.code == "success") {
            $(".captcha-row").hide("slow", function(){
                showMessage(data.msg);
            });
        } else {
            console.log(data.code);
            showMessage(data.msg);
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
let msg = msgText;
$(".alert").find('.message').text(msg);
$(".alert").fadeIn("slow", function() {
    setTimeout(function(){
        $(".alert").fadeOut("slow");
    }, 2000);
});
}
