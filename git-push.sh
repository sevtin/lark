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