---
title: "Fail spectacularly with AI"
date: 2025-06-08T13:16:33+02:00
draft: false
tags:
  - learn
  - failure
---

What can be juicier than reflecting on a failure? Recently, I caused a production downtime.

The incident started small. A performance issue on a small project. CPU usage was high, and mongod was among the top consumers. Easy, right? Just ask your favorite LLM-based tool to guide you through the troubleshooting process. And follow its advice...

I have a DBA background (spent a bit of my time in RDBMS world), though I have little to no experience with MongoDB. So I asked AI how to investigate CPU usage. And its answer was totally logical and expected. "Index addition looks good enough", I thought. Though the application was not 'happy' about that change - not because it's badly engineered or unstable. The root cause is my overconfidence, fueled by AI's guidance.

Lessons learned:
* Team communication is critical.
* AI tools are a great help within your professional domain.
* A junior with a fancy _toy_ tool, is still a junior.

AI-agents are a great tool, which amplifies one's productivity. Though it also works as a magnifying glass for one's shortcomings as well ;)
