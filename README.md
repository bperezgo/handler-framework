# Handler Framework (or HandlerLib?)

This framework is intended to be used in microservices, or in structures of lambda in aws. The idea is:
The client build a function, that will be run in some architecture (for example, lambda functions), and with this framework you can inyect the dependencies before start up the function in, for example, lambda function. But if you want, you can start up this function in a local server, and, with this in mind, you can test your function locally before deploy the function to the cloud.

Examples of the implementation in the next repository:

[handler-framework-examples](https://github.com/bperezgo/handler-framework-examples)