module wombag

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.9.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/schema v1.2.0
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/olekukonko/tablewriter v0.0.4
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.8.1
	github.com/urfave/negroni v1.0.0
	golang.org/x/crypto v0.5.0
	golang.org/x/net v0.7.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.9
//	internal/wombaglib v1.0.0
)

//replace internal/wombaglib => ./internal/wombaglib
