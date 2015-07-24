# WUS
Wid URL Shortener

## Getting Started

Compile the project:

    go build -v -o wus

Create/migrate/seed the database (SQLite is default if the environment variable `DATABASE_URL` is not set):

    ./wus resetDB

Run the app (on `localhost:3000` by default):

    ./wus

Enjoy the server freshly created.

## API documentation

The Swagger API documentation is available on `/api/explorer`.

## Documentation

See the [Zest](https://github.com/solher/zest) doc.
