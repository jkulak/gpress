## Compilation

I like to use Docker 🐳 to work with Golang on my machine.

Run the build

```
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.6 go build -v
```

Work inside the Docker container

```
$ docker run --rm -ti -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.6 bash
```

When mounting the whole $GOPATH directory

```
$ docker run --rm -ti -e "GOPATH=/root/go" -v "$PWD":/root/go -w /root/go golang:1.6 bash
```
