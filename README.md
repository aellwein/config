
github.com/aellwein/config
==========================

What is it?
-----------

This Go module provides a configuration file management defined by [Cfg](pkg/config/api.go) interface to be used across your Go files without a coupling to the actual implementation.

The ``Cfg`` interface mimics the [viper](https://github.com/spf13/viper) interface and also uses
it as the default implementation.

Implementation can be exchanged on demand without doing modification to your app's code.

Usage
-----

The usage is very simple. Refer to [example.go](examples/example.go) file to get an idea on how to use it.


License
-------

[MIT License](LICENSE)