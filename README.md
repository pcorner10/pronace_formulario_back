# pronace_formulario_back
Este repositorio contendrá el código necesario para ejecutar un servidor para el manejo de formularios.
## Estructura basada en 
[Golang Project Structuring](https://medium.com/sellerapp/golang-project-structuring-ben-johnson-way-2a11035f94bc)


Explicación de los directorios y archivos:

*  cmd/: Contiene el punto de entrada de la aplicación, main.go, que inicializa y configura el servidor.
*  internal/: Contiene el código principal de la aplicación organizado por dominio.
*  app/: Contiene los controladores (handler/), servicios (service/) y casos de uso (usecase/) específicos de la aplicación.
*  handler/: Contiene los controladores que manejan las solicitudes HTTP y llaman a los casos de uso correspondientes.
*  repository/: Contiene las interfaces de repositorio que definen las operaciones de acceso a la base de datos para las entidades de dominio.
*  service/: Contiene los servicios de dominio que encapsulan la lógica de negocio y utilizan los repositorios.
*  usecase/: Contiene los casos de uso que representan los diferentes escenarios de uso de la aplicación y coordinan la ejecución de servicios y repositorios.
*  domain/: Contiene las entidades de dominio que representan los objetos centrales de la aplicación.
*  entity/: Contiene las estructuras de datos de las entidades de dominio.
*  repository/: Contiene las implementaciones de repositorio que interactúan con la base de datos para realizar operaciones de CRUD en las entidades de dominio.
*  service/: Contiene los servicios de dominio que encapsulan la lógica de negocio.
*  infrastructure/: Contiene la implementación de la infraestructura externa, como la base de datos, el servidor HTTP y el ORM.
*  database/: Contiene el código para establecer la conexión con la base de datos PostgreSQL.
*  http/: Contiene la configuración y la lógica del servidor HTTP utilizando el framework Gin.
*  orm/: Contiene el código



