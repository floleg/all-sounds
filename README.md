# all-sounds
Hi Deezer people, I'm Florent! This a proposition of a Go implementation for the REST API exercise.

<!-- ----------------------------------------------------------------------------------------------- -->

## Requirements

A few tools can be useful to build, run, and test this application: 
- [go 1.19](https://go.dev/doc/install) to buil the sources;
- [docker](https://docs.docker.com/get-started/) and [docker-compose](https://docs.docker.com/compose/) to run the stack in containers;
- [curl](https://curl.se/docs/manpage.html) to consume our API endpoints;
- [jq](https://stedolan.github.io/jq/) to get nicely formatted json responses.

## Toolchain

### RDBMS

The relational DB system chosen for this solution is postgresql, run as a docker container.

### Libraries

In order not to re-invent the wheel, a little help from the great golang open-source community is always useful. Here's the dependencies I used to implement this API (unexhaustive):

- [gin-gonic](https://gin-gonic.com/docs/) - A powerful http web framework to easily write and maintain REST endpoints;
- [gorm](https://gorm.io/) - The best solution I found to interact with a postgresql instance and generate the relational schema used in this solution;
- [dockertest](https://github.com/ory/dockertest) - Because end-to-end integration is far better than mocking to validate interactions with the RDBMS. Dockertest is used to run a postgresql container for each test suite;
- [viper](github.com/spf13/viper) - Viper is used to inject per-environment configuration;
- [faker](https://github.com/bxcodec/faker) - Integrations tests require data provisioning, faker fits perfectly for this purpose.

## Omitted components and implementations

Due to understandable tme constraint, some server components and software implementations have been voluntarily omitted from this solution. Let's present how this application could be properly industrialized and released.

**Rate limiting** should be implemented on the router layer, despite the fact that gin doesn't provide such feature natively, but I might be wrong on this one.

**Logging / Observability** has to be improved on all levels. Here we're just redirecting runtime informations on the console output. A proper way to do it would be either using elastic search aloginside with Kibana or Grafana's Loki and Fluentd.

**Monitoring** can be easily setup with the combination of prometheus, node exporter and Grafana.

## Implementation

### Relation schema

Here's the relational schema used to store the API data, entirely defined with Gorm entities as declared in the [model](https://gorm.io/docs/migration.html#Auto-Migration) package and persisted on server startup with [auto migrate](https://gorm.io/docs/migration.html#Auto-Migration) feature.

![all-sound relations schema!](/assets/images/all-sounds.png "San Juan Mountains")