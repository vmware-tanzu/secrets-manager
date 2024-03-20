package org.wsecm.client;

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

public class WsecmHttpClient {
    private static final Logger LOGGER = Logger.getLogger(WsecmHttpClient.class.getName());
    private static final String SPIFFE_SOCKET_PATH = "unix:///spire-agent-socket/agent.sock";

    public HttpClient client() {
        try {
            SSLContext sslContext = configureSSLContext();
            return HttpClient.newBuilder().sslContext(sslContext).build();
        } catch (Exception e) {
            LOGGER.log(Level.SEVERE, "Failed to fetch secrets", e);
            throw new RuntimeException(e);
        }
    }

    private static SSLContext configureSSLContext() throws SocketEndpointAddressException, X509SourceException, NoSuchAlgorithmException, KeyManagementException {
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