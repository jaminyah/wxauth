let activateForm = {
    email: "",
    code: ""
}


$(document).ready(function() {
    $("#input-email").val("");
    $("#input-code").val("");
});

function activateUser() {
    console.log('Activate submit');

    const url = '/api/activate';
    activateForm.email = document.getElementById("input-email").value;
    activateForm.code = document.getElementById("input-code").value;

    // Debug
    console.log("email: " + activateForm.email)
    console.log("code: " + activateForm.code)

    let activateData = {
        method: 'post',
        body: JSON.stringify(activateForm),
        headers: new Headers()
    }

    // response code = 200 success, client redirect login on index.html
    // response code = 400 fail, client remains on activation page
    
    fetch(url, activateData)
    .then(function(response) {
        console.log("Activate data response.");
        return response.json();
    })
    .then(function(data) {
        console.log(data.code);
        console.log(data.msg);
        console.log(data.email);

       if (data.code === 200) {
           window.location.href = "http://localhost:8090/"
       }
    })
    .catch(function(error) {
        console.log("fetch error: ")
        console.log(error);
    });
}