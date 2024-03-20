package org.wsecm.task;

import org.wsecm.task.backoff.BackoffStrategy;
import org.wsecm.task.backoff.ExponentialBackoffStrategy;

import java.util.concurrent.ScheduledExecutorService;

public class FetchSecretsTaskFactory implements TaskFactory {
    @Override
    public Task createTask(ScheduledExecutorService scheduler) {
        BackoffStrategy backoffStrategy = new ExponentialBackoffStrategy();
        return new FetchSecretsTask(scheduler, backoffStrategy, 20000);
    }
}