package services

import (
	repoMocks "auth_audit/internal/app/repository/mocks"
	"auth_audit/internal/app/repository/models"
	"auth_audit/internal/app/server/DTO"
	"auth_audit/internal/app/server/services/mocks"
	"auth_audit/pkg/errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"testing"

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
		ur *repoMocks.MockUserRepository
		us *UserService
	)

	BeforeEach(func() {
		ur = &repoMocks.MockUserRepository{}
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

	Context("hashPwd", func() {
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
		sr *repoMocks.MockSessionRepository
		ss *SessionService
	)

	BeforeEach(func() {
		sr = &repoMocks.MockSessionRepository{}
		ss = NewSessionService(sr)
	})

	Context("Create", func() {
		It("Should create a session", func() {
			sr.On("Create", mock.Anything).Return(nil)

			session, err := ss.Create(1)
			Expect(err).ToNot(HaveOccurred())
			Expect(session).ToNot(BeNil())
		})

		It("empty foreignKey", func() {
			session, err := ss.Create(0)
			Expect(session).To(BeNil())
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(errors.NullForeignKey))
		})
	})

	Context("GetByToken", func() {
		It("token is empty", func() {
			session, err := ss.GetByToken("")
			Expect(session).To(BeNil())
			Expect(err).To(Equal(errors.TokenIsEmpty))
		})

		It("with valid token", func() {
			currentTime := time.Now()
			validSession := &models.Session{
				Token:    "valid_token",
				LiveTime: currentTime.Add(time.Hour),
			}

			sr.On("GetByToken", validSession.Token).Return(validSession, nil)
			session, err := ss.GetByToken(validSession.Token)
			Expect(session).To(Equal(validSession))
			Expect(err).To(BeNil())
		})

		It("with expired token should return TokenHasExpired error", func() {
			expiredSession := &models.Session{
				Token:    "expired_token",
				LiveTime: time.Now().Add(-time.Hour),
			}

			sr.On("GetByToken", expiredSession.Token).Return(nil, errors.TokenHasExpired)

			session, err := ss.GetByToken("expired_token")
			Expect(session).To(BeNil())
			Expect(err).To(Equal(errors.TokenHasExpired))
		})
	})

	Context("generateToken", func() {
		It("length should be 32", func() {
			token := ss.generateToken()
			Expect(token).To(HaveLen(32))
		})
	})
})

var _ = Describe("AuthService", func() {
	var (
		us          *mocks.MockUserService
		ss          *mocks.MockSessionService
		aas         *mocks.MockAuthAudit
		as          *AuthService
		registerDto DTO.RegisterUserDTO
		loginDto    DTO.LoginUserDTO
	)

	BeforeEach(func() {
		us = &mocks.MockUserService{}
		ss = &mocks.MockSessionService{}
		aas = &mocks.MockAuthAudit{}
		as = NewAuthService(us, ss, aas)

		registerDto = DTO.RegisterUserDTO{
			Login:    "login",
			Password: "password_test",
		}

		loginDto = DTO.LoginUserDTO{
			Login:    "login",
			Password: "password_test",
		}
	})

	Context("Register", func() {
		var user *models.User

		BeforeEach(func() {
			user = &models.User{}
		})

		It("if such user already exists", func() {
			us.On("GetUserByLogin", mock.Anything).Return(&models.User{
				Login:        "a",
				PasswordHash: "by",
			}, nil)

			session, err := as.Register(registerDto)
			Expect(session).To(BeNil())
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(errors.UserAlreadyExist))
		})

		It("user successfully registered", func() {
			us.On("GetUserByLogin", mock.Anything).Return(nil, nil)
			user.ID = 1
			us.On("CreateUser", mock.Anything).Return(user, nil)

			ss.On("Create", user.ID).Return(&models.Session{UserID: user.ID}, nil)

			session, err := as.Register(registerDto)
			Expect(err).ToNot(HaveOccurred())
			Expect(session).ToNot(BeNil())
			Expect(session.UserID).To(Equal(user.ID))
		})

		It("empty user.ID for session foreignKey", func() {
			us.On("GetUserByLogin", mock.Anything).Return(nil, nil)
			user.ID = 0
			us.On("CreateUser", mock.Anything).Return(user, nil)
			ss.On("Create", user.ID).Return(nil, errors.NullForeignKey)

			session, err := as.Register(registerDto)
			Expect(session).To(BeNil())
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(errors.NullForeignKey))
		})
	})

	Context("Login", func() {
		It("empty login or/and password", func() {
			s, err := as.Login(DTO.LoginUserDTO{})
			Expect(s).To(BeNil())
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(errors.MustBeProvidedLoginAndPwd))
		})

		It("unable to find user", func() {
			us.On("GetUserByLogin", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
			s, err := as.Login(loginDto)
			Expect(s).To(BeNil())
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(errors.InvalidLoginOrPassword))
		})

		It("user has been blocked", func() {
			us.On("GetUserByLogin", mock.Anything).Return(&models.User{
				Blocked: true,
			}, nil)

			s, err := as.Login(loginDto)
			Expect(s).To(BeNil())
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(errors.UserHasBeenBlocked))
		})

		It("more than 5 login attempts", func() {
			us.On("GetUserByLogin", mock.Anything).Return(&models.User{
				FailedLoginAttempts: 5,
			}, nil)
			us.On("BlockUser", mock.Anything).Return(nil)
			aas.On("Create", mock.Anything, mock.Anything).Return(nil)

			s, err := as.Login(loginDto)
			Expect(s).To(BeNil())
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(errors.UserHasBeenBlocked))
		})

		It("wrong passwords when comparing", func() {
			us.On("GetUserByLogin", mock.Anything).Return(&models.User{
				PasswordHash: "something_password",
			}, nil)
			us.On("IncrementFailedLoginAttempts", mock.Anything).Return(0, nil)
			aas.On("Create", mock.Anything, mock.Anything).Return(nil)

			s, err := as.Login(loginDto)
			Expect(s).To(BeNil())
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(errors.InvalidLoginOrPassword))
		})

		It("user successfully logged in", func() {
			bytes, err := bcrypt.GenerateFromPassword([]byte(loginDto.Password), bcrypt.DefaultCost)
			Expect(err).ToNot(HaveOccurred())

			user := &models.User{
				PasswordHash: string(bytes),
			}
			user.ID = 1

			us.On("GetUserByLogin", mock.Anything).Return(user, nil)
			ss.On("Create", user.ID).Return(&models.Session{
				UserID: user.ID,
			}, nil)
			aas.On("Create", mock.Anything, user.ID).Return(nil)

			s, err := as.Login(loginDto)
			Expect(err).ToNot(HaveOccurred())
			Expect(s).ToNot(BeNil())
			Expect(s.UserID).To(Equal(user.ID))
		})
	})

	Context("GetAuthAuditByToken", func() {
		It("should be return 1 audits", func() {
			ss.On("GetByToken", mock.Anything).Return(&models.Session{UserID: 1}, nil)
			var user models.User
			user.ID = 1
			us.On("GetUserByID", mock.Anything).Return(&user, nil)
			aas.On("GetAllAuditsByUserID", mock.Anything).Return([]models.AuthenticationAudit{{
				Time:  time.Now(),
				Event: "test",
			}}, nil)

			audits, err := as.GetAuthAuditByToken("token")
			Expect(err).ToNot(HaveOccurred())
			Expect(audits).To(HaveLen(1))
		})
	})
})
