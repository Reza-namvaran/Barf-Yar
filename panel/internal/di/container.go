package di

import (
	"database/sql"
	"log"
	"sync"

	"github.com/Reza-namvaran/Barf-Yar/panel/internal/auth"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/handlers"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/storage"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/templates"
)

// Container holds all dependencies
type Container struct {
	db       *sql.DB
	auth     auth.Service
	admin    storage.AdminService
	activity storage.ActivityService
	template *templates.TemplateService
	mu       sync.RWMutex
}

func NewContainer() *Container {
	return &Container{}
}

// Set & Get the database connection
func (c *Container) SetDB(db *sql.DB) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.db = db
}

func (c *Container) GetDB() *sql.DB {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.db
}

func (c *Container) GetAuthService() auth.Service {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.auth == nil {
		c.auth = auth.NewService()
	}
	return c.auth
}

func (c *Container) GetAdminService() storage.AdminService {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.admin == nil {
		c.admin = storage.NewAdminService(c.db)
	} else {
	}

	log.Println("  Returning admin service...")
	return c.admin
}

func (c *Container) GetActivityService() storage.ActivityService {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.activity == nil {
		db := c.db
		c.activity = storage.NewActivityService(db)
	}
	return c.activity
}

func (c *Container) GetTemplateService() *templates.TemplateService {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.template == nil {
		var err error
		c.template, err = templates.NewTemplateService()
		if err != nil {
			panic(err)
		}
	}
	return c.template
}

func (c *Container) GetHandlers() *handlers.Handlers {
	authService := c.GetAuthService()
	log.Println("Getting admin service...")
	adminService := c.GetAdminService()
	log.Println("Getting activity service...")
	activityService := c.GetActivityService()
	log.Println("Getting template service...")
	templateService := c.GetTemplateService()
	log.Println("Creating handlers...")
	return handlers.NewHandlers(authService, adminService, activityService, templateService)
}
