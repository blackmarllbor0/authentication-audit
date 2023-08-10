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
		h := NewHandler(as)
		router.POST("/auth/register", h.register)
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
})
