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

title = "ADR-0022: VSecM Shall Use the Latest Stable Dependencies"
weight = 22
+++

- Status: accepted
- Date: 2024-09-02
- Tags: security, dependencies

## Context and Problem Statement

We should use the latest stable dependencies to ensure that we are not using
vulnerable dependencies. This includes the GoLang version that we are using
too.

## Decision

We have decided to:

1. Regularly update all project dependencies to their latest stable versions.
2. Use the latest stable version of GoLang for development and production.
3. Implement an automated dependency checking and updating process.

## Rationale

- Security: Latest versions often include security patches and vulnerability fixes.
- Performance: Newer versions may offer performance improvements.
- Features: Access to new features and improvements in the ecosystem.
- Compatibility: Staying current reduces the risk of future compatibility issues.

## Consequences

### Positive

- Improved security posture
- Access to latest features and optimizations
- Easier to attract contributors who prefer working with up-to-date technologies

### Negative

- Potential for breaking changes requiring code updates
- Increased time spent on dependency management and testing
- Possible learning curve for new features or changes in updated dependencies

## Implementation

1. Set up automated dependency checking tools (e.g., Dependabot, Renovate).
2. Establish a regular schedule for reviewing and applying updates.
3. Implement comprehensive testing to catch any issues introduced by updates.
4. Document the process for handling breaking changes in dependencies.
5. For VSecM core components, always use the latest stable GoLang version and 
   dependencies.
6. Regularly review and update the CI/CD pipeline to use the latest stable 
GoLang version.

## Special Considerations for VSecM SDK

While the core VSecM components should use the latest stable dependencies, the 
**VSecM SDK** may maintain compatibility with older versions:

1. The VSecM SDK can support older dependencies, including older GoLang runtimes, 
   as long as there are no known vulnerabilities that impact the code.
2. This approach allows for wider adoptability, enabling projects using older 
   Go versions to still incorporate and use the VSecM SDK.
3. Maintain a compatibility matrix documenting which SDK versions work with 
   which Go versions.
4. Regularly assess the trade-offs between backward compatibility and 
   security/feature improvements.
5. Clearly communicate the supported Go versions and any security implications 
   in the SDK documentation.

This strategy balances the need for up-to-date, secure core components with the 
practical considerations of SDK adoption across various project environments.

## Alternatives Considered

- Freezing dependencies at specific versions: Rejected due to security risks and 
  missing out on improvements.
- Manual updates on an as-needed basis: Rejected due to potential oversight and 
  inconsistency.

## References

- [Go Release Policy](https://golang.org/doc/devel/release.html)
- [Semantic Versioning](https://semver.org/)
- [OWASP Dependency-Check](https://owasp.org/www-project-dependency-check/)

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
