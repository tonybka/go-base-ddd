# Fundamental DDD elements in Golang
Fundamental elements of DDD implementation in Golang microservice




# Components
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


# Q&A
## Why UUIDv1 for entity's identity?
### Reasons (short)
- Need to have random, unique identity values for entities
- To decouple domain layer from persistence layer, we should not use table row's id for entity's identity
- MySQL support storing UUIDv1 as BINARY(16) in database which saves storage space
- Comparing UUIDv1 values is faster than comparing values of the other versions (they have common dash-separated groups if they are generated on a same machine)

### References
- [Making UUIDs More Performant in MySQL](https://emmer.dev/blog/making-uuids-more-performant-in-mysql/)
- [Storing UUID Values in MySQL Tables](https://dev.mysql.com/blog-archive/storing-uuid-values-in-mysql-tables/)
- [GUID/UUID Performance Breakthrough](http://mysql.rjweb.org/doc.php/uuid)
- [Storing UUID Values in MySQL](https://www.percona.com/blog/2014/12/19/store-uuid-optimized-way/)