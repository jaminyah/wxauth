/*
* Email validation reference:
* https://stackoverflow.com/questions/46155/how-to-validate-an-email-address-in-javascript
*/

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

let emailValidated = 'false';
let passwdConfirmed = 'false';
let captchaConfirmed = 'false';

$(document).ready(function() {
    $("#captcha-solution").val("");
    $("#field-email").val("");
    $('input[type=checkbox]').prop('checked',false);
    $("#reg-btn").prop("disabled", true);
    $("#input-passwd, #input-confirm").keyup(checkPasswordMatch);
    getCaptcha();
});


function displayCaptcha(captcha) {
    let captchaImage = "<img src='" + captcha.data + "'/>";
    $("#captcha-img").html(captchaImage);
}


function getCaptcha() {
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

/*
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
*/


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

            captchaConfirmed = 'true';
            $("#captcha-solution").val("");
            $("#captcha-blk").hide("slow", function(){
                console.log(data.msg);
            });
            checkSubmit();
        } else {
            console.log("verify captcha fail.")
            console.log(data.code);
            showMessage(data.msg);
            $("#captcha-solution").val("");
            captchaConfirmed = 'false';
            getCaptcha();
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

function registerUser() {
    console.log('Register submit');
    const url = '/api/register';

    var address = document.getElementById("field-email").value;
    var password = document.getElementById("field-passwd").value;

    var registerForm = {
        emailAddr: address,
        password: password
    }

    let registerData = {
        method: 'post',
        body: JSON.stringify(registerForm),
        headers: new Headers()
    }

    // http response code = 200 success, client redirect to activation.html
    // http response code = 400 fail, client remains on register page
    /*
    fetch(url, registerData)
    .then(function(response) {
        console.log("registerData response.");
        return response.json();
    })
    .then(function(data) {
        console.log(data.captchaId);
        captchaForm.Id = data.captchaId;
        blob = data.data;
        displayCaptcha(data);
      
    })
    .catch(function(error) {
        console.log("fetch error: ")
        console.log(error);
    });
    */
}

/******************************************************** EMAIL VALIDATION *******/

let email = document.getElementById("input-email");
let emailError = document.getElementById("emailError");
emailError.style.display = "none";

email.addEventListener("keyup", event => {

    console.log('verify email')

    let text = email.value
    emailCheck(text, emailError);
});


function emailCheck(text, emailError) {

    //let condition = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    let condition = /(?!.*\.{2})^([a-z\d!#$%&'*+\-\/=?^_`{|}~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+(\.[a-z\d!#$%&'*+\-\/=?^_`{|}~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]+)*|"((([ \t]*\r\n)?[ \t]+)?([\x01-\x08\x0b\x0c\x0e-\x1f\x7f\x21\x23-\x5b\x5d-\x7e\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|\\[\x01-\x09\x0b\x0c\x0d-\x7f\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]))*(([ \t]*\r\n)?[ \t]+)?")@(([a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|[a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF][a-z\d\-._~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]*[a-z\d\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])\.)+([a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]|[a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF][a-z\d\-._~\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF]*[a-z\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])\.?$/i

    emailError.style.color = "red";
    const inputEmail = document.getElementById('input-email');
    const iconEmail = document.getElementById('emailValidIcon')

    if (!text.match(condition)) {
        inputEmail.classList.remove('is-success')
        inputEmail.classList.add('is-danger')
        iconEmail.classList.remove('fa-check')
        iconEmail.classList.add('fa-exclamation-triangle')
        emailError.style.display = "block";
        emailError.innerText = "Invalid Email Address";
    } else {
            inputEmail.classList.remove('is-danger')
            inputEmail.classList.add('is-success')
            iconEmail.classList.remove('fa-exclamation-triangle')
            iconEmail.classList.add('fa-check')
            emailError.style.display = "none";
            emailValidated = 'true';
            checkSubmit();
    }
    return;
}

/******************************************************** PASSWORD VALIDATION *******/

const passInput = document.getElementById("input-passwd");
const passwordError = document.getElementById("passwdError");

passwordError.style.display = "none";
passInput.addEventListener('keyup', event => {
    let text = passInput.value;
    passwdCheck(text, passwordError);
});

function passwdCheck(text, passwordError) {
    let condition1 = /(?=.*\d)/; //should contain atleast 1 digit
    let condition2 = /(?=.*[a-z])/; //should contain atleast 1 lowercase
    let condition3 = /(?=.*[A-Z])/; //should contain atleast 1 uppercase
    let condition4 = /[a-zA-Z0-9]{8,}/; //should contain atleast 8 characters

    passwdError.style.color = "red";

    if (!text.match(condition1)) {
        passwdError.style.display = "block";
        passwdError.innerText = "Password should contain atleast 1 digit";
        showPassErrorStyle();
    } else

    if (!text.match(condition2)) {
        passwdError.style.display = "block";
        passwdError.innerText = "Password should contain atleast 1 lowercase";
        showPassErrorStyle();
    } else

    if (!text.match(condition3)) {
        passwdError.style.display = "block";
        passwdError.innerText = "Password should contain atleast 1 uppercase";
        showPassErrorStyle();
    } else

    if (!text.match(condition4)) {
        passwdError.style.display = "block";
        passwdError.innerText = "Password should contain atleast 8 characters";
        showPassErrorStyle();
    } else {
        passwdError.style.display = "none";
        showPassSucessStyle();
    }
    return;
}

function showPassErrorStyle() {
    passInput.classList.remove('is-success');
    passInput.classList.add('is-danger');
    passwdValidIcon.classList.remove('fa-check');
    passwdValidIcon.classList.add('fa-exclamation-triangle');
}

function showPassSucessStyle() {
    passInput.classList.remove('is-danger');
    passInput.classList.add('is-success');
    passwdValidIcon.classList.remove('fa-exclamation-triangle');
    passwdValidIcon.classList.add('fa-check');
}

/******************************************************** PASSWORD CONFIRMATION *****/

let confirmInput = document.getElementById("input-confirm");
let confirmError = document.getElementById("confirmError");

function checkPasswordMatch() {

    let inputPass = $("#input-passwd").val();
    let inputConfirm = $("#input-confirm").val();

    if (inputPass !== inputConfirm) {
        $("#confirmError").html("Passwords match fail.");
        showConfirmErrorStyle();
    } else {
        $("#confirmError").html("Passwords match success.");
        showConfirmSucessStyle();
        $("#reg-btn").prop("disabled", false);
    }
} 

function showConfirmErrorStyle() {
    confirmInput.classList.remove('is-success');
    confirmInput.classList.add('is-danger');
    confirmValidIcon.classList.remove('fa-check');
    confirmValidIcon.classList.add('fa-exclamation-triangle');
}

function showConfirmSucessStyle() {
    confirmInput.classList.remove('is-danger');
    confirmInput.classList.add('is-success');
    confirmValidIcon.classList.remove('fa-exclamation-triangle');
    confirmValidIcon.classList.add('fa-check');
}

/******************************************************** TOGGLE PASSWORD *****/

const passwdVisible = document.querySelector('.passwd-chkbox');

passwdVisible.addEventListener('click', function(e) {
    const passtype = passInput.getAttribute('type') === 'password' ? 'text' : 'password';
    const confirmtype = confirmInput.getAttribute('type') === 'password' ? 'text' : 'password';
    passInput.setAttribute('type', passtype);
    confirmInput.setAttribute('type', confirmtype);
});

function checkSubmit() {
    console.log('check submit');

    if ( emailValidated && passwdConfirmed && captchaConfirmed) {
        $("#reg-btn").prop("disableed", false);
    }
}