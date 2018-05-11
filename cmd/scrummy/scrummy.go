package main

import (
	"expvar"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/wlbr/scrummy"
	"github.com/wlbr/scrummy/gotils"
)

var (
	//Version is a linker injected variable for a git revision info used as version info
	Version = "Unknown build"
	/*Buildstamp is a linker injected variable for a buildtime timestamp used in version info */
	Buildstamp = "unknown build timestamp."

	buildversionExpvar, gitversionExpvar, buildstampExpvar, compilerversionExpvar *expvar.String
	gorootExpvar, goosExpvar, goarchExpvar, logFile, activeloglevel               *expvar.String
	numcpuExpvar, numgoroutineExpvar                                              *expvar.Int

	config gotils.Config
)

func addAdditionalExpVars(config gotils.Config) {
	buildversionExpvar = expvar.NewString("buildversion")
	buildstampExpvar = expvar.NewString("buildtimestamp")
	compilerversionExpvar = expvar.NewString("compilerversion")
	gorootExpvar = expvar.NewString("GOROOT")
	goosExpvar = expvar.NewString("GOOS")
	goarchExpvar = expvar.NewString("GOARCH")
	numcpuExpvar = expvar.NewInt("NumCPU")
	numgoroutineExpvar = expvar.NewInt("NumGoroutine")
	activeloglevel = expvar.NewString("ActiveLogLevel")
	logFile = expvar.NewString("LogFile")

	buildversionExpvar.Set(config.GitVersion)
	buildstampExpvar.Set(fmt.Sprintf("%s", config.BuildTimeStamp))
	compilerversionExpvar.Set(runtime.Version())
	gorootExpvar.Set(runtime.GOROOT())
	goosExpvar.Set(runtime.GOOS)
	goarchExpvar.Set(runtime.GOARCH)
	numcpuExpvar.Set(int64(runtime.NumCPU()))
	numgoroutineExpvar.Add(int64(runtime.NumGoroutine()))
	activeloglevel.Set(config.ActiveLogLevel.String())
	logFile.Set(config.LogFileName)
}

func init() {
	cwd, _ := os.Getwd()
	var logfilename string
	config = gotils.Config{}

	var loglevel = "unknown"
	var err error
	flag.StringVar(&loglevel, "LogLevel", "All", "Determines logging verbosity. [Off|Info|Debug|Warnings|Error|Fatal|All].")
	flag.StringVar(&logfilename, "LogFile", "", "Sets the name of the logfile. Uses STDOUT if empty.")
	flag.Parse()
	config.ActiveLogLevel, err = gotils.LogLevelString(loglevel)
	if err != nil {
		config.ActiveLogLevel = gotils.All
	}
	config.Logger = gotils.NewLogger(logfilename, config.ActiveLogLevel)
	config.Logger.SetConvenienceLogger()
	gotils.LogInfo("Current working directory is '%s'.", cwd)
	if err != nil {
		config.Logger.Warn("Error in config, Loglevel '%s' not existing in tools/loglevel.go. Setting LogLevel to 'All'", loglevel)
	}

	btime, err := time.Parse("2006-01-02_15:04:05_MST", Buildstamp)
	if err != nil {
		config.BuildTimeStamp = time.Now()
	} else {
		config.BuildTimeStamp = btime
	}
	config.GitVersion = Version
	config.Logger.Info("Version: %s of %s \n", config.GitVersion, config.BuildTimeStamp)

	addAdditionalExpVars(config)
}

func main() {
	flag.Parse()

	scrummy.DrawGraphs("data/scrummastergilde/planning/planning-doneteam5.xlsx")
}
