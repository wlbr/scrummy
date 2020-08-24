package main

import (
	"github.com/wlbr/scrummy"
	"github.com/wlbr/scrummy/tools"
)

var (
	//Version is a linker injected variable for a git revision info used as version info
	Version = "Unknown build"
	/*Buildstamp is a linker injected variable for a buildtime timestamp used in version info */
	Buildstamp = "unknown build timestamp."

	config *tools.CommonConfig
)

func main() {

	config = new(tools.CommonConfig)
	defer config.CleanUp()
	config.Initialize(Version, Buildstamp)

	scrummy.DrawGraphs("data/scrummastergilde/planning/planning-doneteam5.xlsx")
}
