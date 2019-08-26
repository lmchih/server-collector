#!/bin/bash


# The script listens to the shutdown signal from the container `server-collector`. To make this works,
# do the following steps:
# 1. crontab
#
# 2. 
# Add this script to /etc/init.d/ directory
# chmod +x /etc/init.d/check-shutdown-signal.sh
echo "waiting" > /var/run/shutdown_signal
while sleep 30; do 
  signal=$(cat /var/run/shutdown_signal)
  if [ "$signal" == "true" ]; then 
    echo "done" > /var/run/shutdown_signal
    sudo shutdown -h now
  fi
done