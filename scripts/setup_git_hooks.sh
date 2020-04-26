#!/bin/bash

echo "Creating link to custom hooks from .git/hook"
rm .git/hooks/pre-commit
rm .git/hooks/prepare-commit-msg
rm .git/hooks/commit-msg

ln -s -f ../../githooks/pre-commit .git/hooks/pre-commit
ln -s -f ../../githooks/prepare-commit-msg .git/hooks/prepare-commit-msg
ln -s -f ../../githooks/commit-msg .git/hooks/commit-msg
