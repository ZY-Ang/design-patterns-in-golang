# 7. Extending functionality by subclassing

Customizing an object by subclassing often isn’t easy. Every new class has a fixed implementation overhead (initialization, finalization, etc.). Defining a subclass also requires an in-depth understanding of the parent class.
 
For example, overriding one operation might require overriding another. An overridden operation might be required to call an inherited operation. And subclassing can lead to an explosion of classes, because you might have to introduce many new subclasses for even a simple extension.
 
Object composition in general and delegation in particular provide flexible alternatives to inheritance for combining behavior. New functionality can be added to an application by composing existing objects in new ways rather than by defining new subclasses of existing classes. On the other hand, heavy use of object composition can make designs harder to understand. Many design patterns produce designs in which you can introduce customized functionality just by defining one subclass and composing its instances with existing ones.

Design patterns:
- [Bridge](./bridge),
- [Chain of Responsibility](./chainofresponsibility),
- [Composite](./composite),
- [Decorator](./decorator),
- [Observer](./observer),
- [Strategy](./strategy)
