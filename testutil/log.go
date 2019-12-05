package testutil

import (
	"github.com/hamba/logger"
)

// Logger is a common discard logger for testing.
var Logger = logger.New(logger.DiscardHandler())
