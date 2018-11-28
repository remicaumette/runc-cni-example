#!/usr/bin/env bash
ctr task delete -f redis
ctr container rm redis
