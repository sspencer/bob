# BOB

Bob is a snippet organizer.  Use it to take notes organized under folders... with tags.

Bob is a vehicle to test technologies like SQLBoiler and best practices to writing APIs in Go.

## MySQL DSN

As Bob is not meant for production use, just dev fun, shortcuts were taken.  Instead of a configuration file, just set an ENV VAR for your (only choice for now...) MySQL DSN:

    export BOB_MYSQL="USER:PASSWORD@unix(/tmp/mysql.sock)/DATABASE_NAME?loc=Local&parseTime=true"

## Usage

To start the server, create the database:

    mysql -e 'create database bob'
    cat dist/schema.sql | mysql bob

And run the server:

    go run main.go

The server starts on `:8080`, but that can be changed with `go run main.go -p 1234`.

## API

### GET /v1/version

    curl http://localhost:8080/v1/version
    {"name":"bob","uptime":"4.910443246s","version":"0.0.1"}

