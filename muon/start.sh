#!/bin/sh
nginx
cd /srv/api
gunicorn --bind unix:/run/sockets/gunicorn.socket --env WEBAPI_SETTINGS=./webapi.json webapi.wsgi:app
