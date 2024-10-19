#!/usr/bin/env bash

cd app || exit

docker build -t localhost:5000/eso-webhook:v1 .
