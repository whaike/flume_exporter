# flume_exporter

[中文说明](https://github.com/whaike/flume_exporter/blob/master/README_zh.md)

Prometheus exporter for flume.

To run it:

```bash
make build

./flume_exporter [flags]
```

Help on flags:
```bash
./flume_exporter --help
```

Configuration: config.yml
```
agents:
- name: "flume-agents"
  enabled: true
# multiple urls can be separated by ,  
  urls: ["http://localhost:36001/metrics"]   
```

### Using Docker
Default
```
docker run -d -p 9360:9360 zhijunwoo/flume_exporter:latest
```

Specified configuration
```
docker run -d \
    -p 9360:9360 \
    -v `pwd`/config.yml:/etc/flume_exporter/config.yml \
    -name flume_exporter \
    zhijunwoo/flume_exporter:latest
```

### monitoring metrics
#### Sources
- AppendAcceptedCount
- AppendBatchAcceptedCount
- AppendBatchReceivedCount
- AppendReceivedCount
- ChannelWriteFail
- EventAcceptedCount
- EventReadFail
- EventReceivedCount
- GenericProcessingFail
- KafkaCommitTimer
- KafkaEmptyCount
- KafkaEventGetTimer
- OpenConnectionCount

#### Channels
- ChannelCapacity
- ChannelSize
- CheckpointBackupWriteErrorCount
- CheckpointWriteErrorCount
- EventPutAttemptCount
- EventPutErrorCount
- EventPutSuccessCount
- EventTakeAttemptCount
- EventTakeErrorCount
- EventTakeSuccessCount
- KafkaCommitTimer
- KafkaEventGetTimer
- KafkaEventSendTimer
- Open
- RollbackCounter
- Unhealthy

- RollbackCount
- ChannelFillPercentage

#### Sinks
- BatchCompleteCount
- BatchEmptyCount
- BatchUnderflowCount
- ChannelReadFail
- ConnectionClosedCount
- ConnectionCreatedCount
- ConnectionFailedCount
- EventDrainAttemptCount
- EventDrainSuccessCount
- EventWriteFail
- KafkaEventSendTimer
- RollbackCount

#### Grafana Dashboard
Grafana Dashboard ID: 10736  
name: Flume Exporter Metrics Overview For Prometheus
For details of the dashboard please see [Flume Exporter Metrics](https://grafana.com/grafana/dashboards/10736)

### New Branch
    [中文说明](https://github.com/whaike/flume_exporter/blob/master/README_zh.md)
