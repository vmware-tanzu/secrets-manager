/**
 * Provides a scheduling service for executing tasks that monitor or interact with specified resources, such as URIs.
 * Utilizes a {@link TaskFactory} to create tasks that can be scheduled for execution.
 * This service is designed to be flexible, allowing for different types of tasks to be executed based on the implementation
 * provided by the {@link TaskFactory}.
 */
package org.vsecm.service;

import org.vsecm.task.Task;
import org.vsecm.task.TaskFactory;

import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.concurrent.ScheduledExecutorService;

/**
 * A utility class for scheduling and executing tasks. This class cannot be instantiated and is used
 * to schedule tasks provided by a {@link TaskFactory}. The tasks can perform any operations defined
 * by their implementation, such as monitoring changes to a resource or processing data.
 */
public final class WatchService {
    private WatchService() {}

    /**
     * Schedules and executes a task created by the provided {@link TaskFactory}. The task is executed
     * using a single-threaded {@link ScheduledExecutorService}.
     *
     * @param taskFactory The factory to create a task for execution.
     * @param secretUri The URI or resource identifier that the task will operate on or monitor.
     * @return A {@link Future} representing pending completion of the task. The future's get method will
     *         return the task's result upon completion.
     */
    public static Future<String> watch(TaskFactory taskFactory, String secretUri) {
        ScheduledExecutorService scheduler = Executors.newScheduledThreadPool(1);
        Task task = taskFactory.createTask(scheduler);
        return task.execute(secretUri);
    }
}