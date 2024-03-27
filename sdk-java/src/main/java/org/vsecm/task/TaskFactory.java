package org.vsecm.task;

import java.util.concurrent.ScheduledExecutorService;

/**
 * Defines a factory interface for creating instances of {@link Task}. This interface ensures
 * that implementing classes provide a method to instantiate tasks, facilitating the creation
 * of tasks with specific configurations or dependencies, such as a {@link ScheduledExecutorService}.
 */
public interface TaskFactory {

    /**
     * Creates and returns a new {@link Task} instance, configured with the provided
     * {@link ScheduledExecutorService}. This method allows for the dynamic creation of tasks,
     * providing flexibility in how tasks are instantiated and scheduled.
     *
     * @param scheduler The {@link ScheduledExecutorService} that the created task will use
     *                  for scheduling its execution. This service controls the timing and
     *                  concurrency of the task execution.
     * @return A new instance of a {@link Task}, ready to be executed with the given scheduler.
     */
    Task createTask(ScheduledExecutorService scheduler);
}

