# /bin/bash

GO111MODULE=on go mod download
GO111MODULE=on go mod vendor
GO111MODULE=on go mod verify
