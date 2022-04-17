#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
mkdir -p $CURDIR/log
echo "$CURDIR/bin/entry_task   -conf-dir=$CONF_DIR/conf"
exec  $CURDIR/bin/entry_task   -conf-dir=$CONF_DIR/conf