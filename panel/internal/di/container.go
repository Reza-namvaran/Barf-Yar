package di

import (
	"database/sql"
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
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.auth == nil {
		c.mu.RUnlock()
		c.mu.Lock()
		if c.auth == nil {
			c.auth = auth.NewService()
		}
		c.mu.Unlock()
		c.mu.RLock()
	}
	return c.auth
}

func (c *Container) GetAdminService() storage.AdminService {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.admin == nil {
		c.mu.RUnlock()
		c.mu.Lock()
		if c.admin == nil {
			c.admin = storage.NewAdminService(c.db)
		}
		c.mu.Unlock()
		c.mu.RLock()
	}
	return c.admin
}

func (c *Container) GetTemplateService() *templates.TemplateService {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.template == nil {
		c.mu.RUnlock()
		c.mu.Lock()
		if c.template == nil {
			var err error
			c.template, err = templates.NewTemplateService()
			if err != nil {
				panic(err)
			}
		}
		c.mu.Unlock()
		c.mu.RLock()
	}
	return c.template
}

func (c *Container) GetHandlers() *handlers.Handlers {
	return handlers.NewHandlers(c.GetAuthService(), c.GetAdminService(), c.GetTemplateService())
}
