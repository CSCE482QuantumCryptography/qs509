package qs509

var isInitialized bool
var openSSLPath string
var openSSLConfigPath string

func Init(sslPath string, configPath string) {
	openSSLPath = sslPath
	openSSLConfigPath = configPath
	isInitialized = true
}

func checkInit() {
	if !isInitialized {
		panic("qs509 must be initialized before use!")
	}
}
