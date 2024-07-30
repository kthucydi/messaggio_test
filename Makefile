SHELL := /bin/bash

.PHONY:	up down push 

up:
	docker compose up -d

down:
	docker compose down

push:
	scp build/binary/messaggio_test anton@87.242.118.172:~/messaggio_test

