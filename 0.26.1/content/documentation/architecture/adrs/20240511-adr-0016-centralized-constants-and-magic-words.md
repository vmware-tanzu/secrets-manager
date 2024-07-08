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

title = "ADR-0016: Centralize Magic Words, Numbers, and Configuration Constants in Common Modules"
weight = 16
+++

- Status: accepted 
- Date: 2024-05-12 
- Tags: code-quality, maintainability, readability, configuration

## Context and Problem Statement

'Magic' words and numbers—values with unexplained meaning or context used directly 
in code—can lead to confusion and errors during development and maintenance. 

These values often represent thresholds, specific identifiers, or configuration 
options that are reused across different parts of the application. 

Without a central place to define them, updating these values can become 
error-prone, and understanding their purpose can be challenging for both new 
and experienced developers.

The maintainability, readability, and overall management of configuration data are 
crucial for the long-term health and efficiency of our software projects. In the 
codebase, 'magic' words and numbers—values directly used in code that lack clear 
context or explanation—can lead to confusion and maintenance challenges. 

Additionally, important configuration data like environment variable names 
often represent system-critical information that is used across various parts 
of the application. Managing these through hard-coded values scattered throughout 
the code can result in errors and inconsistencies when changes are made.

Here is an example of how such variables may be defined as constants:

```go
package constants

type Identifier string

// CorrelationId is the identifier for the correlation ID.
// It is used to correlate log messages and other data across
// services while logging.
const CorrelationId Identifier = "correlationId"

type EnvVarName string

const AppVersion EnvVarName = "APP_VERSION"
const VSecMLogLevel EnvVarName = "VSECM_LOG_LEVEL"
const VSecMSafeSpiffeIdPrefix EnvVarName = "VSECM_SPIFFEID_PREFIX_SAFE"
const VSecMSafeEndpointUrl EnvVarName = "VSECM_SAFE_ENDPOINT_URL"
const VSecMKeyGenDecryptMode EnvVarName = "VSECM_KEYGEN_DECRYPT"

type FieldName string

const RootKeyText FieldName = "KEY_TXT"
```

### Exceptions

#### URLs and API Endpoints 

Given the potential for over-centralization and 
the specific contextual use of API endpoints within different modules or 
services, we will exclude API endpoints from being defined in the common 
constants module. This decision is based on:

* *Reduced Complexity*: Keeping API endpoints defined within the modules or 
services where they are used maintains simplicity and avoids the overhead of 
managing a large number of rarely changed or highly specific constants.
* *Local Context Clarity*: Defining endpoints closer to their usage context 
helps in understanding and maintaining the service-specific logic without 
navigating away to a central constants file.

## Decision Drivers

- **Configuration Management**: Centralizing configuration data for easier management.
- **Error Prevention**: Reducing the risk of errors and inconsistencies in the codebase.

## Considered Options

1. Centralize 'magic' words, numbers, and configuration constants in a common module.
2. Keep 'magic' words, numbers, and configuration constants distributed across the codebase.

## Decision Outcome

Chosen option: "option 1", because centralizing 'magic' words, numbers, and
configuration constants in a common module improves maintainability, readability,
and consistency across the codebase. It will also make it less error-prone to
update these values and provide a clear reference point for developers.

### Positive Consequences

- Improved maintainability and readability of the codebase.
- Reduced risk of errors and inconsistencies in configuration data.
- Easier management and updating of 'magic' words, numbers, and configuration 
  constants.
- Clear reference point for developers to understand the purpose and context of 
  these values.

### Negative Consequences

- Increased complexity in managing the common constants module.
- Potential for over-centralization leading to confusion or misuse of constants.
- Possible need for additional documentation to explain the purpose and usage of 
  these centralized constants.
- Increased dependency on the common constants module for various parts of the 
  application.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}

