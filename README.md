# Email Database Indexer App

This project consists of a web application that indexes an email database and allows searching for information within it. It is divided into several parts, each with its specific objective.

## Problem Definition

The main goal of this application is to create an interface for searching information in email databases. The project is divided into the following parts:

### Part 1: Email Database Indexing

In this stage, we download the Enron Corp email database and then create a program that indexes its contents using the ZincSearch tool.

### Part 2: Profiling

We profile our indexer using the tools provided by Go. Then, we generate a chart to visualize the application's performance.

### Part 3: Visualizer

Create a simple interface that allows users to search and query the contents of the email database.

### Optional Part 4: Optimization

We use the information obtained from the profiling in Part 2 to optimize the code.

## Technologies Used

- Backend Language: Go
- Database: ZincSearch
- API Router: chi
- Interface: Vue 3
- CSS: Tailwind

**Note:** No other external libraries were used in the backend.

## How to Run the Application

Below are instructions to execute each part of the application.

### Part 1: Email Database Indexing

To index the Enron Corp email database, follow these steps:

1. Download the email database from [http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz) (423MB).

2. Run the following command:

```bash
$ ./indexer enron_mail_20110402
```


### Part 2: Profiling
To profile the application, follow the instructions at https://go.dev/doc/diagnostics.

### Part 3: Visualizer
To run the search interface, use the following command:
```bash
$ ./mail-indexer -port your_port
```

The search application will be available at [http://localhost:tu_puerto](http://localhost:your_port).

### Optional Part 4: Optimization
Use the profiling information obtained in Part 2 to optimize your code.


<<<<<<< HEAD
## Contacto

Para cualquier pregunta o comentario, puedes ponerte en contacto conmigo: [correo electrónico](mailto:arivalladares2.0@gmail.com).
=======
## Contact
For any questions or comments, you can get in touch with me through: [my email](mailto:arivalladares2.0@gmail.com).
>>>>>>> 987c1bf (Updated readme, translated to english)



