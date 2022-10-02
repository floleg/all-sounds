# all-sounds
Hi Deezer people, I'm Florent! This a proposition of a Go implementation for the REST API exercise.

<!-- ----------------------------------------------------------------------------------------------- -->

## Requirements

A few tools can be useful to build, run, and test this application: 
- [go 1.19](https://go.dev/doc/install) to buil the sources;
- [docker](https://docs.docker.com/get-started/) and [docker-compose](https://docs.docker.com/compose/) to run the stack in containers;
- [jq](https://stedolan.github.io/jq/) to get nicely formatted json responses.

## Toolchain

### Database

The relational DB system chosen for this solution is postgresql, run as a docker container.

### Libraries

In order not to re-invent the wheel, a little help from the great golang open-source community is always useful. Here's the dependencies I used to implement this API (unexhaustive):

- [gin-gonic](https://gin-gonic.com/docs/) - A powerful http web framework to easily write and maintain REST endpoints;
- [gorm](https://gorm.io/) - The best solution I found to interact with a postgresql instance and generate the relational schema used in this solution;
- [dockertest](https://github.com/ory/dockertest) - Because end-to-end integration is far better than mocking to validate interactions with the RDBMS. Dockertest is used to run a postgresql container for each test suite;
- [viper](github.com/spf13/viper) - Viper is used to inject per-environment configuration.