# Git Select
> `git checkout` with ease

This little tool improves UX of the `git checkout` command.
Right now it supports only basic operations.

### Usage:
```shell
    git select -b develop
    # This will create the branch `develop` if it doesn't exists.
    
    git select
    # This will open a select menu where you can search 
    # branches and checkout on Enter. 
```

### Installation:
Download the latest release [here](https://github.com/Rawnly/git-select/releases/latest)

### Features
- [x] Checkout branch via interactive prompt
- [x] Create and checkout a new branch
- [ ] Check for git installation with `cli/safeexec`.
- [ ] Support all the `git checkout` flags/parameters.