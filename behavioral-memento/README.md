# Memento

> Memento is a behavioral design pattern that lets you save and restore the previous state of an object without revealing the details of its implementation.

## Assumption
Imagine  you are creating a text editor which can take snapshot and get back to the state of latest snapshot

there are 3 object need to be created:
* text editor
    contains following fields: text, x, y, and selectionWidth
    owns following methods: setText, setCursor, setSelectionWidth and createSnapshot
* snapshot:
    contains the same fields to text editor, and link to text editor who created itself
    owns method restore let editor restore this state
* Command:
    manage snapshots, execute commands of creating snapshot and restore snapshot