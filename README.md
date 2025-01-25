# valid

A command-line tool that validates whether input values meet specified rules.

## Description

`valid` is a lightweight and flexible tool for validating input values against formats, patterns, ranges, and character sets, ensuring they meet defined rules.
It supports a variety of validations, including:

- **Character validation**: Checks if input uses specific character sets like ASCII, digits, or alphanumeric characters
- **Range validation**: Verifies that values fall within specified numerical or date ranges
- **Pattern matching**: Validates formats such as email addresses, URLs, or semantic versions

Ideal for scripts, workflows, and applications, `valid` simplifies the validation process and ensures data consistency by applying reliable rules.

## Usage

```shell
Validates that input values meet specified rules

Usage:
  valid [flags]

Flags:
      --alpha                 validates if the value contains only English letters (a-zA-Z)
      --alphanumeric          validates if the value contains only English letters and digits (a-zA-Z0-9)
      --ascii                 validates if the value contains only ASCII characters
      --base64                validates if the value is valid Base64
      --digit                 validates if the value contains only digits (0-9)
      --domain                validates if the value is a valid domain
      --email                 validates if the value is a valid email address
      --enum string           validates if the value matches any of the specified enumerations
      --exact-length string   validates if the value's length is exactly equal to the specified number
      --float                 validates if the value is a floating-point number
  -h, --help                  help for valid
      --int                   validates if the value is an integer
      --json                  validates if the value is valid JSON
      --lower-case            validates if the value contains only lower case unicode letters
      --max string            validates if the value is less than or equal to the specified maximum
      --max-length string     validates if the value's length is less than or equal to the specified maximum
      --min string            validates if the value is greater than or equal to the specified minimum
      --min-length string     validates if the value's length is greater than or equal to the specified minimum
      --not-empty             validates if the value is not empty
      --pattern string        validates if the value matches the specified regular expression
      --printable-ascii       validates if the value contains only printable ASCII characters
      --semver                validates if the value is a valid semantic version
      --timestamp string      validates if the value matches a timestamp format [rfc3339, datetime, date, time]
      --upper-case            validates if the value contains only upper case unicode letters
      --url                   validates if the value is a valid URL
      --uuid                  validates if the value is a valid UUID
      --value string          the value to validate
  -v, --version               version for valid
```

## Related projects

- [validation-action](https://github.com/tmknom/validation-action): Composite Action for validating input values using `valid`.
