package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/inconshreveable/log15"
	prompt "github.com/segmentio/go-prompt"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "platform-starter"
	app.Usage = "Initialize platform projects with common configuration."
	app.Version = "1.0.0"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "dir",
			Usage: "directory to initialize",
			Value: ".",
		},
		cli.BoolFlag{
			Name:  "npm",
			Usage: "forces the usage of npm for installing dependencies",
		},
	}

	app.Run(os.Args)
}

var requirements = []string{
	"csscomb",
	"editorconfig-tools",
}

type file struct {
	asset    *asset
	dest     []string
	fromRoot bool
}

func (f file) path(root, wd string) string {
	if f.fromRoot {
		return filepath.Join(append([]string{root}, f.dest...)...)
	}

	return filepath.Join(append([]string{wd}, f.dest...)...)
}

var files = []file{
	{mustAsset(configCsscombJson()), mkPath(".csscomb.json"), false},
	{mustAsset(configEditorconfig()), mkPath(".editorconfig"), true},
	{mustAsset(hooksPreCommit()), mkPath(".git", "hooks", "pre-commit"), true},
}

func run(ctx *cli.Context) error {
	log15.Info("Starting platform-starter")

	log15.Info("Installing requirements...")
	for _, r := range requirements {
		ensureInstalled(r, ctx.Bool("npm"))
	}

	root, err := os.Getwd()
	if err != nil {
		log15.Crit("unable to get current working directory", "err", err)
		os.Exit(1)
	}

	if !isDir(".git") {
		log15.Warn("Current directory is not a git repository.")
		log15.Info("Initializing git repository...")
		if err := cmd("git", "init"); err != nil {
			log15.Crit("unable to initialize git repository", "err", err)
			os.Exit(1)
		}
	}

	dir := ctx.String("dir")
	dir, err = filepath.Abs(dir)
	if err != nil {
		log15.Crit("unable to get absolute path for directory", "dir", dir, "err", err)
		os.Exit(1)
	}

	log15.Info("Copying assets...")
	for _, f := range files {
		log15.Info("Copying", "file", filepath.Join(f.dest...))
		if err := copyFile(root, dir, f); err != nil {
			log15.Crit("error copying asset", "file", f.path(root, dir), "err", err)
			os.Exit(1)
		}
	}

	log15.Info("Everything ready!")
	return nil
}

func ensureInstalled(program string, npm bool) {
	_, err := exec.LookPath(program)
	if err != nil {
		log15.Warn(fmt.Sprintf("Looks like `%s` is not installed", program))
		if err := install(program, npm); err != nil {
			log15.Crit(fmt.Sprintf("Unable to install `%s`", program), "err", err)
			os.Exit(1)
		}
	}
}

func install(program string, npmForce bool) error {
	log15.Info(fmt.Sprintf("Installing %s...", program))
	if !npmForce {
		yarn, err := exec.LookPath("yarn")
		if err == nil {
			return cmd(yarn, "global", "add", program)
		}

		log15.Warn("yarn is not installed, resorting to install using npm")
	}

	npm, err := exec.LookPath("npm")
	if err == nil {
		return cmd(npm, "install", "-g", program)
	}

	log15.Crit("npm and yarn are not installed. Aborting process.")
	os.Exit(1)
	return nil
}

func copyFile(root, pwd string, file file) error {
	path := file.path(root, pwd)
	_, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	} else if err == nil {
		log15.Warn(fmt.Sprintf("file %s already exists", filepath.Join(file.dest...)))
		if !prompt.Confirm("Do you want to overwrite it?") {
			log15.Warn("Skipped copy of file.", "file", filepath.Join(file.dest...))
			return nil
		}

		if err := os.Remove(path); err != nil {
			return fmt.Errorf("unable to remove file: %s", err)
		}
	}

	return ioutil.WriteFile(path, file.asset.bytes, file.asset.info.Mode())
}

func cmd(bin string, args ...string) error {
	cmd := exec.Command(bin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	return cmd.Wait()
}

func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fi.IsDir()
}

func mkPath(args ...string) []string {
	return args
}

func mustAsset(asset *asset, err error) *asset {
	if err != nil {
		log15.Crit(err.Error())
		os.Exit(1)
	}

	return asset
}
