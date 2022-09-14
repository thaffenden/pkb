# pkb

**P**ersonal **K**nowledge **B**ase

Config driven cli to manage notes and todo items to power a personal knowledge base.

## Expanding values in templates

The following values will be automatically expanded in templates:

- `{{.Name}}` - the name of the created document
- `{{.Date}}` - the current date in the format YYYY-MM-DD
- `{{.Time}}` - the time the file was created in the format HH:MM

## TODO

- [] flesh out README
- [] add workflow to build and release binary on merge to master
- [] add install script
- [] add unit tests for actual file creation
- [] add some actual end user tests using demo config
- [] add logic to check for updates when running
- [] add templating expansion for common values (date, tags, etc)
