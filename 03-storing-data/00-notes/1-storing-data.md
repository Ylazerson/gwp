# B"H


### Storing data

Here are the general places where you would store data:
- In memory (while the program is running)
- In files on the filesystem
- In a database, fronted by a server program

---

####  In-memory Storage

- This refers to storing data in the running application itself
- For Go, this primarily means with `arrays`, `slices`, `maps`, and **most importantly**, `structs`.

- Very often we start off with using relational databases and then as we scale, we realize that we need to cache the data that we retrieve from the database in order to improve performance. 
- Instead of using an external in-memory database like Redis, we have the option of refactoring our code and storing the cache data in memory.

---

#### File Storage

The book explores two ways of storing data to files in Go:
1. Using CSV files. 
2. Using the `gob` package.
    - Gob is a **binary format** that can be saved in a file, providing a quick and effective means of serializing in-memory data to one or more files. 

See chapter 6 for details.