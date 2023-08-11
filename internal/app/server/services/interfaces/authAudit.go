package interfaces

type AuthAuditService interface {
	Create(event string, userID uint) error
}
