#!/bin/bash

#
# Just some commands to test the API with.
#

# Add host
curl -X POST -H 'Content-Type:application/json' --data '{"endpoint": "google.com", "regex": "blah"}' localhost:8081/add
curl -X POST -H 'Content-Type:application/json' --data '{"endpoint": "womp.com", "regex": "blah"}' localhost:8081/add