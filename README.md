# WUS
Wid URL Shortener

## Getting Started

Compile the project:

    go build -v -o wus

Create/migrate/seed the database:

    ./wus resetDB

SQLite is default if the environment variable `DATABASE_URL` is not set. PostgreSQL is used otherwise.

Run the app (on `localhost:3000` by default):

    ./wus

## Main models

### Users

```go
type User struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Password  string  `json:"password"`
	Email     string  `json:"email"`
	AccountID int     `json:"accountId" sql:"index"`
	Account   Account `json:"account,omitempty"`
}
```

### Accounts

```go
type Account struct {
	Users        []User        `json:"users,omitempty"`
	Sessions     []Session     `json:"sessions,omitempty"`
	RoleMappings []RoleMapping `json:"roleMappings,omitempty"`
}
```

### Urls

```go
type Url struct {
	ShortHandle string         `json:"shortHandle" sql:"unique"`
	LongUrl     string         `json:"longUrl"`
	Notes       string         `json:"notes"`
	Enabled     bool           `json:"enabled"`
	AccountID   int            `json:"accountId" sql:"index"`
	Account     domain.Account `json:"account,omitempty"`
}
```

## Auth token

To authenticate your requests to the API, the session token must be passed with the Authorization header.

    Authorization: AuthToken authToken=OzWe9nWpOzRLis2E9XO720viC6S1gn7


## API documentation

The Swagger API documentation is available on `/api/explorer`.
