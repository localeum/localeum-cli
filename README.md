# Localeum CLI tool

Using the Localeum CLI, you can integrate the Localeum platform into your CI/CD process to download actual translations.

Read more docs on https://docs.localeum.com

## Getting started

## Install

### MacOS
Install using Homebrew:

```
$ brew tap localeum/tools
$ brew install localeum-cli
```

Direct links:
- [Darwin x86_64](https://github.com/localeum/localeum-cli/releases/latest/download/localeum-cli_Darwin_x86_64.tar.gz)
- [Darwin arm64](https://github.com/localeum/localeum-cli/releases/latest/download/localeum-cli_Darwin_arm64.tar.gz)

### Linux

Direct links:
- [Linux arm64](https://github.com/localeum/localeum-cli/releases/latest/download/localeum-cli_Linux_arm64.tar.gz)
- [Linux arm v6](https://github.com/localeum/localeum-cli/releases/latest/download/localeum-cli_Linux_arm_v6.tar.gz)
- [Linux arm v7](https://github.com/localeum/localeum-cli/releases/latest/download/localeum-cli_Linux_arm_v7.tar.gz)
- [Linux i386](https://github.com/localeum/localeum-cli/releases/latest/download/localeum-cli_Linux_i386.tar.gz)
- [Linux x86_64](https://github.com/localeum/localeum-cli/releases/latest/download/localeum-cli_Linux_x86_64.tar.gz)

### FreeBSD

Direct links:
- [FreeBSD arm64](https://github.com/localeum/localeum-cli/releases/latest/download/localeum-cli_freebsd_arm64.tar.gz)
- [FreeBSD arm v6](https://github.com/localeum/localeum-cli/releases/latest/download/localeum-cli_freebsd_arm_v6.tar.gz)
- [FreeBSD arm v7](https://github.com/localeum/localeum-cli/releases/latest/download/localeum-cli_freebsd_arm_v7.tar.gz)

### Windows

Direct links:
- [Windows i386](https://github.com/localeum/localeum-cli/releases/latest/download/localeum-cli_Windows_i386.zip)
- [Windows x86_64](https://github.com/localeum/localeum-cli/releases/latest/download/localeum-cli_Windows_x86_64.zip)

## Configure

### Initialization

Run the command in the root directory of your application.

```
$ localeum-cli init
```

Follow a few steps to create a config file.

```console
Enter the API key for your project: 
```

Your API key can be found on the project settings page (https://app.localeum.com/projects/{YOUR_PROJECT}/settings)

```console
Enter the path to the directory for the localization files:
```

The directory where the localization files will be saved, for example **locales**

```console
Enter the file format for export (json, json_nested, csv, yaml): 
```

If you did everything correctly, you will see the inscription.

```console
Success! You finished initializing your application!
```

The config file localeum.yml should appear at the root of your project.

### Command: pull

To download the latest translations, run the following command.

```
$ localeum-cli pull
```





