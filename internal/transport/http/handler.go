package v1

import (
	"net/http"

	v1 "github.com/YnMann/go-sister-design-site/internal/transport/http/v1"

	"github.com/YnMann/go-sister-design-site/internal/config"
	"github.com/YnMann/go-sister-design-site/internal/service"
	"github.com/YnMann/go-sister-design-site/internal/status"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"

	"github.com/gin-contrib/cors"
)

type Handler struct {
	service *service.Service
	cfg     *config.Config
	health  *status.Status
}

func New(s *service.Service, config *config.Config) *Handler {
	return &Handler{
		service: s,
		cfg:     config,
		health:  status.NewStatus(),
	}
}

// init gin handler
func (h *Handler) Init() *gin.Engine {
	if h.cfg.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Recovery())

	//for dev
	if !h.cfg.IsProd {
		router.Use(gin.Logger()) //only for development purposes! do not use in prod!
	}

	secureMiddleware := secure.New(secure.Options{
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
	})

	// from https://github.com/unrolled/secure
	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			err := secureMiddleware.Process(c.Writer, c.Request)
			// If there was an error, do not continue.
			if err != nil {
				c.Abort()
				return
			}

			// Avoid header rewrite if response is a redirection.
			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()

	// Set up CORS middleware options
	corsConfig := cors.Config{
		AllowAllOrigins: true, // Allow all origins
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
	}

	router.Use(
		cors.New(corsConfig),
		secureFunc,
	)

	// Ping
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Static("/static", h.cfg.Environment.PathToPublicFiles)

	router.LoadHTMLGlob("templates/*")

	h.initAPI(router)

	return router
}

// initAPI initializes API routes
func (h *Handler) initAPI(router *gin.Engine) {
	v1Handler := v1.New(h.service, h.cfg)

	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	api := router.Group(h.cfg.Environment.RootRouter)
	{
		v1Handler.Init(api)
	}
}
