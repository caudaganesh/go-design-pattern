# Factory

## Description

The purpose of this pattern is to abstract the user from the knowledge of the creation of the implementation of an interface. With the Factory method, we delegate the creation of families of objects to a different package or object to abstract us from the knowledge of the pool of possible objects we could use. 

## Usage
 - When you have different logic for a same method, i.e. you may have `authorize` method, but the logic is different for each implementations.
 - To abstract the instance creation of an object, i.e. you may have a `shape` interface, and you want to create multiple implementation for it, it can be different instance of the `shape`, `circle`, `rectangle`, and so on. 
 - Etc