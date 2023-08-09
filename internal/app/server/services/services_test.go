package services

import (
	"auth_audit/internal/app/repository/mocks"
	"auth_audit/internal/app/repository/models"
	"auth_audit/internal/app/server/DTO"
	"auth_audit/pkg/errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"testing"

	"github.com/stretchr/testify/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

var _ = Describe("UserService", func() {
	const (
		login = "login"
		pwd   = "test_password"
	)

	var (
		ur *mocks.MockUserRepository
		us *UserService
	)

	BeforeEach(func() {
		ur = &mocks.MockUserRepository{}
		us = NewUserService(ur)
	})

	Context("CreateUser", func() {
		It("Should create a user", func() {
			ur.On("Create", mock.Anything).Return(nil)

			createUserDTO := DTO.RegisterUserDTO{
				Login:    login,
				Password: pwd,
			}

			user, err := us.CreateUser(createUserDTO)
			Expect(err).ToNot(HaveOccurred())
			Expect(user).ToNot(BeNil())
		})

		It("Should return an error for short password", func() {
			createUserDTO := DTO.RegisterUserDTO{
				Login:    login,
				Password: "short",
			}

			user, err := us.CreateUser(createUserDTO)
			Expect(err).To(HaveOccurred())
			Expect(user).To(BeNil())
			Expect(err).To(Equal(errors.ShortPassword))
		})
	})

	Context("GetUserByLogin", func() {
		It("should get user by login", func() {
			mockUser := &models.User{
				Login:        login,
				PasswordHash: pwd,
			}

			ur.On("GetByLogin", login).Return(mockUser, nil)

			user, err := us.GetUserByLogin(login)
			Expect(err).ToNot(HaveOccurred())
			Expect(user).ToNot(BeNil())
			Expect(user).To(Equal(mockUser))
		})

		It("should return an error for nonexistent user", func() {
			ur.On("GetByLogin", "nonexistentuser").Return(nil, gorm.ErrRecordNotFound)

			user, err := us.GetUserByLogin("nonexistentuser")
			Expect(err).To(HaveOccurred())
			Expect(user).To(BeNil())
			Expect(err).To(Equal(gorm.ErrRecordNotFound))
		})
	})

	Context("HashPwd", func() {
		It("should hash the password", func() {
			hash, err := us.hashPwd(pwd)
			Expect(err).ToNot(HaveOccurred())
			Expect(hash).ToNot(BeEmpty())

			// Verify that the hash matches the original password
			err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
			Expect(err).ToNot(HaveOccurred())
		})

		It("should return an error for empty password", func() {
			hash, err := us.hashPwd("")
			Expect(err).To(Equal(errors.ShortPassword))
			Expect(hash).To(BeEmpty())
		})
	})
})

var _ = Describe("SessionService", func() {
	var (
		sr *mocks.MockSessionRepository
		ss *SessionService
	)

	BeforeEach(func() {
		sr = &mocks.MockSessionRepository{}
		ss = NewSessionService(sr)
	})

	Context("Create", func() {
		It("Should create a session", func() {
			sr.On("Create", mock.Anything).Return(nil)

			session, err := ss.Create(1)
			Expect(err).ToNot(HaveOccurred())
			Expect(session).ToNot(BeNil())
		})
	})
})
