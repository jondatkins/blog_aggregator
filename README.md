# Blog Aggregator CLI Tool

blog_aggregator / gator is a RSS feed aggregator app for the terminal.

## Requirements

It requires Go and Postgres to run the app.

## Installation

```bash
git clone git@github.com:jondatkins/blog_aggregator.git
cd blog_aggregator
go install
```

## Config and Usage

You will need a config file '~/gatorconfig.json'. This will have the form:

```json
{
  "db_url": "postgres://postgres:@localhost:5432/gator?sslmode=disable",
}
```

## Example Usage

```bash
# Register a new user
blog_aggregator register myUsername
# Login as user
blog_aggregator login myUsername
# Add a RSS feed for the current user to follow
blog_aggregator addfeed "Ycombinator Hacker News RSS" "https://news.ycombinator.com/rss"
blog_aggregator follow "https://hnrss.org/newest"
blog_aggregator unfollow "https://hnrss.org/newest"
# List all users
blog_aggregator users
# List all feeds
blog_aggregator feeds
# Show all feeds being followed by current user
blog_aggregator following
blog_aggregator 
# quit out of app
blog_aggregator exit

```
