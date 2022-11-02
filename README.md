# all-sounds
This repository features a simple REST API, inspired by [the go REST API tutorial](https://go.dev/doc/tutorial/web-service-gin)

## Requirements

A few tools can be useful to build, run, and test this application: 
- [go 1.19](https://go.dev/doc/install) to buil the sources;
- [docker](https://docs.docker.com/get-started/) and [docker-compose](https://docs.docker.com/compose/) to run the stack in containers;
- [curl](https://curl.se/docs/manpage.html) to consume / test our API endpoints;
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

## Relation schema

Here's the relational schema used to store the API data, entirely defined with Gorm entities as declared in the [model](https://gorm.io/docs/migration.html#Auto-Migration) package and persisted on server startup with [auto migrate](https://gorm.io/docs/migration.html#Auto-Migration) feature.

<img src="./assets/images/all-sounds.png" width="1000" height="500">

## Run it!

### With docker compose

The compose file contained in this repository features two services:
- **db**: the postgresql db container;
- **server**: the Go API server container, based on the root Dockerfile of this project.

To get the stack up and running, simply run:
`➜  ~ docker-compose up`

### As a standalone application

If you wish to start the application server in Go native mode, simply:
`➜  ~ ENV=dev go run ./cmd/server`

Note that the `ENV` must be set to load a specific config file to interact with the database outside the docker created subnet. Obviously, a postgresql instance must be running in order to get a functional API server.

## Test it!

### Artist

<details>
  <summary>GET /artist - Returns a list of artists.</summary>

Query parameters:
- `offset`: mandatory - Sets the offset in the select query;
- `limit`: mandatory -  Set the fetched records limit in the select query;
- `query`: optional - A free text field compared to the `name` column.

```
➜  ~ curl -s http://127.0.0.1:8080/artist\?offset\=0\&limit\=10 | jq
[
  {
    "ID": 1,
    "CreatedAt": "2022-10-03T09:54:57.562Z",
    "UpdatedAt": "2022-10-03T09:54:59.128Z",
    "DeletedAt": null,
    "Name": "Artist One",
    "Tracks": null
  },
  {
    "ID": 2,
    "CreatedAt": "2022-10-03T09:54:57.562Z",
    "UpdatedAt": "2022-10-03T09:54:59.128Z",
    "DeletedAt": null,
    "Name": "Artist Two",
    "Tracks": null
  }
]

```

```
➜  ~ curl -s http://127.0.0.1:8080/artist\?query\=One\&offset\=0\&limit\=10 | jq
[
  {
    "ID": 1,
    "CreatedAt": "2022-10-03T09:54:57.562Z",
    "UpdatedAt": "2022-10-03T09:54:59.128Z",
    "DeletedAt": null,
    "Name": "Artist One",
    "Tracks": null
  }
]

```
</details>

<details>
  <summary>GET /artist/:id - Returns a single artist with a collection of albums, based on its unique id</summary>

```
{
  "ID": 1,
  "CreatedAt": "2022-10-03T09:54:57.562Z",
  "UpdatedAt": "2022-10-03T09:54:59.128Z",
  "DeletedAt": null,
  "Name": "Artist One",
  "Tracks": [
    {
      "ID": 1,
      "CreatedAt": "2022-10-03T10:22:55.079Z",
      "UpdatedAt": "2022-10-03T10:22:56.58Z",
      "DeletedAt": null,
      "Title": "Track One",
      "ArtistID": 1,
      "Users": null,
      "Albums": null
    },
    {
      "ID": 2,
      "CreatedAt": "2022-10-03T10:22:55.079Z",
      "UpdatedAt": "2022-10-03T10:22:56.58Z",
      "DeletedAt": null,
      "Title": "Track Two",
      "ArtistID": 1,
      "Users": null,
      "Albums": null
    }
  ]
}

```
</details>

### Album

<details>
  <summary>GET /album - Returns a list of albums.</summary>

Query parameters:
- `offset`: mandatory - Sets the offset in the select query;
- `limit`: mandatory -  Set the fetched records limit in the select query;
- `query`: optional - A free text field compared to the `name` column.

```
➜  ~ curl -s http://127.0.0.1:8080/album\?offset\=0\&limit\=10 | jq             
[
  {
    "ID": 1,
    "CreatedAt": "2022-10-03T10:40:35.447Z",
    "UpdatedAt": "2022-10-03T10:40:36.855Z",
    "DeletedAt": null,
    "Title": "Album One",
    "ReleaseYear": 2000,
    "Tracks": null
  },
  {
    "ID": 2,
    "CreatedAt": "2022-10-03T10:40:35.447Z",
    "UpdatedAt": "2022-10-03T10:40:36.855Z",
    "DeletedAt": null,
    "Title": "Album Two",
    "ReleaseYear": 2001,
    "Tracks": null
  }
]


```

```
➜  ~ curl -s http://127.0.0.1:8080/album\?query\=One\&offset\=0\&limit\=10 | jq
[
  {
    "ID": 1,
    "CreatedAt": "2022-10-03T10:40:35.447Z",
    "UpdatedAt": "2022-10-03T10:40:36.855Z",
    "DeletedAt": null,
    "Title": "Album One",
    "ReleaseYear": 2000,
    "Tracks": null
  }
]

```
</details>

<details>
  <summary>GET /album/:id - Returns a single album with a collection of tracks, based on its unique id</summary>

```
➜  ~ curl -s http://127.0.0.1:8080/album/1 | jq
{
  "ID": 1,
  "CreatedAt": "2022-10-03T10:40:35.447Z",
  "UpdatedAt": "2022-10-03T10:40:36.855Z",
  "DeletedAt": null,
  "Title": "Album One",
  "ReleaseYear": 2000,
  "Tracks": [
    {
      "ID": 1,
      "CreatedAt": "2022-10-03T10:22:55.079Z",
      "UpdatedAt": "2022-10-03T10:22:56.58Z",
      "DeletedAt": null,
      "Title": "Track One",
      "ArtistID": 1,
      "Users": null,
      "Albums": null
    },
    {
      "ID": 2,
      "CreatedAt": "2022-10-03T10:22:55.079Z",
      "UpdatedAt": "2022-10-03T10:22:56.58Z",
      "DeletedAt": null,
      "Title": "Track Two",
      "ArtistID": 1,
      "Users": null,
      "Albums": null
    }
  ]
}
```
</details>

### Track

<details>
  <summary>GET /track - Returns a list of tracks.</summary>

Query parameters:
- `offset`: mandatory - Sets the offset in the select query;
- `limit`: mandatory -  Set the fetched records limit in the select query;
- `query`: optional - A free text field compared to the `name` column.

```
➜  ~ curl -s http://127.0.0.1:8080/track\?offset\=0\&limit\=10 | jq
[
  {
    "ID": 1,
    "CreatedAt": "2022-10-03T10:22:55.079Z",
    "UpdatedAt": "2022-10-03T10:22:56.58Z",
    "DeletedAt": null,
    "Title": "Track One",
    "ArtistID": 1,
    "Users": null,
    "Albums": null
  },
  {
    "ID": 2,
    "CreatedAt": "2022-10-03T10:22:55.079Z",
    "UpdatedAt": "2022-10-03T10:22:56.58Z",
    "DeletedAt": null,
    "Title": "Track Two",
    "ArtistID": 1,
    "Users": null,
    "Albums": null
  }
]
```

```
➜  ~ curl -s http://127.0.0.1:8080/album\?query\=One\&offset\=0\&limit\=10 | jq
[
  {
    "ID": 1,
    "CreatedAt": "2022-10-03T10:40:35.447Z",
    "UpdatedAt": "2022-10-03T10:40:36.855Z",
    "DeletedAt": null,
    "Title": "Album One",
    "ReleaseYear": 2000,
    "Tracks": null
  }
]

```
</details>

<details>
  <summary>GET /track/:id - Returns a single track with a collection of albums, based on its unique id</summary>

```
➜  ~ curl -s http://127.0.0.1:8080/track/1 | jq 
{
  "ID": 1,
  "CreatedAt": "2022-10-03T10:22:55.079Z",
  "UpdatedAt": "2022-10-03T10:22:56.58Z",
  "DeletedAt": null,
  "Title": "Track One",
  "ArtistID": 1,
  "Users": null,
  "Albums": [
    {
      "ID": 1,
      "CreatedAt": "2022-10-03T10:40:35.447Z",
      "UpdatedAt": "2022-10-03T10:40:36.855Z",
      "DeletedAt": null,
      "Title": "Album One",
      "ReleaseYear": 2000,
      "Tracks": null
    }
  ]
}
```
</details>

### User

<details>
  <summary>GET /user - Returns a list of users.</summary>

Query parameters:
- `offset`: mandatory - Sets the offset in the select query;
- `limit`: mandatory -  Set the fetched records limit in the select query;
- `query`: optional - A free text field compared to the `name` column.

```
➜  ~ curl -s http://127.0.0.1:8080/user\?offset\=0\&limit\=10 | jq
[
  {
    "ID": 1,
    "CreatedAt": "2022-10-03T10:50:47.013Z",
    "UpdatedAt": "2022-10-03T10:50:48.412Z",
    "DeletedAt": null,
    "Login": "login1",
    "Tracks": null
  },
  {
    "ID": 2,
    "CreatedAt": "2022-10-03T10:50:47.013Z",
    "UpdatedAt": "2022-10-03T10:50:48.412Z",
    "DeletedAt": null,
    "Login": "login2",
    "Tracks": null
  }
]

```

```
➜  ~ curl -s http://127.0.0.1:8080/user\?query\=1\&offset\=0\&limit\=10 | jq
[
  {
    "ID": 1,
    "CreatedAt": "2022-10-03T10:50:47.013Z",
    "UpdatedAt": "2022-10-03T10:50:48.412Z",
    "DeletedAt": null,
    "Login": "login1",
    "Tracks": null
  }
]

```
</details>

<details>
  <summary>GET /user/:id - Returns a single track with a collection of liked tracks, based on its unique id</summary>

```
➜  ~ curl -s http://127.0.0.1:8080/user/1 | jq 
{
  "ID": 1,
  "CreatedAt": "2022-10-03T10:50:47.013Z",
  "UpdatedAt": "2022-10-03T10:50:48.412Z",
  "DeletedAt": null,
  "Login": "login1",
  "Tracks": [
    {
      "ID": 1,
      "CreatedAt": "2022-10-03T10:22:55.079Z",
      "UpdatedAt": "2022-10-03T10:22:56.58Z",
      "DeletedAt": null,
      "Title": "Track One",
      "ArtistID": 1,
      "Users": null,
      "Albums": null
    },
    {
      "ID": 2,
      "CreatedAt": "2022-10-03T10:22:55.079Z",
      "UpdatedAt": "2022-10-03T10:22:56.58Z",
      "DeletedAt": null,
      "Title": "Track Two",
      "ArtistID": 1,
      "Users": null,
      "Albums": null
    }
  ]
}
```
</details>

<details>
  <summary>POST /user/:id/track/:id - Appends a track to a user, and returns the updated user entity</summary>

```
➜  ~ curl -X POST -s http://127.0.0.1:8080/user/1/track/2 | jq
{
  "ID": 1,
  "CreatedAt": "2022-10-03T10:50:47.013Z",
  "UpdatedAt": "2022-10-03T09:05:27.480880658Z",
  "DeletedAt": null,
  "Login": "login1",
  "Tracks": [
    {
      "ID": 2,
      "CreatedAt": "2022-10-03T10:22:55.079Z",
      "UpdatedAt": "2022-10-03T10:22:56.58Z",
      "DeletedAt": null,
      "Title": "Track Two",
      "ArtistID": 1,
      "Users": null,
      "Albums": null
    }
  ]
}

```
</details>