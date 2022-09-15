# Composite

A mechanism for treating individual (scalar) objects and compositions of objects
in a uniform manner.


### Motivation

- Objects using other object's fields/methods through embedding
- Composition lets us make compound objects
- Is used to treat both single (scalar) and composite objects uniformly 
  - (they will have the same interface)
  - Foo and []Foo has to have the same interface for instance

Examples:
- mathematical expressions composed of simple expressions
- A shape group made of several shapes
