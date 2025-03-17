package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Start timer
		start := time.Now()

		// Process the request
		err := next(c)

		// Calculate latency
		latency := time.Since(start)

		// Gather request details
		status := c.Response().Status
		method := c.Request().Method
		path := c.Request().URL.Path
		clientIP := c.RealIP()
		userAgent := c.Request().UserAgent()

		// Determine color based on status code
		var statusColor string
		switch {
		case status >= http.StatusInternalServerError: // 5xx errors
			statusColor = colorRed
		case status >= http.StatusBadRequest: // 4xx errors
			statusColor = colorYellow
		case status >= http.StatusMultipleChoices: // 3xx redirects
			statusColor = colorCyan
		default: // 2xx success
			statusColor = colorGreen
		}

		// Log the request details with color
		fmt.Printf("%s | %s | %s%s%s | %s%d%s | %s | %s | %s | %v\n",
			colorBlue+clientIP+colorReset,
			colorPurple+method+colorReset,
			colorCyan, path, colorReset,
			statusColor, status, colorReset,
			latency,
			colorYellow+userAgent+colorReset,
			colorWhite+start.Format(time.RFC3339)+colorReset,
			err,
		)

		// Return the error (if any) to propagate it up the chain
		return err
	}
}
