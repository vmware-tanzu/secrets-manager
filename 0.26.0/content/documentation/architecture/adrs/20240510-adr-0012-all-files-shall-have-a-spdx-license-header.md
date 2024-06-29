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

title = "ADR-0012: All Files Shall Have a SPDX License Header"
weight = 12
+++

- Status: accepted 
- Date: 2024-05-12 
- Tags: legal, licensing, compliance

## Context and Problem Statement

Licensing information not being embedded in source code and other project files
may lead to confusion and legal risks. This could make software distribution and
use challenging. To mitigate this, we have decided to include a SPDX license 
identifier at the top of every file.

This is the license header format that we will use:

```txt
/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/
```

The comment block will change depending on the file type. For example, for a
Go file, the comment block will be `//` instead of `/*` and `*/`.

## Decision Drivers 

The decision to adopt SPDX license headers is driven by the need for clear, 
accessible, and unambiguous licensing information directly within the source 
files. This approach is supported by industry best practices for open-source 
compliance, particularly in environments where software is frequently audited 
or distributed across different legal jurisdictions. Adopting SPDX will also 
facilitate easier integration and reuse of external open-source components that 
are already using SPDX identifiers.

## Considered Options

1. Standardize and clarify the licensing across all code files in our project.
2. Do not enforce a specific license header format.

## Decision Outcome

Chosen option: "option 1", because it ensures that all files in the project
clearly state their licensing, reducing ambiguity and potential legal issues.

### Positive Consequences

- **Consistency in Licensing**: All files will clearly state their licensing, 
  reducing ambiguity and potential legal issues.
- **Ease of Compliance Verification**: Tools that recognize SPDX identifiers 
  can automatically verify the compliance of all files with declared licenses, 
  streamlining audits and compliance checks.

### Negative Consequences

None.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}

