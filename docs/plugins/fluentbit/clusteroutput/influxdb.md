# Influxdb

The linfluxdb output plugin, allows to flush  records into a [InfluxDB](https://www.influxdata.com/time-series-platform/influxdb/) time series database.

| Field       | Description                                                                                                 | Scheme                       |
| ----------- | ----------------------------------------------------------------------------------------------------------- | ---------------------------- |
| host        | IP address or hostname of the target InfluxDB service.                                                      | string                       |
| port        | TCP port of the target InfluxDB service.                                                                    | *int32                       |
| database    | InfluxDB database name where records will be inserted.                                                      | string                       |
| bucket      | InfluxDB bucket name where records will be inserted.                                                        | string                       |
| org         | Specifies the size of files in S3. Maximum size is 50G, minimim is 1M.                                      | string                       |
| sequenceTag | The name of the tag whose value is incremented for the consecutive simultaneous events.                     | string[]strin                |
| httpUser    | Optional username for HTTP Basic Authentication.                                                            | *[plugins.Secret](../secret.md) |
| httpPasswd  | Password for user defined in HTTP_User.                                                                     | *[plugins.Secret](../secret.md) |
| httpToken   | Authentication token used with InfluDB v2.                                                                  | string                       |
| tagKeys     | Space separated list of keys that needs to be tagged.                                                       | string                       |
| autoTags    | Automatically tag keys where value is string.                                                               | string                       |
| tagsListKey | Key of the string array optionally contained within each log record that contains tag keys for that record | string                       |
| tls         |                                                                                                             | *[plugins.TLS](../tls.md)       |
