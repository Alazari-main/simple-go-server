package handler

import (
	"net/http"
)

// IHandler provides common interface for all http handlers.
type IHandler http.Handler