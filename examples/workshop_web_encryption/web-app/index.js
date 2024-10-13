// Generate a key pair
var keyPair = nacl.sign.keyPair();

// Create a nonce (random message)
var nonce = nacl.randomBytes(32); // 32-byte nonce

// Sign the nonce with the private key
var signature = nacl.sign.detached(nonce, keyPair.secretKey);

// Encode data to Base64 for transmission
var publicKeyBase64 = nacl.util.encodeBase64(keyPair.publicKey);
var nonceBase64 = nacl.util.encodeBase64(nonce);
var signatureBase64 = nacl.util.encodeBase64(signature);

// Prepare data to send to the server
var data = {
    publicKey: publicKeyBase64,
    nonce: nonceBase64,
    signature: signatureBase64
};




// function fetchAndVerifyNonce() {
//     fetch('http://localhost:8080/getNonce', {
//         method: 'GET',
//         headers: {
//             'Content-Type': 'application/json'
//         }
//     })
//         .then(response => response.json())
//         .then(data => {
//             // Decode Base64 strings
//             var nonce = nacl.util.decodeBase64(data.nonce);
//             var signature = nacl.util.decodeBase64(data.signature);
//             var publicKey = nacl.util.decodeBase64(data.publicKey);
//
//             // Verify the signature
//             var isValid = nacl.sign.detached.verify(nonce, signature, publicKey);
//
//             if (isValid) {
//                 console.log('Server signature is valid.');
//             } else {
//                 console.error('Invalid server signature.');
//             }
//         })
//         .catch(error => {
//             console.error('Error:', error);
//         });
// }

// Call the function
//fetchAndVerifyNonce();

async function receiveEncryptedData() {
    try {
        const response = await fetch('http://localhost:8080/getEncryptedData', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        const data = await response.json();

        // Decode Base64 encoded data
        const encryptedAESKey = nacl.util.decodeBase64(data.encryptedAESKey);
        const nonceAESKey = nacl.util.decodeBase64(data.nonceAESKey);
        const encryptedMessage = nacl.util.decodeBase64(data.encryptedMessage);
        const nonceMessage = nacl.util.decodeBase64(data.nonceMessage);
        const serverBoxPublicKey = nacl.util.decodeBase64(data.serverBoxPublicKey);
        const serverSignPublicKey = nacl.util.decodeBase64(data.serverSignPublicKey);
        const signature = nacl.util.decodeBase64(data.signature);

        // Verify the server's signature
        const dataToVerify = data.nonceAESKey + data.encryptedMessage + data.nonceMessage;
        const isValidSignature = nacl.sign.detached.verify(
            nacl.util.decodeUTF8(dataToVerify),
            signature,
            serverSignPublicKey
        );

        if (!isValidSignature) {
            console.error('Invalid server signature');
            return;
        }

        console.log('Server signature is valid');

        // Decrypt the AES key using NaCl box
        const sharedSecretAESKey = nacl.box.open(
            encryptedAESKey,
            nonceAESKey,
            serverBoxPublicKey,
            boxKeyPair.secretKey
        );

        if (!sharedSecretAESKey) {
            console.error('Failed to decrypt AES key');
            return;
        }

        // Import the AES key into WebCrypto
        const aesKey = await window.crypto.subtle.importKey(
            'raw',
            sharedSecretAESKey,
            { name: 'AES-GCM' },
            false,
            ['encrypt', 'decrypt']
        );

        // Decrypt the message using AES-GCM
        const decryptedArrayBuffer = await window.crypto.subtle.decrypt(
            {
                name: 'AES-GCM',
                iv: nonceMessage
            },
            aesKey,
            encryptedMessage
        );

        const decryptedBytes = new Uint8Array(decryptedArrayBuffer);
        const decryptedMessage = nacl.util.encodeUTF8(decryptedBytes);
        console.log('Decrypted message from server:', decryptedMessage);

        // --- Encrypt a message to send to the server ---

        // Message to send
        const messageToServer = 'Hello from client';

        // Convert message to Uint8Array
        const messageBytes = nacl.util.decodeUTF8(messageToServer);

        // Generate a nonce for AES-GCM encryption
        const nonceMessageToServer = window.crypto.getRandomValues(new Uint8Array(12)); // AES-GCM nonce is 12 bytes

        // Encrypt the message using AES-GCM
        const encryptedMessageToServerBuffer = await window.crypto.subtle.encrypt(
            {
                name: 'AES-GCM',
                iv: nonceMessageToServer
            },
            aesKey,
            messageBytes
        );

        const encryptedMessageToServer = new Uint8Array(encryptedMessageToServerBuffer);

        // Prepare data to sign
        const nonceMessageToServerBase64 = nacl.util.encodeBase64(nonceMessageToServer);
        const encryptedMessageToServerBase64 = nacl.util.encodeBase64(encryptedMessageToServer);

        const dataToSign = nonceMessageToServerBase64 + encryptedMessageToServerBase64;

        // Sign the data with the client's signing private key
        const signatureToServer = nacl.sign.detached(
            nacl.util.decodeUTF8(dataToSign),
            signKeyPair.secretKey
        );

        // Prepare payload
        const payload = {
            nonceMessageToServer: nonceMessageToServerBase64,
            encryptedMessageToServer: encryptedMessageToServerBase64,
            signature: nacl.util.encodeBase64(signatureToServer)
        };

        // Send the data to the server
        const sendResponse = await fetch('http://localhost:8080/receiveEncryptedData', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        });

        const result = await sendResponse.text();
        console.log('Server response:', result);

    } catch (err) {
        console.error('Failed to process:', err);
    }
}


