# Pub/Sub Service

This service is the middle point between the front end and GCP's pub/sub service for a seismograph simulator.

There are two "Devices". One written in Golang, and another written in C#. Both devices create sensor data representing seismic activity.
The data sent is in the form of:

```json
{
    "id": 1,
    "scale": 0, // between 1 and -1
    "location": "37.772239, -122.422889",
    "time": "2020-12-13T19:58:16"
}
```

This service subscribes to a topic on GCP's pub/sub service.
For every message recieved from the topic, the same message is sent through a websocket to a react-based front end application.

The application uses Auth0 for user authentication, and the same service is used to authenticate

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
| PUB_SUB_SERVICE_LOGGING_PATH  |   the path to your config.yaml | [project_path]/config/config.yaml |
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
      SUBSCRIPTION: ""

    DEVICE:
      NAME: ""
      PATH: ""
      LOCATION: ""

SIMULATION:
    INTERVAL: 1 #seconds
    JOB_SIZE: 10

SERVER:
  PORT: 8000
```
