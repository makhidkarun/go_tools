# go_tools
Various tools written in Go

There are no binaries included in the repo. 

From the base go_tools/ directory you build the tools. There
are currently just a few, status_report, chargen, and vng 
the Vargn name generator.

  go build -o bin/status_report  cmd/status_report/main.go

Currently working on the issues of locating the names.db file.
It is a SQLite database that currently must be in the directory
you run chargen from.

  Addendum: Added a -d <database_name_with_path> flag to chargen.
    Defaults to data/names.db


