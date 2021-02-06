# [scratch-container](https://www.youtube.com/watch?v=8fi7uSYlOdc)

[Inspired by Liz Rice's "from scratch" series](https://www.youtube.com/watch?v=8fi7uSYlOdc)

## [Liz Rice - Inspiration Repo](https://github.com/lizrice/containers-from-scratch)

Writing a container in a few lines of Go code, as seen at DockerCon 2017 and on O'Reilly Safari

You need root permissions for this version to work. Or you can adapt it to be a rootless container by as shown in these slides.

Note that the Go code uses some syscall definitions that are only available when building with GOOS=linux.

## Usage

```bash
# cd into source
cd src

# go build and run main
go run main.go
```
