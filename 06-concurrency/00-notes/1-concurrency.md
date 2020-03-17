# B"H


### Concurrency

> Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.
>
> -- <cite>Rob Pike, co-creator of Go</cite>

![](img/concurrency.png)

![](img/parallelism.png)

---

**Parallel** programs that have to run tasks at the same time will need the environment variable `GOMAXPROCS` to be set to more than 1. Since Go 1.5, `GOMAXPROCS` is set to the number of CPUs available in the system. 

**Concurrent** programs can run within a single CPU and tasks scheduled to run independently. 

What’s important to note now is that although Go can be used to create parallel programs, **it was created with concurrency in mind and not parallelism**.

---

Go's support for concurrency is supported by two main constructs:
1. `goroutines` 
2. `channels`. 

