package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gitlab.com/Wuriyanto/go-codebase/pkg/logutil"
)

// Logger function for writing all request log into console
func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		start := time.Now()
		req := c.Request()
		res := c.Response()

		remoteAddr := req.RemoteAddr
		if ip := req.Header.Get(echo.HeaderXRealIP); ip != "" {
			remoteAddr = ip
		} else if ip = req.Header.Get(echo.HeaderXForwardedFor); ip != "" {
			remoteAddr = ip
		} else {
			remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
		}

		var bodyBytes []byte
		if req.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(req.Body)
		}
		// Restore the io.ReadCloser to its original state
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// Use the content
		bodyString := string(bodyBytes)

		entry := logutil.Logger.WithFields(logrus.Fields{
			"request":   req.RequestURI,
			"method":    req.Method,
			"remote":    remoteAddr,
			"requestId": req.Header.Get("X-Request-Id"),
			"payload":   bodyString,
		})

		entry.Info("start handling request")

		if err := next(c); err != nil {
			c.Error(err)
		}

		latency := time.Since(start)

		entry.WithFields(logrus.Fields{
			"size":       res.Size,
			"status":     res.Status,
			"textStatus": http.StatusText(res.Status),
			"took":       latency,
			fmt.Sprintf("#%s.latency", logutil.AppName): latency.Nanoseconds(),
		}).Info("completed handling request")

		return nil
	}
}
