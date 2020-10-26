# 6. Tight Coupling

Classes that are tightly coupled are hard to reuse in isolation, since they depend on each other. Tight coupling leads to monolithic systems, where you can’t change or remove a class without understanding and changing many other classes. The system becomes a dense mass that’s hard to learn, port, and maintain.

Loose coupling increases the probability that a class can be reused by itself and that a system can be learned, ported, modified, and extended more easily. Design patterns use techniques such as abstract coupling and layering to promote loosely coupled systems.

Design patterns: 
- [Abstract Factory](./abstractfactory),
- [Bridge](./bridge),
- [Chain of Responsibility](./chainofresponsibility),
- [Command](./command),
- [Facade](./facade),
- [Mediator](./mediator),
- [Observer](./observer)
