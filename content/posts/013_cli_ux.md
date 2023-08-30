---
title: "User eXperience for CLI applications"
date: 2023-08-29T20:15:08+02:00
draft: false
tags:
  - cli
  - software development
  - methodology
---
History goes in cycles. Decades ago, we interfaced computers via command prompts. Then graphical user interface brought computing to everage Joe's home. And command line interface faded to the background. But few years ago command line utilities became fashinoble again. There are a lot of debates around GUI vs CLI. But it's obvious that UX is applicable to both.

While CLI tools are back in town, they have changed. Modern CLI is a bit different from classic tools like `cp`, `mkdir`, `tar`, ... They are designed for humans first and machines second.

Below are my favourite tricks To provide a better user experience for CLI tools:
* Preffer flags over positional arguments.
* Use 'standard' names for flags.
* Provide both interactive and non-interactive modes.
* Print progress bar / spinner on long operations.
* Timeout operation if it takes too much time.
* Make timeouts configurable.
* Design a concise and useful help text.
* Provide full documentation, when help is not enough.
* Have sensible defaults.
* Allow configuration via conf file and/or environment variables.

And the last but not least: choose program name carefully, [labeling a concept changes how we perceive it](https://en.wikipedia.org/wiki/Labeling_theory).
BTW, don't follow any recommendations blindly, see what works for your project ;)

References:
* [Wikipedia: Command-line interface](https://en.wikipedia.org/wiki/Command-line_interface)
* [CLI Guildlines](https://clig.dev)

