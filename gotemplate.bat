@echo off
REM this ensure the script run without echoing every line to the console

REM -- 1. Creating the cmd folder --
REM MD stand

MD cmd
REM creating a file
COPY NUL "cmd/api.go" > NUL
COPY NUL "cmd/main.go" > NUL


MD internal
MD "internal/adapters"
MD "internal/env"
MD "internal/json"
MD "internal/business"
MD "internal/routers"


echo Script finished. Press any key to exit.
pause > NUL