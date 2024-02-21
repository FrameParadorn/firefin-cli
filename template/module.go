package template

var Controller string = `
package {{module_name_lower}}

import (
	_ "github.com/FrameParadorn/foodfin-backend/infrastructure/validate"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	module *{{module_name}}Module
}

func NewController(module *{{module_name}}Module) *Controller {
	controller := Controller{
		module: module,
	}

	// Router
	// app.Get("/users", c.FindAll)
	route := controller.module.app.Group("/api/v1/{{module_name_lower}}")

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

	return nil
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
	app        *fiber.App
	db         *gorm.DB
	controller *Controller
	service    *Service
	repositry  *Repositry
}

func NewModule(app *fiber.App, db *gorm.DB) any {
	module := {{module_name}}Module{
		app: app,
		db:  db,
	}
	module.repositry = NewReposiry(&module)
	module.service = NewService(&module)
	module.controller = NewController(&module)

	db.AutoMigrate(&{{module_name}}Entity{})

	return &module
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
	module *{{module_name}}Module
}

func NewService(module *{{module_name}}Module) *Service {
	return &Service{
		module: module,
	}
}

func (s *Service) Create(request *dto.CreateRequestDto) (*{{module_name}}Entity, error) {

	{{module_name_lower}} := {{module_name}}Entity{}

	copier.Copy(&{{module_name_lower}}, request)

	result := s.module.db.Create(&{{module_name_lower}})
	if result.Error != nil {
		return nil, result.Error
	}

	return &{{module_name_lower}}, nil

}

`
