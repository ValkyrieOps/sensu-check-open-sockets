[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/ValkyrieOps/check-open-sockets)
![Go Test](https://github.com/ValkyrieOps/check-open-sockets/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/ValkyrieOps/check-open-sockets/workflows/goreleaser/badge.svg)

# check-open-sockets

## Table of Contents
- [Overview](#overview)
- [Files](#files)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definition](#check-definition)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview

The check-open-sockets is a [Sensu Check][6] that determines total open system sockets using ss

## Files

bin/sensu-check-open-sockets

## Usage examples

```
The Sensu Go Open Sockets plugin

Usage:
  sensu-check-open-sockets [flags]
  sensu-check-open-sockets [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -c, --crit int   Critical threshold - count of open sockets required for critical state (default 20)
  -h, --help       help for sensu-check-open-sockets
  -w, --warn int   Warning threshold - count of open sockets required for warning state (default 10)

Use "sensu-check-open-sockets [command] --help" for more information about a command.

```
## Configuration

### Asset registration

[Sensu Assets][10] are the best way to make use of this plugin. If you're not using an asset, please
consider doing so! If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add the asset:

```
sensuctl asset add ValkyrieOps/sensu-check-open-sockets
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index][https://bonsai.sensu.io/assets/ValkyrieOps/check-open-sockets].

### Check definition

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: check-open-sockets
  namespace: default
spec:
  command: sensu-check-open-sockets -w 500 -c 800
  subscriptions:
  - system
  runtime_assets:
  - ValkyrieOps/check-open-sockets
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an Asset. If you would
like to compile and install the plugin from source or contribute to it, download the latest version
or create an executable script from this source.

From the local path of the check-open-sockets repository:

```
go build
```

## Additional notes

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu-community/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/sensu-community/check-plugin-template/blob/master/.github/workflows/release.yml
[5]: https://github.com/sensu-community/check-plugin-template/actions
[6]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[7]: https://github.com/sensu-community/check-plugin-template/blob/master/main.go
[8]: https://bonsai.sensu.io/
[9]: https://github.com/sensu-community/sensu-plugin-tool
[10]: https://docs.sensu.io/sensu-go/latest/reference/assets/
