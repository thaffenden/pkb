# Schema

The documents in this directory define the [JSON schema](https://json-schema.org/)
for the configuration and files generated by `pkb`.

You can annoate your config file with the `$schema` keyword for a better
editing experience (provided your editor integrates with JSON schema tooling).

e.g.:

```json
{
  "$schema": "https://github.com/thaffenden/pkb/blob/main/schema/config.json",
  "directory": "/home/user",
  ...
}
```

Changes to the schame should be validated through the `make lint-schema`
command to ensure they are correct.

## WIP

The following schemas are for work in progress functionality so may be subject
to change:

- `reading-list.json`
- `todo-list.json`

## Examples

Some example files have been added to display what the configuration should
look like. If you try editing these with an editor that supports JSON schemas
you should see the field names in your auto complete and the description of
what each field is for.