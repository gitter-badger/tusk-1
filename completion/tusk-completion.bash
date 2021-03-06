#!/bin/bash

_tusk_bash_autocomplete() {
    local cur words opts meta
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    words="$( "${COMP_WORDS[@]:0:$COMP_CWORD}" --generate-bash-completion | cut -f1 -d":")"

    # Split words into completion type and options
    meta="$( echo "${words}" | head -n1 )"
    opts="$( echo "${words}" | tail -n +2 )"

    case "${meta}" in
        normal)
            declare -a values tasks flags
            values=( ${opts} )
            for option in "${values[@]}"; do
                if [[ "${option}" = --* ]]; then
                    flags+=("${option}")
                else
                    tasks+=("${option}")
                fi
            done

            if [[ "${cur}" = --* ]]; then
                COMPREPLY=( $(compgen -W "${flags[*]}" -- "${cur}") )
            else
                COMPREPLY=( $(compgen -W "${tasks[*]}" -- "${cur}") )
            fi
            ;;
        file)
            COMPREPLY=( $(compgen -f -- "${cur}") )
            ;;
    esac

    return 0
}

complete -o filenames -o bashdefault -F _tusk_bash_autocomplete tusk
