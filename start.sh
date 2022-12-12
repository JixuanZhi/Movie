#!/bin/bash

docker run --rm -it -d -v "$(pwd)":/go_source go:1.11 /bin/bash
