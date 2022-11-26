# Template Method

> Template Method is a behavioral design pattern that allows you to defines a skeleton of an algorithm in a base class and let subclasses override the steps without changing the overall algorithm’s structure.

## Assumption
Imagine  you are building a game, including 3 roles, but has the same actions:
    * attack
    * defend

3 roles:
    * warrior, has following property:
        power and health point
        warrior attack other role can cause damage of its power
        if warrior is attacked by other role will be substract health point according damage cause by other role
    * knight
        power, defense and health point
        knight attack other role can cause damage of its power
        if knight is attacked by other role will be substract health point according difference between its defense and damage cause by other role
    * magician
        power, mana and health point
        magician attack other role can cause damage of its power + mana
        if magician is attacked by other role will be substract health point according damage cause by other role