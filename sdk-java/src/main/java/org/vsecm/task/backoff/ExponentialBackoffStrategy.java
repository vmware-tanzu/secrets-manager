package org.vsecm.task.backoff;

import static java.lang.Math.abs;

/**
 * Implements an exponential backoff strategy that increases the backoff interval based on the difference
 * between the number of successes and errors. If the error count is greater than the success count, the interval
 * is doubled to allow more time before the next retry. Otherwise, the interval remains unchanged.
 */
public class ExponentialBackoffStrategy implements BackoffStrategy {

    /**
     * Calculates the backoff interval using an exponential strategy. The interval is doubled when
     * the number of errors exceeds the number of successes.
     *
     * @param interval The base interval or the current interval before applying the backoff strategy.
     * @param successCount The number of successful operations.
     * @param errorCount The number of failed operations.
     * @return The calculated backoff interval. Doubles the interval if errors exceed successes, otherwise
     *         returns the original interval.
     */
    @Override
    public long calculateBackoff(long interval, long successCount, long errorCount) {
        long delta = abs(successCount - errorCount);

        if (delta > successCount){
            return interval * 2;
        }
        return interval;
    }
}

