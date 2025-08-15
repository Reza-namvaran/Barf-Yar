package di

import (
	"database/sql"
	"log"
	"sync"

	"github.com/Reza-namvaran/Barf-Yar/panel/internal/auth"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/handlers"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/repository"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/service"
	"github.com/Reza-namvaran/Barf-Yar/panel/internal/templates"
)

// Container holds all dependencies
type Container struct {
	db *sql.DB

	// Repositories
	adminRepo    repository.AdminRepository
	activityRepo repository.ActivityRepository

	// Services
	authService     auth.Service
	adminService    service.AdminService
	activityService service.ActivityService

	// Others
	templateService *templates.TemplateService

	mu sync.RWMutex
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

// --- Repositories ---
func (c *Container) GetAdminRepository() repository.AdminRepository {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.adminRepo == nil {
		c.adminRepo = repository.NewAdminRepository(c.db)
	}
	return c.adminRepo
}

func (c *Container) GetActivityRepository() repository.ActivityRepository {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.activityRepo == nil {
		c.activityRepo = repository.NewActivityRepository(c.db)
	}
	return c.activityRepo
}

// --- Services ---
func (c *Container) GetAuthService() auth.Service {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.authService == nil {
		c.authService = auth.NewService()
	}

	log.Printf("After auth")
	return c.authService
}
func (c *Container) GetAdminService() service.AdminService {
    c.mu.RLock()
    if c.adminService != nil {
        defer c.mu.RUnlock()
        return c.adminService
    }
    c.mu.RUnlock()

    // build service outside lock to avoid deadlock
    repo := c.GetAdminRepository()
    newService := service.NewAdminService(repo)

    c.mu.Lock()
    defer c.mu.Unlock()
    if c.adminService == nil { // double-check in case of race
        c.adminService = newService
    }

	log.Printf("After admin")
    return c.adminService
}

func (c *Container) GetActivityService() service.ActivityService {
    c.mu.RLock()
    if c.activityService != nil {
        defer c.mu.RUnlock()
        return c.activityService
    }
    c.mu.RUnlock()

    // build service outside lock to avoid deadlock
    repo := c.GetActivityRepository()
    newService := service.NewActivityService(repo)

    c.mu.Lock()
    defer c.mu.Unlock()
    if c.activityService == nil { // double-check in case of race
        c.activityService = newService
    }

	log.Printf("After activity")
    return c.activityService
}

// --- Template ---
func (c *Container) GetTemplateService() *templates.TemplateService {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.templateService == nil {
		var err error
		c.templateService, err = templates.NewTemplateService()
		if err != nil {
			panic(err)
		}
	}

	log.Printf("After templates")

	return c.templateService
}

// --- Handlers ---
func (c *Container) GetHandlers() *handlers.Handlers {
	log.Println("Resolving handlers...")
	return handlers.NewHandlers(
		c.GetAuthService(),
		c.GetAdminService(),
		c.GetActivityService(),
		c.GetTemplateService(),
	)
}
