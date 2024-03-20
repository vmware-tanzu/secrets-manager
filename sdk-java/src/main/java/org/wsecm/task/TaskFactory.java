package org.wsecm.task;

import java.util.concurrent.ScheduledExecutorService;

public interface TaskFactory {
    Task createTask(ScheduledExecutorService scheduler);
}

