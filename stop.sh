#!/bin/bash
killall service_app # kill service_app service
echo "stop service_app success"
ps -aux | grep service_app