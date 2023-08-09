package postgres

import (
	"auth_audit/config"
	"auth_audit/internal/app/repository/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"testing"
	"time"
)

const testSchemaName = "tests"

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var _ = Describe("Database", func() {
	var (
		cvg *config.Config
		r   *Repository
	)

	BeforeEach(func() {
		cvg = config.NewConfig(viper.New())
		err := cvg.LoadConfig("../../../../config", "yaml", "config")
		Expect(err).ToNot(HaveOccurred())

		r = NewRepository(cvg)
		_, err = r.Connect()
		Expect(err).ToNot(HaveOccurred())
	})

	Context("Connect", func() {
		It("successful connection and disconnection from the database", func() {
			Expect(r.db).ToNot(BeNil())
			r.db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", testSchemaName))
		})
	})

	Context("autoMigrate", func() {
		It("successful table creation in test database schema", func() {
			r.db.Exec(fmt.Sprintf("SET search_path TO %s", testSchemaName))

			r.db.DropTableIfExists(
				&models.Session{},
				&models.AuthenticationAudit{},
				&models.User{},
			)

			err := r.autoMigrate()
			Expect(err).ToNot(HaveOccurred())

			Expect(r.db.HasTable(&models.User{})).To(BeTrue())
			Expect(r.db.HasTable(&models.Session{})).To(BeTrue())
			Expect(r.db.HasTable(&models.AuthenticationAudit{})).To(BeTrue())
		})
	})

	AfterEach(func() {
		err := r.Disconnect()
		Expect(err).ToNot(HaveOccurred())
	})
})

var _ = Describe("UserRepository", func() {
	const (
		login    = "login"
		password = "test_password"
	)

	var (
		DB   *gorm.DB
		ur   *UserRepository
		user *models.User
	)

	BeforeEach(func() {
		DB = connectToDB()

		ur = NewUserRepository(DB)

		user = &models.User{
			Login:        login,
			PasswordHash: password,
		}
	})

	Context("Create", func() {
		It("user successfully created", func() {
			err := ur.Create(user)
			Expect(err).ToNot(HaveOccurred())
			Expect(user).ToNot(BeNil())
			Expect(user.ID).ToNot(Equal(0))
		})
	})

	Context("GetById", func() {
		It("managed to find the user", func() {
			err := ur.Create(user)
			findUser, err := ur.GetById(user.ID)
			Expect(err).ToNot(HaveOccurred())
			Expect(findUser).ToNot(BeNil())
			Expect(findUser.ID).To(Equal(user.ID))
		})

		It("unable to find user", func() {
			_, err := ur.GetById(111111)
			Expect(err).To(Equal(gorm.ErrRecordNotFound))
		})
	})

	Context("GetByLogin", func() {
		It("managed to find the user", func() {
			err := ur.Create(user)
			findUser, err := ur.GetByLogin(user.Login)
			Expect(err).ToNot(HaveOccurred())
			Expect(findUser).ToNot(BeNil())
			Expect(findUser.Login).To(Equal(user.Login))
		})

		It("unable to find user", func() {
			_, err := ur.GetByLogin("user.Login")
			Expect(err).To(Equal(gorm.ErrRecordNotFound))
		})
	})

	AfterEach(func() {
		err := DB.Close()
		Expect(err).ToNot(HaveOccurred())
	})
})

var _ = Describe("SessionRepository", func() {
	var (
		DB      *gorm.DB
		sr      *SessionRepository
		session *models.Session
	)

	BeforeEach(func() {
		DB = connectToDB()

		sr = NewSessionRepository(DB)

		session = &models.Session{
			Token:    "uwhrbwrubrbq467y978h%^&*3uhgKjhsb",
			LiveTime: time.Now().Add(time.Hour * 1),
		}
	})

	Context("Create", func() {
		It("session successfully created", func() {
			ur := NewUserRepository(DB)

			var user models.User
			err := ur.Create(&user)
			Expect(err).ToNot(HaveOccurred())
			Expect(user.ID).ToNot(Equal(0))

			session.UserID = user.ID

			err = sr.Create(session)
			Expect(err).ToNot(HaveOccurred())
			Expect(session).ToNot(BeNil())
			Expect(session.ID).ToNot(Equal(0))
		})
	})

	AfterEach(func() {
		err := DB.Close()
		Expect(err).ToNot(HaveOccurred())
	})
})

func connectToDB() *gorm.DB {
	cvg := config.NewConfig(viper.New())
	err := cvg.LoadConfig("../../../../config", "yaml", "config")
	Expect(err).ToNot(HaveOccurred())

	DB, err := gorm.Open("postgres", cvg.GetValueByKeys("app.db.dsn"))
	Expect(err).ToNot(HaveOccurred())
	DB.Exec(fmt.Sprintf("SET search_path TO %s", testSchemaName))

	return DB
}
