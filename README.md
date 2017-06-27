# platform-starter
Starter and basic configuration for platform frontend projects.

### Requirements

* Go installed
* Node.js installed
* NPM or yarn installed (yarn preferred)

### Install

```bash
go get -v github.com/src-d/platform-starter/...
```

### Usage

Just execute the following command at the root of your project:

```
platform-starter
```

It will install the following things:

* CSSComb config
* Editorconfig
* A hook to make sure everything is formatted

If your project is not a git repository it will initialize one.

If the frontend of your project is not at the root of your project, you can use the `--dir` flag.

```
platform-starter --dir ./client
```

If you want to force the install of all the dependencies with NPM use the `--npm` flag.

```
platform-starter --npm
```

### Contributing

If you changed any of the files in `config` and `hooks` remember to run `make` before committing to update the bindata.
