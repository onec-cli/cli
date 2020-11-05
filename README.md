#cli

## Syntax

$CLI command-name global-options command-options arguments

where:

$CLI - program name (draft = onec-cli)  

|||
| --- | --- |
| command-name | The command to execute. Note that you can use either the full command name or its alias. |
| global-options | A set of global options that may be used for all commands. |
| command-options | A set of options corresponding to the command. |
| arguments | A set of arguments corresponding to the command. |

| global-options | Description |
| --- | --- |
| `-u, --user, --db-user` | [Default: ""] db-user |
| `-p, --password, --db-password, --db-pwd` | [Default: ""] db-password |
...

## Environment Variables
...

## Commands

### config

|  | |
| --- | --- |
| Command aliases | c |
| Command options | |
| Command arguments | |
| `ID` | A unique ID for the new configuration. |
| `show` |This argument should be followed by a configured server ID. The configuration for this server ID will be shown. |
| `delete` | This argument should be followed by a configured server ID. The configuration for this server ID will be deleted. |
| `clear` | Clears all stored configuration. |

#### Example 1

Configure the user and password by passing them in as command options for the `DEV-1` configuration.

```
$CLI c DEV-1 --user=admin --password=password
```

#### Example 2

Delete the `DEV-1` configuration.

```
$CLI c delete DEV-1
```

#### Example 3

Show details of the `DEV-1` configuration

```
$CLI c show DEV-1
```

## platform

```
$CLI **p**latform **i**nstall ?
$CLI **p**latform **c**reate -flags
$CLI **p**latform **r**un -flags
```

## storage

## feature

* $CLI **f**eature **s**tart -flags
* $CLI **f**eature **e**xport -flags

## test

| | |
| --- | --- |
| Command aliases | t |
| Command options | |
| `--id` | Settings ID configured using the config command. If not specified, the default configured settings is used. |
| `--spec` | Path to a file spec. |
| Command arguments | |
| `Framework path` | Path to a ".epf" file with test framework. |

#### Example 1

.......... `DEV-1` .........

```
$CLI t --id=DEV-1 --spec=./test.json ./vanessa-automation.epf
```

## bootstrap

* $CLI **b**ootstrap -flags

## sync

* $CLI **s**ync -flags
