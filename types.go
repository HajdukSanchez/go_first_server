package main

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc
