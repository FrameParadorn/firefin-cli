package template

var Controller string = `
package {{module_name_lower}}

import (
	_ "github.com/FrameParadorn/foodfin-backend/infrastructure/validate"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service *Service
}

func NewController(app *fiber.App) *Controller {
	controller := new(Controller)

	// Router
	// Ex.
	// app.Get("/users", c.FindAll)

	return controller
}

func (c *Controller) FindAll(ctx *fiber.Ctx) error {

	// Handler logic here

	// Ex.
	// var request UserCreateRequestDto
	// if err := validate.ValidateWithBodyParser(ctx, &request); err != nil {
	// 	return err
	// }
	// return c.service.Create(&request)
}
`

var Dto string = `
package {{module_name_lower}}

type {{module_name}}CreateRequestDto struct {
}
`

var Etity string = `
package {{module_name_lower}} 

type {{module_name}}Entity struct {
}

func (e *{{module_name}}Entity) TableName() string {
	return "{{module_name_lower}}"
}
`

var Module string = `
package {{module_name_lower}}

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type {{module_name}}Module struct {
	controller *Controller
	service    *Service
}

func NewModule(app *fiber.App, db *gorm.DB) *{{module_name}}Module {
	return &{{module_name}}Module{
		controller: NewController(app),
		service:    NewService(db),
	}
}
`

var Repository string = `
package {{module_name_lower}}

import "gorm.io/gorm"

type Repositry struct {
	db *gorm.DB
}

func NewReposiry(db *gorm.DB) *Repositry {
	return &Repositry{
		db: db,
	}
}
`

var Service string = `
package {{module_name_lower}}

import (
	"gorm.io/gorm"
)

type Service struct {
	db         *gorm.DB
	repository *Repositry
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db:         db,
		repository: NewReposiry(db),
	}
}

func (s *Service) Create(request *{{module_name}}CreateRequestDto) error {
	{{module_name_lower}} := &{{module_name}}Entity{}
	return s.db.Create({{module_name_lower}}).Error
}

`
