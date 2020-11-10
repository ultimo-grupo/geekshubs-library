# 📚 GeeksHubs Library Demo App

API REST para gestionar libros.

## Cómo funciona

La API se ha implementado con una pequeña aplicación en Go. Para persistir la información utiliza una Base de Datos MySQL.
Hay un archivo `docker-compose` que al ejecutarse levantará dos servicios:

* La aplicacíon, que escucha en el puerto `:8080`
* La Base de Datos MySQL que escucha en el puerto `:3306`

Si necesitas ralizar cambios en la aplicación y quieres ejecutarla directamente en tu máquina local necesitas instalar el intérprete de Go y después (desde la carpeta `app`)

* **IMPORTANTE:** Deberás ir al archivo `repository.go` y cambiar los parámetros de conexión de MySQL

`make run`

y para ejecutar los tests

`make test`

## API 

La documentación de la API se encuentra [aquí](https://documenter.getpostman.com/view/255227/TVejgpWn)
