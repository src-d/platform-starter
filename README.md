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
* Eslint config
* CSSComb, Editorconfig, Eslint (with plugins) and Prettier binaries
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

### Tooling guide

* We use **yarn** as package manager. You can use NPM to install global binaries, but installations of specific projects **must** be done with yarn.
* All other configurations should be installed by this command.
* All JavaScript code should be formatted using prettier and adhere to the rules defined in the ESLint configuration (airbnb-base with customizations).
* All CSS code should be formatted using CSSComb with the provided configuration.

### Editor plugins

**Vim/NeoVim**

* CSSComb: [vim-csscomb](https://github.com/csscomb/vim-csscomb)
* Editorconfig: [editorconfig-vim](https://github.com/editorconfig/editorconfig-vim)
* CSSComb and Prettier: [neoformat](https://github.com/sbdchd/neoformat)
* Prettier: [several options](https://github.com/prettier/prettier/tree/master/editors/vim)

**Visual Studio Code**

* CSSComb: [vscode-csscomb](https://marketplace.visualstudio.com/items?itemName=mrmlnc.vscode-csscomb)
* Editorconfig: [EditorConfig for VS Code](https://marketplace.visualstudio.com/items?itemName=mrmlnc.vscode-csscomb)
* Prettier + ESLint: [Prettier ESLint](https://github.com/prettier/prettier/tree/master/editors/vim)

You're welcome to add more editors.

### Contributing

If you changed any of the files in `config` and `hooks` remember to run `make` before committing to update the bindata.
