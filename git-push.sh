#!/usr/bin/env bash

read -p "Input commit message: " commit_msg

if [ -z "$commit_msg" ]; then
  echo "commit message is empty"
  exit 1
fi

git add .

git commit -m "$commit_msg"

current_branch=$(git symbolic-ref --short HEAD)
echo "current branch: $current_branch"

git pull origin $current_branch
if [ $? -ne 0 ]; then
  echo "pull failed"
  exit 1
fi

git push origin $current_branch
if [ $? -ne 0 ]; then
  echo "push failed"
  exit 1
fi

#hint: You have divergent branches and need to specify how to reconcile them.
#hint: You can do so by running one of the following commands sometime before
#hint: your next pull:
#hint:
#hint:   git config pull.rebase false  # merge
#hint:   git config pull.rebase true   # rebase
#hint:   git config pull.ff only       # fast-forward only
#hint:
#hint: You can replace "git config" with "git config --global" to set a default
#hint: preference for all repositories. You can also pass --rebase, --no-rebase,
#hint: or --ff-only on the command line to override the configured default per
#hint: invocation.