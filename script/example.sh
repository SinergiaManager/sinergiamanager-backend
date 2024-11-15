#!/bin/bash
sh -c "docker-entrypoint.sh mongod --quiet --replSet rs0" &

echo "Running example script"

mongosh <<EOF
load("/file.js");
EOF

wait