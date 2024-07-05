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
    private static final String SPIFFE_SOCKET_PATH = "unix:///spire-agent-socket/spire-agent.sock";

    /**
     * Creates an instance of {@link HttpClient} with SPIFFE-based SSL context.
     * This client can be used for secure HTTP communication.
     *
     * @return A configured {@link HttpClient} instance ready for secure communication.
     * @throws RuntimeException if there's an issue configuring the SSL context,
     *                          encapsulating any underlying exceptions.
     */
    public HttpClient client() {
        try {
            SSLContext sslContext = configureSSLContext();
            return HttpClient.newBuilder().sslContext(sslContext).build();
        } catch (Exception e) {
            LOGGER.log(Level.SEVERE, "Failed to fetch secrets", e);
            throw new RuntimeException(e);
        }
    }

    /**
     * Configures and returns an {@link SSLContext} suitable for SPIFFE based secure communication.
     *
     * @return An {@link SSLContext} configured with SPIFFE X.509 SVIDs for mutual TLS.
     * @throws SocketEndpointAddressException If the SPIFFE Workload API socket endpoint address is incorrect.
     * @throws X509SourceException If there's an issue fetching or processing the X.509 SVIDs.
     * @throws NoSuchAlgorithmException If the SSL context cannot be instantiated due to a missing algorithm.
     * @throws KeyManagementException If there's an issue initializing the {@link SSLContext} with SPIFFE SVIDs.
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