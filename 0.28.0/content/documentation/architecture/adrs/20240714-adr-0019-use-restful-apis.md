+++
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title = "ADR-0019: Use RESTful APIs for Communication Between Components"
weight = 19
+++

- Status: accepted
- Date: 2024-09-02
- Tags: api, rest, communication, protocols

## Context and Problem Statement

Although there are other communication protocols that we can use such as
gRPC, we have decided to use RESTful APIs for communication between components,
especially when talking to VSecM Safe. 

There are benefits and liabilities that this decision brings as we'll discuss
while we finalize this ADR.

## Decision Drivers

- Compatibility with existing systems and tools
- Ease of use and development
- Scalability and performance requirements
- Long-term maintainability and extensibility

## Considered Options

* RESTful APIs
* gRPC
* GraphQL
* WebSockets

## Decision Outcome

Chosen option: "RESTful APIs", because it provides the best balance of 
compatibility, ease of use, and maintainability for our current needs.

### Positive Consequences

* Broad compatibility and accessibility
* Ease of use and simplicity
* Backward compatibility
* Rich ecosystem and tooling support
* Standardization and governance
* Interoperability across platforms

### Negative Consequences

* Potentially less efficient for high-frequency, real-time communication compared 
  to WebSockets or gRPC
* Lack of built-in type safety compared to gRPC
* May require more bandwidth for data transfer compared to more efficient protocols

## Pros and Cons of the Options

### RESTful APIs

* Good, because it's widely understood and adopted in the industry
* Good, because it's stateless, which simplifies server implementation
* Good, because it leverages standard HTTP methods and status codes
* Good, because it's easy to debug and test using standard tools
* Good, because it supports caching mechanisms out of the box
* Good, because it's platform and language agnostic
* Good, because it has excellent documentation and community support
* Bad, because it may not be as efficient for high-frequency data exchange
* Bad, because it lacks built-in state management for more complex operations
* Bad, because it may lead to over-fetching of data in some scenarios

### gRPC

* Good, because it offers high performance and efficiency
* Good, because it provides strong typing and code generation
* Bad, because it has limited browser support
* Bad, because it may require additional infrastructure for HTTP/2

### GraphQL

* Good, because it allows clients to request exactly what they need
* Good, because it reduces over-fetching and under-fetching of data
* Bad, because it can be complex to implement and optimize
* Bad, because it may introduce security concerns if not properly implemented

### WebSockets

* Good, because it provides real-time, bidirectional communication
* Good, because it's efficient for frequent, small data transfers
* Bad, because it requires maintaining persistent connections
* Bad, because it may not be suitable for all types of API interactions

## Implementation Details

1. All new API endpoints for VMware Secrets Manager will be designed following 
   RESTful principles.
2. API documentation using OpenAPI Specification (formerly Swagger) could be 
   considered for future implementation to enhance API documentation and 
   maintainability.
4. Security measures such as authentication, authorization, and rate limiting 
   will be implemented for all RESTful APIs.
5. Versioning strategy will be implemented to manage API changes over time.

## Exceptions

While RESTful APIs are the preferred choice for most interactions with VMware 
Secrets Manager, there are specific scenarios where alternative protocols may 
be more suitable:

1. Message Queues: For scenarios involving replication or cross-cluster secrets 
  federation, where real-time, high-throughput, or guaranteed message delivery 
  is required, we may opt to use message queue systems. In such cases, we will 
  use the native protocols supported by the chosen message queue technology 
  (e.g., AMQP for RabbitMQ, Kafka protocol for Apache Kafka).

2. Streaming Data: In situations where continuous data streams are necessary, 
  such as real-time monitoring or log streaming, we may consider using 
  WebSockets or other streaming protocols that are more efficient for these 
  use cases.

3. Internal Communication: For internal microservices communication where 
  performance is critical, we may consider using gRPC or other binary protocols 
  for improved efficiency.

These exceptions will be evaluated on a case-by-case basis, considering factors 
such as performance requirements, scalability needs, and the specific use case. 
Any deviation from RESTful APIs will be thoroughly documented and justified.

## More Information

For more details on RESTful API design and best practices, refer to:

* [REST API Tutorial](https://restfulapi.net/)
* [Microsoft REST API Guidelines](https://github.com/microsoft/api-guidelines/blob/vNext/Guidelines.md)
* [VMware API Style Guide](https://github.com/vmware/api-style-guide)

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
