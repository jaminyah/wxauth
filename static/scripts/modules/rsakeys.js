
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

export { getPublicKey, publicKeyEncrypt };