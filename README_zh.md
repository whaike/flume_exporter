# flume_exporter
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

### 新分支说明
- feature/read-flume-process
    主要作用：通过在Linux主机上执行`ps -ef | grep Dflume.monitoring.port`命令获取当前正在运行的有监控的flume进程，
    然后自动把他们监控起来，不需要每次启停一些程序后还要修改配置文件

- feature/configreturn
    主要作用：加一个web server ，通过http返回所有flume的配置，以便于汇总分析所有已启动的flume.
  
    /conf -> 返回所有的配置文件
    ```
  [{
  "Name": "agent_log_aly_oss",
  "Port": "9615",
  "ConfigName": "/opt/apache-flume-1.9.0-bin/conf/sdklog-aly/oss-sdklog-aly-3.conf",
  "ConfigDetail": ""
  },
  {}]
  ```
    ConfigDetail是配置文件的字符串形式，经过了base64编码

    /fuck -> 返回配置文件的主要分析结果 

    ```
      [{
            "topic": "",
            "group": "",
            "path": "",
            "logExample": ""
        },
        {
            "topic": "scene-server-log-new",
            "group": "flume-scene-server-log-group-aly-oss",
            "path": "oss://geek-hadoop/log/geek/scenelog/%{geekYear}/%{geekMonth}/%{geekDate}/%{geekHour}",
            "logExample": "log.xxx..snappy"
        }
        ]
    ```
之所以保留空，是想着可以通过返回直接判断解析过程是否存在异常，也就是flume运行数量和解析回来的数量是否一致。单独添加字段麻烦

另外：
* 多个flume主机需要自行聚合数据
* 本打算做成自动监听进程数的，但后来觉得，自动更新监控数不利于监控，如果某个flume挂掉，原来的监控指标可以看出来，但是如果没了，处理起来就没那么优雅了，所以保留每次更改flume重启一次flume_exporter的习惯。