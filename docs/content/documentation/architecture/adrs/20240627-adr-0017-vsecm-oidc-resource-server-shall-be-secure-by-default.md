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

title = "ADR-0017: VSecM OIDC Resource Server Shall Be Secure by Default"
weight = 17
+++

- Status: accepted
- Date: 2024-09-02
- Tags: security, threat model, oidc, resource-server

## Context and Problem Statement

VSecM has an experimental OIDC resource server functionality. 

This ADR aims to capture important security considerations and decisions to 
ensure that the OIDC resource server is secure by default.

Here are some of the key security considerations that need to be addressed:

For starters, the OIDC Resource Server shall be disabled by default. 
This is both to ensure that the system is secure by default, and also
preserve resources by not running unnecessary services.

The OIDC Resource Server functionality may be considered experimental at the time
of writing this ADR. It is disabled by default and requires thorough testing and
security review before being enabled in production environments. The status and
stability of this feature may evolve over time, so it's important to refer to the
most current documentation for its current state and best practices for usage.

It is worth reiterating that the OIDC Resource Server creates a RESTful API
and exposes the functionality of VSecM Sentinel to the outside world. Yes,
only authenticated users can access the API, but it is still a security risk
if not properly configured.

When this functionality is enabled, the threat model of VSecM changes, and
it will be up to the operator to ensure that the OIDC Provider is properly
configured and that the Resource Server is properly secured.

That said, using OIDC as an authentication mechanism does not inherently
decrease the security of the system. In fact, OIDC can provide a secure way
to authenticate users since it is built on top of the OAuth 2.0 protocol,
allowing clients to verify the identity of users based on the authentication
performed by an authorization server.

However, as with any system that opens up additional points of access, there
are a number of factors that need to be considered:

1. Quality of the OIDC Provider: If you're not implementing OIDC yourself,
   the security of your system will depend largely on the security measures
   taken by the OIDC provider. Make sure you use a reputable and secure
   provider.

2. Access Control: If you allow OIDC authentication for sentinel, it's crucial
   that you have a robust access control system in place. We should make sure
   that only authorized users can write secrets. Additionally, consider the
   Principle of Least Privilege, granting only the necessary access to users
   and nothing more.

3. Secure Token Handling: OIDC uses tokens to establish and confirm user
   identity. Ensure these tokens are securely handled at all times. They
   should not be logged or exposed in any way, and should be securely stored
   if they need to be stored at all. It's often best practice to handle these
   tokens in-memory and never persist them.

 4. Validation of Tokens: Always validate ID tokens on your backend service.
    This includes verifying the signature, validating the claims, and confirming
    that the 'aud' claim contains your app’s client ID.

 5. Use of Secure Communication: All communication should happen over secure
    channels (HTTPS).

 6. Session Management: Ensure proper session management. Invalidate sessions
    server-side after logout, and consider implementing token expiration and
    revocation.Monitoring and Auditing: Regularly monitor and audit the
    activities within your system to detect any unauthorized access or
    anomalous activities.

Some of these are to be fixed within the realm of VSecM, and some of these
will be the things that the user/operator have to consider and secure
themselves.

Remember that adding new access points to a system always involves risk, but
with the proper implementation and continuous monitoring, it's possible to
maintain a high level of security.

We should at least test these with a provider like Keycloak and create a
recommended setup guide for the users before making this feature public.

### Exceptions

There are no exceptions to this decision. 

The OIDC Resource Server functionality must always be secure by default and 
disabled by default. This applies to all deployment methods, including 
Helm Charts and generated manifests. 

Users must explicitly enable and configure the OIDC Resource Server if they 
choose to use this feature, ensuring they are aware of the security implications 
and follow the provided guidelines for secure implementation.

## Decision Drivers

- VSecM shall always be secure by default ([**ADR-0006**][ADR-0006])
- Protect sensitive data from unauthorized access
- Comply with security best practices for OIDC implementations
- Maintain the integrity and confidentiality of secrets managed by VSecM

[ADR-0006]: @/documentation/architecture/adrs/20240510-adr-0006-be-secure-by-default.md

## Considered Options

1. Enable OIDC Resource Server by default with basic security measures
2. Disable OIDC Resource Server by default and provide comprehensive security guidelines
3. Implement OIDC Resource Server with advanced security features enabled by default

## Decision Outcome

Chosen option: "Disable OIDC Resource Server by default and provide 
comprehensive security guidelines"

This decision ensures that VSecM remains secure by default while allowing users 
to opt-in to the OIDC Resource Server functionality with proper guidance. It 
aligns with the principle of secure-by-default and gives users control over 
enabling additional access points to their secrets.

### Positive Consequences

- VSecM remains secure out-of-the-box without additional configuration
- Users are forced to make a conscious decision to enable the OIDC Resource Server
- Comprehensive security guidelines help users implement the feature safely
- Reduced attack surface for default installations
- Aligns with the principle of least privilege

### Negative Consequences

- Additional setup required for users who want to use the OIDC Resource Server
- Potential for misconfiguration if users don't follow the provided guidelines carefully
- May limit immediate usability for some users expecting OIDC functionality by default

## Implementation Details

1. The OIDC Resource Server will be disabled by default in the VSecM configuration.
2. Detailed documentation will be provided on how to securely enable and configure 
   the OIDC Resource Server.
3. A recommended setup guide using a provider like Keycloak will be created and 
   included in the documentation.
4. Security considerations and best practices will be clearly outlined in the 
   public documentation.
5. Warning messages will be displayed when enabling the OIDC Resource Server, 
   reminding users of the security implications.
6. Regular security audits of the OIDC Resource Server implementation will 
   be conducted.

## References

- [OAuth 2.0 Threat Model and Security Considerations](https://datatracker.ietf.org/doc/html/rfc6819)
- [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html)
- [OWASP Authentication Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html)

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
