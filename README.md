1password-linux-to-bitwarden
=

This tool allows you to take an `.1pux` export file of 1Password 8 & convert it to a JSON file importable by Bitwarden.

It has only been tested on Linux by me, but nothing suggests it shouldn't work on macOS & Windows too.

```
Usage of ./1password-linux-to-bitwarden:
  -1pass-backup-file-path string
        The absolute path (location) of a 1Password .1pux backup file
  -bitwarden-export-file-path string
        The absolute path to where you want to export the JSON to import into Bitwarden (defaults to your current working directory)
```
