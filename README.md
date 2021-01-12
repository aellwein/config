
github.com/aellwein/config
==========================

What is it?
-----------

This Go module provides a configuration file management defined by stable, versioned [Cfg](v1/config/api.go) interface to be used across your Go files without a coupling to the actual implementation.

The ``Cfg`` interface mimics the [viper](https://github.com/spf13/viper) interface and also uses
it as its default implementation.

Implementation can be exchanged on demand without doing modification to your app's code.

Usage
-----

The usage is very simple. Refer to [example.go](examples/example.go) file to get an idea on how to use it.

Bonus
-----

Config object can be passed via [context.Context](https://golang.org/pkg/context/#Context) interface. 

This is useful, for example, if you work with commandline parsers like [spf13/cobra](https://github.com/spf13/cobra) or [urfave/cli](https://github.com/urfave/cli).


License
-------

[MIT License](LICENSE)