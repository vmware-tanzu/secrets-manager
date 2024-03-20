package org.wsecm;

import com.sun.net.httpserver.HttpServer;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.util.logging.Logger;


public class HealthCheckServer {

    private static final Logger LOGGER = Logger.getLogger(HealthCheckServer.class.getName());

    public static void startServer() {
        try {
            HttpServer server = HttpServer.create(new InetSocketAddress(8080), 0);
            server.createContext("/", exchange -> {
                String response = "OK";
                exchange.sendResponseHeaders(200, response.getBytes().length);
                OutputStream os = exchange.getResponseBody();
                os.write(response.getBytes());
                os.close();
            });
            new Thread(server::start).start();
            LOGGER.info("Health check server sta1rted on port 8080");
        } catch (IOException e) {
            LOGGER.info("Failed to start health check server");
        }
    }
}
