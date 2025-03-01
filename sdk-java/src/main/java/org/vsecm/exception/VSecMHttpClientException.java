package org.vsecm.exception;



/**
 * Custom exception class for handling errors in the VSecMHttpClient class.
 * Each error code corresponds to a specific type of failure encountered during
 * the HTTP client setup and SSL context configuration.
 */

public class VSecMHttpClientException extends RuntimeException {

    // Error codes for different failure cases
    public static final int ERROR_CODE_SOCKET_PATH = 1000;
    public static final int ERROR_CODE_X509_FETCH = 1001;
    public static final int ERROR_CODE_SSL_CONTEXT = 1002;
    public static final int ERROR_CODE_KEY_MANAGEMENT = 1003;

    private final int errorCode;

    /**
     * Constructor for the custom exception.
     *
     * @param errorCode The specific error code representing the type of failure.
     * @param message The detailed error message.
     */


    public VSecMHttpClientException(int errorCode, String message) {
        super(message);
        this.errorCode = errorCode;
    }


    /**
     * Returns the error code associated with this exception.
     *
     * @return The error code.
     */

    public int getErrorCode() {
        return errorCode;
    }


    @Override
    public String toString() {
        return "VSecMHttpClientException{" +
                "errorCode=" + errorCode +
                ", message=" + getMessage() +
                '}';
    }

    // Convenience methods for specific error codes
    public static VSecMHttpClientException socketPathError(String message) {
        return new VSecMHttpClientException(ERROR_CODE_SOCKET_PATH, message);
    }


    public static VSecMHttpClientException x509FetchError(String message) {
        return new VSecMHttpClientException(ERROR_CODE_X509_FETCH, message);
    }


    public static VSecMHttpClientException sslContextError(String message) {
        return new VSecMHttpClientException(ERROR_CODE_SSL_CONTEXT, message);
    }


    public static VSecMHttpClientException keyManagementError(String message) {
        return new VSecMHttpClientException(ERROR_CODE_KEY_MANAGEMENT, message);
    }

}