minio-cf - Minio CF Tool
========================

Minio tool to be used with Pivotal "cf" tool to generate the config for create-service command.

EXAMPLES:
--------
Deploy Minio Server instance
```
$ minio-cf --access-key minio --secret-key minio123 > minio-server.conf
$ cf create-service minio standard server-instance-1 -c minio-server.conf
```

Deploy Minio GCS gateway instance
```
$ minio-cf --access-key minio --secret-key minio123 --gcs /path/to/credentials.json > minio-gcs.conf
$ cf create-service minio standard server-instance-2 -c minio-gcs.conf
```

Deploy Minio Azure gateway instance
```
$ minio-cf --access-key azureaccountname --secret-key azureaccountkey > minio-azure.conf
$ cf create-service minio standard server-instance-3 -c minio-azure.conf
```
