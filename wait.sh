#!/bin/sh
set -e
postgres="$1"
shift
cmd="$@"

until PGPASSWORD=$PWDPG psql -h "$postgres" -U "root" -c '\q'; do
  sleep 1
done

>&2 echo "all ok"
exec $cmd