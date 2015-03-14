Ganglia Sink
------------

Assuming you've installed nolo into `/opt/nolo`, you should be able to run the nolo `load` plugin and pass the input into gmetric by running:

    /opt/nolo/bin/nolo-ganglia /opt/nolo/plugins/load

You can run this plugin regularly by entering adding it to your crontab:

    * * * * * /opt/nolo/bin/nolo-ganglia /opt/nolo/plugins/load

Please note: the ganglia adapter does not yet support any metadata flags, so it will only pass along the identifier and value.

Installation
============

At the moment, this metric sink does not have lovely packaging. To install, you'll need to install git and the go development environment, then run:

    go get github.com/nolo-metrics/nolo-ganglia

This will give you a copy of `nolo-ganglia` in your `$GOPATH`. Copy the binary to its final home:

    mkdir -p /opt/nolo/bin
    cp $GOPATH/bin/nolo-ganglia /opt/nolo/bin/nolo-ganglia

And you're ready to roll!
