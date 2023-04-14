# Fundamental DDD elements in Golang
Fundamental elements of DDD implementation in Golang microservice




# Components
## Application Layer
- DTOs, Mapper & Presenter
- Request handlers
- Event handlers
- Message Publisher
## Domain Layer
### Entity
- Mutable, has identity
### Value Object
- Immutable, no identity
### Domain Event
- Utility for communication between bounded contexts, sub-domains, and between microservices

## Persistence Layer
- Data Model
- Repository
