export XDG_CONFIG_HOME="$HOME/.config"
export BASE_PATH="/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin"
export PATH="$HOME/.mashina/bin:$BASE_PATH"

export MASHINA_CONFIG="$XDG_CONFIG_HOME/mashina"
export MASHINA_INIT="$MASHINA_CONFIG/init"
export MASHINA_BIN="$MASHINA_CONFIG/bin"
export MASHINA_OPT="$MASHINA_CONFIG/opt"
export MASHINA_FUNCTIONS="$MASHINA_CONFIG/functions"
export MASHINA_TMP="$MASHINA_CONFIG/tmp"

export PATH="$MASHINA_BIN:$PATH"

function mashina__is_variable_defined {
  if [[ -n "$1" ]]; then
    return 0
  else
    return 1
  fi
}

function mashina__is_bash {
  mashina__is_variable_defined "$BASH_VERSION"
}

function mashina__is_zsh {
  mashina__is_variable_defined "$ZSH_VERSION"
}

export STARSHIP_CONFIG="$HOME/.config/starship/starship.toml"

if $(mashina__is_bash); then eval "$(starship init bash)"; fi
if $(mashina__is_zsh); then eval "$(starship init zsh)"; fi
