package org.vsecm.task.backoff;

/**
 * Defines a strategy for calculating backoff intervals. Implementations of this interface
 * can be used to determine the delay before retrying an operation, typically after a failure.
 */
public interface BackoffStrategy {

    /**
     * Calculates the backoff interval based on the number of successes, the number of errors,
     * and the current interval.
     *
     * @param interval The base interval or the current interval before applying the backoff strategy.
     * @param successCount The number of successful operations.
     * @param errorCount The number of failed operations.
     * @return The calculated backoff interval.
     */
    long calculateBackoff(long interval, long successCount, long errorCount);
}

