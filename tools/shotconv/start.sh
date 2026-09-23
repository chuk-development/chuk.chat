#!/bin/sh
# Run by the nginx image's entrypoint before nginx starts. Keeps shotconv
# running in the background and restarts it if it ever exits.
(while true; do /usr/local/bin/shotconv; sleep 1; done) &
