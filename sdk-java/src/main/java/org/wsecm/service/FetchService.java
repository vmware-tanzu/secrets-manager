package org.wsecm.service;

import org.wsecm.client.WsecmHttpClient;
import org.wsecm.fileops.SaveToFile;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.logging.Level;
import java.util.logging.Logger;

public class FetchService {

    private final WsecmHttpClient wsecmHttpClient;
    private static final Logger LOGGER = Logger.getLogger(FetchService.class.getName());
    private static final String SECRET_URI = "https://vsecm-safe.vsecm-system.svc.cluster.local:8443/sentinel/v1/secrets?reveal=true";
    private static final String DEFAULT_SECRETS_PATH = "/opt/vsecm/secrets.json";

    public FetchService() {
        this.wsecmHttpClient = new WsecmHttpClient();
    }

    public void fetch() {
        try {
            HttpRequest request = HttpRequest.newBuilder(new URI(SECRET_URI)).GET().build();
            HttpResponse<String> response = wsecmHttpClient.client().send(request, HttpResponse.BodyHandlers.ofString());
            logResponse(response);
        } catch (URISyntaxException | IOException | InterruptedException e) {
            LOGGER.log(Level.SEVERE, "Failed processes when request to URI : " + SECRET_URI, e);
        }
    }

    private static void logResponse(HttpResponse<String> response) {
        LOGGER.info("Response status code: " + response.statusCode());

        if (response.statusCode() == 200) {
            SaveToFile.saveData(response.body(), DEFAULT_SECRETS_PATH);
        }

        LOGGER.info("Response body: " + response.body());
    }
}
