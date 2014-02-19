Ganglia Sink
------------

Assuming you've installed nolo into `/opt/nolo`, you should be able to
run the nolo `load` plugin and pass the input into gmetric by running:

    /opt/nolo/bin/nolo-ganglia /opt/nolo/plugins/load

You can run this plugin regularly by entering adding it to your crontab:

    * * * * * /opt/nolo/bin/nolo-ganglia /opt/nolo/plugins/load

Please note: the ganglia adapter does not yet support any metadata
flags, so it will only pass along the identifier and value.
