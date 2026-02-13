# Clone Organization Repositories

Clone or update repositories from a GitHub organization to local disk.

## Prerequisites

- Set a `GITHUB_TOKEN` environment variable.
- The token must have access to read repositories in the target organization.

## Usage

`clone-org-repos` has 2 arguments:

- `org`, `-o`: Name of the organization to clone. Required.
- `path`, `-p`: Path to clone repositories into. Optional. Defaults to the user's home directory.

```bash
$ ./clone-org-repos -o sous-chefs

Cloning repository: rvm
SSH Clone URL: git@github.com:sous-chefs/rvm.git
git pull origin
commit 84ce9add8a421f278830ffb7192fc7d9b0e82438
Author: Lance Albertson <lance@osuosl.org>
Date:   Thu Oct 08 09:09:37 2020 -0700

    Merge pull request #406 from sous-chefs/automated/standardfiles

    Automated PR: Standardizing Files
```

```bash
$ ./clone-org-repos -o sous-chefs -p ~/mydevfolder/sous-chefs

Cloning repository: rvm
SSH Clone URL: git@github.com:sous-chefs/rvm.git
git pull origin
commit 84ce9add8a421f278830ffb7192fc7d9b0e82438
Author: Lance Albertson <lance@osuosl.org>
Date:   Thu Oct 08 09:09:37 2020 -0700

    Merge pull request #406 from sous-chefs/automated/standardfiles

    Automated PR: Standardizing Files
```
