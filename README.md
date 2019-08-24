# 

*# server-collector*
An application smartly turn off your unused machine

*## Go 1.11 Modules*

Please upgrade your go version to v1.11+ so that you can use go module. You have to set `GO111MODULE=on`. For more information, please see [golang/go](_https://github.com/golang/go/wiki/Modules_)
```sh
    $ export GO111MODULE=on
    $ env
```


*## Build binary*

Executes the following commands right under the root directory of this repository:
```sh
    $ go build -o server-collector cmd/main.go
```
or
```sh
    $ go build -o server-collector github.com/lmchih/server-collector/cmd
```

This both generate the executable binary


*## Build image*

To build an image, use Dockerfile at the directory:

```sh
    $ docker build -f build/package/Dockerfile -t {your_image_path_with_tags} .
```

*## Run via Docker*

You can run like this:
```sh
    $ docker run -d --name server-collector -p 8080:8080 dockerhub.com/lmchih/server-collector
```


*## Usage*



*## Contact*

Any questions and feedbacks are so welcome.
* ricklin126@gmail.com

