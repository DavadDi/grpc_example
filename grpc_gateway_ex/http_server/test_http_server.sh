#!/bin/bash

curl -v http://127.0.0.1:9080/v1/hello/echo -X POST -d '{"HelloReq": {"name": "Dave"}}' -H  "content-type: application/json"
