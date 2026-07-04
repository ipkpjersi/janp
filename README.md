# janp

Just Another Notifcation Program.

janp displays a desktop notification (via GTK) on a chosen monitor for a
configurable amount of time.

## Usage

```
janp "Your alert message"
```

If no message is passed, a default test alert is shown.

Other flags:

- `-v`, `--version` - print the version and exit.

## Build

Build the application with the following command:

```
go build
```

## Configuration

janp reads a JSON config file named `janp.json`. On startup it looks for the
config in the following locations, in order, and uses the first one it finds:

1. `~/.janp.json`
2. `janp.json` next to the janp executable
3. `/etc/janp.json`
4. `/usr/local/etc/janp.json`
5. `$XDG_CONFIG_HOME/janp/janp.json` (defaults to `~/.config/janp/janp.json`)

### Example config

An example `.janp.json`:

```json
{
	"XineramaHead": 1,
	"NotificationTime": 3,
	"SaveFolder": "janp/temp"
}
```

### Settings

- `XineramaHead` (number) - the monitor (Xinerama head) index to show the
  notification on. Defaults to `0`.
- `NotificationTime` (number) - how many seconds the notification stays on
  screen. Defaults to `3`.
- `SaveFolder` (string) - storage directory. It is currently used only for
  janp's temporary lock files: janp writes a small lock file (`.notifyf0`,
  `.notifyf1`, ...) here for each visible notification so that stacked
  notifications do not overlap, and removes it when the notification is
  dismissed or expires. If the value starts with `/` it is used as an absolute
  path, otherwise it is treated as relative to your home folder. When unset,
  janp uses a `temp` subfolder inside the `janp` folder in your home directory
  (`~/janp/temp`).
