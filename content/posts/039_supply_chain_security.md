---
title: "Supply Chain Security"
date: 2026-04-15T12:10:28+01:00
draft: false
tags:
  - git
  - security
---

Supply chain attacks are nothing new, unfortunately. Though the last few got me scared. Like [this](https://futuresearch.ai/blog/litellm-attack-transcript/) one, around LiteLLM python library. The LiteLLM one was particularly nasty one, as it was enough to just install the library.

I love open-source community and what it provides to all of us. But social engineering makes it too hard to keep your keys private.

The good news is, private ssh keys are more or less safe when used via ssh-agent, and keylogger is hard to setup without privilege escalation (at least on linux). So please, please, please do not blindly type your password in a random pop-up window.

References:
* [AI-Enabled Social Engineering](https://cloud.google.com/blog/topics/threat-intelligence/unc1069-targets-cryptocurrency-ai-social-engineering)
* [Trivy ecosystem supply chain temporarily compromised](https://github.com/aquasecurity/trivy/security/advisories/GHSA-69fq-xp46-6x23)
