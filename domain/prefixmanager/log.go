package prefixmanager

import (
	"github.com/Nexellia-Network/nexelliad/infrastructure/logger"
	"github.com/Nexellia-Network/nexelliad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
