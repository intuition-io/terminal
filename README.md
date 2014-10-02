Terminal
========

[![Build Status](https://api.shippable.com/projects/53f671c7fc1e7b4903f36573/badge/master)](https://www.shippable.com/projects/53f671c7fc1e7b4903f36573)
[![Coverage Status](https://img.shields.io/coveralls/intuition-io/terminal.svg)](https://coveralls.io/r/intuition-io/terminal?branch=develop)
[![Gitter chat](https://badges.gitter.im/intuition-io.png)](https://gitter.im/intuition-io)

> [Intuition][telepathy] HTTP API client


Getting started
---------------

Compile from source

```console
go get github.com/intuition-io/terminal
```

Or download a single binary

```console
PLATFORM="linux/amd64"
# See gobuild.io for other supported platforms
wget http://gobuild.io/github.com/intuition-io/terminal/master/${PLATFORM} -O output.zip
unzip output.zip

./terminal --help
./terminal <cmd> --help
```


Documentation
-------------

Check it out on [gowalker][walker], [godoc][doc], or browse it locally:

```console
$ make doc
$ $BROWSER http://localhost:6060/pkg/github.com/intuition-io/terminal/
```


Contributing
------------

> Fork, implement, add tests, pull request, get my everlasting thanks and a
> respectable place here [=)][jondotquote]

```console
make devinstall
make test TESTARGS=-v
```


Conventions
-----------

*Terminal* follows some wide-accepted guidelines

* [Semantic Versioning known as SemVer][semver]
* [Git commit messages][commit]


Authors
-------

| Selfie               | Name            | Twitter                     |
|----------------------|-----------------|-----------------------------|
| <img src="https://avatars.githubusercontent.com/u/1517057" alt="text" width="40px"/> | Xavier Bruhiere | [@XavierBruhiere][xbtwitter] |


Licence
-------

Copyright 2014 Xavier Bruhiere.

*Terminal* is available under the MIT Licence.


---------------------------------------------------------------

<p align="center">
  <img src="https://raw.github.com/hivetech/hivetech.github.io/master/images/pilotgopher.jpg" alt="gopher" width="200px"/>
</p>


[telepathy]: https://github.com/intuition-io/telepathy
[semver]: http://semver.org
[commit]: https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit#
[xbruhiere]: https://avatars.githubusercontent.com/u/1517057
[xbtwitter]: https://twitter.com/XavierBruhiere
[jondotquote]: https://github.com/jondot/groundcontrol
[walker]: http://gowalker.org/github.com/intuition-io/terminal
[doc]: http://godoc.org/github.com/intuition-io/terminal
