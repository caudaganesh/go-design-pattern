![alt text](https://codecov.io/gh/caudaganesh/go-design-pattern/branch/master/graph/badge.svg)

# go-design-pattern

Contains all the design pattern implementation written in go, alongside the Clean Architecture

## Clean Architecture

Source : [cleancoders](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

According to Uncle Bob (Robert C. Martin), the Clean Architecture should achieve separation by dividing the software into layers. Each has at least one layer for business rules, and another for interfaces. In result it should produce system with these specifications:

1. Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
2. Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
3. Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
4. Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
5. Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

###  Clean Architecture Diagram
![alt text](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)
