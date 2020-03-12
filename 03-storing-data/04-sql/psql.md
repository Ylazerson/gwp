# B"H


### Prerequisites:
1. Install local PostgreSQL.


### Some PostgreSQL CLI from shell:

```sh
which psql

# -- -----------------------------------
# Create DB user:
#     -P : prompt for 
#     -d : allow gwp to create databases
createuser -P -d gwp

# Create DB:
createdb gwp

# Execute SQL script:
#    -h is the hostname/IP of the local server, thus avoiding Unix domain sockets
psql -h 127.0.0.1 -U gwp -f setup.sql -d gwp
```

---

### Some PostgreSQL CLI from within `psql`:
Run `psql` then can use following:

```sh

# Show databases:
\l

# Show tables in the current database:
\dt 


# Switch to another database:
\c <database_name>

# Describe a particular table:
\d <table_name>

# Ger version:
select version();

# Show all the available commands:
\?

# Exit:
\q
```


---

### Get this Go module up and running:


```sh
# Create a new module (initialize the go.mod file) 
go mod init sandbox/testsql

go run store.go
```

---

### Notes on GO SQL Usage:
- Go thru chapter 6 for details.
- Here will just be some notes.

Notes:
- The `init` function is called automatically for every package.
- The `sql.DB` struct is a handle to the database and represents a **pool** of zero or database connections that’s maintained by the `sql` package. 
- The connection will be set up **lazily** when it’s needed.
- Notice that when you import the database driver, you set the name of the package to be an underscore `_`. 
    - This is because you shouldn’t use the database driver directly; you should use `database/sql` only. 
    - This way, if you upgrade the version of the driver or change the implementation of the driver, you don’t need to make changes to all your code.

---

#### `QueryRow` 
- Returns only a single reference to an `sql.Row` struct. 
- If more than one `sql.Row` is returned by the SQL statement, only the first is returned by `QueryRow`. 
- The rest are discarded. 

---

#### `Exec`
- Use when no need to returns query results. 
- The result just gives the number of rows affected and the last inserted id, if applicable. 