#!/bin/bash

BOLDGREEN="\033[1;32m"
RESET="\033[0m"

psql -U postgres -d recordings -f init-commands.sql

printf "${BOLDGREEN}Finished initializing the database!${RESET}"