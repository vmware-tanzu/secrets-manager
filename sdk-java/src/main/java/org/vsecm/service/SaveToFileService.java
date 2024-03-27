/**
 * Provides functionality to save data to a file, creating the necessary directories if they do not exist.
 * This service is designed to facilitate the saving of data to the file system, ensuring that the file
 * path's parent directories are created if necessary.
 */
package org.vsecm.service;

import java.io.BufferedWriter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.logging.Level;
import java.util.logging.Logger;

/**
 * Utility class for saving data to a file. It ensures that the parent directories of the specified
 * file path are created before writing the data. This class cannot be instantiated.
 */
public class SaveToFileService {
    private static final Logger LOGGER = Logger.getLogger(SaveToFileService.class.getName());

    /**
     * Saves the provided data to a file at the specified path. If the file's parent directories do not exist,
     * they are created. If writing to the file fails, an error is logged.
     *
     * @param data The string data to be saved to the file.
     * @param pathStr The file system path where the data should be saved. Must be a valid path string.
     */
    public static void saveData(String data, String pathStr) {
        Path path = Paths.get(pathStr);

        // Attempt to create the directory(ies) if they don't exist
        try {
            Files.createDirectories(path.getParent());
        } catch (IOException e) {
            LOGGER.log(Level.SEVERE, "Failed to create directory for secrets", e);
            return;
        }

        // Attempt to write the data to the file
        try (BufferedWriter writer = Files.newBufferedWriter(path)) {
            writer.write(data);
        } catch (IOException e) {
            LOGGER.log(Level.SEVERE, "Error saving data to " + pathStr, e);
        }
    }
}
