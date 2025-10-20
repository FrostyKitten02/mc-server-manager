# MC Server Manager Backend
This is a backend service for managing MC servers, viewing logs, running commands on the server.
Assign what commands users are allowed to run on the server via RCON protocol.
## Requirements
MC servers need to run in docker containers and have their RCON port exposed, if running this service on the same server as the MC servers then just use localhost, this is the preferred way of doing it.
## System Variables
- MCSM_DATA_DIR: Specifies where your config files are located by default is the data directory.