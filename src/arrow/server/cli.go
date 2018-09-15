package server

import "flag"

type Options struct {
	httpAddr   string
	httpsAddr  string
	tunnelAddr string
	domain     string
	tlsCrt     string
	tlsKey     string
	logTo      string
	logLevel   string
}

func parseArgs() *Options {

	httpAddr := flag.String("httpAddr", ":80", "")
	httpsAddr := flag.String("httpsAddr", ":443", "")
	tunnelAddr := flag.String("tunnelAddr", "4443", "")
	domain := flag.String("domain", "ngrok.com", "")
	tlsCrt := flag.String("tlsCrt", "", "")
	tlsKey := flag.String("tlsKey", "", "")
	logTo := flag.String("log", "stdout", "")
	logLevel := flag.String("log-level", "DEBUG", "")
	flag.Parse()

	return &Options{
		httpAddr:   *httpAddr,
		httpsAddr:  *httpsAddr,
		tunnelAddr: *tunnelAddr,
		domain:     *domain,
		tlsCrt:     *tlsCrt,
		tlsKey:     *tlsKey,
		logTo:      *logTo,
		logLevel:   *logLevel,
	}
}
