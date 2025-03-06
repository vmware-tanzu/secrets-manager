/**
 * Provides an HTTP client configured with SPIFFE based SSL context for secure communication.
 * This client is specifically designed for environments that utilize SPIFFE for workload identification
 * and secure communication.
 *
 * <p>This implementation leverages the SPIFFE Workload API to fetch X.509 SVIDs for establishing TLS connections,
 * ensuring that the communication is both secure and aligned with the SPIFFE standards.</p>
 */
package org.vsecm.client;

import io.spiffe.exception.SocketEndpointAddressException;
import io.spiffe.exception.X509SourceException;
import io.spiffe.provider.SpiffeSslContextFactory;
import io.spiffe.workloadapi.DefaultX509Source;
import org.vsecm.exception.VSecMHttpClientException;

import javax.net.ssl.SSLContext;
import java.net.http.HttpClient;
import java.security.KeyManagementException;
import java.security.NoSuchAlgorithmException;
import java.util.logging.Level;
import java.util.logging.Logger;

/**
 * Represents a HTTP client that is secured with SPIFFE-based mutual TLS,
 * utilizing the SPIFFE Workload API for fetching SVIDs.
 */
public class VSecMHttpClient {
    private static final Logger LOGGER = Logger.getLogger(VSecMHttpClient.class.getName());

    /**
     * The path to the SPIFFE socket used to communicate with the SPIFFE Workload API.
     * This socket provides access to the X.509 SVID (SPIFFE Verifiable Identity Document)
     * used for mutual TLS authentication.
     */
    private static final String SPIFFE_SOCKET_PATH = "unix:///spire-agent-socket/spire-agent.sock";

    /**
     * Creates and configures an {@link HttpClient} instance with SPIFFE-based mutual TLS.
     * This client is used for secure communication with SPIFFE-enabled services.
     *
     * @return An instance of {@link HttpClient} configured with mutual TLS using SPIFFE credentials.
     *
     * @throws VSecMHttpClientException.SocketPathError If the SPIFFE socket path is inaccessible.
     * @throws VSecMHttpClientException.X509FetchError If fetching X.509 SVIDs fails.
     * @throws VSecMHttpClientException.SSLContextError If SSLContext configuration fails.
     * @throws RuntimeException If an unknown error occurs during client initialization.
     *
     * @see #configureSSLContext()
     */
    public HttpClient client() {
        try {
            SSLContext sslContext = configureSSLContext();
            return HttpClient.newBuilder().sslContext(sslContext).build();
        } catch (Exception e) {
            LOGGER.log(Level.SEVERE, "Failed to fetch secrets", e);

            if (e instanceof SocketEndpointAddressException) {
                throw VSecMHttpClientException.socketPathError("SPIFFE socket path is inaccessible: " + e.getMessage());
            } else if (e instanceof X509SourceException) {
                throw VSecMHttpClientException.x509FetchError("Failed to fetch X.509 SVIDs: " + e.getMessage());
            } else if (e instanceof NoSuchAlgorithmException || e instanceof KeyManagementException) {
                throw VSecMHttpClientException.sslContextError("SSLContext configuration failed: " + e.getMessage());
            } else {
                throw new RuntimeException("Unknown error occurred: " + e.getMessage(), e);
            }
        }
    }

    /**
     * Configures and returns an {@link SSLContext} instance using SPIFFE credentials.
     * This method creates a secure SSL context that is configured with X.509 SVIDs
     * obtained from the SPIFFE Workload API to enable mutual TLS for communication.
     *
     * @return A configured {@link SSLContext} that supports mutual TLS using SPIFFE identities.
     * @throws SocketEndpointAddressException If there is an issue with the SPIFFE socket path,
     *                                         typically indicating that the SPIFFE Workload API is inaccessible.
     * @throws X509SourceException If there is an error fetching or processing the X.509 SVIDs from the SPIFFE Workload API.
     * @throws NoSuchAlgorithmException If the SSLContext cannot be instantiated due to a missing or unsupported algorithm.
     * @throws KeyManagementException If there is an issue initializing the SSLContext, typically indicating a problem
     *                                with key management or protocol setup.
     * @see DefaultX509Source
     * @see SpiffeSslContextFactory
     */

    private SSLContext configureSSLContext() throws SocketEndpointAddressException, X509SourceException, NoSuchAlgorithmException, KeyManagementException {
        DefaultX509Source.X509SourceOptions sourceOptions = DefaultX509Source.X509SourceOptions
                .builder()
                .spiffeSocketPath(SPIFFE_SOCKET_PATH)
                .build();

        DefaultX509Source x509Source = DefaultX509Source.newSource(sourceOptions);
        LOGGER.info("SPIFFE_ID: " + x509Source.getX509Svid().getSpiffeId());

        SpiffeSslContextFactory.SslContextOptions sslContextOptions = SpiffeSslContextFactory.SslContextOptions
                .builder()
                .acceptAnySpiffeId()
                .x509Source(x509Source)
                .build();

        return SpiffeSslContextFactory.getSslContext(sslContextOptions);
    }

}