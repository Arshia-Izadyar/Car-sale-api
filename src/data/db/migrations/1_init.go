package migrations

import (
	"database/sql"

	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/constants"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/db"
	"github.com/Arshia-Izadyar/Car-sale-api/src/data/models"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Up_1() {
	database := db.GetDB()
	createTable(database)
}

var logger = logging.NewLogger(config.GetConfig())

func createTable(database *gorm.DB) {
	tables := []interface{}{}

	tables = addNewTable(database, &models.BaseModel{}, tables)
	tables = addNewTable(database, &models.Role{}, tables)
	tables = addNewTable(database, &models.User{}, tables)
	tables = addNewTable(database, &models.UserRole{}, tables)

	tables = addNewTable(database, &models.PropertyCategory{}, tables)
	tables = addNewTable(database, &models.Property{}, tables)

	tables = addNewTable(database, &models.Country{}, tables)
	tables = addNewTable(database, &models.City{}, tables)
	tables = addNewTable(database, &models.PersianYear{}, tables)
	tables = addNewTable(database, &models.Color{}, tables)
	tables = addNewTable(database, &models.File{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(err, logging.Postgres, logging.Insert, "cant add tables", nil)
		panic(err)
	}
	createDefaultInfo(database)
}

func addNewTable(db *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !db.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createIfNotExists(db *gorm.DB, r *models.Role) {
	exists := 0
	db.Model(&models.Role{}).Select("1").Where("name = ?", r.Name).First(&exists)
	if exists == 0 {
		db.Create(r)
	}
}
func createDefaultInfo(db *gorm.DB) {
	admin := models.Role{Name: "admin"}
	createIfNotExists(db, &admin)
	defaultRole := models.Role{Name: "default"}
	createIfNotExists(db, &defaultRole)

	u := models.User{
		BaseModel:   models.BaseModel{},
		Username:    constants.AdminRoleName,
		FirstName:   "test",
		LastName:    sql.NullString{Valid: true, String: "test"},
		PhoneNumber: sql.NullString{Valid: true, String: "09108624707"},
		Enable:      true,
	}
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte("a123"), bcrypt.DefaultCost)
	u.Password = string(hashedPass)
	createAdmin(db, &u, admin.ID)
}

func createAdmin(db *gorm.DB, u *models.User, roleID int) {
	exists := 0
	db.Model(&models.User{}).Select("1").Where("username = ?", u.Username).First(&exists)
	if exists == 0 {
		db.Create(u)
		ur := models.UserRole{UserId: u.ID, RoleId: roleID}
		db.Create(&ur)
	}
}
