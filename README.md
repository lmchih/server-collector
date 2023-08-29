# server-collector

**server-collector** smartly turns off your unused machines. Deployed on top of feature servers(especially those internal testing of features), the application constantly(configurable) checks the latest commit date of your Github repository. If the repository was inactive for long time, **server-collector** marks your servers as unused and turns them off.

## Go 1.11 Modules

Please upgrade your **go** version to **v1.11+** so that you can use go module. You have to set `GO111MODULE=on`. For more information, please see [golang/go](https://github.com/golang/go/wiki/Modules)

```sh
export GO111MODULE=on
```

## Executable

### Usage

```sh
$ server-collector -h

Usage: main [-h] [-b value] [-c value] [-f value] [-i value] [-o value] [-r value] [-t value] [-u value] [parameters ...]
 -b, --branch=value
                    Github repo branch (Support master only)
 -c value           Seconds between every check
 -f, --from-file=value
                    The path of configuration file. Support yaml only
 -h, --help         Help
 -i, --ip=value     Support localhost only
 -o, --owner=value  Github repo owner: https://github.com/{owner}/{repo}
 -r, --repo=value   Github repo name: https://github.com/{owner}/{repo}
 -t, --token=value  Your personal/organization Github Personal Access Token
 -u, --unused-days=value
                    Days considered unused
```

### Github Access Token

Github access token is **required**. Run with the option `--token=value` or `-t value`

```sh
server-collector --token=mygithubaccesstoken
```

### Configuration

#### configuration file(optional)

You may want to create a yaml file with the following format. Use the option `--from-file=<YOUR_YAML_PATH>` to pass the file in. (reference: `./configs/config.yaml`)

```yaml
# Application version
version: 0.0.3
# Your personal/organization access token
accessToken: ""
# Support localhost only
serverIP: 127.0.0.1
# https://github.com/{sourceOwner}/{sourceRepo}
sourceOwner: lmchih
# https://github.com/{sourceOwner}/{sourceRepo}
sourceRepo: server-collector
# Support master branch only
sourceBranch: master
# Seconds between every check
checkFrequency: 120
# Inactive days for your monitored repository
unusedDays: 3
```

### Runs at background

```sh
cp ./server-collector /usr/local/bin/
server-collector &
```

## Run with container

### Step 1: Start the container

#### Environment Variables

| key             | type   | example value                            | description                                   |
| --------------- | ------ | ---------------------------------------- | --------------------------------------------- |
| TARGET_SERVER   | string | 127.0.0.1                                | target server IP: support localhost only      |
| ACCESS_TOKEN    | string | {YOUR_PERSONAL_ACCESS_TOKEN}             | Github access token                           |
| SOURCE_OWNER    | string | {YOUR_GITHUB_USERNAME}                   | https://github.com/{sourceOwner}/{sourceRepo} |
| SOURCE_REPO     | string | server-collector                         | https://github.com/{sourceOwner}/{sourceRepo} |
| SOURCE_BRANCH   | string | master                                   | Support master branch only                    |
| CHECK_FREQUENCY | int64  | 120                                      | Seconds between every check                   |
| UNUSED_DAYS     | int64  | 3                                        | Inactive days for your monitored repository   |

```sh
# Docker (pass environment variables with -e flag)
docker run -idt -v /var/run:/var/run \
-e ACCESS_TOKEN={yourgithubaccesstoken} \
-e SOURCE_OWNER={repo_owner} \
-e SOURCE_REPO={repo_name} \
-e CHECK_FREQUENCY=120 \
-e UNUSED_DAYS=3  \
rickming/server-collector:0.0.4

# Kubernetes
cd deployments
$ kubectl apply -f server-collector.yaml
deployment.extensions "server-collector" created
configmap "server-collector-conf" created
```

### Step 2: Start the listening process at the host machine

```sh
cd scripts
chmod +x check-shutdown-signal.sh

# Add sudo before sh command if you are not superuser
sh ./check-shutdown-signal.sh &

# NOTICE: In Kubernetes cluster, run the script at the worker node where your pod is located.
```

## Build

### Binary

Executes the following commands right under the root directory of this repository:

```sh
go build -o server-collector cmd/server-collector/binary/main.go
```

This generates an executable named `server-collector`

### Docker image

To build a Docker image, use Dockerfile at the `build/package/` directory:

```sh
docker build -f build/package/Dockerfile -t {your_image_path_with_tags} .
```

## Contact

Any questions and feedbacks are so welcome.

* ricklin126@gmail.com