async function receiveEncryptedDataTMinusOne() {
    try {
        const response = await fetch('http://localhost:8080/getEncryptedData', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });
        const data = await response.json();

        // Decode Base64 encoded data
        const encryptedAESKey = nacl.util.decodeBase64(data.encryptedAESKey);
        const nonceAESKey = nacl.util.decodeBase64(data.nonceAESKey);
        const encryptedMessage = nacl.util.decodeBase64(data.encryptedMessage);
        const nonceMessage = nacl.util.decodeBase64(data.nonceMessage);
        const serverBoxPublicKey = nacl.util.decodeBase64(data.serverBoxPublicKey);
        var serverSignPublicKey = nacl.util.decodeBase64(data.serverSignPublicKey);
        const signature = nacl.util.decodeBase64(data.signature);


        console.log(data.serverSignPublicKey)
        console.log(serverSignPublicKey)

        // Verify the server's signature
        const dataToVerify = data.nonceAESKey + data.encryptedMessage + data.nonceMessage //+ data.serverBoxPublicKey;

        console.log('#########')
        console.log("dataToVerify:", dataToVerify);
        console.log('#########')


        const isValidSignature = nacl.sign.detached.verify(
            nacl.util.decodeUTF8(dataToVerify),
            signature,
            serverSignPublicKey
        );

        if (!isValidSignature) {
            console.error('Invalid server signature');
            return;
        }

        console.log('Server signature is valid');




        // Decrypt the AES key using NaCl box
        const sharedSecretAESKey = nacl.box.open(
            encryptedAESKey,
            nonceAESKey,
            serverBoxPublicKey,
            boxKeyPair.secretKey
        );

        if (!sharedSecretAESKey) {
            console.error('Failed to decrypt AES key');
            return;
        }

        // Import the AES key into WebCrypto
        const aesKey = await window.crypto.subtle.importKey(
            'raw',
            sharedSecretAESKey,
            { name: 'AES-GCM' },
            false,
            ['decrypt']
        );

        // Decrypt the message using AES-GCM
        const decryptedArrayBuffer = await window.crypto.subtle.decrypt(
            {
                name: 'AES-GCM',
                iv: nonceMessage
            },
            aesKey,
            encryptedMessage
        );

        const decryptedBytes = new Uint8Array(decryptedArrayBuffer);
        const decryptedMessage = nacl.util.encodeUTF8(decryptedBytes);
        console.log('Decrypted message:', decryptedMessage);


        // TODO: send an encrypted message for server to decrypt.
    } catch (err) {
        console.error('Failed to decrypt message:', err);
    }
}


