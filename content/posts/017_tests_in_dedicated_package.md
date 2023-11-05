---
title: "Tests in dedicated package"
date: 2023-11-05T16:15:18+01:00
draft: false
tags:
  - go
  - programming
  - methodology
---

Documentation for the Go testing module [says](https://pkg.go.dev/testing):
> The test file can be in the same package as the one being tested, or in a corresponding package with the suffix "_test".

In other words, we are free to use dedicated packages for tests, and this is one of the supported ways to structure a project. A simple search on GitHub shows the following adoption results:
* [~1M use a dedicated test package](https://github.com/search?q=language%3Ago++path%3A*_test.go+%2F%5Epackage+.*_test%2F&type=code)
* [~3M use a production package for tests](https://github.com/search?q=language%3Ago++path%3A*_test.go+%2F%5Epackage+.*%2F&type=code)

Why use a dedicated test package at all? What does it bring?
* **Tests resiliency**: Tests are not tightly coupled with internal implementation as we have access to exported functionality. Changes to the actual implementation do not break the tests. But changes to exposed APIs do.
* **Improved API for packages**: We are forced to use our packages as clients from the beginning. Bad design will surface much earlier in the process (ideally before implementation, as we follow the red/green/refactor cycle of TDD).
* **Separation of test dependencies**: Fewer dependencies in production code translate to fewer worries. Test dependencies are not even supposed to run in production.

Isn't it breaking TDD's approach? A little... if you are a purist. But don't forget about the magical 20/80 ratio, 20% effort will give us 80% code coverage. A higher code coverage ratio brings more problems than benefits, IMHO.

References:
* [GO: Testing package](https://pkg.go.dev/testing)
