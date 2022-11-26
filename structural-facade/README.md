# Facade

## Assumption
Imagine you want to provide a method to client to read and print file contents in console.
We can use io package in standrad liabray.
However, there is limited method we need to use in io package, and we have to expose all method from io pakcage to client
at this time, we can use facade pattern to wrap io package into an structral, avoiding expose all method from io package to client