# Adapter

## Description

Adapter pattern is a pattern that allows you to wrap an implementation with modified objects. By adding this layer, we can maintain an open closed principle in case we want to use and outdated interface and not possible to replace it easily. It's also being useful when you have multiple usage, but each usage has different way to construct the object we need to execute a use case.


## Usage
 - When you have multiple delivery, grpc and rest for example, will have different way to construct the request. Here we can use the adapters to interact with the use case, and let the adapter construct the response, that can be used for the delivery.
 - When you have to use an old interface, that realty difficult to change and you want to keep the open close principle
 - Etc