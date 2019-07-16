#!/usr/bin/env zsh

RED="\033[1;31m"
GREEN="\033[1;32m"
NOCOLOR="\033[0m"

# Candidates
delimiters=("'\t'" "' '" "':'" "','" )
iconDelimiters=("' '" "':'" "'::'")
tagDelimiters=("'\t'" "' '" "':'" "','")


function test_line () {
    line=${1}
    delimiter=${2}
    iconDelimiter=${3}
    tagDelimiter=${4}

    # Construct Command
    taggo_cmd="taggo -d ${delimiter} -p ${iconDelimiter} -i '2,3,4' -t '|TEST|' -c yellow -s ${tagDelimiter} -b 2"
    taggo_revert_cmd="${taggo_cmd} -r"
    fzf_cmd="fzf -0 -1 --ansi"

    # check argument
    err=$(echo hoge | eval ${taggo_cmd} 2>&1)
    if [[ $err == *"delimiter must not be substring of iconDelimiter"* ]]; then
        echo "${GREEN}[pass]${NOCOLOR} Invalid argument"
        return 0
    fi

    # Run
    after_taggo_line=$(echo ${line} | eval ${taggo_cmd})
    after_fzf_selected_line=$(echo ${line} | eval ${taggo_cmd} | eval ${fzf_cmd})
    after_taggo_reverted_line=$(echo ${line} | eval ${taggo_cmd} | eval ${fzf_cmd} | eval ${taggo_revert_cmd})

    # Foe Debug
    # echo "-- Original line -- "
    # echo ${line}
    # echo "-- After taggo line -- "
    # echo ${after_taggo_line}
    # echo "-- After fzf select line -- "
    # echo ${after_fzf_selected_line}
    # echo "-- After taggo reverted line -- "
    # echo ${after_taggo_reverted_line}
    if [[ ${line} == ${after_taggo_reverted_line} ]] ; then
        echo "${GREEN}[pass]${NOCOLOR} ${after_taggo_line}"
    else
        echo "${RED}[fail]${NOCOLOR}"
        echo "expected:${line}"
        echo "actual  :${after_taggo_reverted_line}"
        echo "Test for delimiter:${delimiter}, iconDelimiter: ${iconDelimiter}, tagDelimiter: ${tagDelimiter}"
    fi
}

cat ./resource/test_cases.txt | while read line
do
    for delimiter in ${delimiters}
    do
        for iconDelimiter in ${iconDelimiters}
        do
            for tagDelimiter in ${tagDelimiters}
            do
                test_line ${line} ${delimiter} ${iconDelimiter} ${tagDelimiter}
            done
        done
    done
done
