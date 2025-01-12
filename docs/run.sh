#!/usr/bin/env bash

set -xe

sudo /usr/local/zeek/bin/zeek -i lo -C alerts.zeek
