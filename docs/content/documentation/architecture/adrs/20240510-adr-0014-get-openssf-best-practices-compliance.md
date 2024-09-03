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

title = "ADR-0014: Get OpenSSF Best Practices Compliance"
weight = 14
+++

- Status: accepted
- Date: 2024-05-12
- Tags: security, compliance, quality 

## Context and Problem Statement

[You can check the current status of the project here][openssf].

Security and sustainability of **VMware Secrets Manager** are of paramount 
importance. In this light, achieving compliance with established industry 
standards is crucial. 

[The Open Source Security Foundation (*OpenSSF*) Best Practices][openssf-ref] 
provide a framework for maintaining high standards of security and quality in 
open-source projects. 

[openssf]: https://www.bestpractices.dev/en/projects/7793
[openssf-ref]: https://www.bestpractices.dev/

## Decision Drivers

- **Security and Compliance**: Ensuring that our project meets industry standards 
  for security and compliance.
- **Credibility and Trust**: Enhancing the credibility and trustworthiness of 
  our project among users, contributors, and partners.
- **Risk Mitigation**: Reducing the risk of security vulnerabilities and 
  compliance issues in our project.
- **Industry Standards**: Aligning with established industry standards for 
  open-source security and quality.
- **Continuous Improvement**: Embracing a culture of continuous improvement and 
  learning in our project.
- **Sustainability**: Ensuring the long-term sustainability and success of our 
  project.
- **Project Resilience**: Building resilience against security threats and 
  vulnerabilities in our project.
- **Developer Experience**: Improving the developer experience by implementing 
  security best practices and quality assurance measures.
- **Feedback Loop**: Establishing a feedback loop for continuous improvement 
  based on security assessments and quality metrics.
- **Adaptability**: Ensuring that our project can adapt to changing security 
  threats, compliance requirements, and quality standards.
- **Documentation**: Documenting security practices, quality assurance processes, 
  and compliance measures for transparency and accountability.
- **Monitoring and Reporting**: Monitoring security vulnerabilities, quality 
  issues, and compliance gaps in our project and reporting on them regularly.
- **Incident Response**: Establishing incident response procedures for security 
  breaches, quality incidents, and compliance violations in our project.
- **Alignment with Best Practices**: Aligning with recognized best practices 
  for open-source security and quality assurance.
- **Inclusivity**: Ensuring that security and quality practices are inclusive 
  and accessible to all contributors and users of our project.

## Considered Options

1. Get OpenSSF Best Practices Compliance.
2. Do not pursue OpenSSF Best Practices Compliance.
3. Postpone the decision on OpenSSF Best Practices Compliance.

## Decision Outcome

Chosen option: "option 1", because it aligns with our project's goals of
enhancing security, credibility, and trust, and it demonstrates our commitment to
adhering to industry standards for open-source security and quality.

### Positive Consequences

- **Enhanced Security Measures**: By aligning with OpenSSF standards, our project 
  will adopt robust security protocols, reducing vulnerabilities.
- **Increased Credibility and Trust**: Compliance will likely increase trust among 
  users, contributors, and potential partners, showcasing our commitment to security.

### Negative Consequences

- **Operational Disruptions**: Implementing new practices may temporarily disrupt 
  existing workflows and slow down development as the core team and the
  contributors adjust.
- **Ongoing Maintenance**: Once compliance is achieved, ongoing efforts will be 
  required to maintain these standards, necessitating continuous monitoring and 
  updates.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
