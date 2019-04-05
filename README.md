# git-ng

Git subcommand for Angular style commit message.

(actually, only add a prefix to your original commit message)

~~~
go get -u github.com/liangzuobin/git-ng
~~~

~~~
Useage:
  git ng --type-flag subject [-o scope] [-b body] [-e footer]
Example:
  git ng -f new feat
Flags:
  type flag must be (one and only) one of the following:
  -f --feat: A new feature
  -x --fix: A bug fix
  -d --docs: Documentation only changes
  -s --style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
  -r --refactor: A code change that neither fixes a bug or adds a feature
  -p --perf: A code change that improves performance
  -t --test: Adding missing tests
  -c --chore: Changes to the build process or auxiliary tools and libraries such as documentation generation

  optional flags:
  -o scope: Scope can be anything specifying place of the commit change
  -b body: Motivation for the change and contrasts with previous behavior
  -e footer: Breaking changes, referencing issues, etc.
~~~

Example: (in fish shell)
~~~
➜  git-ng master ✓ git ng -f 'more flag' -b '- body
                   - scope
                   - footer'
➜  git-ng master ✓ git log
commit 6e4a2fb41495f41795ed81068f178a7bab9528c1 (HEAD -> master)
Author: liangzuobin <liangzuobin123@gmail.com>
Date:   Sat Mar 16 15:44:02 2019 +0800

    feat: more flag

    - body
    - scope
    - footer
~~~
