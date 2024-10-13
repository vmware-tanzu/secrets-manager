const appConfig = {
  url: {
    REGISTER: 'http://localhost:8080/register',
    ENROLL: 'http://localhost:8080/getEncryptedData',
    POLL: 'http://localhost:8080/receiveEncryptedData'
  }
};

const appState = { nonceMessage: null, aesKey: null };

const respond = (message) => document.getElementById('encrypted-response-txt').innerText = message;
const respondDecrypted = (message) => document.getElementById('decrypted-response-txt').innerText = message;
const textToDecrypt = () => document.getElementById('encrypted-response-txt').innerText;

async function register() {
  const signKeyPair = nacl.sign.keyPair();
  const boxKeyPair = nacl.box.keyPair();

  const signPublicKeyBase64 = nacl.util.encodeBase64(signKeyPair.publicKey);
  const boxPublicKeyBase64 = nacl.util.encodeBase64(boxKeyPair.publicKey);

  const signatureToServer = nacl.sign.detached(
    nacl.util.decodeUTF8(boxPublicKeyBase64), signKeyPair.secretKey
  );

  return fetch(appConfig.url.REGISTER, {
    method: 'POST',
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify({
      signPublicKey: signPublicKeyBase64,
      boxPublicKey: boxPublicKeyBase64,
      signature: nacl.util.encodeBase64(signatureToServer)
    })
  }).then(() => ({
      signKeyPair,
      boxKeyPair,
      signPublicKeyBase64,
      boxPublicKeyBase64
    }))
    .catch((err) => {
      respond('Error:' + err);
      return {};
    })
}

const prepareEncryptionKeys = (data) => ({
  encryptedAESKey: nacl.util.decodeBase64(data.encryptedAESKey),
  nonceAESKey: nacl.util.decodeBase64(data.nonceAESKey),
  encryptedMessage: nacl.util.decodeBase64(data.encryptedMessage),
  nonceMessage: nacl.util.decodeBase64(data.nonceMessage),
  serverBoxPublicKey: nacl.util.decodeBase64(data.serverBoxPublicKey),
  serverSignPublicKey: nacl.util.decodeBase64(data.serverSignPublicKey),
  signature: nacl.util.decodeBase64(data.signature)
});

const verifySignature = ({
  nonceAESKey, encryptedMessage, nonceMessage, signature, serverSignPublicKey
  }) => {
  const dataToVerify = nonceAESKey + encryptedMessage + nonceMessage;
  return nacl.sign.detached.verify(
    nacl.util.decodeUTF8(dataToVerify), signature, serverSignPublicKey);
};

async function extractAesKey({ encryptedAESKey, nonceAESKey,
    serverBoxPublicKey, boxKeyPairSecretKey}) {
  const sharedSecretAESKey = nacl.box.open(
    encryptedAESKey, nonceAESKey, serverBoxPublicKey, boxKeyPairSecretKey
  );

  if (!sharedSecretAESKey) {
    respond('Failed to decrypt AES key');
    return null;
  }

  return await window.crypto.subtle.importKey(
    'raw', sharedSecretAESKey, {name: 'AES-GCM'}, false, ['encrypt', 'decrypt']
  ).catch((err) => {
    respond('Failed to import AES key:' + err);
    return null;
  })
}

async function verifyServer(cryptoParams) {
  try {
    const boxKeyPair = cryptoParams.boxKeyPair;
    const response = await fetch(appConfig.url.ENROLL, {
      method: 'GET', headers: {'Content-Type': 'application/json'}
    });
    const data = await response.json();
    respond(data.encryptedMessage);

    const {
      encryptedAESKey, nonceAESKey, encryptedMessage, nonceMessage,
      serverBoxPublicKey, serverSignPublicKey, signature
    } = prepareEncryptionKeys(data);

    appState.nonceMessage = nonceMessage;

    if (!verifySignature({
      nonceAESKey: data.nonceAESKey, encryptedMessage: data.encryptedMessage,
      nonceMessage: data.nonceMessage, signature, serverSignPublicKey
    })) {
      respond('Invalid server signature');
      return {};
    }

    const boxKeyPairSecretKey = boxKeyPair.secretKey;

    const aesKey = await extractAesKey({encryptedAESKey, nonceAESKey,
      serverBoxPublicKey, boxKeyPairSecretKey});
    appState.aesKey = aesKey;

    return {...cryptoParams, aesKey};
  } catch (err) {
    respond('Failed to process: ' + err);
    return {};
  }
}

const serverPayload = () => document.getElementById('payload-txt').value;

async function encryptAndSend({aesKey, signKeyPair}) {
  const messageBytes = nacl.util.decodeUTF8(serverPayload());
  const nonceMessageToServer = window.crypto.getRandomValues(new Uint8Array(12));
  const encryptedMessageToServerBuffer = await window.crypto.subtle.encrypt(
    {name: 'AES-GCM', iv: nonceMessageToServer},
    aesKey, messageBytes
  );

  const encryptedMessageToServer = new Uint8Array(encryptedMessageToServerBuffer);

  const nonceMessageToServerBase64 = nacl.util.encodeBase64(nonceMessageToServer);
  const encryptedMessageToServerBase64 = nacl.util.encodeBase64(encryptedMessageToServer);

  const dataToSign = nonceMessageToServerBase64 + encryptedMessageToServerBase64;
  const signatureToServer = nacl.sign.detached(
    nacl.util.decodeUTF8(dataToSign), signKeyPair.secretKey
  );

  fetch(appConfig.url.POLL, {
    method: 'POST', headers: {'Content-Type': 'application/json'},
    body: JSON.stringify({
      nonceMessageToServer: nonceMessageToServerBase64,
      encryptedMessageToServer: encryptedMessageToServerBase64,
      signature: nacl.util.encodeBase64(signatureToServer)
    })
  }).then(() => {
    console.log('Message sent');
  }).catch((err) => {
    console.error('Failed to send message: ' + err);
  });
}

const onEncryptClick = (e) => {
  e.stopPropagation();
  e.preventDefault();

  register()
    .then(verifyServer)
    .then(encryptAndSend)
    .catch(error => respond('Error:' + error));
}

const onDecryptClick = (e) => {
  e.stopPropagation();
  e.preventDefault();

  window.crypto.subtle.decrypt(
    { name: 'AES-GCM', iv: appState.nonceMessage},
    appState.aesKey, nacl.util.decodeBase64(textToDecrypt()),
  ).then((decryptedArrayBuffer) => {
    const decryptedBytes = new Uint8Array(decryptedArrayBuffer);
    const decryptedMessage = nacl.util.encodeUTF8(decryptedBytes);
    respondDecrypted(decryptedMessage);
  });
};

(function () {
  const btnEncrypt = document.getElementById('encrypt-btn');
  const btnDecrypt = document.getElementById('decrypt-btn');
  btnEncrypt.addEventListener('click', onEncryptClick, false);
  btnDecrypt.addEventListener('click', onDecryptClick, false);
}());
