# gorestable

gorestable is a simple go reservation table system

## What is the scope of this project?

gorestable wants to provide a simple table reservation system.
AAA (Authorization, Accounting, Authentication - reorder as you wish) is *intentionally* put out
of the scope. gorestable *intentionally* trusts the messages it receives, and takes them verbatim.
gorestable is intended to work as _backend_. A separate _frontend_ is expected to take care of
AAA and to craft well-formed messages.

## What is not in scope for this project?

* AAA
* handling *thousands* of reservations [1]
* handling *thousands*  of operations per second [1]

[1] but hey, if you have that kind of load and you are still interested in gorestable, get in
touch with us!

## Current Status

Initial development

## Documentation

TODO, look under docs/

## License

MIT

## Author(s):
    Francesco Romani fromani - gmail
    Alessio Pascucci feeder81 - gmail
