# Docker Vault volume extension

## State of the project

I just wrote all this code before going to sleep last night.

## Installation

`sudo docker-volume-vault -token <your-root-vault-token> -url <url-to-vault>`

## Usage 

`docker run --volume-driver vault --volume secret/foo:/etc/secret`

