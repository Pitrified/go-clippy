# Clippy

Minimal tool to manage clipboard registers.

## Usage

```
git clone https://github.com/Pitrified/go-clippy.git
cd go-clippy
go run cmd/main/main.go
```

Launch the app and start copying text:
a preview of the content stored in the registers will be shown on each button.
Click on a button to copy the content of that register back to the clipboard.

## Build

On Windows,
to build an executable that does not open a terminal window,
use:

```
go build -ldflags="-H=windowsgui" main.go
```

## TODO

- [x] If a copy of the incoming content is already in the registers,
      only move that to the top.
- [ ] Dynamic number of registers.
- [ ] After clicking a button, move the window to background.
- [ ] Button to clear registers.
- [ ] Initialize the registers from a file.
- [ ] Save the registers to a file.
- [ ] Deactivate empty registers.
