# B"H



---

### Prerequisite Step 1 - get this Go module up and running:


```sh
# Create a new module (initialize the go.mod file) 
go mod init sandbox/testrest
```

---

### Prerequisite Step 2: Follow instructions in `/03-storing-data/04-sql/psql/md` 

---



### Step 1 - Build Package:
```sh
go install
```


### Step 2: Start up the server:
```sh
testrest
```

---

## Run Tests

### Run basic CRUD tests in `curl`
- C: Create
- R: Read
- U: Update
- D: Delete


---
### CREATE

```sh

# Create:
curl -i -X POST -H "Content-Type: application/json"  -d '{"content":"My first post","author":"Sau Sheong"}' http://127.0.0.1:8080/post/

```


```sh
psql

\c gwp
```


```sql
select   *
from     posts
;
```


---
### READ

```sh

# Read
curl -i -X GET http://127.0.0.1:8080/post/1

```


---

### UPDATE

```sh

# Update:
curl -i -X PUT -H "Content-Type: application/json"  -d '{"content":"Updated post","author":"Sau Sheong"}' http://127.0.0.1:8080/post/1

```

```sql
select   *
from     posts
;
```


---

### DELETE

```sh

# Delete
curl -i -X DELETE http://127.0.0.1:8080/post/1

```