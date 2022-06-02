# card-deck-api

An API that deals with decks of cards.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)

*i.e., no need to have Go / PostgreSQL installed locally*

Go to [Quick Start](#quick-start)!

## Endpoints

- POST /decks
  - create a new deck
  - optional query param:
    - shuffle
      - indicates whether the deck to be created should be shuffled or not
      - `value:` boolean
      - `default:` false
    - cards
      - if the deck to be created should only contain specific cards
      - `value:` comma-delimited card codes
      - `default:` standard 52-card deck
  - sample request:
    ```
    POST /decks?shuffle=true&cards=AS,KD,AC,2C,KH
    ```
- GET /decks/{id}
  - open a deck (i.e., get details of a particular deck)
  - sample request:
    ```
    GET /decks/e0fc3775-78a2-4217-a892-22ca8abce79d
    ```
- DELETE /decks/{id}/cards
  - draw card(s) from a deck
  - optional query param:
    - count
      - number of cards to be drawn
      - `value:` integer
      - `default:` 1
  - sample request:
    ```
    DELETE /decks/e0fc3775-78a2-4217-a892-22ca8abce79d/cards?count=2
    ```

## Usage

#### configure
```bash
$ make .env
```

* see generated `.env` file for configuration

#### tidy dependencies
```bash
$ make deps
```

#### run unit tests
```bash
$ make test
```

#### run all tests (unit + integration tests)
```bash
$ make testIncludeInt
```

#### build docker image of the API
```bash
$ make build
```

#### start local database
```bash
$ make startDB
```

#### stop docker-compose containers
```bash
$ make stop
```

### Quick Start
```bash
$ make quickStart
# this will make the API accessible on http://localhost:8080
# (i.e., http://localhost:8080/decks)
#
# NOTE: might take a while when run for the first time
```

### Helpers during development:

#### format all .go files in project (using go fmt)
```bash
$ make fmt
```

#### generate test mocks for all interfaces in project
```bash
$ make genMocks
```
