const appState = {
    keys: {
        rsa: {
            encryption: {
                publicKey: null,
                privateKey: null,
            },
            signing: {
                publicKey: null,
                privateKey: null,
            }
        },
        aes: null
    },
    nonce: "",
    signedNonce: "",
    serverPublicKey: null,
    sharedAESKey: null
}

async function initializeWithServer() {
    try {
        // Generate client's RSA key pair for encryption
        const clientEncryptionKeyPair = await window.crypto.subtle.generateKey(
            {
                name: "RSA-OAEP",
                modulusLength: 2048,
                publicExponent: new Uint8Array([1, 0, 1]),
                hash: "SHA-256",
            },
            true,
            ["encrypt", "decrypt"]
        );

        // Generate client's RSA key pair for signing
        const clientSigningKeyPair = await window.crypto.subtle.generateKey(
            {
                name: "RSA-PSS",
                modulusLength: 2048,
                publicExponent: new Uint8Array([1, 0, 1]),
                hash: "SHA-256",
            },
            true,
            ["sign", "verify"]
        );

        // Store the keys in appState
        appState.keys.rsa.encryption.publicKey = await exportPublicKey(clientEncryptionKeyPair.publicKey);
        appState.keys.rsa.encryption.privateKey = await exportPrivateKey(clientEncryptionKeyPair.privateKey);
        appState.keys.rsa.signing.publicKey = await exportPublicKey(clientSigningKeyPair.publicKey);
        appState.keys.rsa.signing.privateKey = await exportPrivateKey(clientSigningKeyPair.privateKey);

        // Generate client nonce
        appState.nonce = generateNonce();

        // Sign client nonce
        appState.signedNonce = await signMessage(
            appState.nonce,
            await window.crypto.subtle.importKey(
                "pkcs8",
                base64ToArrayBuffer(appState.keys.rsa.signing.privateKey),
                {
                    name: "RSA-PSS",
                    hash: "SHA-256",
                },
                false,
                ["sign"]
            )
        );

        console.log("Sending client public key:", appState.keys.rsa.encryption.publicKey);
        console.log("Sending client nonce:", appState.nonce);

        const response = await fetch('http://localhost:8080/init', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                clientPublicKey: appState.keys.rsa.encryption.publicKey,
                clientNonce: appState.nonce,
                signedNonce: appState.signedNonce
            }),
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        console.log("Received from server:", data);

        // Store the server's public key
        appState.serverPublicKey = data.serverPublicKey;

        // Verify the signed nonce
        const isValidSignature = await verifySignature(data.nonce, data.signedNonce, data.serverPublicKey);

        if (isValidSignature) {
            console.log("Server signature verified successfully");

            // Decrypt the AES key
            const decryptedAESKey = await decryptAESKey(data.encryptedAESKey);
            appState.sharedAESKey = arrayBufferToBase64(decryptedAESKey);

            console.log("AES key decrypted successfully");
        } else {
            console.error("Invalid server signature");
        }

        console.log("Initialization with server complete");
    } catch (error) {
        console.error('Error during server initialization:', error);
    }
}

async function signMessage(message, privateKey) {
    const encoder = new TextEncoder();
    const data = encoder.encode(message);

    // Hash the message using SHA-256 before signing it
    const hashedData = await window.crypto.subtle.digest('SHA-256', data);

    const signature = await window.crypto.subtle.sign(
        {
            name: "RSA-PSS",
            saltLength: 32,
        },
        privateKey,
        hashedData // Sign the hashed data
    );
    return arrayBufferToBase64(signature);
}

async function verifySignature(data, signature, publicKeyString) {
    console.log("Verifying signature:");
    console.log("Data (base64):", data);
    console.log("Signature (base64):", signature);
    console.log("Public Key (base64):", publicKeyString);

    try {
        const publicKey = await window.crypto.subtle.importKey(
            "spki",
            base64ToArrayBuffer(publicKeyString),
            {
                name: "RSA-PSS",
                hash: "SHA-256",
            },
            false,
            ["verify"]
        );

        const dataBuffer = base64ToArrayBuffer(data);
        const signatureBuffer = base64ToArrayBuffer(signature);

        console.log("Data buffer length:", dataBuffer.byteLength);
        console.log("Signature buffer length:", signatureBuffer.byteLength);

        // Hash the data first
        const hashedData = await window.crypto.subtle.digest('SHA-256', dataBuffer);
        console.log("Hashed data length:", hashedData.byteLength);

        const result = await window.crypto.subtle.verify(
            {
                name: "RSA-PSS",
                saltLength: 32,
            },
            publicKey,
            signatureBuffer,
            hashedData
        );

        console.log("Signature verification result:", result);
        return result;
    } catch (error) {
        console.error("Error during signature verification:", error);
        return false;
    }
}

async function decryptAESKey(encryptedKeyBase64) {
    const privateKey = await window.crypto.subtle.importKey(
        "pkcs8",
        base64ToArrayBuffer(appState.keys.rsa.encryption.privateKey),
        {
            name: "RSA-OAEP",
            hash: "SHA-256",
        },
        false,
        ["decrypt"]
    );

    const encryptedKey = base64ToArrayBuffer(encryptedKeyBase64);
    return await window.crypto.subtle.decrypt(
        {
            name: "RSA-OAEP"
        },
        privateKey,
        encryptedKey
    );
}

async function exportPublicKey(publicKey) {
    const exported = await window.crypto.subtle.exportKey(
        "spki",
        publicKey
    );
    return arrayBufferToBase64(exported);
}

async function exportPrivateKey(privateKey) {
    const exported = await window.crypto.subtle.exportKey(
        "pkcs8",
        privateKey
    );
    return arrayBufferToBase64(exported);
}

function generateNonce(length = 32) {
    const charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let nonce = '';
    const randomValues = new Uint8Array(length);
    crypto.getRandomValues(randomValues);
    for (let i = 0; i < length; i++) {
        nonce += charset[randomValues[i] % charset.length];
    }
    return nonce;
}

function arrayBufferToBase64(buffer) {
    const binary = String.fromCharCode.apply(null, new Uint8Array(buffer));
    return window.btoa(binary);
}

function base64ToArrayBuffer(base64) {
    const binaryString = window.atob(base64);
    const bytes = new Uint8Array(binaryString.length);
    for (let i = 0; i < binaryString.length; i++) {
        bytes[i] = binaryString.charCodeAt(i);
    }
    return bytes.buffer;
}

// Call this function to start the initialization process
initializeWithServer().then(() => console.log("Initialization process completed"));