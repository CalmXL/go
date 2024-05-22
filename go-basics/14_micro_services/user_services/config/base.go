package config

import "flag"

const (
	DBName = "micro_services"
	// Port   = "50001"
)

var (
	IP   = flag.String("ip", "0.0.0.0", "IP Address")
	Port = flag.String("port", "50001", "IP Port")
)

const (
	REGPhoneNumber = `^1[3-9]\d{9}$`
)
