#! /bin/bash

curl -s https://zone01normandie.org/assets/superhero/all.json | jq ' .[] | select( .id == '$HERO_ID') | .connections .relatives' | cut -d '"' -f2
