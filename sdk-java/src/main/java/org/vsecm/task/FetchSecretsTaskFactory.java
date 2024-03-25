package org.vsecm.task;

import org.vsecm.task.backoff.BackoffStrategy;
import org.vsecm.task.backoff.ExponentialBackoffStrategy;

import java.util.concurrent.ScheduledExecutorService;

/**
 * A factory for creating {@link FetchSecretsTask} instances. This factory configures the tasks
 * with an {@link ExponentialBackoffStrategy} and a predefined initial interval, encapsulating
 * the creation logic and dependencies of {@code FetchSecretsTask}.
 */
public class FetchSecretsTaskFactory implements TaskFactory {

    /**
     * Creates and returns a new instance of {@link FetchSecretsTask}, configured with
     * an {@link ExponentialBackoffStrategy} for retry logic and a predefined initial interval.
     * This method provides a convenient way to instantiate {@code FetchSecretsTask} with
     * recommended settings.
     *
     * @param scheduler The {@link ScheduledExecutorService} to be used by the created task for scheduling.
     * @return A new instance of {@link FetchSecretsTask}, ready to be executed.
     */
    @Override
    public Task createTask(ScheduledExecutorService scheduler) {
        BackoffStrategy backoffStrategy = new ExponentialBackoffStrategy();
        return new FetchSecretsTask(scheduler, backoffStrategy, 20000);
    }
}