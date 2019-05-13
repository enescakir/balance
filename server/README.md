## HTTP Server
It's a simple HTTP server that shows example usage for `balance` package.

It checks parenthesis balance of given string and save request history to memory or MySQL database for calculating some metrics.

![Dashboard](https://github.com/EnesCakir/balance/blob/master/dashboard.png)

### Usage
```shell
# Enter to server directory
$ cd server

# If you want to use in memory database, you don't have to copy config file. 
# Default: in memory database
# Copy config.example.json to config.json
$ cp config/config.example.json config/config.json
# Put database credentials into config.json

# Download goland MySQL driver
$ go get -u github.com/go-sql-driver/mysql

# Run server 
$ go run main.go

# Visit `http://localhost:8080/`
```

#### Testing
**Test coverage:** ~95%

`querylog.MysqlRepository` and `database.db` tests require MySQL connection.

Create `config/config.mysql.json` file, set driver to `mysql` then put your database credentials.

You can use `go test ./...` command for testing.

#### Docker
You can use balance server with Docker

**Notice:** `Dockerfile` is at project root, not the server root. Because this repository isn't public right now. 
When we put it to the server root, it can't download `github.com/enescakir/balance` package.

```shell
###
# Running balance server with MySQL server
$ docker-compose -f docker-compose.yml up balance

# Visit `http://localhost:8080/`

###
# Running balance server with another MySQL server
# Build docker image
$ docker build -t balance .

# Run docker image with your MySQL database credentials
$ docker run -it -p 8080:8080 \
    -e DATABASE_HOST=127.0.0.1 \
    -e DATABASE_PORT=3306 \
    -e DATABASE_NAME=balance \
    -e DATABASE_USER=root \
    -e DATABASE_PASSWORD=secret \
    balance
```

### Endpoints     
#### **POST** /isbalanced

Checks the parentheses balance of given expression 
```json
// Example request body:
{
    "expr": "{[()]}"
}
```
```json
// Response for "{[()]}":
{
    "valid": true
}

// Response for "[())]":
{
    "valid": false,
    "error": "Mismatch at index: 3"
}
```

#### **GET** /logs

Returns the collection of logs at given date range.

It accepts `start` and `end` as query parameter. They are optional.

*Example:* `/logs?start=2019-05-11+12:17:00`

```json
// Example response body
[
  {
    "id": 148,
    "query": "[()]",
    "status": 1,
    "response_time": 230647,
    "created_at": "2019-05-12T12:58:52Z"
  }, {
    "id": 147,
    "query": "[(()]",
    "status": 2,
    "response_time": 89170,
    "created_at": "2019-05-12T12:57:39Z"
   }
]
```

#### **GET** /logs/status

Returns the collection of status:count pairs at given date range.

It accepts `start` and `end` as query parameter. They are optional.

*Example:* `/logs/status?start=2019-05-11+12:17:00`

```json
// Example response body
[
    {
        "status": 1,
        "count": 6
    },
    {
        "status": 2,
        "count": 2
    }
]
```

#### **GET** /logs/histogram

Returns the collection of label:responseTime bins at given date range.
Labels are milliseconds.

It accepts `start` and `end` as query parameter. They are optional.

*Example:* `/logs/histogram?start=2019-05-11+12:17:00`

```json
// Example response body
[
    {
        "label": "0-10",
        "count": 3
    },
    {
        "label": "10-20",
        "count": 4
    },
    {
        "label": "20-30",
        "count": 2
    },
    {
        "label": "30-40",
        "count": 3
    },
    {
        "label": "40-50",
        "count": 5
    }
]
```

### Credits
- [Bootstrap 4](https://getbootstrap.com)
- [jQuery](https://jquery.com)
- [Date Range Picker](http://www.daterangepicker.com)
- [Moment.js](https://momentjs.com)
- [Chart.js](https://www.chartjs.org)

