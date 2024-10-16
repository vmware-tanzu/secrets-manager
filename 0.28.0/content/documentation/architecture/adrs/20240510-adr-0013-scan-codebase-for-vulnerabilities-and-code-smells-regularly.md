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

title = "ADR-0013: Scan the Codebase for Vulnerabilities and Code Smells Regularly"
weight = 13
+++

- Status: accepted
- Date: 2024-09-02
- Tags: security, static-analysis, quality

## Context and Problem Statement

As our software project grows in complexity and scale, the risk of introducing
security vulnerabilities and code smells increases. Currently, our codebase
lacks a consistent and systematic approach to identifying these issues early
in the development cycle, leading to higher maintenance costs and potential
security breaches in production.

## Decision Drivers

- Improve code quality and security
- Detect potential vulnerabilities early in the development process
- Ensure consistent code analysis across the project
- Integrate seamlessly with existing CI/CD pipelines

## Considered Options

1. Use `go vet` and `govulncheck`
2. Use `go vet`, `govulncheck`, and Snyk
3. Use `go vet`, `govulncheck`, `codesweep`, and `gosec`
4. Use `go vet`, `govulncheck`, Snyk, and `golangci-lint`
5. Maintain status quo (no regular scanning)

## Decision Outcome

Chosen option: "Use go vet, govulncheck, Snyk, and golangci-lint", because this 
combination of tools provides a comprehensive set of tools for code analysis, 
vulnerability detection, and security monitoring while also including a 
powerful linter for Go code.

### Implementation Details

1. `go vet`:
    - Run `go vet ./...` as part of the CI/CD pipeline to catch common coding 
      mistakes.
    - Before release: Manually run `go vet ./...` and review results.

2. `govulncheck`:
    - Install: `go install golang.org/x/vuln/cmd/govulncheck@latest`
    - Run: `govulncheck ./...` in the CI/CD pipeline to check for known 
      vulnerabilities.
    - Before release: Manually run `govulncheck ./...` and review results.

3. Snyk:
    - Install Snyk CLI: `npm install -g snyk`
    - Navigate to the project directory: `cd $WORKSPACE/secrets-manager`
    - Authenticate (if not already done): `snyk auth`
    - Run: `snyk monitor` in the CI/CD pipeline.
    - Before release: Manually run `snyk test` and review results.
    - Monitor projects on https://app.snyk.io/org/$username/projects
      (replace $username with your actual username)

4. `golangci-lint`:
    - Install `golangci-lint` by following the official documentation 
    - Create a configuration file `.golangci.yml` in the project root with 
      desired linters and settings.
    - Run: `golangci-lint run` before every release cut and review results.

5. Pre-release manual check:
    - Before cutting a release, a designated team member should run all the above
      commands manually and review the results.
    - Any new vulnerabilities or issues found should be addressed before proceeding
      with the release.
    - Document the results of this manual check in the release notes.

### Positive Consequences

- Early detection of potential vulnerabilities and code smells
- Improved overall code quality and security
- Consistent code analysis across the project
- Comprehensive linting with `golangci-lint` catches a wide range of issues

### Negative Consequences

- Slight increase in release preparation time due to additional scanning steps
  during release cuts
- Potential false positives that may require manual review
- Initial setup time for configuring `golangci-lint`

## Additional Notes

- Consider implementing codesweep and gosec in the future if more comprehensive
  static analysis is needed.
- Regularly review and update the scanning tools to ensure they remain effective
  and relevant.
- Provide documentation and training on how to interpret and act on the results
  of these scans.
- Periodically review and update the `.golangci.yml` configuration to ensure it
  aligns with project needs and best practices.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
