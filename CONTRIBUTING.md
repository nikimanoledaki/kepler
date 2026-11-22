# Contributing to Kepler 🌱

Thanks for your interest in contributing to Kepler! This guide will help you get started with contributing to our project.

## What Should I Know Before I Get Started? 🤔

Kepler is a Go-based project focused on sustainable computing. Before contributing, it's helpful to:

- Have basic knowledge of Go programming
- Understand container technologies and Kubernetes
- Familiarize yourself with the project structure by exploring the repository
- Read our [GOVERNANCE.md](GOVERNANCE.md) and [MAINTAINERS.md](MAINTAINERS.md) documents

## Gen AI policy

Our project adheres to the Linux Foundation's Generative AI Policy, which can be viewed at [https://www.linuxfoundation.org/legal/generative-ai](https://www.linuxfoundation.org/legal/generative-ai).

Disclose AI assistance in the pull request description. You are responsible for every line you submit and must be able to explain it in review.

Pull request descriptions, issue comments, and replies to review should be in your own words, not pasted from an AI tool. AI-assisted code is fine; the conversation with maintainers should come from you.

## How Do I Start Contributing? 🚀

1. Find or open an issue, and wait for a maintainer to triage it (see [Pull Request Policy](#pull-request-policy-))
2. Fork the repository on GitHub
3. Clone your fork locally
4. Set up your development environment (see below)
5. Create a new branch for your work
6. Make your changes
7. Submit a pull request that links the triaged issue

Don't forget to sign your commits according to our [DCO](DCO) requirements!

### Keep Your Company Affiliation Up to Date 🏢

The CNCF tracks project health on [Kepler DevStats](https://kepler.devstats.cncf.io/d/5/companies-table?orgId=1) using affiliation data from [cncf/gitdm](https://github.com/cncf/gitdm). DevStats attributes contributions to a company only through a `cncf/gitdm` entry, so contributors without one count toward no company.

If you are affiliated with a company, university, or research institution, add or update your affiliation in [cncf/gitdm](https://github.com/cncf/gitdm) following the instructions in its README.

## How Can I Contribute? 💡

### Continuous Integration 🔄

Our project uses GitHub Actions for CI. Each PR will trigger automated builds and tests. Check the `.github/workflows` directory to understand our CI pipeline. Ensure your contributions pass all CI checks before requesting a review.

### Pull Request Policy 🚦

Maintainers agree on the direction of a change in an issue before they review
code in a pull request. A large diff for a direction the project has not agreed
on costs reviewers time and costs you work.

#### Triage First

Every pull request links an issue that a maintainer has triaged. An issue counts
as triaged when it carries one of these labels:

- `triage/accepted`
- `help wanted`
- `good first issue`

Only maintainers can apply these labels. Opening an issue and a pull request
together does not count as triage.

A `good first issue` also needs a maintainer to assign it to you before you open
a pull request. Comment on the issue to ask for it.

Link the issue in the pull request description, for example `Closes #123` or
`Part of #123`. Features and design changes also need an accepted
[enhancement proposal](docs/developer/proposal/EP-TEMPLATE.md) before
implementation.

These pull requests skip the triage requirement:

- Pull requests from maintainers
- Pull requests from bots, such as Dependabot
- Pull requests that change at most 20 lines across at most 2 files, such as
  typo fixes
- Pull requests a maintainer labels `policy/exempt`

#### Closing Rules

Any maintainer may close a pull request under these rules. The closing comment
names the rule.

Enforced by the [`pr-policy`](.github/workflows/pr-policy.yaml) workflow:

1. **No triaged issue.** The pull request does not link an issue with one of
   the labels above.
2. **Too many open pull requests.** Contributors without write access may have
   at most 3 open pull requests at once, drafts included.

Enforced by the [`stale`](.github/workflows/stale.yml) workflow:

3. **Waiting on author.** A maintainer labels the pull request
   `waiting-on-author` when review comments or failing CI need the author's
   action. With no activity for 14 days, the pull request becomes stale, and
   it closes 7 days later.

Applied by maintainers:

4. **Duplicate or conflicting.** The pull request overlaps an open pull request
   or proposal and does not explain how they relate.
5. **Claimed by someone else.** The linked issue is assigned to another
   contributor who is active on it.
6. **No accepted proposal.** The pull request adds a feature or changes the
   design without an accepted enhancement proposal.
7. **Too large.** The pull request changes more than 500 lines, excluding
   generated files, without an accepted proposal or a plan to split it.

Contributors without write access may hold at most 3 assigned issues at once.
Maintainers unassign an issue after 14 days without a linked pull request, so
others can pick it up.

#### After a Close

A close is not a judgment of the work. Once the linked issue carries one of the
labels above, open a new pull request or ask a maintainer to reopen it.

### Code Review 👀

All submissions require review. We use GitHub pull requests for this purpose:

1. Submit a pull request from your fork to our main repository
2. Ensure all CI checks pass
3. Address feedback from maintainers
4. Once approved, a maintainer will merge your changes

### Commit Hygiene

Kepler uses merge commits, so the individual commits in a PR remain in the
project history. A coherent commit sequence helps future contributors understand
changes and makes debugging and reverting easier.

- Aim for each commit to represent a coherent change with a meaningful message.
- Keep separate commits for distinct, useful logical steps within a PR.
- Generally fold intermediate fixes, such as link corrections and review
  follow-ups, into the commits they belong to before merging.
- Order commits so the progression of the change is easy to follow.

These are commit hygiene principles, not strict rules or a requirement to squash
every PR into a single commit.

### Reporting Bugs 🐛

Found a bug? Please report it by creating an issue with the bug template. Include:

- A clear and descriptive title
- Steps to reproduce the issue
- Expected vs actual behavior
- Screenshots or logs if applicable
- Your environment details (OS, Go version, etc.)

### Suggesting Enhancements ✨

#### How Do I Submit A (Good) Enhancement Suggestion?

Enhancement suggestions are tracked as GitHub issues. Create an issue with the enhancement template and provide:

- A clear and descriptive title
- A detailed description of the proposed enhancement
- Any potential implementation ideas you have
- Why this enhancement would be useful to most Kepler users

#### How Do I Submit A (Good) Improvement Item?

For smaller improvements to existing functionality:

- Focus on a single, specific improvement
- Explain how your improvement makes Kepler better
- Provide context around why this improvement matters
- Link to any related issues or discussions

## Development 💻

### Set up your dev environment

1. Install the Go version specified in the repository's [`go.mod`](go.mod) file
2. Install pre-commit hooks:

```bash
pre-commit install
```

See [docs/developer/pre-commit.md](docs/developer/pre-commit.md) for more information.

### Testing

Run tests with:

```bash
make test
```

Generate coverage reports with:

```bash
make coverage
```

Our CI automatically runs tests and uploads coverage to Codecov.

## Your First Code Contribution 🎉

Looking for a place to start?

1. Look for issues labeled `good first issue` or `help wanted`
2. Comment on the issue before starting work, and ask a maintainer to assign it to you
3. Read the code in the area you want to work on to understand patterns
4. Ask questions in issues or on Slack if you need help

Remember that even small contributions like fixing typos or improving documentation are valuable!

Thanks for contributing to Kepler! 💚
