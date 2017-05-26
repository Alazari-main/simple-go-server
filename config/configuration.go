package config

import (
	"net/http"
)

var (
	Instance = &configuration {}
)

type configuration struct {
	WorkingDirectory	http.Dir
	Port				int
}