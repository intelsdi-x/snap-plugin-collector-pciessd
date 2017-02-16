# SNAP collector plugin - PCIe SSD

This plugin collects PCIe SSD metrics. The implementation is based on the NVMe specification and Intel(R) Solid State Drive DC P3700 Series.
Currently it works only on Intel(R) Solid State State Drive DC P3700 Series. To support more SSD model, please check the `./conf/config.json` or contact [me](https://github.com/WangJialei-A)

1. [Getting Started](#getting-started)
	* [Installation](#installation)
	* [Configuration and Usage](#configuration-and-usage)
2. [Documentataion](#documentation)
3. [License](#license)
4. [Acknowledgements](#acknowledgements)

## Getting Started

### Operating systems
* Linux/amd64

### Installation

Clone repository into `$GOPATH/src/github.com/intelsdi-x/`
```
$ git clone https://github.com/intelsdi-x/snap-plugin-collector-pciessd.git
```
Build the Snap PCIe SSD plugin by running make within the cloned repository:
```
$ make
```
It may take a while to pull dependencies if you haven't had them ready'
This builds the plugin in `./build/`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Load the plugin and create a task, see example in `./example`

`./conf/config.json` describes how to parse the NVMe result and generate the metrics.

For snap configuration
* conf - defines the path of configuration file.
See more details in `example/snap-config.json` and `./conf/config.json`

## Documentation 

Metrics are all defined in `./conf/config.json`. Metrics listed below are for the default configuration.

```
/intel/pciessd/drive-name/CriticalWarning
/intel/pciessd/drive-name/Temperature
/intel/pciessd/drive-name/AvailableSpare
/intel/pciessd/drive-name/AvailableSpareThreshold
/intel/pciessd/drive-name/PercentageUsedEstimate
/intel/pciessd/drive-name/DataUnitsRead
/intel/pciessd/drive-name/DataUnitsWrite
/intel/pciessd/drive-name/HostReadCommands
/intel/pciessd/drive-name/HostWriteCommands
/intel/pciessd/drive-name/ControllerBusyTime
/intel/pciessd/drive-name/PowerCycles
/intel/pciessd/drive-name/PowerOnHours
/intel/pciessd/drive-name/UnsafeShutdowns
/intel/pciessd/drive-name/MediaErrors
/intel/pciessd/drive-name/NumberOfErrorInformationLogEntries

/intel/pciessd/drive-name/ProgramFailCount/SmartID
/intel/pciessd/drive-name/ProgramFailCount/NormalizedValue
/intel/pciessd/drive-name/ProgramFailCount/CurrentRawValue
/intel/pciessd/drive-name/EraseFailCount/SmartID
/intel/pciessd/drive-name/EraseFailCount/NormalizedValue
/intel/pciessd/drive-name/EraseFailCount/CurrentRawValue
/intel/pciessd/drive-name/WearLevelingCount/SmartID
/intel/pciessd/drive-name/WearLevelingCount/NormalizedValue
/intel/pciessd/drive-name/WearLevelingCount/Min
/intel/pciessd/drive-name/WearLevelingCount/Avg
/intel/pciessd/drive-name/WearLevelingCount/Max
/intel/pciessd/drive-name/EndToEndErrorDetectionCount/SmartID
/intel/pciessd/drive-name/EndToEndErrorDetectionCount/NormalizedValue
/intel/pciessd/drive-name/EndToEndErrorDetectionCount/CurrentRawValue
/intel/pciessd/drive-name/CRCErrorCount/SmartID
/intel/pciessd/drive-name/CRCErrorCount/NormalizedValue
/intel/pciessd/drive-name/CRCErrorCount/CurrentRawValue
/intel/pciessd/drive-name/TimedWorkloadMediaWear/SmartID
/intel/pciessd/drive-name/TimedWorkloadMediaWear/NormalizedValue
/intel/pciessd/drive-name/TimedWorkloadMediaWear/CurrentRawValue
/intel/pciessd/drive-name/TimedWorkloadHostReadPercentage/SmartID
/intel/pciessd/drive-name/TimedWorkloadHostReadPercentage/NormalizedValue
/intel/pciessd/drive-name/TimedWorkloadHostReadPercentage/CurrentRawValue
/intel/pciessd/drive-name/TimedWorkloadTimer/SmartID
/intel/pciessd/drive-name/TimedWorkloadTimer/NormalizedValue
/intel/pciessd/drive-name/TimedWorkloadTimer/CurrentRawValue
/intel/pciessd/drive-name/ThermalThrottleStatus/SmartID
/intel/pciessd/drive-name/ThermalThrottleStatus/NormalizedValue
/intel/pciessd/drive-name/ThermalThrottleStatus/Percentage
/intel/pciessd/drive-name/ThermalThrottleStatus/EventCount
/intel/pciessd/drive-name/RetryBufferOverflowCounter/SmartID
/intel/pciessd/drive-name/RetryBufferOverflowCounter/NormalizedValue
/intel/pciessd/drive-name/RetryBufferOverflowCounter/CurrentRawValue
/intel/pciessd/drive-name/PLLLockLossCount/SmartID
/intel/pciessd/drive-name/PLLLockLossCount/NormalizedValue
/intel/pciessd/drive-name/PLLLockLossCount/CurrentRawValue
/intel/pciessd/drive-name/NANDBytesWritten/SmartID
/intel/pciessd/drive-name/NANDBytesWritten/NormalizedValue
/intel/pciessd/drive-name/NANDBytesWritten/CurrentValue
/intel/pciessd/drive-name/HostBytesWritten/SmartID
/intel/pciessd/drive-name/HostBytesWritten/NormalizedValue
/intel/pciessd/drive-name/HostBytesWritten/CurrentValue

/intel/pciessd/drive-name/CurrentTemperature
/intel/pciessd/drive-name/OvertempShutdownFlagForLifeCriticalComponentTemperature
/intel/pciessd/drive-name/OvertempShutdownFlagForLastCriticalComponentTemperature
/intel/pciessd/drive-name/HighestTemperature
/intel/pciessd/drive-name/LowestTemperature
/intel/pciessd/drive-name/SpecifiedMaximumOperatingTemperature
/intel/pciessd/drive-name/SpecifiedMinimumOperatingTemperature
/intel/pciessd/drive-name/EstimatedOffset

/intel/pciessd/drive-name/MarketingName

/intel/pciessd/drive-name/ReadCommandLatency/MajorRevision
/intel/pciessd/drive-name/ReadCommandLatency/MinorRevision
/intel/pciessd/drive-name/ReadCommandLatency/0us
/intel/pciessd/drive-name/ReadCommandLatency/32us
/intel/pciessd/drive-name/ReadCommandLatency/64us
/intel/pciessd/drive-name/ReadCommandLatency/96us
/intel/pciessd/drive-name/ReadCommandLatency/128us
/intel/pciessd/drive-name/ReadCommandLatency/160us
/intel/pciessd/drive-name/ReadCommandLatency/192us
/intel/pciessd/drive-name/ReadCommandLatency/224us
/intel/pciessd/drive-name/ReadCommandLatency/256us
/intel/pciessd/drive-name/ReadCommandLatency/288us
/intel/pciessd/drive-name/ReadCommandLatency/320us
/intel/pciessd/drive-name/ReadCommandLatency/352us
/intel/pciessd/drive-name/ReadCommandLatency/384us
/intel/pciessd/drive-name/ReadCommandLatency/416us
/intel/pciessd/drive-name/ReadCommandLatency/448us
/intel/pciessd/drive-name/ReadCommandLatency/480us
/intel/pciessd/drive-name/ReadCommandLatency/512us
/intel/pciessd/drive-name/ReadCommandLatency/544us
/intel/pciessd/drive-name/ReadCommandLatency/576us
/intel/pciessd/drive-name/ReadCommandLatency/608us
/intel/pciessd/drive-name/ReadCommandLatency/640us
/intel/pciessd/drive-name/ReadCommandLatency/672us
/intel/pciessd/drive-name/ReadCommandLatency/704us
/intel/pciessd/drive-name/ReadCommandLatency/736us
/intel/pciessd/drive-name/ReadCommandLatency/768us
/intel/pciessd/drive-name/ReadCommandLatency/800us
/intel/pciessd/drive-name/ReadCommandLatency/832us
/intel/pciessd/drive-name/ReadCommandLatency/864us
/intel/pciessd/drive-name/ReadCommandLatency/896us
/intel/pciessd/drive-name/ReadCommandLatency/928us
/intel/pciessd/drive-name/ReadCommandLatency/960us
/intel/pciessd/drive-name/ReadCommandLatency/992us
/intel/pciessd/drive-name/ReadCommandLatency/1ms
/intel/pciessd/drive-name/ReadCommandLatency/2ms
/intel/pciessd/drive-name/ReadCommandLatency/3ms
/intel/pciessd/drive-name/ReadCommandLatency/4ms
/intel/pciessd/drive-name/ReadCommandLatency/5ms
/intel/pciessd/drive-name/ReadCommandLatency/6ms
/intel/pciessd/drive-name/ReadCommandLatency/7ms
/intel/pciessd/drive-name/ReadCommandLatency/8ms
/intel/pciessd/drive-name/ReadCommandLatency/9ms
/intel/pciessd/drive-name/ReadCommandLatency/10ms
/intel/pciessd/drive-name/ReadCommandLatency/11ms
/intel/pciessd/drive-name/ReadCommandLatency/12ms
/intel/pciessd/drive-name/ReadCommandLatency/13ms
/intel/pciessd/drive-name/ReadCommandLatency/14ms
/intel/pciessd/drive-name/ReadCommandLatency/15ms
/intel/pciessd/drive-name/ReadCommandLatency/16ms
/intel/pciessd/drive-name/ReadCommandLatency/17ms
/intel/pciessd/drive-name/ReadCommandLatency/18ms
/intel/pciessd/drive-name/ReadCommandLatency/19ms
/intel/pciessd/drive-name/ReadCommandLatency/20ms
/intel/pciessd/drive-name/ReadCommandLatency/21ms
/intel/pciessd/drive-name/ReadCommandLatency/22ms
/intel/pciessd/drive-name/ReadCommandLatency/23ms
/intel/pciessd/drive-name/ReadCommandLatency/24ms
/intel/pciessd/drive-name/ReadCommandLatency/25ms
/intel/pciessd/drive-name/ReadCommandLatency/26ms
/intel/pciessd/drive-name/ReadCommandLatency/27ms
/intel/pciessd/drive-name/ReadCommandLatency/28ms
/intel/pciessd/drive-name/ReadCommandLatency/29ms
/intel/pciessd/drive-name/ReadCommandLatency/30ms
/intel/pciessd/drive-name/ReadCommandLatency/31ms
/intel/pciessd/drive-name/ReadCommandLatency/32ms
/intel/pciessd/drive-name/ReadCommandLatency/64ms
/intel/pciessd/drive-name/ReadCommandLatency/96ms
/intel/pciessd/drive-name/ReadCommandLatency/128ms
/intel/pciessd/drive-name/ReadCommandLatency/160ms
/intel/pciessd/drive-name/ReadCommandLatency/192ms
/intel/pciessd/drive-name/ReadCommandLatency/224ms
/intel/pciessd/drive-name/ReadCommandLatency/256ms
/intel/pciessd/drive-name/ReadCommandLatency/288ms
/intel/pciessd/drive-name/ReadCommandLatency/320ms
/intel/pciessd/drive-name/ReadCommandLatency/352ms
/intel/pciessd/drive-name/ReadCommandLatency/384ms
/intel/pciessd/drive-name/ReadCommandLatency/416ms
/intel/pciessd/drive-name/ReadCommandLatency/448ms
/intel/pciessd/drive-name/ReadCommandLatency/480ms
/intel/pciessd/drive-name/ReadCommandLatency/512ms
/intel/pciessd/drive-name/ReadCommandLatency/544ms
/intel/pciessd/drive-name/ReadCommandLatency/576ms
/intel/pciessd/drive-name/ReadCommandLatency/608ms
/intel/pciessd/drive-name/ReadCommandLatency/640ms
/intel/pciessd/drive-name/ReadCommandLatency/672ms
/intel/pciessd/drive-name/ReadCommandLatency/704ms
/intel/pciessd/drive-name/ReadCommandLatency/736ms
/intel/pciessd/drive-name/ReadCommandLatency/768ms
/intel/pciessd/drive-name/ReadCommandLatency/800ms
/intel/pciessd/drive-name/ReadCommandLatency/832ms
/intel/pciessd/drive-name/ReadCommandLatency/864ms
/intel/pciessd/drive-name/ReadCommandLatency/896ms
/intel/pciessd/drive-name/ReadCommandLatency/928ms
/intel/pciessd/drive-name/ReadCommandLatency/960ms
/intel/pciessd/drive-name/ReadCommandLatency/992ms
/intel/pciessd/drive-name/WriteCommandLatency/MajorRevision
/intel/pciessd/drive-name/WriteCommandLatency/MinorRevision
/intel/pciessd/drive-name/WriteCommandLatency/0us
/intel/pciessd/drive-name/WriteCommandLatency/32us
/intel/pciessd/drive-name/WriteCommandLatency/64us
/intel/pciessd/drive-name/WriteCommandLatency/96us
/intel/pciessd/drive-name/WriteCommandLatency/128us
/intel/pciessd/drive-name/WriteCommandLatency/160us
/intel/pciessd/drive-name/WriteCommandLatency/192us
/intel/pciessd/drive-name/WriteCommandLatency/224us
/intel/pciessd/drive-name/WriteCommandLatency/256us
/intel/pciessd/drive-name/WriteCommandLatency/288us
/intel/pciessd/drive-name/WriteCommandLatency/320us
/intel/pciessd/drive-name/WriteCommandLatency/352us
/intel/pciessd/drive-name/WriteCommandLatency/384us
/intel/pciessd/drive-name/WriteCommandLatency/416us
/intel/pciessd/drive-name/WriteCommandLatency/448us
/intel/pciessd/drive-name/WriteCommandLatency/480us
/intel/pciessd/drive-name/WriteCommandLatency/512us
/intel/pciessd/drive-name/WriteCommandLatency/544us
/intel/pciessd/drive-name/WriteCommandLatency/576us
/intel/pciessd/drive-name/WriteCommandLatency/608us
/intel/pciessd/drive-name/WriteCommandLatency/640us
/intel/pciessd/drive-name/WriteCommandLatency/672us
/intel/pciessd/drive-name/WriteCommandLatency/704us
/intel/pciessd/drive-name/WriteCommandLatency/736us
/intel/pciessd/drive-name/WriteCommandLatency/768us
/intel/pciessd/drive-name/WriteCommandLatency/800us
/intel/pciessd/drive-name/WriteCommandLatency/832us
/intel/pciessd/drive-name/WriteCommandLatency/864us
/intel/pciessd/drive-name/WriteCommandLatency/896us
/intel/pciessd/drive-name/WriteCommandLatency/928us
/intel/pciessd/drive-name/WriteCommandLatency/960us
/intel/pciessd/drive-name/WriteCommandLatency/992us
/intel/pciessd/drive-name/WriteCommandLatency/1ms
/intel/pciessd/drive-name/WriteCommandLatency/2ms
/intel/pciessd/drive-name/WriteCommandLatency/3ms
/intel/pciessd/drive-name/WriteCommandLatency/4ms
/intel/pciessd/drive-name/WriteCommandLatency/5ms
/intel/pciessd/drive-name/WriteCommandLatency/6ms
/intel/pciessd/drive-name/WriteCommandLatency/7ms
/intel/pciessd/drive-name/WriteCommandLatency/8ms
/intel/pciessd/drive-name/WriteCommandLatency/9ms
/intel/pciessd/drive-name/WriteCommandLatency/10ms
/intel/pciessd/drive-name/WriteCommandLatency/11ms
/intel/pciessd/drive-name/WriteCommandLatency/12ms
/intel/pciessd/drive-name/WriteCommandLatency/13ms
/intel/pciessd/drive-name/WriteCommandLatency/14ms
/intel/pciessd/drive-name/WriteCommandLatency/15ms
/intel/pciessd/drive-name/WriteCommandLatency/16ms
/intel/pciessd/drive-name/WriteCommandLatency/17ms
/intel/pciessd/drive-name/WriteCommandLatency/18ms
/intel/pciessd/drive-name/WriteCommandLatency/19ms
/intel/pciessd/drive-name/WriteCommandLatency/20ms
/intel/pciessd/drive-name/WriteCommandLatency/21ms
/intel/pciessd/drive-name/WriteCommandLatency/22ms
/intel/pciessd/drive-name/WriteCommandLatency/23ms
/intel/pciessd/drive-name/WriteCommandLatency/24ms
/intel/pciessd/drive-name/WriteCommandLatency/25ms
/intel/pciessd/drive-name/WriteCommandLatency/26ms
/intel/pciessd/drive-name/WriteCommandLatency/27ms
/intel/pciessd/drive-name/WriteCommandLatency/28ms
/intel/pciessd/drive-name/WriteCommandLatency/29ms
/intel/pciessd/drive-name/WriteCommandLatency/30ms
/intel/pciessd/drive-name/WriteCommandLatency/31ms
/intel/pciessd/drive-name/WriteCommandLatency/32ms
/intel/pciessd/drive-name/WriteCommandLatency/64ms
/intel/pciessd/drive-name/WriteCommandLatency/96ms
/intel/pciessd/drive-name/WriteCommandLatency/128ms
/intel/pciessd/drive-name/WriteCommandLatency/160ms
/intel/pciessd/drive-name/WriteCommandLatency/192ms
/intel/pciessd/drive-name/WriteCommandLatency/224ms
/intel/pciessd/drive-name/WriteCommandLatency/256ms
/intel/pciessd/drive-name/WriteCommandLatency/288ms
/intel/pciessd/drive-name/WriteCommandLatency/320ms
/intel/pciessd/drive-name/WriteCommandLatency/352ms
/intel/pciessd/drive-name/WriteCommandLatency/384ms
/intel/pciessd/drive-name/WriteCommandLatency/416ms
/intel/pciessd/drive-name/WriteCommandLatency/448ms
/intel/pciessd/drive-name/WriteCommandLatency/480ms
/intel/pciessd/drive-name/WriteCommandLatency/512ms
/intel/pciessd/drive-name/WriteCommandLatency/544ms
/intel/pciessd/drive-name/WriteCommandLatency/576ms
/intel/pciessd/drive-name/WriteCommandLatency/608ms
/intel/pciessd/drive-name/WriteCommandLatency/640ms
/intel/pciessd/drive-name/WriteCommandLatency/672ms
/intel/pciessd/drive-name/WriteCommandLatency/704ms
/intel/pciessd/drive-name/WriteCommandLatency/736ms
/intel/pciessd/drive-name/WriteCommandLatency/768ms
/intel/pciessd/drive-name/WriteCommandLatency/800ms
/intel/pciessd/drive-name/WriteCommandLatency/832ms
/intel/pciessd/drive-name/WriteCommandLatency/864ms
/intel/pciessd/drive-name/WriteCommandLatency/896ms
/intel/pciessd/drive-name/WriteCommandLatency/928ms
/intel/pciessd/drive-name/WriteCommandLatency/960ms
/intel/pciessd/drive-name/WriteCommandLatency/992ms
```

## License
[Snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE)

## Acknowledgements
* Author: [Wang Jialei A](https://github.com/WangJialei-A)
