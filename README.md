# git-ng

Git subcommand for Angular style commit message.

(actually, only add a prefix to your original commit message)

~~~
go get -u github.com/liangzuobin/git-ng
~~~

~~~
Useage:
  git ng -flag message.
Example:
  git ng -f new feat.
Flags:
  must be one (and the only one) of the following:
  -f --feat: A new feature
  -x --fix: A bug fix
  -d --docs: Documentation only changes
  -s --style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
  -r --refactor: A code change that neither fixes a bug or adds a feature
  -p --perf: A code change that improves performance
  -t --test: Adding missing tests
  -c --chore: Changes to the build process or auxiliary tools and libraries such as documentation generation
~~~

~~~
➜  git-ng master ✗ git ng -r main
➜  git-ng master ✓
commit ddf2556a514521af5fb166d8358e6b20fa18d634 (HEAD -> master)
Author: liangzuobin <liangzuobin123@gmail.com>
Date:   Fri Mar 15 09:40:40 2019 +0800

    Refactor: main
~~~