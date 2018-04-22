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

See `main.yaml.example`

If the file doesn't exist or does not contain at least one application name, stagehand will exit with an error message

## TODO
Currently being built in stages, starting with small features and slowly adding features from the TODO list

1. ~~Basic cobra skeleton~~
2. ~~Open terminal~~
3. ~~Enable arbitrary applications to be opened via YAML file~~
4. Position terminal window based on YAML file (handle multiple screens?)
5. Update YAML file with `add` command
6. Implement `position` command.
7. Record window position with command line (Is this possible?)
8. Allow user to specify yaml filename (still in directory)
9. Allow user to specify default directory for yaml files
  - My idea was to write something to the default yaml file telling the program to look elsewhere

Eventually
a. Handle multiple windows from same application

## Proposed usage/commands
```
Usage:
    stagehand                     Opens specified applications and moves their windows
    stagehand [command]           Executes command as described below

Available Commands:
    add    [application]          Add a new application to YAML file
    record [application]          records desired position of application window
    remove [application]          removes application from YAML file
    position [application]        moves application window to defined position

Use "stagehand [command] --help" for more information about a command.

$ stagehand add Terminal
Added "Terminal" to YAML file <filename>

$ stagehand record Terminal
Saved current position of "Terminal" window

$ stagehand remove Terminal
Removed "Terminal" from YAML file <filename>

$ stagehand position Terminal
"Terminal" window moved to desired location
```
