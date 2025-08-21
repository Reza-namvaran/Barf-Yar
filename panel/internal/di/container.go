package di

import (
	"database/sql"
	"log"
	"sync"

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
	supporterRepo repository.SupporterRepository

	// Services
	authService     service.AuthService
	adminService    service.AdminService
	activityService service.ActivityService
	supporterService service.SupporterService

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

func (c *Container) GetSupporterRepository() repository.SupporterRepository {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.supporterRepo == nil {
		c.supporterRepo = repository.NewSupporterRepository(c.db)
	}
	return c.supporterRepo
}  

// --- Services ---
func (c *Container) GetAuthService() service.AuthService {
    c.mu.Lock()
    defer c.mu.Unlock()
    if c.authService == nil {
        repo := repository.NewSessionRepository(c.db)
        c.authService = service.NewAuthService(repo)
    }
    return c.authService
}

func (c *Container) GetAdminService() service.AdminService {
    if c.adminService != nil {
        return c.adminService
    }

     repo := c.GetAdminRepository()

    c.mu.Lock()
    defer c.mu.Unlock()
    if c.adminService == nil {
        c.adminService = service.NewAdminService(repo)
        log.Printf("After admin")
    }
    return c.adminService
}

func (c *Container) GetActivityService() service.ActivityService {
    if c.activityService != nil {
        return c.activityService
    }

    repo := c.GetActivityRepository()

    c.mu.Lock()
    defer c.mu.Unlock()
    if c.activityService == nil {
        c.activityService = service.NewActivityService(repo)
        log.Printf("After activity")
    }
    return c.activityService
}

func (c *Container) GetSupporterService() service.SupporterService {
	if c.supporterService != nil {
		return c.supporterService
	}

	repo := c.GetSupporterRepository()

	c.mu.Lock()
	defer c.mu.Unlock()
	if c.supporterService == nil {
		c.supporterService = service.NewSupporterService(repo)
		log.Printf("After supporter")
	}
	return c.supporterService
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
		c.GetSupporterService(),
		c.GetTemplateService(),
	)
}
