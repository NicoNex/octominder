# octominder
Octominder reminds you about stuff.

## Configuration
You can configure octominder with its config file.
The default path for the config file is `$HOME/.config/octominder.toml`, but you can specify alternative config files by using the `-c` flag argument.

You can run octominder in userspace at startup and it will use your native notifications.
Octominder is compatible with all major OSes and desktop environments.

### Example configuration
```toml
[reminder.water]
message = "Don't forget to drink water!"
repeat = "20m"

[reminder.break]
message = "Time for a break!"
repeat = "30s"

[reminder.eat]
message = "Eat something!"
repeat = "5h30m"
```
