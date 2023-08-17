package handlers

import (
	"auth_audit/internal/app/repository/models"
	"auth_audit/internal/app/server/DTO"
	"auth_audit/internal/app/server/services/mocks"
	"auth_audit/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}

var _ = Describe("Handler", func() {
	var (
		router *gin.Engine
		as     *mocks.MockAuthService
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		router = gin.Default()
		as = &mocks.MockAuthService{}
		ss := &mocks.MockSessionService{}
		h := NewHandler(as, ss)
		router.POST("/auth/register", h.register)
		router.POST("/auth/login", h.login)
		router.GET("/auth/audit", h.getAuthAuditByToken)
	})

	Context("register", func() {
		It("successful registration", func() {
			as.On("Register", mock.Anything).Return(&models.Session{
				Token: "8ytvuybr49r2849bg92b4g2850h4nnjfb",
			}, nil)

			// Create a test JSON request.
			requestBody := `{"login": "test_login", "password": "test_password"}`

			// Create an HTTP request with test data.
			req, err := http.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(requestBody))
			Expect(err).ToNot(HaveOccurred())

			// Create an HTTP ResponseRecorder to record the response.
			w := httptest.NewRecorder()

			// Send an HTTP request to the router
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusCreated))
		})

		It("empty login and/or password", func() {
			as.On("Register", DTO.RegisterUserDTO{
				Login:    "",
				Password: "",
			}).Return(nil, errors.MustBeProvidedLoginAndPwd)

			requestBody := `{"login": "", "password": ""}`

			req, err := http.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(requestBody))
			Expect(err).ToNot(HaveOccurred())

			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})
	})

	Context("login", func() {
		It("with valid request body should return a token when login is successful", func() {
			body := `{"login": "test_login", "password": "test_password"}`

			as.On("Login", DTO.LoginUserDTO{
				Login:    "test_login",
				Password: "test_password",
			}).Return(&models.Session{
				Token: "test_token",
			}, nil)

			req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(body))
			Expect(err).ToNot(HaveOccurred())

			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(Equal(`{"token":"test_token"}`))
		})

		It("with invalid request body should return a 400 status with error message", func() {
			invalidBody := `{"password": ""}`

			as.On("Login", DTO.LoginUserDTO{
				Login:    "",
				Password: "",
			}).Return(nil, errors.MustBeProvidedLoginAndPwd)

			req, err := http.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(invalidBody))
			Expect(err).ToNot(HaveOccurred())

			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(Equal(`{"error":"login and password must be provided"}`))
		})
	})

	Context("getAuthAuditByToken", func() {
		It("with invalid token should return a 500 status with error message", func() {
			as.On("GetAuthAuditByToken", mock.Anything).Return(nil, errors.TokenHasExpired)

			req, err := http.NewRequest(http.MethodGet, "/auth/audit", nil)
			Expect(err).ToNot(HaveOccurred())
			req.Header.Set("X-Token", "invalid-token")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusInternalServerError))
			Expect(w.Body.String()).To(Equal(`{"error":"the token has expired"}`))
		})

		It("with valid token should return a 200 status with audit data", func() {
			token := "token"
			as.On("GetAuthAuditByToken", mock.Anything).Return([]DTO.AuthAuditDTO{{
				Timestamp: time.Now(),
				Event:     "event",
			}}, nil)

			req, err := http.NewRequest(http.MethodGet, "/auth/audit", nil)
			Expect(err).ToNot(HaveOccurred())
			req.Header.Set("X-Token", token)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
})
