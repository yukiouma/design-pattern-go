# Mediator

> Mediator is a behavioral design pattern that lets you reduce chaotic dependencies between objects. The pattern restricts direct communications between the objects and forces them to collaborate only via a mediator object
>
>  * Use the Mediator pattern when it’s hard to change some of the classes because they are tightly coupled to a bunch of other classes.
>  * Use the pattern when you can’t reuse a component in a different program because it’s too dependent on other components.
>  * Use the Mediator when you find yourself creating tons of component subclasses just to reuse some basic behavior in various contexts.

## Assumption
Imagine  you are building an text editor application, including 4 parts:
* text area: edit text
* clean button: clean all text
* lock buttion: freeze/unfreeze contents in text area

text area has 2 properties:
    content: to save text contents

clean botton:
    empty contents in text area
    if lock is true in text area, clean button should not work as expect

lock button:
    lock: if text area can be edited
    change lock to true(or false)


