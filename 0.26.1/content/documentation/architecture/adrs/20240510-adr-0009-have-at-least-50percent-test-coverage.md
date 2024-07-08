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

title = "ADR-0009: Have at Least 50% Test Coverage Across the Project"
weight = 9
+++

- Status: accepted
- Date:  2024-05-11
- Tags: testing, quality, coverage

## Context and Problem Statement

An uneven or low code coverage across a project can lead to several critical bugs 
making it into production. This can impact the client's trust and increase the 
number of post-release hotfixes and patches. There is a clear need to standardize 
and increase the test coverage to ensure higher quality releases and reduce 
maintenance costs.

## Decision Drivers

- Without enough coverage, changes to the codebase can introduce bugs that are 
  not caught by the existing tests.
- Low test coverage can lead to a lack of confidence in the codebase, making 
  developers hesitant to make changes.
- High test coverage can help identify areas of the codebase that are more 
  complex or error-prone, allowing for targeted refactoring and improvement.
- High test coverage can help ensure that new features do not inadvertently 
  break existing functionality.
- High test coverage can help maintain a consistent level of quality across 
  the codebase.
- High test coverage can help reduce the number of bugs that make it to 
  production, improving customer satisfaction and reducing maintenance costs.
- High test coverage can help improve the overall quality of the codebase, 
  making it easier to maintain and extend over time.

## Considered Options

1. Have at least 50% test coverage across the project.
2. Do not enforce a specific test coverage threshold.

## Decision Outcome

Chosen option: "option 1", because increasing the test coverage to at least 50%
across the project is expected to improve the overall quality of the codebase,
reduce the number of bugs that make it to production, and align with our
strategic goals of delivering higher quality software and improving user
satisfaction.

### Positive Consequences

- Improved code quality.
- Reduced number of bugs.
- Increased developer confidence in the codebase.
- Reduced regression issues.
- Easier refactoring and maintenance.
- Better user satisfaction.
- Lower maintenance costs.

### Negative Consequences

- Increased development time.
- Possible impact on feature development speed.
- Possible need for investment in better testing tools and continuous 
  integration systems.
- Possible need for adjustments to the test coverage threshold based on project 
  requirements and team feedback.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}

