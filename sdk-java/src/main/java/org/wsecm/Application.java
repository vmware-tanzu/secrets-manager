package org.wsecm;

import org.wsecm.service.WatchService;
import org.wsecm.task.FetchSecretsTaskFactory;
import org.wsecm.task.TaskFactory;

import static org.wsecm.HealthCheckServer.startServer;

public class Application {
    public static void main(String[] args) {
        startServer();
        WatchService watchService = WatchService.getInstance();
        TaskFactory taskFactory = new FetchSecretsTaskFactory();
        watchService.scheduleTask(taskFactory);
    }
}
