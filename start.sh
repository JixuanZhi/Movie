#!/bin/bash

docker run  -p 80:5001 --rm -it -d -v "$(pwd)":/go_source go:1.11 /bin/bash