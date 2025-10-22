package seeder

import (
	"github.com/arfanxn/welding/internal/infrastructure/database/factory"
	employeeRepository "github.com/arfanxn/welding/internal/module/employee/domain/repository"
	roleRepository "github.com/arfanxn/welding/internal/module/role/domain/repository"
	roleUserRepository "github.com/arfanxn/welding/internal/module/role_user/domain/repository"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/internal/module/user/domain/repository"
	"go.uber.org/fx"
)

var _ Seeder = (*UserSeeder)(nil)

type UserSeeder struct {
	userRepository     repository.UserRepository
	employeeRepository employeeRepository.EmployeeRepository
	roleRepository     roleRepository.RoleRepository
	roleUserRepository roleUserRepository.RoleUserRepository
}

type NewUserSeederParams struct {
	fx.In

	UserRepository     repository.UserRepository
	EmployeeRepository employeeRepository.EmployeeRepository
	RoleRepository     roleRepository.RoleRepository
	RoleUserRepository roleUserRepository.RoleUserRepository
}

func NewUserSeeder(params NewUserSeederParams) Seeder {
	return &UserSeeder{
		userRepository:     params.UserRepository,
		employeeRepository: params.EmployeeRepository,
		roleRepository:     params.RoleRepository,
		roleUserRepository: params.RoleUserRepository,
	}
}

func (s *UserSeeder) Seed() error {
	superAdmin := factory.UserFactory.MustCreateWithOption(map[string]any{
		"Id":    "01K7HR3Z46X1B3X7XQ2JFKF2DM",
		"Name":  "Super Admin",
		"Email": "super_admin@gmail.com",
	}).(*entity.User)
	admin := factory.UserFactory.MustCreateWithOption(map[string]any{
		"Id":    "01K7HR9A6WKK0W67W8S7F8755X",
		"Name":  "Admin",
		"Email": "admin@gmail.com",
	}).(*entity.User)
	head := factory.UserFactory.MustCreateWithOption(map[string]any{
		"Id":    "01K7HR8N4Z9TE0ENB0K1275MNW",
		"Name":  "Head",
		"Email": "head@gmail.com",
	}).(*entity.User)
	customerServiceAdmin := factory.UserFactory.MustCreateWithOption(map[string]any{
		"Id":    "01K7HR8WZZBSS957Q9NSQ85MYA",
		"Name":  "Customer Service Admin",
		"Email": "customer_service_admin@gmail.com",
	}).(*entity.User)

	users := []*entity.User{
		superAdmin,
		admin,
		head,
		customerServiceAdmin,
	}

	err := s.userRepository.SaveMany(users)
	if err != nil {
		return err
	}

	{
		// ========== Employee ==========
		employees := []*entity.Employee{}
		for _, user := range users {
			employee := factory.EmployeeFactory.MustCreateWithOption(map[string]any{
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
		roleMap := make(map[string]*entity.Role)
		for _, role := range roles {
			roleMap[role.Name] = role
		}

		// Create role-user relationships for each user with their respective roles
		roleUsers := []*entity.RoleUser{
			// Super Admin role assignment
			factory.RoleUserFactory.MustCreateWithOption(map[string]any{
				"RoleId": roleMap["super admin"].Id,
				"UserId": superAdmin.Id,
			}).(*entity.RoleUser),

			// Admin role assignment
			factory.RoleUserFactory.MustCreateWithOption(map[string]any{
				"RoleId": roleMap["admin"].Id,
				"UserId": admin.Id,
			}).(*entity.RoleUser),

			// Head role assignment
			factory.RoleUserFactory.MustCreateWithOption(map[string]any{
				"RoleId": roleMap["head"].Id,
				"UserId": head.Id,
			}).(*entity.RoleUser),

			// Customer Service Admin role assignment
			factory.RoleUserFactory.MustCreateWithOption(map[string]any{
				"RoleId": roleMap["customer service admin"].Id,
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
