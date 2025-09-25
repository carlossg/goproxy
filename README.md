# goproxy

```
export GOPROXY=...
```

then

```
go mod download -x
```

or

```
docker run -ti --rm -e GOPROXY -v `pwd`:/build -w /build \
    golang:1.25.0-alpine go mod download -x;
```
