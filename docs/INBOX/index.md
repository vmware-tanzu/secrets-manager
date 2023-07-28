---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: aegis, secrets management, secrets store, Kubernetes

title: Aegis
description: keep your secrets… secret
buttons:
- icon: home
  content: Quickstart
  url: '/docs'
  external_url: false
- icon: github
  content:  Source Code
  url: 'https://github.com/shieldworks/aegis'
  external_url: true
- icon: slack
  content:  Community
  url: '/contact#community'
  external_url: false

grid_navigation:
- title: <strong>Aegis</strong> Quickstart
  excerpt: Get your hands dirty. Install and use Aegis on your Kubernetes cluster.
  cta: get started
  url: '/docs'
- title: Production Tips
  excerpt: Production deployment recommendations to let your ops teams <code>#sleepmore</code>.
  cta: prepare your clusters
  url: '/production'
- title: Using <strong>Aegis</strong> SDK
  excerpt: Use <strong>Aegis Go SDK</strong> for a tighter integration with <strong>Aegis</strong> components.
  cta: dive in; water is warm
  url: '/docs/sdk'
- title: Keeping Secrets
  excerpt: A tutorial on how to dispatch secrets to workloads.
  cta: get your hands dirty
  url: '/docs/register'
- title: <strong>Aegis</strong> Community
  excerpt: Join us on <strong>Slack</strong>. It’s nice and cozy in here.
  cta: welcome to the jungle
  url: '/contact#community'
- title: Coming Up Next…
  excerpt: What we are planning to do in the near (<em>and far</em>) future.
  cta: see what’s cooking
  url: '/timeline'
---

## 🎉 Important Announcement: Aegis is Transitioning to VMware Secrets Manager for Cloud-Native Apps 🎉

Dear Aegis Community,

I am excited to announce that Aegis, your trusted open-source project for secrets 
management, is transitioning to a new home under VMware’s umbrella, with the new 
name “*VMware Secrets Manager for Cloud-Native Apps*” 
(*or Secrets Manager for short*).

This decision was not made lightly, but it is the best step forward for our 
community. Under VMware, the project will gain increased visibility, robust 
support, and a thriving ecosystem to accelerate its growth. 

<p style="background: #edc910; color: #000000;padding:1em;margin:1em 0 1em 0;">
Rest assured, <strong>the project will remain open-source</strong>, and I 
(<em>Volkan Özçelik</em>) will continue to serve as the core maintainer.</p>

Here’s what to expect in the coming weeks:

1. We will continue contributing to the current Aegis repository
   (<https://github.com/shieldworks/aegis>) and synchronizing it with the 
   private VMware repository. During this period, we will modify the code and rename 
   instances of “Aegis” as “VMware Secrets Manager.”
2. The public repo will be synchronized with the private repo under VMware’s GitHub.
3. Our documentation website (<https://aegis.ist/>) will undergo name and branding 
   changes. While the content will remain the same, you may notice differences in 
   the theme and references to “VMware Secrets Manager for Cloud-Native Apps” 
   instead of Aegis.
4. The documentation GitHub repository (<https://github.com/shieldworks/aegis-web/>) 
   will be merged into the main repository (<https://github.com/shieldworks/aegis>) 
   to expedite the transition.
5. Once the code is ready, we will submit it for VMware Open Source’s due 
   diligence and implement any changes based on their feedback.
6. Upon completing these steps, the project under VMware will become open to 
   public contribution. We will then archive the existing repositories with a 
   note directing users to the current code at VMware’s GitHub.

To ensure a seamless transition, we will continue to publish artifacts to 
DockerHub during this process.

We are grateful for your support and understanding during this transition. 

This move will give Aegis the resources to grow and serve our community better. 

Please feel free to reach out if you have any questions or concerns.

During this process, I’ll be transparent and inform you using various channels 
I have used.

Best Regards,

Volkan Özçelik

Core Maintainer, Aegis


<!--div style='padding:56.25% 0 0 0;position:relative;'>
  <iframe src='https://vimeo.com/showcase/10074951/embed' 
    allowfullscreen frameborder='0' 
    style='position:absolute;top:0;left:0;width:100%;height:100%;'></iframe>
</div-->

[spiffe]: https://spiffe.io/
[age]: https://age-encryption.org/

[contact]: /contact
[contribute]: /contributing
[coffee]: /coffee
[slack-invite]: https://join.slack.com/t/aegis-6n41813/shared_invite/zt-1myzqdi6t-jTvuRd1zDLbHX0gN8VkCqg "Join aegis.slack.com"
