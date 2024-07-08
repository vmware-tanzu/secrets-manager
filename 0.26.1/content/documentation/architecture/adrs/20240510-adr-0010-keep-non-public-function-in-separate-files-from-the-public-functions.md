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

title = "ADR-0010: Keep Non-Public Function in Separate Files from the Public Functions"
weight = 10
+++

- Status: accepted 
- Date: 2024-05-12 
- Tags: code-organization, maintainability 

## Context and Problem Statement

This co-location of public and non-public functions may lead to confusion about 
the intended use and accessibility of various functions, complicates the code 
review process, and makes maintenance more challenging as the project scales. 

Clear separation can enhance understandability and manageability of the codebase.

## Decision Drivers

- **Code Clarity and Organization**: Enhancing readability and maintainability 
  by clearly separating public interfaces from internal implementation details.
- **Encapsulation and Security**: Strengthening encapsulation by physically 
  separating public and non-public code, which can also help prevent accidental 
  usage of non-public APIs.
- **Team Collaboration**: Simplifying code reviews and collaboration by making 
  it clearer what parts of the codebase are stable for external use and which 

## Considered Options

1. Keep non-public functions in separate files from the public functions.
2. Keep public and non-public functions in the same file, but separate them 
   with clear comments or other visual cues.
3. Keep public and non-public functions in the same file without any explicit 
   separation.

## Decision Outcome

Chosen option: "option 1", because it provides the clearest separation between
public and non-public functions, enhancing code clarity and organization.

### Positive Consequences

- **Improved Code Organization:** This separation will enhance the clarity of 
   which functions are intended for external use and which are meant for internal 
   use only.
- **Easier Maintenance and Scalability:** With functions more clearly organized, 
   new team members can more easily understand the architecture of the codebase, 
   and maintaining separate areas of concern becomes more straightforward.
- **Potential Initial Overhead:** The initial reorganization of the codebase may 
   require significant effort, especially if the project is large and functions 
   are deeply intertwined.
- **Enhanced Security and Encapsulation:** Keeping non-public functions separate 
   can prevent accidental exposure of internal logic to the public API, aligning 
   with best practices in software encapsulation and security.

### Negative Consequences

- **Increased Complexity in Project Structure**: Having separate files for public 
  and non-public functions can complicate the project structure, potentially 
  making it harder to navigate the codebase.
- **Overhead in Management**: Maintaining a clear division between files can 
  add overhead to the management of the project.
- **Potential for Misplacement**: There is a risk of developers placing functions 
  in the incorrect files, which could lead to improper use or exposure of 
  functions meant to be restricted.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
