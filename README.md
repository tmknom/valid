# valid

A command-line tool for validating input values against specified rules.

## Description

`valid` is a lightweight and flexible tool that validates input values based on character sets, ranges, formats, and custom rules.
It helps ensure data correctness and integrity.

Supported validation rules include:

- **Character sets**: Restricts input to specific character sets, such as ASCII, digits, or alphanumeric characters.
- **Ranges**: Ensures values fall within specified numerical, length, or date ranges.
- **Formats**: Checks if input matches formats, such as timestamps, URLs, or semantic versions.
- **Custom rules**: Matches user-defined rules like enumerations or regular expressions.

`valid` simplifies validation in scripts, workflows, and applications.

## Usage

Use the `--value` flag to specify the input to validate. For example:

### Validation success example

If the input meets all specified rules, the command **exits successfully with no output**:

```shell
valid --value 1234567890 --digit --min-length 10
```

In this case:

- The exit code is `0` (indicating success).
- No error message or output is displayed.

To confirm the exit code, run:

```shell
$ valid --value 1234567890 --digit --min-length 10
$ echo $?
0
```

### Validation failure example

If the input does not meet the specified rules, the command fails with a non-zero exit code and an error message.
For example, the following command fails because `"invalid"` does not meet the specified rules (`--digit` and `--min-length 10`):

```shell
valid --value invalid --digit --min-length 10
```

This command fails with a non-zero exit code, displaying an error message:

```shell
Error: Validation error: The specified value "invalid" is invalid. Issues: the length must be no less than 10, must contain digits only.
```

## Command Help Output

The following is the output of `valid --help`, showing all available flags and usage details.

```shell
Validates that input values meet specified rules

Usage:
  valid [flags]

Flags:
      --alpha                 validates that the value contains only English letters (a-zA-Z)
      --alphanumeric          validates that the value contains only English letters and digits (a-zA-Z0-9)
      --ascii                 validates that the value contains only ASCII characters
      --base64                validates that the value is a valid Base64 string
      --digit                 validates that the value contains only digits (0-9)
      --domain                validates that the value is a valid domain
      --email                 validates that the value is a valid email address
      --enum string           validates that the value matches one of the specified enumerations (comma-separated list)
      --exact-length string   validates that the length of value is exactly the specified number
      --float                 validates that the value is a floating-point number
      --format string         specifies the output format (default, github-actions) (default "default")
  -h, --help                  help for valid
      --int                   validates that the value is an integer
      --json                  validates that the value is a valid JSON string
      --lower-case            validates that the value contains only lowercase Unicode letters
      --mask-value            masks the value in error messages to protect sensitive data
      --max string            validates that the value is less than or equal to the specified maximum
      --max-length string     validates that the length of value is less than or equal to the specified maximum
      --min string            validates that the value is greater than or equal to the specified minimum
      --min-length string     validates that the length of value is greater than or equal to the specified minimum
      --not-empty             validates that the value is not empty
      --pattern string        validates that the value matches the specified regular expression
      --printable-ascii       validates that the value contains only printable ASCII characters
      --semver                validates that the value is a valid semantic version
      --timestamp string      validates that the value matches the timestamp format specified in the timestamp input (rfc3339, datetime, date, or time)
      --upper-case            validates that the value contains only uppercase Unicode letters
      --url                   validates that the value is a valid URL
      --uuid                  validates that the value is a valid UUID
      --value string          the value to validate against the specified rules
      --value-name string     the name of the value to include in error messages
  -v, --version               version for valid
```

## FAQ

### What happens if validation succeeds or fails?

If validation **succeeds**, the command exits with a status code of `0` and produces no output.
If validation **fails**, the command exits with a non-zero status code (`1`) and displays an error message.
This behavior follows standard CLI conventions.

### How can I use `valid` in a script?

Since `valid` returns `0` on success and a non-zero exit code on failure,
you can use it in shell scripts as follows:

```shell
if valid --value 12345 --digit --min-length 10; then
  echo "Validation passed"
else
  echo "Validation failed"
fi
```

This is useful in automation, CI/CD pipelines, or when enforcing validation rules in deployment scripts.

### What happens if I specify multiple validation rules?

All specified rules are evaluated independently, and the value must satisfy **all** of them to be considered valid.
For example, if you specify `--digit` and `--min-length 10`, the value must consist only of digits (0-9) and have a length of at least 10 characters.

Validation rules do not have a fixed evaluation order.
If multiple rules are not met, all relevant errors will be included in the final error message, but their order is not guaranteed.

### Can I customize error messages to make debugging easier?

Yes, use the `--value-name` flag to specify a custom name for the validated value.

Standard error message:

```shell
Error: Validation error: The specified value "example" is invalid. Issues: must contain digits only.
```

Customized error message:

```shell
Error: Validation error: The specified account-id "example" is invalid. Issues: must contain digits only.
```

This replaces "value" in error messages, making it easier to identify which input caused the issue when validating multiple values in the same script or workflow.

### Can I hide sensitive values in error messages?

Yes, use the `--mask-value` to obscure the input as `***`.

Standard error message:

```shell
Error: Validation error: The specified value "example" is invalid. Issues: must contain English letters only.
```

Customized error message:

```shell
Error: Validation error: The specified value "***" is invalid. Issues: must contain English letters only.
```

This is useful in CI/CD environments, where it prevents sensitive data, such as tokens or API keys, from being exposed in logs.

### Can I define a custom error message?

No, you cannot specify a fully custom error message.
The error message format is predefined.
However, you can modify certain parts, such as changing the value name with `--value-name` or masking sensitive data with `--mask-value`.
While these options allow for some customization, the overall structure and wording of the error message remain unchanged.

### What happens if no validation rules are specified?

If no validation rules are provided, the command completes successfully without validation.
No errors or warnings are generated.

This behavior is **intentional** for now but may change in the future.
Be mindful of this to avoid unintentionally skipping validation.

## Related projects

- [validation-action](https://github.com/tmknom/validation-action): Composite Action for validating input values against specified rules.
    - This action internally uses `valid` to perform input validation.
