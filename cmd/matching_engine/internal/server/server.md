# Server

- The process of receiving a request to the gRPC Gateway.

- Send the task to the Event Process, Action Process.

- Save the order book in Process memory

- Send the orders to be settled to the Action process
 
- Send the placed order to the Event process.



## TODO 

- [ ] redis place order
- [x] order book serving
- [ ] order book snapshot
- [ ] order book recovery 
- [ ] order book recovery test