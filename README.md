# Actions Service

## Descripción
El servicio de likes gestiona la funcionalidad de "me gusta" en álbumes o fotos. Optimiza el cálculo y almacenamiento mediante Redis y genera eventos para notificaciones.

## Características
- Registro de likes en álbumes o fotos.
- Prevención de duplicación de likes por usuario en una misma publicación.
- Cálculo y actualización rápida del contador de likes sin recargar la base de datos.
- Uso de Redis para cachear los álbumes más populares.
- Generación de eventos para notificaciones (evento.like.nuevo).

## Tecnologías Utilizadas
- **Lenguaje**: Go
- **Base de Datos**: PostgreSQL
- **Cache**: Redis
- **Orquestación**: Docker
- **Mensajería**: NATS

## Instalación
1. Clona el repositorio:
   ```sh
   git clone <repo-url>
   cd likes-service
   ```
2. Configura las variables de entorno en `.env`.
3. Construye y ejecuta el servicio con Docker:
   ```sh
   docker-compose up --build -d
   ```

## Endpoints
| Método | Ruta          | Descripción |
|--------|--------------|-------------|
| POST   | /likes       | Da like a una publicación |
| DELETE | /likes/:id   | Remueve un like |
| GET    | /likes/count | Obtiene el conteo de likes de una publicación |
| GET    | /popular     | Obtiene las publicaciones más populares |

## Seguridad
- Autenticación mediante JWT.
- Protección contra duplicación de likes.
- Uso de Redis para optimización.

## Contribución
1. Realiza un fork del repositorio.
2. Crea una rama con tu feature (`git checkout -b feature-nueva`).
3. Haz commit de tus cambios (`git commit -m 'Agrega nueva funcionalidad'`).
4. Sube la rama (`git push origin feature-nueva`).
5. Abre un Pull Request.

## Licencia
Este proyecto está bajo la licencia MIT.
