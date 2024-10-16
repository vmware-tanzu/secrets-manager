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

title = "ADR-0021: VSecM Shall Use Environment Variables for Configuration"
weight = 21
+++

- Status: accepted
- Date: 2024-09-02
- Tags: configuration, environment-variables

## Context and Problem Statement

VSecM (VMware Secrets Manager) needs a flexible and secure way to configure its 
components across different environments. The challenge is to find a method that 
is both easy to manage and secure for handling configuration settings, especially 
those containing sensitive information.

## Decision

We have decided to use environment variables as the primary method for 
configuration in VSecM.

## Rationale

1. **Security**: Environment variables provide a level of security by keeping 
  sensitive information out of the codebase and configuration files. This reduces 
  the risk of accidentally committing secrets to version control.

2. **Flexibility**: Environment variables can be easily set and modified across 
  different environments (development, staging, production) without changing the 
  application code.

3. **Container-friendly**: In containerized environments, which are common in 
  modern deployments, environment variables are a standard way to pass 
  configuration to applications.

4. **Simplicity**: Using environment variables simplifies the deployment process 
   as there's no need to manage multiple configuration files for different 
   environments.

5. **Integration**: Many CI/CD and cloud platforms provide built-in support 
   for managing environment variables, making it easier to integrate with 
   existing infrastructure.

6. **Runtime configuration**: Environment variables allow for runtime 
   configuration changes without requiring application restarts in many cases.

## Consequences

### Positive

- Improved security by keeping sensitive information out of the codebase
- Easier management of configurations across different environments
- Better compatibility with containerized deployments and cloud platforms
- Simplified deployment process

### Negative

- Potential for environment variable sprawl if not managed carefully
- Debugging might be more challenging as configuration is not in a centralized file
- Limited structure for complex configurations (compared to configuration files)

## Implementation

When implementing this approach:

1. Use clear and consistent naming conventions for environment variables
2. Provide documentation on all available configuration options
3. Use sensible defaults where possible to minimize the number of required variables
4. Implement validation for required environment variables at application startup
5. Implement a fallback mechanism to use default values when environment 
   variables are not set

## Default Configuration

VSecM codebase will contain sane defaults for all environment variable 
configurations as much as possible. This approach ensures that:

1. VSecM will run out of the box even when most (or all) of the environment 
   variables are not defined or configured.
2. Users can get started quickly without needing to configure every aspect 
   of the system.
3. The system remains functional with minimal configuration, reducing the 
   chance of misconfigurations or startup failures.

When implementing default configurations:

1. Choose secure and generally applicable default values.
2. Document all default values clearly in the codebase and user documentation.
3. Implement a clear hierarchy: environment variables should override defaults 
   when present.
4. Log the use of default values at startup to inform users of the active 
   configuration.
5. Regularly review and update default values to ensure they remain appropriate 
   and secure as the system evolves.

This approach balances ease of use with the flexibility to customize the 
configuration as needed, making VSecM both user-friendly and adaptable to 
various deployment scenarios.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
