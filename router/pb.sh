#!/bin/bash

protoc30 --go_out=plugins=grpc:. router.proto
