#!/bin/sh
envsubst < /app/config.yaml.dist > /app/config.yaml
exec /app/app