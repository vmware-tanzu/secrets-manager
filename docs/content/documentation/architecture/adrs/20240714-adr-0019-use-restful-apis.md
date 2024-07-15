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

- Status: draft
- Date: 2024-07-14
- Tags: api, rest, communication, protocols

## Context and Problem Statement

Although there are other communication protocols that we can use such as
gRPC, we have decided to use RESTful APIs for communication between components,
especially when talking to VSecM Safe. 

There are benefits and liabilities that this decision brings as we'll discuss
while we finalize this ADR.

Here are some advantages of sticking with RESTful APIs:

### Broad Compatibility and Accessibility

REST is Ubiquitous: REST is a well-established standard used widely across various 
platforms and languages. This broad adoption makes it easier for developers to 
interact with Kubernetes using tools and libraries they are already familiar with.

HTTP/1.1 Compatibility: REST, based on HTTP/1.1, ensures compatibility with a 
wide range of clients, including browsers, curl, and other HTTP clients, which 
are not natively compatible with protocols such as gRPC.

### Ease of Use and Simplicity

Simplicity: REST APIs are straightforward to understand and use. They rely on 
simple HTTP methods (GET, POST, PUT, DELETE), making them easy to test and debug 
using common tools.

Human Readability: REST APIs use JSON or XML, which are human-readable formats. 
This makes it easier for developers to manually inspect and interact with the API.

### Backward Compatibility

Stable and Predictable: REST APIs offer stability and backward compatibility, 
crucial for a system like Kubernetes, which has a large and diverse user base. 
Changes to the API need to be carefully managed to avoid breaking existing clients.

Versioning: REST APIs have well-established conventions for versioning, allowing 
systems (such as Kubernetes) to evolve the API without disrupting existing users.

### Ecosystem and Tooling

Rich Ecosystem: There is a rich ecosystem of tools and libraries for working 
with REST APIs, including API gateways, documentation generators 
(like Swagger/OpenAPI), and client libraries for numerous languages.

Monitoring and Logging: Established tools for monitoring, logging, and analyzing 
HTTP traffic can be directly applied to REST APIs.

### Standardization and Governance:

API Governance: The RESTful API model aligns well with our needs for API 
governance and lifecycle management. It provides clear guidelines and best 
practices for designing, exposing, and managing APIs.

Standards Compliance: REST APIs adhere to established web standards, making 
them more interoperable and compliant with various web technologies and protocols.

### Interoperability:

Cross-Platform Compatibility: REST APIs work seamlessly across different 
platforms and environments, ensuring that Kubernetes can be accessed and managed 
from various systems without compatibility issues.

Firewall and Proxy Friendly: REST APIs, using HTTP/1.1, are more firewall and 
proxy-friendly, which is important for enterprise environments with strict 
network policies.