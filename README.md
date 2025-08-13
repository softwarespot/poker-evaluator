# Poker evaluator

![Go Tests](https://github.com/softwarespot/poker-evaluator/actions/workflows/go.yml/badge.svg)

This project showcases my proficiency in Go by creating a clear and readable
poker hand evaluator. The primary focus is on writing clean, maintainable code
that effectively demonstrates the logic behind evaluating various poker hands.

## Game play

- Provide 2 poker hands with 5 cards in each hand.
- Each card must start with the rank and then the suit e.g. `5C` as in the
  `5 of Clubs`.
- Cards are case-insensitive. Therefore, `aH` is the same as `Ah` or `AH`.

### Ranks (values)

- 2-9 are the same, e.g. `2` or `8`
- 10 is `T`
- Jack is `J`
- Queen is `Q`
- King is `K`
- Ace is `A`

### Suits

- Clubs are `C`
- Diamonds are `D`
- Hearts are `H`
- Spades are `S`

## Prerequisites

- Go 1.25.0 or above
- make (if you want to use the `Makefile` provided)
- Docker

## Dependencies

**IMPORTANT:** No 3rd party dependencies are used.

I could easily use [Cobra](https://github.com/spf13/cobra) (and usually I do,
because it allows me to write powerful CLIs), but I felt it was too much for
such a tiny project. I only ever use dependencies when it's say an adapter for
an external service e.g. Redis, MySQL or Prometheus.

## Run not using Docker

```bash
go run . "AH AC AD QC QD" "QD QH QC AD AC"
```

Output the result as JSON, instead of text

```bash
go run . --json "AH AC AD QC QD" "QD QH QC AD AC"
```

or when using `make`

```bash
make

# As text
./bin/poker-evaluator "AH AC AD QC QD" "QD QH QC AD AC"

# As JSON
./bin/poker-evaluator --json "AH AC AD QC QD" "QD QH QC AD AC"
```

### Version

Display the version of the application and exit.

```bash
# As text
./bin/poker-evaluator --version

# As JSON
./bin/poker-evaluator --json --version
```

### Help

Display the help text and exit.

```bash
./bin/poker-evaluator --help
```

## Run using Docker

1. Build the Docker image with the tag `poker-evaluator`.

```bash
docker build -t poker-evaluator .
```

2. Run the Docker image with the provided arguments.

```bash
# As text
docker run --rm poker-evaluator "AH AC AD QC QD" "QD QH QC AD AC"

# As JSON
docker run --rm poker-evaluator --json "AH AC AD QC QD" "QD QH QC AD AC"
```

### Version

Display the version of the application and exit.

```bash
# As text
docker run --rm poker-evaluator --version

# As JSON
docker run --rm poker-evaluator --json --version
```

### Help

Display the help text and exit.

```bash
docker run --rm poker-evaluator --help
```

## Tests

Tests are written as [Table-Driven Tests](https://go.dev/wiki/TableDrivenTests).

```bash
go test -cover -v ./...
```

or when using `make`

```bash
make test
```

### Coverage

For the `poker` package (found under `internal`), the test coverage is about
97%. This isn't to say there aren't any "bugs" 😀, but that it's better than 0%
and is an indication that it's passing the
[Poker rules](https://en.wikipedia.org/wiki/List_of_poker_hands).

```bash
go test -cover -v github.com/softwarespot/poker-evaluator/internal/poker
```

### Linting

Docker

```bash
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run --tests=false --default=none -E durationcheck,errorlint,exhaustive,gocritic,ineffassign,misspell,predeclared,revive,staticcheck,unparam,unused,whitespace --max-issues-per-linter=10000 --max-same-issues=10000
```

Local

```bash
golangci-lint run --tests=false --default=none -E durationcheck,errorlint,exhaustive,gocritic,ineffassign,misspell,predeclared,revive,staticcheck,unparam,unused,whitespace --max-issues-per-linter=10000 --max-same-issues=10000
```

## Additional information

This section documents any additional information which might be deemed important for the reviewer.

### Decisions made

- Despite using 1.25.0+ and the `slices` pkg being available, I have opted not
  to use it, and instead went for how I've been writing Go code before the
  `slices` pkg existed. Although for production code, I have started to use it
  where applicable.
- I haven't used an assertion library, as I have never used one in production
  code and have opted for creating my own simple test helpers.
- The `Hand.Compare()` func returns an `int` as an alias type of `Winner`, as I
  thought since it returns 3 exclusive values, that it should be similar to that
  of [https://pkg.go.dev/cmp#Compare].
- Naming is hard, so I have tried my best to name funcs, structs, variables
  etc... as best I can.

### Project setup

Commands used to setup the project's directory.

```bash
mkdir poker-evaluator
cd poker-evaluator
go mod init github.com/softwarespot/poker-evaluator
touch README.md
touch LICENSE
touch Dockerfile
```

### License

The code has been licensed under the [MIT](https://opensource.org/license/mit) license.
