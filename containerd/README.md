# [Getting started with containerd](https://containerd.io/docs/getting-started/)

There are many different ways to use containerd:

If you are a developer working on containerd you can use the ctr tool to quickly test features and functionality without writing extra code
If you want to integrate containerd into your project, you can use a simple client package. In this guide, we will pull and run a Redis server with containerd using that client package.
We will assume that you are running a modern Linux host for this example with a compatible build of runc.

This project requires Go 1.9.x or above.

## Usage

```bash
# download containerd
wget https://github.com/containerd/containerd/archive/v1.4.3.zip

# unzip
unzip v1.4.3.zip

# start
sh start
```

```go
// connecting to the container
package main

import (
	"log"
  // containerd
	"github.com/containerd/containerd"
)

func main() {
	if err := containerExample(); err != nil {
		log.Fatal(err)
	}
}

func containerExample() error {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()
	return nil
}
```
