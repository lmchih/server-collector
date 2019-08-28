# server-collector

An application smartly turns off your unused machines. On those feature servers(especially for those internal testing of features), the app constantly(for like every couple minute, configurable) checking the commit date from your Github repo. If the commit was not updated for long time, **server-collector** marks your server as unused and then turns it off.

## Go 1.11 Modules

Please upgrade your go version to v1.11+ so that you can use go module. You have to set `GO111MODULE=on`. For more information, please see [golang/go](https://github.com/golang/go/wiki/Modules)

```sh
export GO111MODULE=on
```

## Usage

## Run as binary

Before running the program, you will need to complete the configuration yaml file. Put your Github information as the following format: `https://github.com/{{sourceOwner}}/{{sourceRepo}}` and decide your checking frequency and the commit expirations.

```yaml
# image version
version: 0.0.3
# your personal/organization access token
accessToken: ""
#localhost
serverIP: 127.0.0.1
# https://github.com/{{sourceOwner}}/{{sourceRepo}}
sourceOwner: lmchih
# https://github.com/{{sourceOwner}}/{{sourceRepo}}
sourceRepo: server-collector
# not used so far
sourceBranch: master
# how often (in seconds) the program runs a check
checkFrequency: 120
# your repo non-active days
unusedDays: 3
```

```sh
cp ./server-collector /usr/local/bin/
server-collector &
```

## Run with container

### Step 1: start the container

```sh
# Docker
docker run -idt -v /var/run:/var/run rickming/server-collector:0.0.3

# Kubernetes
cd deployments
$ kubectl apply -f server-collector.yaml
deployment.extensions "server-collector" created
configmap "server-collector-conf" created
```

### Step 2: start the listening process

```sh
cd scripts
chmod +x check-shutdown-signal.sh

# Add sudo before sh command if you are not superuser
sh ./check-shutdown-signal.sh &

# NOTE: If you are in Kubernetes cluster, run this script at the worker where you pod is located.
```

## Build

### binary

Executes the following commands right under the root directory of this repository:

```sh
go build -o server-collector cmd/server-collector/binary
```

This generates an executable named `server-collector`

### image

To build a Docker image, use Dockerfile at the `build/package/` directory:

```sh
docker build -f build/package/Dockerfile -t {your_image_path_with_tags} .
```

## Contact

Any questions and feedbacks are so welcome.

* ricklin126@gmail.com
