# IoT-Final-Project

An IoT seismograph simulator using Google Cloud IoT Core, Pub/Sub and Dataflow

This IoT "device" simulates seismograph activity, publishing a message with scale (-1 <=> 1), location (gps cords), date, and device id.


## Environment Variables

### Linux or macOS

```bash
export VARIABLE_NAME="[VALUE]"
```

### Windows

powershell:

```powershell
$env:VARIABLE_NAME="[VALUE]"
```

command prompt:

```powershell
set VARIABLE_NAME=[VALUE]
```

| Name  | Description | Default |
|---|---|---|
| CONFIG_PATH  |   the path to your config.yaml | [project_path]/config/config.yaml |
| GOOGLE_APPLICATION_CREDENTIALS | The path to a service account key-file used for local authentication  |  |


## Configuration

Your configruation should look something like this:

```yaml
GCP:
PROJECT_ID: ""
BUCKET_NAME: ""

PUB_SUB:
    TOPIC_ID: ""
    TOPIC_NAME: ""

IOT_CORE:
    GCP_CLOUD_IOT_CORE_REGION: us-central1 #us-central1, europe-west1, asia-east1
    REGISTRY_ID: ""

DEVICE:
    NAME: ""
    PATH: "" #[registry_id]/devices/[device_name]
    LOCATION: "" # The devices "physical" gps location 

SIMULATION:
    INTERVAL: 1  # seconds between publishes.
    JOB_SIZE: 10 # Number of messages to publish
```

## Building the project

```bash
go build -v bin/iot_device .
```

## Usage

Define the message interval and job size in the configuration file.
Intervals are in seconds.

Execute the binary:

```bash
bin/iot_device
```