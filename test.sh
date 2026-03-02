#!/usr/bin/env bash
go run . reset
# Expecting exit code: 0
go run . register kahya
# Expecting exit code: 0
go run . addfeed "Hacker News RSS" "https://hnrss.org/newest"
# Expecting exit code: 0
go agg . 5s
