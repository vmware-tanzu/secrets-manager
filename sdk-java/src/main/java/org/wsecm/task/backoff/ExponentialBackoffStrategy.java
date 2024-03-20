package org.wsecm.task.backoff;

import static java.lang.Math.abs;

public class ExponentialBackoffStrategy implements BackoffStrategy {
    @Override
    public long calculateBackoff(long interval, long successCount, long errorCount) {
        long delta = abs(successCount - errorCount);

        if (delta > successCount){
            return interval * 2;
        }
        return interval;
    }
}

