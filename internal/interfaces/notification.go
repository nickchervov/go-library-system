package interfaces

import "github.com/nickchervov/go-library-system/internal/models"

type NotificationService interface {
	NotifyOverdue(loan models.LoanInfo) error
}
