SHELL := /bin/bash

.PHONY:	up down push 

up: 
	docker compose up -d

down:
	docker compose down

set_env:
	@echo "set config file"
	cp messaggio_test/.env_example messaggio_test/.env
	cp second_handler/.env_example second_handler/.env

push:
	scp messaggio_test/build/binary/messaggio_test anton@87.242.118.172:~/messaggio_test/messaggio_test
	scp messaggio_test/.env anton@87.242.118.172:~/messaggio_test/.env
	scp second_handler/build/binary/second_handler anton@87.242.118.172:~/second_handler/second_handler
	scp second_handler/.env anton@87.242.118.172:~/second_handler/.env
