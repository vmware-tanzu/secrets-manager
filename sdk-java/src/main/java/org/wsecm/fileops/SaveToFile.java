package org.wsecm.fileops;

import java.io.BufferedWriter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.logging.Level;
import java.util.logging.Logger;

public class SaveToFile {
    private static final Logger LOGGER = Logger.getLogger(SaveToFile.class.getName());

    public static void saveData(String data, String pathStr) {
        Path path = Paths.get(pathStr);

        try {
            Files.createDirectories(path.getParent());
        } catch (IOException e) {
            LOGGER.log(Level.SEVERE, "Failed to create directory for secrets", e);
            return;
        }

        try (BufferedWriter writer = Files.newBufferedWriter(path)) {
            writer.write(data);
        } catch (IOException e) {
            LOGGER.log(Level.SEVERE, "Error saving data to " + pathStr, e);
        }
    }
}
