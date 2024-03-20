package org.wsecm.task;

import org.wsecm.service.FetchService;
import org.wsecm.task.backoff.BackoffStrategy;

import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import java.util.logging.Level;
import java.util.logging.Logger;

public class FetchSecretsTask implements Task {

    private static final Logger LOGGER = Logger.getLogger(FetchSecretsTask.class.getName());
    private final ScheduledExecutorService scheduler;
    private final BackoffStrategy backoffStrategy;
    private final FetchService fetchService;
    private long interval;
    private long successCount = 0;
    private long errorCount = 0;

    public FetchSecretsTask(ScheduledExecutorService scheduler, BackoffStrategy backoffStrategy, long initialInterval) {
        this.scheduler = scheduler;
        this.backoffStrategy = backoffStrategy;
        this.interval = initialInterval;
        fetchService = new FetchService();
    }

    @Override
    public void execute() {
        scheduler.schedule(this::run, interval, TimeUnit.MILLISECONDS);
    }

    private void run() {
        try {
            fetchService.fetch();
            successCount++;
        } catch (Exception e) {
            errorCount++;
            interval = backoffStrategy.calculateBackoff(interval, successCount, errorCount);
            LOGGER.log(Level.WARNING, "Could not fetch secrets. Will retry in " + interval + "ms.");
        } finally {
            execute();
        }
    }
}