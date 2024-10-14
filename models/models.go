package models
import (
	"gorm.io/gorm"
"time"
)

const (
	Monthly = "1 Month Plan"
	TriMonthly = "3 Month Plan"
	HalfYearly = "6 Month Plan"
	Yearly = "12 Month Plan"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Subscription struct {
    gorm.Model
    PlanName  string `json:"plan_name"`
    ExpiresAt time.Time `json:"expires_at"`
    
    // One-to-one relationship
    UserID uint  `json:"user_id"`
    User   User  `gorm:"OnDelete:CASCADE;" json:"user,omitempty"`
}
