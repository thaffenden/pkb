# pkb

**P**ersonal **K**nowledge **B**ase

Config driven CLI to manage notes and todo items to power a personal knowledge base.

## Commands

Run `pkb --help` for a full list of available commands.

## Configuration

Config driven means you can control how `pkb` works with the options in a
config file.

By default `pkb` checks for the config file in your `XDG_CONFIG_DIR`, or
`$HOME/.config`.

### Fields

You can see an example of the config file format in the
[example.json](./example.json) in the root of this repo.

- `directory` - the directory you want to store all of your notes in.
- `editor` - the editor to use to open notes after they have been created.
- `templates` - the options for the different types of notes you might want
to create.

### Expanding values in templates

The following values will be automatically expanded in templates:

- `{{.CustomDateFormat}}` - the current date in a custom format specified in the
template config. This must be a valid golang date format, with the exception
of day suffixes (e.g. 1st, 3rd etc). If your format contains a day suffix this
will be be handled so the suffix is correctly displayed. The format string needs
to be included in the template config to be able to use a custom date format.
- `{{.Date}}` - the current date in the format YYYY-MM-DD
- `{{.Name}}` - the name of the created document
- `{{.Time}}` - the time the file was created in the format HH:MM

### Custom name formats

By default when you create a new file from a template you will be prompted to
enter the name value, however you can customise with `name_format`.

The following values are currently supported:

- `DATE` - the current date in the format YYYY-MM-DD
- `PROMPT` - prompt for user input

You can combine the supported formats to use them both, e.g.:

```json
"name_format": "DATE-PROMPT"
```

In this example the document would be created with the current date then the
value you typed in the prompt, e.g. `2022-09-19-typed-value.md`.

### Selecting an output directory

If you don't want to hardcode the output directory in your config you can use
the dynamic values to prompt you for input at time of creation.

The following values are currently supported:

- `{{Prompt}}` - will let you type in a new directory name. If the directory
does not already exist it will be created.
- `{{Select}}` - select from existing directories inside the parent.
