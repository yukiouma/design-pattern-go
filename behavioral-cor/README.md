# Chain of Responsibility

> Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

* all handler classes implement the same interface

## Assumption
Imagine you are building an app with roles
one person can be assigned one role
some opertaion can be done by some specify roles

roles list:
    * reader
    * appender
    * updater
    * deleter

operation and correspoding roles:
    * delete a file: deleter only
    * update a file: updater + deleter
    * append contents to a file: appender + updater + deleter
    * read a file: all roles can read files

we're going to set a series handler to make these decisions, each handler only make one of it. 
when one handler decide this user can not do this operation:
    * if it's not the end of the chain, pass the parameter to next handler
    * if it's the end of the chain, return false result directly