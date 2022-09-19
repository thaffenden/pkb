# pkb

**P**ersonal **K**nowledge **B**ase

Config driven cli to manage notes and todo items to power a personal knowledge base.

## Expanding values in templates

The following values will be automatically expanded in templates:

- `{{.CustomDateFormat}}` - the current date in a custom format specified in the
template config. This must be a valid golang date format, with the exception
of day suffixes (e.g. 1st, 3rd etc). If your format contains a day suffix this
will be be handled so the suffix is correct displayed. The format string needs
to be included in the template config to be able to use a custom date format.
- `{{.Date}}` - the current date in the format YYYY-MM-DD
- `{{.Name}}` - the name of the created document
- `{{.Time}}` - the time the file was created in the format HH:MM

## TODO

- [] flesh out README
- [] add workflow to build and release binary on merge to master
- [] add install script
- [] add unit tests for actual file creation
- [] add some actual end user tests using demo config
- [] add logic to check for updates when running
