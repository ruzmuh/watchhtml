# WatchHTML Tool

WatchHTML is a Go-based utility designed to track changes in the HTML of a given webpage. It operates like a cron job, periodically checking the webpage and storing an initial snapshot of the HTML elements specified by an XPath. On subsequent checks, it compares the current HTML to the stored snapshot. If a change is detected, a notification is sent to a specified Slack channel through a webhook.

## Getting Started

These instructions will guide you through getting a copy of the project and running it on your local machine for development and testing purposes.

### Prerequisites

You will need Go installed on your machine. For instructions on how to install Go, please refer to the [official Go documentation](https://golang.org/doc/install).

### Installation

1. Clone the repository into your workspace:

```
git clone https://github.com/ruzmuh/watchhtml.git
```

2. Navigate to the cloned repository:

```
cd watchhtml
```

3. Compile and install the package:

```
go install
```

## Usage

The tool is run as a standalone task, akin to a cron job:

```
watchhtml --url <url> --xpath <xpath> --slackwebhook <webhook> --storedir <dir>
```

You can also use environment variables prefixed with `WATCHHTML_` in place of flags.

### Flags / Environment Variables

* `--slackwebhook string` / `WATCHHTML_SLACKWEBHOOK`: Slack webhook URL for notifications.
* `--storedir string` / `WATCHHTML_STOREDIR`: Directory for storing HTML snapshots. By default, this is the current directory (`"./"`).
* `--url string` / `WATCHHTML_URL`: URL of the webpage to track for changes.
* `--xpath string` / `WATCHHTML_XPATH`: XPath of the HTML elements to track on the webpage.
* `--extramessage string` / `WATCHHTML_EXTRAMESSAGE`: extra message content. e.g mention user

**Note**: The `--version` and `--help` flags do not have corresponding environment variables.

### Print Version and Help

* `--version`: Prints the version of the tool and exits.
* `--help`: Displays usage information.

## Versioning

To check the version of the tool, use the `--version` flag:

```
watchhtml --version
```
