package interfaces

import "github.com/nickchervov/go-library-system/internal/models"

type AnalyticsProvider interface {
	GetMostPopularGenre() (string, int)
	FindOverdueLoans() []models.LoanInfo
}
