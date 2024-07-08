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

title = "ADRs"
weight = 1
sort_by = "weight"
insert_anchor_links = "left"
redirect_to = "documentation/architecture/adrs/20240509-adr-0001-use-log4brains-to-manage-the-adrs"
+++

## Introduction

Welcome 👋 to the architecture knowledge base of **VMware Secrets Manager**.

You will find here all the Architecture Decision Records (*ADR*) of the project.

## Definition and Purpose

### Architectural Decision

An Architectural Decision (*AD*) is a software design choice that addresses a
functional or non-functional requirement that is architecturally significant.

### Architectural Decision Record

An Architectural Decision Record (*ADR*) captures a single AD, such as often
done when writing personal notes or meeting minutes; the collection of ADRs
created and maintained in a project constitutes its decision log.

An ADR is immutable: only its status can change (*i.e., become deprecated or
superseded*). That way, you can become familiar with the whole project history
just by reading its decision log in chronological order.

Moreover, maintaining this documentation aims at:

- 🚀 Improving and speeding up the onboarding of a new team member.
- 🔭 Avoiding blind acceptance/reversal of a past decision.
- 🤝 Formalizing the decision process of the team.

## Usage

This website is automatically updated after a change on the `main` branch of
the project's Git repository.

In fact, the developers manage this documentation directly with markdown files
located next to their code, so it is more convenient for them to keep it up-to-date.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
