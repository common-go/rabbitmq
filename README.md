# rabbitmq
RabbitMQ is an open-source message broker software that implements the Advanced Message Queuing Protocol (AMQP). It facilitates communication between different components of a system by sending and receiving messages. RabbitMQ is known for its reliability, flexibility, and support for various messaging protocols, making it a popular choice for building robust and scalable messaging systems.
### Libraries for RabbitMQ
- GO: [rabbitmq](https://github.com/core-go/rabbitmq). Example is at [go-subscription](https://github.com/project-samples/go-subscription).
- nodejs: [rabbitmq-ext](https://www.npmjs.com/package/rabbitmq-ext). Example is at [rabbitmq-sample](https://github.com/typescript-tutorial/rabbitmq-sample).
#### A common flow to consume a message from a message queue
![A common flow to consume a message from a message queue](https://cdn-images-1.medium.com/max/800/1*Y4QUN6QnfmJgaKigcNHbQA.png)
- The libraries to implement this flow are:
    - [mq](https://github.com/core-go/mq) for GOLANG. Example is at [go-subscription](https://github.com/project-samples/go-subscription)
    - [mq-one](https://www.npmjs.com/package/mq-one) for nodejs. Example is at [rabbitmq-sample](https://github.com/typescript-tutorial/rabbitmq-sample)

### Key Features of RabbitMQ
#### Multi-Protocol Support
- Primarily implements AMQP but also supports other protocols like STOMP, MQTT, and HTTP.
#### Flexible Routing
- Provides powerful routing capabilities using exchanges, bindings, and queues.
#### High Availability
- Supports clustering and replication for high availability and fault tolerance.
#### Durability
- Messages can be persisted to ensure they are not lost in case of broker failure
#### Management and Monitoring
- Includes a web-based management interface and supports various monitoring tools for managing and observing the broker.
#### Plugins and Extensions
- Highly extensible with a wide range of plugins for additional functionality.
#### Message Acknowledgments
- Ensures message delivery reliability through acknowledgments and confirmations.

### Use Cases of RabbitMQ
#### Task Queues
- Distributing tasks among worker processes, ensuring load balancing and reliability.
#### Asynchronous Processing
- Decoupling components of an application to handle long-running tasks asynchronously.
#### RPC (Remote Procedure Call)
- Implementing RPC over messaging to allow for asynchronous method calls between services.
#### Complex Routing
- Implementing sophisticated message routing logic using exchanges, bindings, and queues.
#### Integration with Microservices
- Enabling communication between microservices in a distributed architecture.
 
### How RabbitMQ Works
RabbitMQ operates using the following core concepts
#### Producer
- An application that sends messages to the broker.
#### Consumer
- An application that receives messages from the broker.

![Microservice Architecture](https://cdn-images-1.medium.com/max/800/1*vKeePO_UC73i7tfymSmYNA.png)

#### Queue
- A buffer that stores messages until they are consumed. Queues are the primary form of message storage in RabbitMQ

![A typical micro service](https://cdn-images-1.medium.com/max/800/1*d9kyekAbQYBxH-C6w38XZQ.png)

#### Exchange
- Receives messages from producers and routes them to queues based on routing rules. Types of exchanges include direct, topic, fanout, and headers.
#### Binding
- Defines the relationship between an exchange and a queue. Binding keys are used to determine how messages are routed.
#### Channel
- A virtual connection within a connection. Channels allow multiple simultaneous interactions with the broker over a single network connection.
#### Connection
- A network connection between a producer or consumer and the RabbitMQ broker.
#### Virtual Host (vHost)
- A virtual cluster of exchanges, queues, and bindings. Provides a way to segregate applications within a broker.

### RabbitMQ vs. Kafka
#### Message Persistence
- <b>RabbitMQ</b>: Primarily designed for transient messages but supports persistent messages for reliability.
- <b>Kafka</b>: Optimized for persistent message storage with configurable retention policies.
#### Use Case Suitability
- <b>RabbitMQ</b>: Suitable for task distribution, RPC, and systems that need complex routing and message transformation.
- <b>Kafka</b>: Best for high-throughput, real-time data streaming, and log aggregation.
#### Message Delivery
- <b>RabbitMQ</b>: Guarantees message delivery with acknowledgments, retries, and dead-letter exchanges.
- <b>Kafka</b>: Provides at-least-once delivery by default, with exactly-once semantics available in specific configurations.
#### Scalability
- <b>RabbitMQ</b>: Scales well for many messaging scenarios but may require more complex configurations for very high-throughput needs.
- <b>Kafka</b>: Designed for horizontal scalability and handles large-scale data streams efficiently.
#### Complexity
- <b>RabbitMQ</b>: Easier to set up and manage for simpler use cases and offers extensive support for various messaging patterns.
- <b>Kafka</b>: Supports various messaging patterns and protocols, making it adaptable to different use cases.

### Advantages of RabbitMQ
#### Ease of Use
- Simple to set up and use with comprehensive documentation and a user-friendly web management interface.
#### Flexibility
- Supports various messaging patterns and protocols, making it adaptable to different use cases.
#### Reliable Message Delivery
- Ensures messages are delivered reliably with support for acknowledgments, retries, and dead-letter exchanges.
#### Extensibility
- Offers a wide range of plugins and extensions to add additional functionality as needed.
#### Clustering and High Availability
- Supports clustering and federation for fault tolerance and high availability.
#### Community and Support
- Backed by a large community and commercial support options, providing ample resources for troubleshooting and assistance.

### Disadvantages of RabbitMQ
#### Performance Overhead
- May introduce performance overhead compared to more lightweight brokers in high-throughput scenarios.
#### Complex Routing Configurations
- Advanced routing configurations can become complex and may require a deep understanding of AMQP.
#### Resource Intensive
- Can be resource-intensive, especially when handling large volumes of messages or complex configurations.
#### Latency
- Slightly higher latency compared to some other message brokers optimized for low-latency operations.

### Example Scenario: Task Distribution System
In a task distribution system, RabbitMQ can be used to distribute tasks among multiple worker processes to ensure load balancing and reliability.
#### Producer
- A web application sends tasks to a RabbitMQ queue.
#### Queue
- A task queue stores tasks until they are processed by worker processes.
#### Consumer
- Multiple worker processes consume tasks from the queue, process them, and send acknowledgment back to RabbitMQ.
#### Acknowledgments
- Ensures tasks are not lost and are re-queued if a worker fails to process a task.
#### Load Balancing
- Tasks are evenly distributed among available workers, ensuring efficient resource utilization

### Conclusion
RabbitMQ is a versatile and reliable message broker that supports a wide range of messaging protocols and patterns. Its ease of use, flexibility, and support for various messaging scenarios make it a popular choice for building robust and scalable messaging systems. While it may introduce some performance overhead and complexity in advanced configurations, its benefits in terms of reliability, extensibility, and community support make it a valuable tool for modern software architectures. Understanding RabbitMQ's core concepts and capabilities can help organizations build efficient and resilient messaging solutions.

## Installation

Please make sure to initialize a Go module before installing core-go/rabbitmq:

```shell
go get -u github.com/core-go/rabbitmq
```

Import:

```go
import "github.com/core-go/rabbitmq"
```
