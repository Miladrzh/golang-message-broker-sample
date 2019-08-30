# Golang Message Broker
###### Tags: Go, Broker, Docker
###### Clean Code Status: Almost clean but some bad practices like hard coding broker URL (Except reading from Env. variables)
In this project we used web-socket protocol and Go-Routines to implemented a simple message broker like [RabbitMQ](https://www.rabbitmq.com/)
 
I wrote this code when I'm at beginning of learning Go and it may be helpful for you to try implement something like this to implement a project in go from scratch.

Code is consisted of 3 main part:

##### Publisher:
Is the sample application that write in queue. I wrote it just for testing my broker.

##### Subscriber(Consumer):
Is the sample application that Consume from queue. I wrote it just for testing my broker.


##### Broker: 
Is the main component that make connection between consumers and publishers.


### How to run project:
I wrote a docker-compose so easily you can run code with:
```
docker-compose up
```
docker-compose configuration can be find in root of project and it use Docker configs in Dockerfiles folder.


Don't hesitate to ask question if you have problems with understanding or running code.
You can find my email from my github profile.
