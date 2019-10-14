<p align="center">
  <img width="460" height="auto" src="./fgfbanner.png">
</p>

# Flutter Google Fonts

Download and link Google Fonts in your flutter project automatically

```
Contribution: This project is ⚠️ Work In Progress, if you have and idea to complete it only write an issue.
```

## Install

### From binaries

If you're using MacOS, you can to install using [homebrew](https://brew.sh/)

```
brew tap antoniott15/fgf
brew install fgf
```

As alternative way, you can download the binaries directly from [here](https://github.com/antoniott15/fgf/releases)

### From Source Code

Using `go get`, don't forget to add your $GOBIN add your $PATH to it works.

```
go get -u github.com/antoniott15/fgf
```

## Usage

First, you need to be in your flutter project dir.
Automatically, fgf search your pubspec.yaml and modify it adding the new font description.
Only execute:

```
fgf
```

and search your font by family name.

Also, you can use directly:

```
fgf add <YOUR FONT NAME>
```

That's all.
