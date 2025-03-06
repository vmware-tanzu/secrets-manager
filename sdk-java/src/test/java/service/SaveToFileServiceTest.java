package service;

import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.vsecm.service.SaveToFileService;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

import static org.junit.jupiter.api.Assertions.*;

class SaveToFileServiceTest {
    private Path tempFile;
    private Path tempDir;

    @BeforeEach
    void setUp() throws IOException {

        tempDir = Files.createTempDirectory("testDir");
        tempFile = tempDir.resolve("testFile.txt");
    }

    @AfterEach
    void tearDown() throws IOException {
        if (Files.exists(tempDir)) {
            Files.walk(tempDir)
                    .sorted((a, b) -> b.compareTo(a))
                    .forEach(path -> {
                        try {
                            Files.deleteIfExists(path);
                        } catch (IOException e) {
                            System.err.println("Failed to delete: " + path + " - " + e.getMessage());
                        }
                    });
        }
    }


    @Test
    void testSaveData_FileCreatedAndDataWritten() throws IOException {
        String testData = "Hello, World!";


        SaveToFileService.saveData(testData, tempFile.toString());


        assertTrue(Files.exists(tempFile), "Folder has been created");


        String content = Files.readString(tempFile);
        assertEquals(testData, content, "Folder's content is valid");
    }

    @Test
    void testSaveData_CreatesMissingDirectories() {
        Path nestedFile = tempDir.resolve("nested/directory/testFile.txt");


        SaveToFileService.saveData("Test", nestedFile.toString());


        assertTrue(Files.exists(nestedFile), "Folder has been created");
        assertTrue(Files.exists(nestedFile.getParent()), "Folder has been created");
    }

    @Test
    void testSaveData_InvalidPath_ShouldNotThrowException() {

        String invalidPath = "/root/protected/test.txt";


        assertDoesNotThrow(() -> SaveToFileService.saveData("Test", invalidPath));
    }
}
