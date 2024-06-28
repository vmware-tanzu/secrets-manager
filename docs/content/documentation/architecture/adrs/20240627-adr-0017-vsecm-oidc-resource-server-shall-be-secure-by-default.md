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

- Status: draft
- Date: 2024-06-27
- Tags: security, threat model, oidc, resource-server

## Context and Problem Statement

VSecM has an experimental OIDC resource server functionality. 

This ADR aims to capture important security considerations and decisions to 
ensure that the OIDC resource server is secure by default.

Here are some of the key security considerations that need to be addressed:

For starters, the OIDC Resource Server shall be disabled by default. 
This is both to ensure that the system is secure by default, and also
preserve resources by not running unnecessary services.

Also, we may decide to move some of these decisions to the production setup
documentation (keeping it here, as an inline comment for now because this feature
is experimental and discussing implementation details in the public documentation
at this point might be premature — Once we make this feature public, we can
move some of these comments here to the production setup documentation).

The OIDC Resource Server functionality is experimental and is disabled by
default. It needs to be thoroughly tested before enabling it in production.

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

TBD

## Decision Drivers

TBD

## Considered Options

TBD

## Decision Outcome

TBD


### Positive Consequences

TBD

### Negative Consequences

TBD

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}

