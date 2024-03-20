package org.wsecm.task.backoff;

public interface BackoffStrategy {
    long calculateBackoff(long interval, long successCount, long errorCount);
}