// Function to receive encrypted message and AES key
function receiveEncryptedDataOld() {
    fetch('http://localhost:8080/getEncryptedData', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            // Decode Base64 encoded data
            var encryptedAESKey = nacl.util.decodeBase64(data.encryptedAESKey);
            var nonceAESKey = nacl.util.decodeBase64(data.nonceAESKey);
            var encryptedMessage = nacl.util.decodeBase64(data.encryptedMessage);
            var nonceMessage = nacl.util.decodeBase64(data.nonceMessage);
            var serverBoxPublicKey = nacl.util.decodeBase64(data.serverBoxPublicKey);
            var serverSignPublicKey = nacl.util.decodeBase64(data.serverSignPublicKey);



            // Verify the server's signature
            const dataToVerify = data.encryptedAESKey + data.nonceAESKey + data.encryptedMessage + data.nonceMessage + data.serverBoxPublicKey;
            const isValidSignature = nacl.sign.detached.verify(
                nacl.util.decodeUTF8(dataToVerify),
                signature,
                serverSignPublicKey
            );

            if (!isValidSignature) {
                console.error('Invalid server signature');
                return;
            }

            console.log('Server signature is valid');



            // Decrypt the AES key
            var sharedSecretAESKey = nacl.box.open(
                encryptedAESKey,
                nonceAESKey,
                serverBoxPublicKey,
                boxKeyPair.secretKey
            );

            if (!sharedSecretAESKey) {
                console.error('Failed to decrypt AES key');
                return;
            }

            // Decrypt the message
            var aesKey = sharedSecretAESKey; // AES key is the shared secret

            // Use AES decryption (need an AES library)
            // Assuming aesjs is included in your project
            var aesCtr = new aesjs.ModeOfOperation.ctr(aesKey);
            var decryptedBytes = aesCtr.decrypt(encryptedMessage);

            var decryptedMessage = nacl.util.encodeUTF8(decryptedBytes);
            console.log('Decrypted message:', decryptedMessage);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}


// Generate Ed25519 key pair for signing
var signKeyPair = nacl.sign.keyPair();

// Generate Curve25519 key pair for encryption
var boxKeyPair = nacl.box.keyPair();


// Encode public keys for transmission
var signPublicKeyBase64 = nacl.util.encodeBase64(signKeyPair.publicKey);
var boxPublicKeyBase64 = nacl.util.encodeBase64(boxKeyPair.publicKey);


const dataToSign = boxPublicKeyBase64;

// Sign the data with the client's signing private key
const signatureToServer = nacl.sign.detached(
    nacl.util.decodeUTF8(dataToSign),
    signKeyPair.secretKey
);

// // Prepare payload
// const payload = {
//     nonceMessageToServer: nonceMessageToServerBase64,
//     encryptedMessageToServer: encryptedMessageToServerBase64,
//     signature: nacl.util.encodeBase64(signatureToServer)
// };
//


// Prepare data to send to the server
var data = {
    signPublicKey: signPublicKeyBase64,
    boxPublicKey: boxPublicKeyBase64,
    signature: nacl.util.encodeBase64(signatureToServer)
};

// Send data to the server
fetch('http://localhost:8080/register', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json'
    },
    body: JSON.stringify(data)
})
    .then(response => response.text())
    .then(result => {
        console.log('Server response:', result);
    })
    .then(() => {
        // Call the function to receive encrypted data
        receiveEncryptedData();
    })
    .catch(error => {
        console.error('Error:', error);
    });

// // Send data to the server
// fetch('http://localhost:8080/verify', { // Ensure the URL and port are correct
//     method: 'POST',
//     headers: {
//         'Content-Type': 'application/json'
//     },
//     body: JSON.stringify(data)
// })
//     .then(response => response.text())
//     .then(result => {
//         console.log('Server response:', result);
//     })
//     .then(() => {
//         // Call the function to receive encrypted data
//         receiveEncryptedData();
//     })
//     .catch(error => {
//         console.error('Error:', error);
//     });

