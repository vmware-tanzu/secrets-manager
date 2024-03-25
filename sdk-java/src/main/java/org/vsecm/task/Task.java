package org.vsecm.task;

import java.util.concurrent.Future;

/**
 * Represents a general task that can be executed, typically involving asynchronous operations.
 * This interface defines a single method for executing a task, which is intended to perform
 * operations such as fetching or processing data from a specified URI and returning a future result.
 */
public interface Task {
    /**
     * Executes the task using the provided URI and returns a {@link Future} containing the result.
     * The implementation of this method should define the specific operation to be performed,
     * such as fetching data from the given URI. The method is designed to be asynchronous, allowing
     * the calling thread to continue execution while the task is processed in the background.
     *
     * @param secretUri The URI or resource identifier that the task will operate on or fetch data from.
     * @return A {@link Future} representing pending completion of the task. The Future's {@code get} method
     *         will return the task's result as a {@link String} upon successful completion, or {@code null}
     *         if the operation fails or does not produce a result.
     */
    Future<String> execute(String secretUri);
}