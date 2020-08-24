package tools

import (
	"expvar"
	"runtime"
)

//CheckErr is a convenience function makes error handling dangerously simple.
func CheckErr(err error) {
	if err != nil {
		LogDebug("%s", err)
	}
}

// Minf64 returns the minimum of a slice of float64
func Minf64(v []float64) float64 {
	m := v[0]
	for _, e := range v {
		if e < m {
			m = e
		}
	}
	return m
}

// Maxf64 returns the maximum of a slice of float64
func Maxf64(v []float64) float64 {
	m := v[0]
	for _, e := range v {
		if e > m {
			m = e
		}
	}
	return m
}

var (
	buildversionExpvar, buildstampExpvar, compilerversionExpvar     *expvar.String
	gorootExpvar, goosExpvar, goarchExpvar, logFile, activeloglevel *expvar.String
	numcpuExpvar, numgoroutineExpvar                                *expvar.Int
)

func addAdditionalExpVars(config *CommonConfig) {
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
	buildstampExpvar.Set(config.BuildTimeStamp.String())
	compilerversionExpvar.Set(runtime.Version())
	gorootExpvar.Set(runtime.GOROOT())
	goosExpvar.Set(runtime.GOOS)
	goarchExpvar.Set(runtime.GOARCH)
	numcpuExpvar.Set(int64(runtime.NumCPU()))
	numgoroutineExpvar.Add(int64(runtime.NumGoroutine()))
	activeloglevel.Set(config.ActiveLogLevel.String())
	logFile.Set(config.LogFileName)
}
