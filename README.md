#The Core-Tech Challenge

## Create two microservices for real time messages using websockets, and at least one demo subscriber.

### Functionality
1. The first service should listen for incoming messages through the websocket protocol and when a new one arrives, the message should be published into message queue
1. The second service should listen for incoming messages through the message queue and when a new message arrives, the message should be published to all the subscribers through the websocket protocol

### Other Requirements
* At least one of the services should have tests
* Make sure your code is well structured and maintainable (including tests)
* You can use frameworks and technologies by your choice, but the language for the microservices should be Javascript (node) or Go (or both).
* The source code should be hosted online using github (or similar service)

Note, this is your chance to impress us, take this opportunity to show us what you can do! You can include authentication, use any CI, container based deployment or whatever you think will grab our attention.
