# server-collector

An application smartly turn off your unused machines

## Go 1.11 Modules

Please upgrade your go version to v1.11+ so that you can use go module. You have to set `GO111MODULE=on`. For more information, please see [golang/go](https://github.com/golang/go/wiki/Modules)

```sh
export GO111MODULE=on
```

## Build

### binary

Executes the following commands right under the root directory of this repository:

```sh
go build -o server-collector cmd/server-collector/binary
```

This generates an executable named `server-collector`

### image

To build a Docker image, use Dockerfile at the directory:

```sh
docker build -f build/package/Dockerfile -t {your_image_path_with_tags} .
```

## Usage

## Run as binary

```sh
cp ./server-collector /usr/local/bin/
server-collector &
```

## Run with container

### Step 1: start the container

```sh
# Linux/MacOS
docker run -v /shutdown_signal:/var/run/shutdown_signal -e SOURCE_REPO=server-collector -idt --name=server-collector rickming/server-collector:0.0.1

# Windows
docker run -v C:\\shutdown_signal:/var/run/shutdown_signal -e SOURCE_REPO=server-collector -idt --name=server-collector rickming/server-collector:0.0.1

```

### Step 2: start the listening process

```sh
cd scripts
chmod +x check-shutdown-signal.sh

# Add sudo before sh command if you are not superuser
sh ./check-shutdown-signal.sh &
```

## Contact

Any questions and feedbacks are so welcome.

* ricklin126@gmail.com
