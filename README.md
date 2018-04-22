# Stagehand

Stagehand is a command-line tool to open applications and position windows based on recorded preferences.

## Quickstart

To start using stagehand, first install stagehand using `go get`:

```bash
go get -u github.com/spf13/cobra/cobra
```

Then create `~/stagehand/workspaces/main.yml` which will store the configuration for your main workspace.

```yaml
[
  name: <Application Name>
]
```

If the file doesn't exist or does not contain at least one application name, stagehand will exit with an error message

## TODO
Currently being built in stages, starting with small features and slowly adding features from the TODO list

1. ~~Basic cobra skeleton~~
2. ~~Open terminal~~
3. Enable arbitrary applications to be opened via YAML file
4. Position terminal window based on YAML file (handle multiple screens?)
5. Replace YAML file with `add` command and a database
6. Record window position with command line (Is this possible?)
7. Implement `position` command.

Eventually
a. Allow multiple sets of applications
b. Handle multiple windows from same application

## Proposed usage/commands
```
Usage:
    stagehand                     Opens specified applications and moves their windows
    stagehand [command]           Executes command as described below

Available Commands:
    add    [application]          Add a new application
    record [application]          records desired position of application window
    remove [application]          removes application from list
    list                          List all applications to open
    position [application]        moves application window to defined position

Use "stagehand [command] --help" for more information about a command.

$ stagehand add Terminal
Added "Terminal" to list of applications to open.

$ stagehand record Terminal
Saved current position of "Terminal" window

$ stagehand remove Terminal
Removed "Terminal" to list of applications to open.

$ stagehand list
Stagehand will open the following applications:
1. Terminal
2. Sublime

$ stagehand position Terminal
"Terminal" window moved to desired location
```
