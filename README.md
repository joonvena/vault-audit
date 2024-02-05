# vault-audit-log-fluentd

This repository contains example how to audit logging device for Hashicorp Vault. It will write the audit log to a file. There is two sidecars containers running along Vault.

## fluentd

This container will tail the audit log and also pushes the log to Azure Blob Storage. There is separate [Dockerfile](./fluentd/Dockerfile) for this that extends the Fluentd base image by installing the [fluent-plugin-azure-storage-append-blob](https://github.com/microsoft/fluent-plugin-azure-storage-append-blob) plugin. Fluentd will also write following files to the mounted audit path (default. `/vault/audit`)  `audit.log.pos` to record the position it last read from the audit log. Also by default the Azure Blob Storage plugin will buffer the content of the log to a file path `/vault/audit/fluentdazblob` and will send them daily to Azure Blob Storage.

## log-rotate

This is used to rotate the logs daily. You can check the program structure from [main.go](./main.go). It has following environment variables available to configure.

| Environment Variable | Default Value | Description                                                                               |
| -------------------- | ------------- | ----------------------------------------------------------------------------------------- |
| `CRON_SCHEDULE`      | `"@daily"`    | Used to set the cron schedule. If not set, the schedule defaults to daily.                |
| `DEBUG`              | ""            | If set, the `-d` flag is added to the `logrotate` command arguments, enabling debug mode. |
