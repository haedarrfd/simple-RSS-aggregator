## RSS Aggregator App

An RSS aggregator is a tool that allows you to view and read all the latest posts from your favorite blogs and websites in one place. This simple RSS Aggregator project was built in Go programming language that allows users to create, retrieve, and manage feeds using REST APIs. It uses PostgreSQL as the database and implements authentication using API keys.

## 🔧 Requirements

To run this project, you must have Go, PostgreSQL, Goose and Sqlc installed on your machine.

- Go: Version 1.20 or later
- PostgreSQL: 15.0 or later
- Goose: 3.22 or later
- Sqlc: 1.27 or later

Make sure that PostgreSQL service is running and you should have the necessary credentials for database connection (e.g., username, password, and database name).

## 🔌 Installation

- **If you haven't install goose, use this command to install it (see the documentation [goose](http://pressly.github.io/goose/))**

```
$ go install github.com/pressly/goose/v3/cmd/goose@latest
```

- **If you haven't install sqlc, use this command to install it (see the documentation [sqlc](https://docs.sqlc.dev/en/stable/index.html))**

```
$ go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

- **Clone this repository**

```
$ git clone https://github.com/haedarrfd/simple-rss-aggregator.git
$ cd simple-rss-aggregator
```

- **Adds any missing module requirements**

```
$ go mod tidy
```

- **Copy the example `.env` file**

```
$ cp .env.example .env
```

**Note**: If the PostgreSQL database user does not have a password, just leave it blank.

- **Run database migrations**

```
$ cd postgresql/schema
$ goose up
```

- **Run the project**

```
$ go run main.go
```

**Or**

- **Build and run the project**

```
$ go build && ./simple-rss-aggregator
```

- **Use the project**
  To use the project you can use tools like _[Postman](https://www.postman.com/)_ or any other tools to interact with the API.

## 📖 API Reference

#### Create a user

```http
  POST /v1/users
```

| Parameter | Type     | Description             |
| :-------- | :------- | :---------------------- |
| `name`    | `string` | **Required**. Your name |

#### Get a user by the API key

```http
  GET /v1/users
```

| Parameter       | Type  | Description                |
| :-------------- | :---- | :------------------------- |
| `authorization` | `key` | **Required**. Your API key |

#### Get item

```http
  GET /api/items/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### add(num1, num2)

Takes two numbers and returns the sum.

## 📁 File Structure

```
.
├── internal
│   ├── auth
│   │   └── auth.go
│   ├── database
│   │   ├── db.go
│   │   ├── feed_follows.sql.go
│   │   ├── feeds.sql.go
│   │   ├── models.sql.go
│   │   ├── posts.sql.go
│   │   └── users.sql.go
├── postgresql
│   ├── queries
│   │   ├── feed_follows.sql
│   │   ├── feeds.sql
│   │   ├── posts.sql
│   │   └── users.sql
│   ├── schema
│   │   ├── 001_users.sql
│   │   ├── 002_feeds.sql
│   │   ├── 003_feed_follows.sql
│   │   └── 004_posts.sql
├── vendor
│   ├── github.com
│   │   ├── go-chi
│   │   │   ├── chi
│   │   │   └── cors
│   │   ├── google
│   │   │   └── uuid
│   │   ├── joho
│   │   │   └── godotenv
│   │   ├── lib
│   │   │   └── pq
│       └── modules.txt
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── handler_feed_follows.go
├── handler_feed.go
├── handler_posts.go
├── handler_user.go
├── json.go
├── main.go
├── middleware_auth.go
├── models.go
├── README.md
├── rss.go
├── scraper.go
└── sqlc.yaml
```
