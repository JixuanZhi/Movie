#! /user/bin/env bash
# Configures the nginx server upon startup.

# Assign configuration values to enviroment values.
export HOST_IP=$1
#export TFX_PORT=$2
export RANK_PORT=$2
#export REVERSE_INDEX_PORT=$4

# Produce nginx configuration from template
/opt/homebrew/Cellar/gettext/0.21.1/bin/envsubst '$HOST_IP $RANK_PORT' < nginx.conf.template > /opt/homebrew/etc/nginx/nginx.conf

# Show generated config for the sake of debugging.
cat /opt/homebrew/etc/nginx/nginx.conf

# Start nginx.
nginx -g 'daemon off;' 2> start_error.log