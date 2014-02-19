package main

import (
	"github.com/jbuchbinder/go-gmetric/gmetric"
	"github.com/nolo-metrics/go-nolo"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"net"
)

func main() {
	args := os.Args

	if 1 == len(args) {
		log.Fatal("usage: nolo-ganglia path-to-plugin")
	}

	name := args[1]

	out, err := exec.Command(name).Output()
	if err != nil {
		log.Fatal(err)
	}

	basename := filepath.Base(name)
	plugin := nolo.Parse(basename, string(out))

	gIP := net.IPv4(127, 0, 0, 1)
	gangliaPort := 1234
	host := "127.0.0.1"
	spoofHost := "127.0.0.1:spoof"

	gm := gmetric.Gmetric{
		Host:  host,
		Spoof: spoofHost,
	}
	gm.AddServer(gmetric.GmetricServer{gIP, gangliaPort})

	for _, metric := range plugin.Metrics {
		m_name := metric.Identifier
		m_value := metric.Value
		m_units := ""
		m_type := uint32(gmetric.VALUE_UNSIGNED_INT)
		m_slope := uint32(gmetric.SLOPE_BOTH)
		m_grp := plugin.Identifier
		m_interval := uint32(300)

		go gm.SendMetric(
			m_name,
			m_value,
			m_type,
			m_units,
			m_slope,
			m_interval,
			m_interval,
			m_grp)
	}
}
