# Expert System
Backward chaining inference engine - Scholar project
Written in Golang (go1.13 darwin/amd64)
# Compilation / Install

Use makefile : make / make run

# Usage

With a file :

./bin/expert_system ./examples/example_input.txt

or dynamically :

./expert_system

## input

example :

A+B=>C

=AB

?C

### rules

and : +

or : |

xor : ^

not : !

implies : =>

if only if : <=>

facts : [A-Z]

### initial facts

=[A-Z]

### query

?[A-Z]

# Authors

* **Fabien Blin** @ 42 Lyon
* **Johann Monnerie** @ 42 Lyon
