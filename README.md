# Email Database Indexer App

Este proyecto consiste en una aplicación web que indexa una base de datos de correos electrónicos y permite buscar información en ella. Está dividido en varias partes, cada una con su objetivo específico.

## Definición del Problema

El objetivo principal de esta aplicación es crear una interfaz para buscar información en bases de datos de correos electrónicos. El proyecto se divide en las siguientes partes:

### Parte 1: Indexar Base de Datos de Correo Electrónico

En esta etapa, descargamos la base de datos de correos de Enron Corp y luego creamos un programa que indexa sus contenidos utilizando la herramienta ZincSearch.

### Parte 2: Profiling

Realizamos el perfilado de nuestro indexer utilizando las herramientas proporcionadas por Go. Luego, generamos un gráfico para visualizar el rendimiento de la aplicación.

### Parte 3: Visualizador

Creacion de interfaz simple que permite a los usuarios buscar y consultar los contenidos de la base de datos de correos electrónicos.

### Opcional Parte 4: Optimización

Utilizamos la información obtenida del perfilado en la Parte 2 para optimizar el código.

## Tecnologías Utilizadas

- Lenguaje Backend: Go
- Base de Datos: ZincSearch
- API Router: chi
- Interfaz: Vue 3
- CSS: Tailwind

**Nota:** No se usaron otras librerías externas en el backend.

## Cómo Ejecutar la Aplicación

A continuación, se proporcionan instrucciones para ejecutar cada una de las partes de la aplicación.

### Parte 1: Indexar Base de Datos de Correo Electrónico

Para indexar la base de datos de correos de Enron Corp, sigue estos pasos:

1. Descarga la base de datos de correos desde [http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz) (423MB).

2. Ejecuta el siguiente comando:

```bash
$ ./indexer enron_mail_20110402
```


### Parte 2: Profiling

Para realizar el perfilado de la aplicación, sigue las instrucciones en [https://go.dev/doc/diagnostics](https://go.dev/doc/diagnostics).

### Parte 3: Visualizador

Para ejecutar la interfaz de búsqueda, utiliza el siguiente comando:
```bash
$ ./mail-indexer -port tu_puerto
```

La aplicación de búsqueda estará disponible en [http://localhost:tu_puerto](http://localhost:tu_puerto).

### Opcional Parte 4: Optimización

Usa la información de perfilado obtenida en la Parte 2 para optimizar tu código.


## Contacto

Para cualquier pregunta o comentario, puedes ponerte en contacto conmigo: [correo electrónico](mailto:arivalladares2.0@gmail.com).



