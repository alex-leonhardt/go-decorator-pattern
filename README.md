## golang and the decorator pattern

a playground for me (and maybe helpful for others?) to train the brain understand the decorator pattern better

- a basic function decorator, see [example1](example1/)
- a basic timing decorator (how long does func X take?), see [example2](example2/)
- decorating a method (think OOP), see [example3](example3/)
- another method decorator, but using a Func type, useful when the func signature is rather long, see [example4](example4/)
- method decorator that takes arguments [example5](example5/)
- return the result from the decorated function, instead of printing it [example6](example6/)
- a logging decorator, see [example7][example7/])

## contribute

pull requests are welcomed 

## todo

- using a struct instead of plain functions
- how does this all work when it's a interface
- pass data to http handlefunc
- nested decorators