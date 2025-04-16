# Proyecto 2. ACID y Concurrencia
Este proyecto consiste en evaluar diferentes niveles de aislamiento para transacciones dentro de Postgres, realizando una simulacion de usuarios concurrentes intentando reservar un mismo asiento.

## Ubicacion de los Archivos
- Docs: Documentacion en cuanto al diseno de la base de datos, analisis y reflexion correspondiente al proyecto
- SQL: Los scripts DDL y SQL para crear tablas e insertar datos
- Simulation: Programa creado para simular consultas concurrentes a la base de datos
## Instrucciones para correr la simulacion
Para correr la simulacion es necesario tener variables de entorno para la conexion a la base de datos, se debe crear un archivo .env dentro del folder de simulation con el siguiente template:

```
DB_USER=usuario
DB_PASSWORD=password
DB_NAME=nombre
DB_HOST=localhost
DB_PORT=5432
```

La unica dependencia para el programa es tener instalado [Golang](https://go.dev), se utilizan librerias estandar. Se requiere navegar al directorio de Simulation y correr los siguientes comandos:
```
go mod tidy
```
```
go run main.go
```
Estos comandos se encargan de descargar las librerias necesarias y correr el programa
