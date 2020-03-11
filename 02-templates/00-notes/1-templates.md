# B"H



### Templates and Template Engines

A **web template** is a predesigned HTML page that’s used repeatedly by a software program, called a **template engine**, to generate one or more HTML pages. 

![](img/template-engine.png)

![](img/go-templates.png)

---

- Go templates are text documents (for web apps, these would normally be HTML files), with certain commands embedded in them, called **actions**.  
- The text is parsed and executed by the **template engine** to produce another piece of text. 
- **Actions** are added between two double braces, `{{` and `}}` 

---

Note, for detailed information on Templates see chapter 8 of the book.
- Over here, the notes are going to be brief.
- See the code under the `02-templates` folder for examples.

---

Using the Go web template engine requires 2 steps:

1.  Parse the text-formatted template source, which can be a string or from a template file, to create a parsed template struct.

2.  Execute the parsed template, passing a ResponseWriter and some data to it. This triggers the template engine to combine the parsed template with the data to generate the final HTML that’s passed to the ResponseWriter.

---

#### Parsing templates

`ParseFiles` (both the function and the method) 
- It can take in one or more filenames as parameters (a variadic function). 
- But it still returns just one template, regardless of the number of files it’s passed. 
- When we pass in more than one file, the returned parsed template has the name and content of the first file. 
- The rest of the files are parsed as a map of templates, which can be referred to later on during the execution. 

---

#### Using `Must`

- The `Must` function wraps around a function that returns a pointer to a template and an error, and panics if the error is not a nil. 
- Panicking refers to a situation where the normal flow of execution is stopped, and if it’s within a function, the function returns to its caller. The process continues up the stack until it reaches the main program, which then crashes.

```go
t  := template.Must(template.ParseFiles("tmpl.html"))
```

---

#### Actions

Go has an extensive set of actions, which are quite powerful and flexible. In this section, we’ll discuss some of the important ones:

- **Conditional actions**
- **Iterator actions**
- **Set actions**
- **Include actions**

---

**Note** the **dot** `.` is an action, and it’s the most important one. 
- The dot is the evaluation of the data that’s passed to the template. 
- The other actions and functions mostly manipulate the dot for formatting and display.


---

#### Important Note
- In general, use **Nested Template Layouts** over **Include Actions** - they're much better!

