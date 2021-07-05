//import { publicKeyEncrypt } from './modules/rsakeys';

/******************************************************** LOGIN USER *****/
$(document).ready(function() {

    if (window.localStorage.getItem('rsaPublic') === null) {
        getPublicKey();
    }
});

function loginUser() {
    console.log('Login user submit');
    const url = '/api/login';

    var addr = document.getElementById("input-email").value;
    var passwd = document.getElementById("input-passwd").value;

    let passEncrypted = publicKeyEncrypt(passwd)

    var loginForm = {
        Email: addr,
        PassRSA: passEncrypted
    }

    let loginData = {
        method: 'post',
        body: JSON.stringify(loginForm),
        headers: new Headers()
    }

    // http response code = 200 success, client redirect to activation.html
    // http response code = 400 fail, client remains on register page
    
    fetch(url, loginData)
    .then(function(response) {
        console.log("registerData response.");
        return response.json();
    })
    .then(function(data) {
        console.log(data.code);
        console.log(data.msg);
        console.log(data.email);

       if (data.code === 200) {
           //window.location.href = "http://localhost:8090/activate.html"
           // window.location.href = "http://auth.wxalert.us/activate.html"
           console.log("Login success.")
       }
    })
    .catch(function(error) {
        console.log("fetch error: ");
        console.log(error);
    });

}

/**************** RSA public key ********************************/

function getPublicKey() {
    console.log("Public Key");

    const url = 'api/public';

    fetch(url) 
        .then(function(response) {
            console.log("Response");
            return response.json();
        })
        .then(function(data) {
            console.log(data.pubkey)
            console.log(data.msg);
            let rsaPubKey = data.pubkey;
            window.localStorage.setItem('rsaPublic', rsaPubKey)
        })
        .catch(function(error) {
            console.log(error);
        });
}

function publicKeyEncrypt(userPass) {

    let encoder = new JSEncrypt({
        default_key_size: 2048
    });

    let rsaPubKey = window.localStorage.getItem('rsaPublic')
    encoder.setPublicKey(rsaPubKey);

    let encoded = encoder.encrypt(userPass);
    return encoded;
}
