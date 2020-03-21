# Composite

## Description

In the Composite design pattern, you will create hierarchies and trees of objects. Objects have different objects with their own fields and methods inside them. This approach is very powerful and solves many problems of inheritance and multiple inheritances.


## Usage
 - When you need to group multiple interfaces, let's say we have a case that `user` and `auth` use case, are commonly used together, we keep them separate because they have different responsibilities, but we can group it when it's commonly used together.
 - Composition also works with hierarchical structs, for example, we have a parent-child relationship for category of users.  
 - Etc