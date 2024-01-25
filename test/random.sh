#!/bin/bash

curl -X GET localhost:8080/random
curl -X POST localhost:8080/match/ -d "home=1&away=2" | jq.exe > result.json