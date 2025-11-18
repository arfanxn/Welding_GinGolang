package seeder

import (
	"fmt"
	"time"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	employeeRepository "github.com/arfanxn/welding/internal/module/employee/domain/repository"
	"github.com/arfanxn/welding/internal/module/role/domain/enum"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	roleUserRepository "github.com/arfanxn/welding/internal/module/role_user/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/user/domain/repository"
	factoryGo "github.com/bluele/factory-go/factory"
	"github.com/guregu/null/v6"
	"github.com/iancoleman/strcase"
	"go.uber.org/fx"
)

var _ Seeder = (*UserSeeder)(nil)

type UserSeeder struct {
	idService          id.IdService
	userFactory        *factoryGo.Factory
	employeeFactory    *factoryGo.Factory
	roleUserFactory    *factoryGo.Factory
	userRepository     repository.UserRepository
	employeeRepository employeeRepository.EmployeeRepository
	roleRepository     roleRepository.RoleRepository
	roleUserRepository roleUserRepository.RoleUserRepository
}

type NewUserSeederParams struct {
	fx.In

	IdService          id.IdService
	UserFactory        *factoryGo.Factory `name:"user_factory"`
	EmployeeFactory    *factoryGo.Factory `name:"employee_factory"`
	RoleUserFactory    *factoryGo.Factory `name:"role_user_factory"`
	UserRepository     repository.UserRepository
	EmployeeRepository employeeRepository.EmployeeRepository
	RoleRepository     roleRepository.RoleRepository
	RoleUserRepository roleUserRepository.RoleUserRepository
}

func NewUserSeeder(params NewUserSeederParams) Seeder {
	return &UserSeeder{
		idService:          params.IdService,
		userFactory:        params.UserFactory,
		employeeFactory:    params.EmployeeFactory,
		roleUserFactory:    params.RoleUserFactory,
		userRepository:     params.UserRepository,
		employeeRepository: params.EmployeeRepository,
		roleRepository:     params.RoleRepository,
		roleUserRepository: params.RoleUserRepository,
	}
}

func (s *UserSeeder) Seed() error {
	var err error
	userFactory := s.userFactory
	employeeFactory := s.employeeFactory
	roleUserFactory := s.roleUserFactory

	superAdmin := userFactory.MustCreateWithOption(map[string]any{
		"Id":              "01K7HR3Z46X1B3X7XQ2JFKF2DM",
		"Name":            strcase.ToCamel(enum.SuperAdmin.String()),
		"EmailVerifiedAt": null.TimeFrom(time.Now()),
		"ActivatedAt":     null.TimeFrom(time.Now()),
		"DeactivatedAt":   null.TimeFromPtr(nil),
	}).(*entity.User)
	admin := userFactory.MustCreateWithOption(map[string]any{
		"Id":   "01K7HR9A6WKK0W67W8S7F8755X",
		"Name": strcase.ToCamel(enum.Admin.String()),
	}).(*entity.User)
	head := userFactory.MustCreateWithOption(map[string]any{
		"Id":   "01K7HR8N4Z9TE0ENB0K1275MNW",
		"Name": strcase.ToCamel(enum.Head.String()),
	}).(*entity.User)
	customerServiceAdmin := userFactory.MustCreateWithOption(map[string]any{
		"Id":   "01K7HR8WZZBSS957Q9NSQ85MYA",
		"Name": strcase.ToCamel(enum.CustomerServiceAdmin.String()),
	}).(*entity.User)

	users := []*entity.User{
		superAdmin,
		admin,
		head,
		customerServiceAdmin,
	}

	for _, user := range users {
		user.Email = fmt.Sprintf("%s@gmail.com", strcase.ToSnake(user.Name))
	}

	err = s.userRepository.SaveMany(users)
	if err != nil {
		return err
	}

	{
		// ========== Employee ==========
		employees := []*entity.Employee{}
		for _, user := range users {
			employee := employeeFactory.MustCreateWithOption(map[string]any{
				"UserId": user.Id,
			}).(*entity.Employee)
			employees = append(employees, employee)
		}

		err = s.employeeRepository.SaveMany(employees)
		if err != nil {
			return err
		}

	}

	{
		// ========== RoleUser ==========
		// Assign roles to the created users by creating role-user relationships.
		// First, fetch all available roles and create a map for easy lookup by name.
		// Then, assign specific roles to each user by creating RoleUser entries.
		// This ensures proper role-based access control in the application.

		// Fetch all available roles from the database
		roles, err := s.roleRepository.All()
		if err != nil {
			return err
		}

		// Create a map of role names to role objects for easy lookup
		roleMap := make(map[enum.RoleName]*entity.Role)
		for _, role := range roles {
			roleMap[role.Name] = role
		}

		// Create role-user relationships for each user with their respective roles
		roleUsers := []*entity.RoleUser{
			// Super Admin role assignment
			roleUserFactory.MustCreateWithOption(map[string]any{
				"RoleId": roleMap[enum.SuperAdmin].Id,
				"UserId": superAdmin.Id,
			}).(*entity.RoleUser),
			// Admin role assignment
			roleUserFactory.MustCreateWithOption(map[string]any{
				"RoleId": roleMap[enum.Admin].Id,
				"UserId": admin.Id,
			}).(*entity.RoleUser),

			// Head role assignment
			roleUserFactory.MustCreateWithOption(map[string]any{
				"RoleId": roleMap[enum.Head].Id,
				"UserId": head.Id,
			}).(*entity.RoleUser),

			// Customer Service Admin role assignment
			roleUserFactory.MustCreateWithOption(map[string]any{
				"RoleId": roleMap[enum.CustomerServiceAdmin].Id,
				"UserId": customerServiceAdmin.Id,
			}).(*entity.RoleUser),
		}

		// Save all role-user relationships to the database
		err = s.roleUserRepository.SaveMany(roleUsers)
		if err != nil {
			return err
		}

	}

	return nil
}
