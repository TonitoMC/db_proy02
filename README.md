# Proyecto 2. ACID y Concurrencia
Este proyecto consiste en evaluar diferentes niveles de aislamiento para transacciones dentro de Postgres, realizando una simulacion de usuarios concurrentes intentando reservar un mismo asiento.

## Notas

- Ejecutar triggers antes de insertar datos, auto-llena los asientos con precios basado en el venue y otros valores insertados (me ahorre bastantes inserciones)
- El script solo se puede ejecutar 1 vez, luego hay que volver a correr los scripts de SQL porque tengo los numeros de los asientos 'quemados' en el codigo.
- Los queries hacen rollback cuando encuentran un error, no se realiza toda la transaccion y luego se espera a que de un error en el commit. Obtuve resultados similares de ambas maneras pero me parecio mejor practica hacerlo de esta manera.
- Los logs de los errores estan comentados porque el output estaba horrible, se pueden descomentar pero es muy poco legible.
- No implemente restricciones en cuanto a la reserva de asientos en la DB, estoy muy al tanto que con un unique en las reservas de asientos se evitan bastantes problemas pero parte del chiste es que truene el sistema.
- 

## Ubicacion de los Archivos
- Docs: Diagrama ERD
- SQL: Los scripts DDL y SQL para crear tablas, insertar datos y agregar triggers. + un nuke.sql porque se me hace de utilidad v:
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
