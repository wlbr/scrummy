package scrummy

import (
	"github.com/wlbr/scrummy/gotils"
)

// Importer x import a boards data into the internal object model.
type Importer interface {
	Read(session Session) Session
}

// Exporter export the internal object model to different data format. E.g. PDF, HTML, CSV ...
type Exporter interface {
	Generate(session Session)
}

// A Session represents a list of configurations of graphs to be generated
type Session interface {
	Config() gotils.Config
	Read() Session
	Generate()
}
