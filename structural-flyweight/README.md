# Flyweight
Also known as: Cache

> Flyweight is a structural design pattern that lets you fit more objects into the available amount of RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object.

## Assumption
Imagine you are designing a car race game:
* each player can select a car to join the game
* each car has two base properties:
    1. color
    2. brand
* during the game, each car has these mutable properties:
    1. location(x, y)
    2. speed km/h

the color and brand properties might be duplilcated a lot