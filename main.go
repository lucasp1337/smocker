package main

import (
	"os"

	"github.com/lucasp1337/smocker/server"
	"github.com/lucasp1337/smocker/server/config"
	"github.com/namsral/flag"
	log "github.com/sirupsen/logrus"
)

var appName, buildVersion, buildCommit, buildDate string // nolint

func parseConfig() (c config.Config) {
	c.Build = config.Build{
		AppName:      appName,
		BuildVersion: buildVersion,
		BuildCommit:  buildCommit,
		BuildDate:    buildDate,
	}

	// Use a prefix for environment variables
	flag.CommandLine = flag.NewFlagSetWithEnvPrefix(os.Args[0], "SMOCKER", flag.ExitOnError)

	flag.StringVar(&c.LogLevel, "log-level", "info", "Available levels: panic, fatal, error, warning, info, debug, trace")
	flag.StringVar(&c.ConfigBasePath, "config-base-path", "/", "Base path applied to Smocker UI")
	flag.IntVar(&c.ConfigListenPort, "config-listen-port", 8081, "Listening port of Smocker administration server")
	flag.IntVar(&c.MockServerListenPort, "mock-server-listen-port", 8080, "Listening port of Smocker mock server")
	flag.StringVar(&c.StaticFiles, "static-files", "client", "Location of the static files to serve (index.html, etc.)")
	flag.IntVar(&c.HistoryMaxRetention, "history-retention", 0, "Maximum number of calls to keep in the history per session (0 = no limit)")
	flag.StringVar(&c.PersistenceDirectory, "persistence-directory", "", "If defined, the directory where the sessions will be synchronized")
	flag.BoolVar(&c.TLSEnable, "tls-enable", false, "Enable TLS using the provided certificate")
	flag.StringVar(&c.TLSCertFile, "tls-cert-file", "/etc/smocker/tls/certs/cert.pem", "Path to TLS certificate file ")
	flag.StringVar(&c.TLSKeyFile, "tls-private-key-file", "/etc/smocker/tls/private/key.pem", "Path to TLS key file")
	flag.Parse()
	return
}

func setupLogger(logLevel string) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	})

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.WithError(err).WithField("log-level", level).Warn("Invalid log level, fallback to info")
		level = log.InfoLevel
	}
	log.WithField("log-level", level).Info("Setting log level")
	log.SetLevel(level)
}

func main() {
	c := parseConfig()
	setupLogger(c.LogLevel)
	server.Serve(c)
}
