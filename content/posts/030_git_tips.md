---
title: "A few GIT tips"
date: 2024-11-16T17:08:39+02:00
draft: false
---

Sooner or later, most of tech guys (and girls) accumulate accounts to git servers (i.e., GitHub, Bitbucket, GitLab, etc.) For example, one can have a personal GitHub account and a work account. It would be nice to set them up independently. Below are a few git setup tips: 

1. **Prefer SSH over HTTPS**

    It's secure and simply more convenient.
    ```shell
    git clone git@github.com:ataraskov/task-dashboard.git # example of clone command
    git remote -v # check what is used as remote
    ```

2. **Use ssh-agent**

    SSH agent help to manage your private keys. Ideally if your password manager supports ssh-agent integration ([1password](https://developer.1password.com/docs/ssh/agent/), [keepassxc](https://keepassxc.org/docs/KeePassXC_UserGuide#_setting_up_ssh_agent_integration), etc.).

    ```shell
    ssh-add -l # list keys from ssh-agent
    ```

3. **Use single-purpose ssh keys**

    One project - one key. 
    
    This is security hygiene 101. Using a single ssh key for all the needs may seem like a good idea at first. At least use a dedicated ssh key for each context (i.e., hobby, employer-abc, project-xyz, etc.).

4. **Keep your keys in a safe place**

    File system is not a good place for private keys. Let's move them into a good password/secret storage (i.e., password manager).

    It's enough to store public keys only in your `~/.ssh` directory.
    ```shell
    ls -l ~/.ssh/project-xyz.pub
    ```

5. **Use ssh_config**
    
    For example, one can choose to set `IdentitiesOnly` in `~/.ssh/config` for all hosts, like below:
    ```shell
    # Defaults
    Host *
        IdentitiesOnly yes
    ```
    This option prevents "leakage" of public ssh keys from your system to the target server (just in case).

6. **Specify key in each context**

    This one may look like a tedious one. But we have a few tricks to aid us.
    
    Let's have an example here. John has two contexts:
    1) `personal` github account (i.e. hobby)
    2) `project-xyz` github account (i.e. work)
    ---
    Our Example John follows below steps:

    6.1) Make sure we have just public keys in `~/.ssh`

    ```shell
    ls ~/.ssh
    ```
    Example output:
    ```shell
    config
    personal.pub
    project-xyz.pub
    ```

    6.2) Update `~/.ssh/config` with custom hosts for each context

    We will use a bit of `ssh_config` magic to configure custom hosts. That allows us to attach different settings to the same target host.
    ```shell
    # Defaults
    Host *
        IdentitiesOnly yes

    # Personal
    Host github.com-personal
        HostName github.com
        User git
        IdentityFile ~/.ssh/personal.pub

    # Project XYZ
    Host github.com-project-xyz
        HostName github.com
        User git
        IdentityFile ~/.ssh/project-xyz.pub
    ```

    6.3) Clone repos using custom hosts

    Now we can use our custom host names in ssh commands (and git as well):
    ```shell
    git clone git@github.com-personal:john/hobby.git
    ```

    6.4) Update gitconfig to force ssh over http

    ```shell
    git config --local \
        url."git@github.com-personal".insteadOf "https://github.com"

    ```

    6.5) What about other commands?

    Yes `go mod tidy` or `go get ...` will fail, unless we say what ssh identity to use.

    Fortunately, that can be done via an environment variable:

    ```shell
    export GIT_SSH_COMMAND="ssh -i ~/.ssh/personal.pub"
    go mod tidy
    ```

    One still can and will face issues here and there.
    But that is totally different journey to take.

For some it may be an overkill, but others can learn something helpful. I hope ;)