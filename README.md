# beep

Beep is designed to create a new Go module or Go package. By default directory name is bebeep and default module name is bebeep too. To change default behavior use ldflags.

## Install

Use default values:
```sh
go install github.com/o7q2ab/beep@latest
```

Change both directory and module names:
```sh
go install \
    -ldflags="-X \
        'github.com/o7q2ab/beep/config.Dir=notbebeep' \
        'github.com/o7q2ab/beep/config.Mod=notnotbebeep' \
    " \
    github.com/o7q2ab/beep@latest
```

## Usage

Create Go module:

```sh
beep
```

Create Go pacakge:

```sh
beep pkg
```
