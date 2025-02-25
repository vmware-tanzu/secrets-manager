/**
 * A service class for fetching data from specified URIs using a SPIFFE-enabled HTTP client.
 * This service leverages {@link org.vsecm.client.VSecMHttpClient} for secure HTTP requests,
 * ensuring communication adheres to the SPIFFE standard for secure and authenticated connections.
 */
package org.vsecm.service;

import org.vsecm.client.VSecMHttpClient;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.logging.Level;
import java.util.logging.Logger;

/**
 * Provides static utility methods for fetching data from a specified URI.
 * This class cannot be instantiated and serves purely as a utility class.
 */
public final class FetchService {
    private static final Logger LOGGER = Logger.getLogger(FetchService.class.getName());

    private FetchService(){}

    /**
     * Fetches data from the specified URI using an HTTP GET request.
     * The request is sent over a secure channel using a SPIFFE-enabled HTTP client,
     * which ensures mutual TLS authentication as part of the communication process.
     *
     * <p>This method constructs an HTTP GET request to the provided URI, sends the request
     * using the {@link VSecMHttpClient} (which handles SPIFFE-based secure communication),
     * and returns the response body as a {@link String}. If any errors occur during
     * the request (such as invalid URI, network issues, or interrupted request), an empty string
     * is returned, and the error is logged.</p>
     *
     * @param secretUri The URI from which to fetch the data. This must be a well-formed URI string
     *                  that points to a resource to be fetched.
     *                  Example: "https://example.com/data".
     * @return The body of the HTTP response as a {@link String}. If the request fails or an error occurs,
     *         an empty string is returned.
     * @throws IllegalArgumentException if the URI is not valid or cannot be parsed.
     * @throws URISyntaxException if the provided URI is malformed or invalid.
     * @throws IOException if an I/O error occurs during communication with the server.
     * @throws InterruptedException if the request is interrupted during execution.
     *
     * @see VSecMHttpClient
     */
    public static String fetch(String secretUri) {
        try {
            HttpRequest request = HttpRequest.newBuilder(new URI(secretUri)).GET().build();
            VSecMHttpClient vSecMHttpClient = new VSecMHttpClient();
            HttpResponse<String> response = vSecMHttpClient.client().send(request, HttpResponse.BodyHandlers.ofString());
            LOGGER.info("Response status code: " + response.statusCode());
            LOGGER.info("Response body: " + response.body());
            return response.body();
        } catch (URISyntaxException | IOException | InterruptedException e) {
            LOGGER.log(Level.SEVERE, "Failed processes when request to URI : " + secretUri, e);
            return "";
        }
    }
}
