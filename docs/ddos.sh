#!/usr/bin/env bash

# ab -n 100000 -c 1000 http://172.21.0.3:1234/
# ab -n 1000000 -c 1000 http://172.21.0.3:1234/

# oha -m POST -d "{\"title\":\"geia\",\"content\":\"sou\"}" -z 120s http://localhost:1234
# oha -m POST -d "{\"title\":\"geia\",\"content\":\"sou\"}" -n 100000 http://localhost:1234
# oha -n 100000 -c 1000 http://localhost:1234

# oha -z 2m -c 1000 http://localhost:1234
oha -z 2m -c 1000 http://localhost:80
