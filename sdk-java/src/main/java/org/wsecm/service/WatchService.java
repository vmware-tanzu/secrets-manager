package org.wsecm.service;

import org.wsecm.task.Task;
import org.wsecm.task.TaskFactory;

import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;

public class WatchService {
    private static WatchService instance;
    private final ScheduledExecutorService scheduler;

    private WatchService() {
        this.scheduler = Executors.newScheduledThreadPool(1);
    }

    public static synchronized WatchService getInstance() {
        if (instance == null) {
            instance = new WatchService();
        }
        return instance;
    }

    public void scheduleTask(TaskFactory taskFactory) {
        Task task = taskFactory.createTask(scheduler);
        task.execute();
    }
}