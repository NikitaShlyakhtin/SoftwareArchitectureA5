# F24 Software Architecture Assignment 5

## SD-01

## Team 15

## Authors:

- Ilia Sardanadze (i.sardanadze@innopolis.university)
- Karim Nasybullin (k.nasybullin@innopolis.university)
- Nikita Shlyakhtin (n.shlyakhtin@innopolis.university)

## How to run:

Step 1: clone the repository

```
git clone https://github.com/NikitaShlyakhtin/SoftwareArchitectureA5
```

Step 2: run the system using docker-compose (note that you must
have running docker engine)

```
docker compose up --build
```

This command will build all required containers with 
PostgreSQL database, Messages Service, User Management Service, Feed Service and API Gateway/CLI frontend.

### How to test:

Attach to container with CLI:

```
docker attach $(docker ps -a -q --filter ancestor=a5-cli_ui --format="{{.ID}}")
```

Or user `docker ps` to view active containers, copy `id` for container with `a5_cli_ui` service and run
`docker attach <id>` with that id

Then you can test application using following commands:

```
register <Name>

createMessage <Username> <Message>

likeMessage <Username> <Message>

showFeed
```