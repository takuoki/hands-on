# sql-injection

## get started

1. Run database

```cmd
$ docker compose up
```

2. Run application

```cmd
$ go run main.go
```

3. Normal request

```cmd
$ curl 'localhost:1323/employees?department=総務'
[{"id":1,"name":"大庭 勇吉","position":"一般社員","department":"総務","salary":300000},{"id":3,"name":"生田 信吉","position":"課長","department":"総務","salary":500000}]
```

4. Malicious request

```cmd
$ curl "localhost:1323/employees?department='or''='" 
[{"id":1,"name":"大庭 勇吉","position":"一般社員","department":"総務","salary":300000},{"id":2,"name":"野崎 竜夫","position":"係長","department":"経理","salary":400000},{"id":3,"name":"生田 信吉","position":"課長","department":"総務","salary":500000},{"id":4,"name":"浜野 一子","position":"部長","department":"経理","salary":820000},{"id":5,"name":"川崎 知治","position":"部長","department":"人事","salary":800000}]
```

## What's wrong

* SQL injection vulnerabilities

```go
	rows, err := h.db.Query(
		"SELECT id, name, position, department, salary " +
			"FROM employees " +
			"WHERE secret = FALSE AND department = '" + dep + "'")
```

* Fixed code

```go
	rows, err := h.db.Query(
		"SELECT id, name, position, department, salary " +
			"FROM employees " +
			"WHERE secret = FALSE AND department = $1", dep)
```

## References

- [SQLインジェクションの説明](https://www.ipa.go.jp/security/vuln/websecurity-HTML-1_1.html)
- [安全なSQLの呼び出し方](https://www.ipa.go.jp/files/000017320.pdf)
- [sql - The Go Programming Language](https://pkg.go.dev/database/sql#DB.Query)
