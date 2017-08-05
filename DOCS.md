Use the KODO plugin to upload files and folders with a KODO bucket. The following parameters are used to configure this plugin:

* `host` - kodo host [default value: https://upload.qbox.me]
* `access_key` - kodo access key
* `secret_key` - kodo secret key
* `bucket` - bucket name
* `source` - location of folder to upload
* `key` - target path in your kodo bucket
* `delete` - deletes files in the bucket if exists

The following is a sample KODO configuration in your .drone.yml file:

```yaml
archive:
  kodo:
    bucket: "release-candidates"
    access_key: $SPOCK_AK,
    secret_key: $SPOCK_SK,
    source: /folder/to/archive
    key: /target/location
    delete: true
  secret: [SPOCK_AK,SPOCK_SK]
```