#! /bin/bash

curl https://zone01normandie.org/assets/superhero/all.json | jq -r '.[] | select ( .id == 170 ) | "(.name) /n (.powerstats.power) /n (.appearence.gender)"'

