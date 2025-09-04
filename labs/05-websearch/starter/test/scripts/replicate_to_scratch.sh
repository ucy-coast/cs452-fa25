#!/bin/bash

SRC_DIR="../data"
DEST_DIR="/local/shards"
NODES="kube1 kube2 kube3"  # List your nodes here

# Check source directory
if [ ! -d "$SRC_DIR" ]; then
  echo "❌ Source directory $SRC_DIR does not exist. Aborting."
  exit 1
fi

# Loop through each node in the NODES list
for NODE in $NODES; do
  echo "=== 🚀 Syncing to $NODE ==="
  
  # Create target directory if missing (no sudo)
  ssh "$NODE" "mkdir -p $DEST_DIR"

  # Sync contents of SRC_DIR into DEST_DIR/data
  rsync -avz "$SRC_DIR/" "$NODE:$DEST_DIR/data/"
done

echo "✅ Data replication complete to /local/scratch on all nodes."
