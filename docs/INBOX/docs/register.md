---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, tutorial, secrets, secret registration
title: Registering Secrets
description: send those bad boys to your workloads
micro_nav: true
page_nav:
  prev:
    content: quickstart
    url: '/docs/'
  next:
    content: overview and prerequisites
    url: '/docs/use-case-overview'
---

<p style="text-align:right;position:relative;top:-40px;"
><a href="https://github.com/ShieldWorks/aegis-web/blob/main/docs/register.md"
style="border-bottom: none;background:#e0e0e0;padding:0.5em;display:inline-block;
border-radius:8px;">
edit this page on <strong>GitHub</strong> ✏️</a></p>

## Introduction 🐢

This document lists various use cases to register secrets to Kubernetes
workloads using **Aegis**, in tutorial form. Each tutorial isolated in 
itself, explaining a specific feature of **Aegis**. 

Following these tutorials, you will have a better understanding of what
**Aegis** is capable of, you will learn core **Aegis** concepts by doing. 

When you complete the tutorials listed here, you will have a fair understanding
of how to use **Aegis** to manage your secrets.

## Follow the White Rabbit 🐇

We advise you to follow these tutorials in the sequence they are presented here. 
We’ve structured them this way to start with simpler use cases and progressively 
introduce more advanced techniques as we build upon our knowledge.

1. [Overview and Prerequisites](/docs/use-case-overview)
2. [Using **Aegis Sidecar**](/docs/use-case-sidecar)
3. [Using **Aegis SDK**](/docs/use-case-sdk)
4. [Using **Aegis Init Container**](/docs/use-case-init-container)
5. [Encrypting Secrets](/docs/use-case-encrypt)
6. [Transforming Secrets](/docs/use-case-transform)
7. [**Aegis** Showcase](/docs/use-case-showcase)

## Further Reading

The use cases above leverage **Aegis Sentinel** and **Aegis SDK**. For the
interested, the following sections cover these tools in greater detail:

* [**Aegis SDK** Documentation](/docs/sdk)
* [**Aegis Sentinel** Command Line Reference](/docs/sentinel)