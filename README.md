# server-collector
An application smartly turn off your unused machines

## Go 1.11 Modules

Please upgrade your go version to v1.11+ so that you can use go module. You have to set `GO111MODULE=on`. For more information, please see [golang/go](https://github.com/golang/go/wiki/Modules)
```sh
    $ export GO111MODULE=on
```


## Build binary

Executes the following commands right under the root directory of this repository:
```sh
    $ go build -o server-collector cmd/server-collector

    # or

    $ go build -o server-collector github.com/lmchih/server-collector/cmd/server-collector
```

They both generate the executable binary


## Build image

To build an image, use Dockerfile at the directory:

```sh
    $ docker build -f build/package/Dockerfile -t {your_image_path_with_tags} .
```

## Run via Docker

You can run like this:
```sh
    $ docker run -d --name server-collector dockerhub.com/lmchih/server-collector
```


## Usage



## Contact

Any questions and feedbacks are so welcome.
* ricklin126@gmail.com
