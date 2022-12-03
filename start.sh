#!/bin/bash

docker run --rm -it -d -v "$(pwd)":/go_source go:1.1 /bin/bash
