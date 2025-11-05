package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type ReverseProxy struct {
	UserServiceURL     *url.URL
	BookingServiceURL  *url.URL
	TrainServiceURL    *url.URL
	ScheduleServiceURL *url.URL
}

func NewReverseProxy(userHost string, userPort int, bookingHost string, bookingPort int, trainHost string, trainPort int, scheduleHost string, schedulePort int) *ReverseProxy {
	userURL, _ := url.Parse(fmt.Sprintf("http://%s:%d", userHost, userPort))
	bookingURL, _ := url.Parse(fmt.Sprintf("http://%s:%d", bookingHost, bookingPort))
	trainURL, _ := url.Parse(fmt.Sprintf("http://%s:%d", trainHost, trainPort))
	scheduleURL, _ := url.Parse(fmt.Sprintf("http://%s:%d", scheduleHost, schedulePort))

	return &ReverseProxy{
		UserServiceURL:     userURL,
		BookingServiceURL:  bookingURL,
		TrainServiceURL:    trainURL,
		ScheduleServiceURL: scheduleURL,
	}
}

func (rp *ReverseProxy) ProxyToUserService() gin.HandlerFunc {
	return rp.createProxy(rp.UserServiceURL, "/api/auth", "/auth")
}

func (rp *ReverseProxy) ProxyToBookingService() gin.HandlerFunc {
	return rp.createProxy(rp.BookingServiceURL, "/api/bookings", "/bookings")
}

func (rp *ReverseProxy) ProxyToTrainService() gin.HandlerFunc {
	return rp.createProxy(rp.TrainServiceURL, "/api/trains", "/trains")
}

func (rp *ReverseProxy) ProxyToScheduleService() gin.HandlerFunc {
	return rp.createProxy(rp.ScheduleServiceURL, "/api/schedules", "/schedules")
}

func (rp *ReverseProxy) createProxy(target *url.URL, stripPrefix, targetPrefix string) gin.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)

		req.URL.Path = strings.Replace(req.URL.Path, stripPrefix, targetPrefix, 1)
		req.Host = target.Host

		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Header.Set("X-Forwarded-Proto", "http")
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(fmt.Sprintf("Service unavailable: %v", err)))
	}

	return gin.WrapH(proxy)
}

func (rp *ReverseProxy) HealthCheckProxy(serviceName string, serviceURL *url.URL) gin.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(serviceURL)

	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = serviceURL.Scheme
		req.URL.Host = serviceURL.Host
		req.URL.Path = "/health"
		req.Host = serviceURL.Host
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(fmt.Sprintf(`{"service": "%s", "status": "unavailable", "error": "%v"}`, serviceName, err)))
	}

	return gin.WrapH(proxy)
}
