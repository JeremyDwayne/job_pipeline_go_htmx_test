package store

type User struct {
	Email    string
	Password string
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}

type Company struct {
	Name string
	Url  string
}

type CompanyStore interface {
	CreateCompany(name string, url string) error
	GetCompany(name string) (*Company, error)
}

type Recruiter struct {
	Name  string
	Email string
}

type RecruiterStore interface {
	CreateRecruiter(name string, email string) error
	GetRecruiter(name string) (*Recruiter, error)
}

type Interviewer struct {
	Name   string
	Email  string
	Title  string
	Tenure string
	Notes  string
}

type InterviewerStore interface {
	CreateInterviewer(name string, email string, title string, tenure string, notes string) error
	GetInterviewer(name string) (*Interviewer, error)
}

type Interview struct {
	Type  string
	Date  string
	Notes string
}

type InterviewStore interface {
	CreateInterview(Type string, Date string, Notes string) error
	GetInterview(Type string) (*Interview, error)
}
