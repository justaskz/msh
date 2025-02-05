#!/usr/bin/env bash

mkdir -p "$HOME/.mashina/bin"

if [[ -n "$MASHINA_DEV" ]]; then
  echo "DEV MODE: on"
  # ln -sf "$PWD/bin/mashina" "$HOME/.mashina/bin/mashina"
  ln -sf "$PWD/bin/mashina" "$HOME/.config/mashina/bin/mashina"
else
  wget https://raw.githubusercontent.com/justaskz/mashina/master/poc/mashina/bin/mashina
fi
