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
     * Fetches data as a {@link String} from the specified URI using an HTTP GET request.
     * This method utilizes a SPIFFE-enabled HTTP client for secure communication.
     *
     * @param secretUri The URI from which to fetch data. Must be a valid URI string.
     * @return The body of the HTTP response as a {@link String}. Returns an empty string if the request fails.
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
