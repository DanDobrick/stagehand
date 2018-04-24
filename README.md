# Stagehand

Stagehand is a command-line tool to open applications and position windows based on recorded preferences.

## Quickstart

If you have `go` installed and your `$GOBIN` in your path, simply run:

```bash
go get -u github.com/DanDobrick/
```

Installing go is out of the scope of these instructions, visit https://golang.org/doc/install#install for more information.

After `stagehand` is installed you need to create `~/stagehand/workspaces/main.yml` which will store the configuration for your main workspace. Add applications to `main.yaml` in the following way:

```yaml
- name: application name
  file: file to open with application (OPTIONAL)
  pos: X-coordinate, Y-coordinate, width, height (OPTIONAL)
```

See `main.yaml.example` for concrete examples

If the file doesn't exist or does not contain at least one application name, stagehand will exit with an error message

# Positioning Error
There are two reasons that you will receive an error with the "See 'README.md#Positioning Error' for more info" message:
1. When resizing the window, something about the application changes the window size or prevents the window from resizing exactly. Since `stagehand` waits for the window to be resized by polling the window size and comparing it to the expected size, you might receive this message although the program continued to execute
2. The application is not Scriptable. Some applications (such as Sublime Text) refuse to be scripted, but there is a workaround using System Events. More details here when the work-around is implemented

## TODO
Currently being built in stages, starting with small features and slowly adding features from the TODO list

1. ~~Basic cobra skeleton~~
2. ~~Open terminal~~
3. ~~Enable arbitrary applications to be opened via YAML file~~
4. ~~Position terminal window based on YAML file~~
5. ~Update YAML file with `add` command~
6. Implement `position` command.
7. "Not scriptable" positioning and directions to README
7. Add testing
8. Record window position with command line
9. Allow user to specify yaml filename (still in directory)
10. Allow user to specify default directory for yaml files
  - My idea was to write something to the default yaml file telling the program to look elsewhere

Eventually
a. Handle multiple windows from same application
b. Add config file for default yaml location and default window size

## Proposed usage/commands
```
Usage:
    stagehand                     Opens specified applications and positions their windows
    stagehand [command]           Executes commands as described below

Available Commands:
    record [application]          records desired position of application window
      - add a flag to write this information to YAML file
    remove [application]          removes application from YAML file
    position [application]        moves application window to defined position

Use "stagehand [command] --help" for more information about a command.

$ stagehand record Terminal
Saved current position of "Terminal" window

$ stagehand remove Terminal
Removed "Terminal" from YAML file <filename>

$ stagehand position Terminal
"Terminal" window moved to desired location
```
