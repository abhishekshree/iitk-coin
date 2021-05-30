## Task 1

To write a program that connects to a database(SQLite).

<u>Packages used:</u>

- [database/sql](https://golang.org/pkg/database/sql/)
- [go-sqlite3](https://github.com/mattn/go-sqlite3)

Note: [Nice reason for that underscore](https://stackoverflow.com/questions/21220077/what-does-an-underscore-in-front-of-an-import-statement-mean)

---

<u>Functions implemented</u>:

1. ```func add(db *sql.DB, rollno string, name string) bool``` : Adds a user to the table.

2. ```func clean(db *sql.DB) bool```: Drops the table for a fresh run.
