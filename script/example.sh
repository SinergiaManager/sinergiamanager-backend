#!/bin/bash
sh -c "docker-entrypoint.sh mongod --quiet" &

sleep 5

echo "Running example script"

mongosh <<EOF
load("/file.js");
EOF

wait