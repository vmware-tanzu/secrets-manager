package service;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.vsecm.service.WatchService;
import org.vsecm.task.Task;
import org.vsecm.task.TaskFactory;


import java.util.concurrent.*;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;
import static org.awaitility.Awaitility.await;
import java.time.Duration;

class WatchServiceTest {
    private TaskFactory mockTaskFactory;
    private Task mockTask;
    private ScheduledExecutorService scheduler;

    @BeforeEach
    void setUp() {
        mockTaskFactory = mock(TaskFactory.class);
        mockTask = mock(Task.class);
        scheduler = Executors.newSingleThreadScheduledExecutor();
    }

    @Test
    void testWatch_ShouldReturnFutureWithExpectedResult() throws Exception {
        // GIVEN
        String expectedResponse = "Mock Response";
        String testUri = "https://example.com/secret";


        CompletableFuture<String> fakeFuture = CompletableFuture.completedFuture(expectedResponse);


        when(mockTaskFactory.createTask(any())).thenReturn(mockTask);
        when(mockTask.execute(testUri)).thenReturn(fakeFuture);

        // WHEN
        Future<String> resultFuture = WatchService.watch(mockTaskFactory, testUri);

        // THEN
        await().atMost(Duration.ofSeconds(2)).until(resultFuture::isDone);


        String actualResult = resultFuture.get();
        assertEquals(expectedResponse, actualResult, "Unexpected condition!");


        verify(mockTaskFactory).createTask(any());
        verify(mockTask).execute(testUri);
    }
}
