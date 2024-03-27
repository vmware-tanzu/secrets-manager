package org.vsecm.task;

import org.vsecm.task.backoff.BackoffStrategy;

import java.util.concurrent.Callable;
import java.util.concurrent.Future;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import java.util.logging.Level;
import java.util.logging.Logger;

import static org.vsecm.service.FetchService.fetch;

/**
 * A task for fetching secrets from a specified URI, with retry logic based on a backoff strategy.
 * This task is designed to periodically fetch data, adjusting its retry interval based on the success
 * and failure outcomes of each attempt.
 */
public class FetchSecretsTask implements Task {

    private static final Logger LOGGER = Logger.getLogger(FetchSecretsTask.class.getName());
    private final ScheduledExecutorService scheduler;
    private final BackoffStrategy backoffStrategy;
    private long interval;
    private long successCount = 0;
    private long errorCount = 0;

    /**
     * Constructs a new {@code FetchSecretsTask} with the specified scheduler, backoff strategy, and initial interval.
     *
     * @param scheduler       The {@link ScheduledExecutorService} to schedule task execution.
     * @param backoffStrategy The {@link BackoffStrategy} to determine the interval between retries.
     * @param initialInterval The initial interval (in milliseconds) before the first execution or retry.
     */
    public FetchSecretsTask(ScheduledExecutorService scheduler, BackoffStrategy backoffStrategy, long initialInterval) {
        this.scheduler = scheduler;
        this.backoffStrategy = backoffStrategy;
        this.interval = initialInterval;
    }

    /**
     * Schedules and executes the task of fetching secrets from the provided URI. If the fetch fails,
     * the task will retry based on the backoff strategy provided during construction.
     *
     * @param secretUri The URI from which to fetch secrets.
     * @return A {@link Future} representing pending completion of the task, containing the fetched data as a {@link String}
     *         upon success, or {@code null} if an error occurs.
     */
    @Override
    public Future<String> execute(String secretUri) {
        Callable<String> task = () -> {
            try {
                String result = fetch(secretUri);
                successCount++;
                return result;
            } catch (Exception e) {
                errorCount++;
                interval = backoffStrategy.calculateBackoff(interval, successCount, errorCount);
                LOGGER.log(Level.WARNING, "Could not fetch secrets. Will retry in " + interval + "ms.", e);
                return null;
            }
        };
        return scheduler.schedule(task, interval, TimeUnit.MILLISECONDS);
    }
}