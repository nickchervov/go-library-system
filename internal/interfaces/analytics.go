package interfaces

import (
	"time"

	"github.com/nickchervov/go-library-system/internal/models"
)

type AnalyticsProvider interface {
	GetMostPopularGenre() (string, int)
	CalculateReadTime(memberId string) time.Duration
	FindOverdueLoans() []models.LoanInfo
	GenerateMonthlyReport(start, end time.Time) (map[string]any, error)
}
