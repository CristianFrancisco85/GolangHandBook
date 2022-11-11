# API REST y WebSockets

Esta es una implementacion con de una API REST con Goriila Mux y WebSockets, con la finalidad de poder realizar una comunicacion bidireccional entre el cliente y el servidor. Tambien se ejemplifica el uso de JWT para la autenticacion de usuarios.

Para la persistencia de datos se utiliza PostgreSQL y para la interaccion con la base de datos se uso el patron de diseño Repository esto para poder tener una abstraccion de la logica de la base de datos.

Todo el codigo esta documentado con comentarios para poder entender el funcionamiento de cada parte del codigo.

## Requisitos

- Go 1.18
- Docker
- Archivo .env con las siguientes variables de entorno

```bash
PORT=5050
JWT_SECRET=****
DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
```

