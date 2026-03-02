#!/usr/bin/env bash
go run . reset
# Expecting exit code: 0
go run . register kahya
# Expecting exit code: 0
go run . addfeed "Ycombinator Hacker News RSS" "https://news.ycombinator.com/rss"
# Expecting exit code: 0
go run . agg 5s
