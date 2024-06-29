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

title = "ADR-0015: Keep Source Code Line Length Around 80 Characters"
weight = 15
+++

- Status: accepted 
- Date: 2024-05-12 
- Tags: code-quality, readability 

## Context and Problem Statement

The readability and maintainability of code are critical to the efficient 
development and long-term viability of our software projects. Historically, 
keeping line lengths to **about 80 characters** has been a standard practice in 
software development. This limit was originally adopted due to the physical 
constraints of hardware, like monitors and printed code reviews, which handled 
80 characters per line optimally. 

While modern technologies have evolved to support longer lines, this standard 
still provides significant benefits in terms of code readability and 
reviewability.

### Practical Reasons

- **Readability**: Shorter lines are generally easier to read than longer lines. 
  When lines are concise, it's easier for the eye to track from the end of one 
  line to the beginning of the next. This can reduce eye strain and mental load, 
  making the code easier to understand at a glance.
- **Code Review and Side-by-Side Comparison**: Keeping lines short is particularly 
  beneficial when doing code reviews or using tools that present two versions of 
  a file in a side-by-side diff. Short lines ensure that differences are clearly
  visible without horizontal scrolling, which can disrupt the review process.
- **Multiple Windows**: Developers often work with multiple files open 
  simultaneously, tiled side by side. Short lines allow more text to be visible 
 in each window, making multitasking more effective without losing context.

### Modern Perspectives

With the advent of modern IDEs and editors that can easily handle longer lines, 
and high-resolution displays that can fit more characters on the screen without 
wrapping, some argue that the 80-character limit is outdated. These tools and 
technologies allow developers to utilize wider screens effectively.

### Arguments for Longer Lines

- **Utilization of Modern Displays**: Allows taking full advantage of widescreen 
  and high-resolution monitors.
- **Fewer Line Breaks**: Can reduce the need for line continuations in languages 
  like Python, making the code look cleaner and less fragmented.

### Arguments for Sticking to the 80-Character Rule

- **Consistency**: Maintaining this standard ensures that code is accessible and 
  readable regardless of the hardware or software environment.
- **Inclusivity**: Adhering to the 80-character limit helps ensure that those 
  using less advanced technology or specific accessibility tools can also work 
  comfortably with the code.

## Decision Drivers

- **Readability**: Ensuring that code is easy to read and understand.
- **Consistency**: Maintaining a consistent style across the codebase.
- **Reviewability**: Facilitating code reviews and side-by-side comparisons.
- **Accessibility**: Ensuring that code is accessible to all developers.
- **Maintainability**: Making code easier to maintain and modify.
- **Inclusivity**: Considering developers with different technology setups.
- **Tooling Support**: Ensuring compatibility with various development tools.

## Considered Options

1. Keep source code line length around 80 characters.
2. Allow longer lines, up to 120 characters.
3. Do not enforce a specific line length limit.

## Decision Outcome

Chosen option: "option 1", because maintaining the 80-character limit provides
significant benefits in terms of readability, reviewability, and accessibility.

When calculating the line length, the tab character will be considered as
**two** characters.

### Positive Consequences

- **Readability**: Shorter lines are easier to read and understand.
- **Consistency**: Maintaining a consistent style across the codebase.
- **Code Organization**: Enforcing a limit on line length means the code will
  not be nested too deeply, which can result in refactoring code into smaller
  meaningful logical units, which will improve the overall code organization
  and readability.

### Negative Consequences

- **Line Continuations**: May require more line continuations in languages that 
  enforce strict line length limits.
- **Modern Display Usage**: May not fully utilize the screen space on modern 
  widescreen monitors.
- **Code Fragmentation**: Longer lines can lead to fragmented code that is harder 
  to follow.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}

