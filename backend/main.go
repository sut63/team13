package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team13/app/controllers"
	_ "github.com/team13/app/docs"
	"github.com/team13/app/ent"
)

type Customer struct {
	Customer []Customer
}

type Customer struct {
	Name  string
	Email string
}

type Paymentchannel struct {
	Paymentchannel []Paymentchannel
}

type Paymentchannel struct {
	bank string
}

type Product struct {
	Product []Product
}

type Product struct {
	NameProduct    string
	BarcodeProduct string
	MFG            string
	Example        string
}

type Typeproducts struct {
	Typeproduct []Typeproduct
}

type Typeproduct struct {
	Typeproduct string
}

type Customers struct{
    Customer []Customer
}

type Zoneproducts struct {
	Zoneproduct []Zoneproduct
}

type Zoneproduct struct {
	Zone string
	
type Paymentchannels struct{
    Paymentchannel []Paymentchannel
}

type Days struct {
	Day []Day
}

type Day struct {
	Day string
}

type Products struct{
    Product []Product
}

type Employees struct
	Employee []Employee
}

type Employee struct {
	Name  string
	Email string
}

type Roles struct {
	Role []Role
}

type Role struct {
	Role string
}

type Companys struct{
    Company []Company
}

type Company struct{
    Name string
}

type Managers struct{
    Customer []Customer
}

type Manager struct{
    Name string
    Email string
}

type Zoneproducts struct{
    Zoneproduct []Zoneproduct
}

type Shifts struct {
	Shift []Shift
}

type Shift struct {
	TimeStart time.Time
	TimeEnd   time.Time
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewCustomerController(v1, client)
	controllers.NewOrderController(v1, client)
	controllers.NewPaymentchannelController(v1, client)
	controllers.NewProductController(v1, client)
	controllers.NewTypeproductController(v1, client)
	controllers.NewZoneproductController(v1, client)
    controllers.NewDayController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewEmployeeWorkingHoursController(v1, client)
	controllers.NewRoleController(v1, client)
	controllers.NewShiftController(v1, client)

	customers := Customers{
		Customer: []Customer{
			Customer{"Dang Dang", "Dang@gmail.com"},
			Customer{"AEK Dang", "AEK@gmail.com"},
			Customer{"PANG Dang", "PANG@gmail.com"},
			Customer{"NW Dang", "NW@gmail.com"},
		},
	}

	for _, c := range customer.Customer {
		client.Customer.
			Create().
			SetName(c.Name).
			SetEmail(c.Email).
			Save(context.Background())
	}

	paymentchannels := paymentchannels{
		Paymentchannel: []Paymentchannel{
   router := gin.Default()
   router.Use(cors.Default())
 
   client, err := ent.Open("sqlite3", "file:ent.db?&cache=shared&_fk=1")
   if err != nil {
       log.Fatalf("fail to open sqlite3: %v", err)
   }
   defer client.Close()
 
   if err := client.Schema.Create(context.Background()); err != nil {
       log.Fatalf("failed creating schema resources: %v", err)
   }
 
   v1 := router.Group("/api/v1")
   controllers.NewCustomerController(v1, client)
   controllers.NewOrderonlineController(v1, client)
   controllers.NewPaymentchannelController(v1, client)
   controllers.NewProductController(v1, client)
   controllers.NewTypeproductController(v1, client)

   controllers.NewZoneproductController(v1, client)
   controllers.NewStockController(v1, client)

   controllers.NewCompanyController(v1, client)
   controllers.NewManagerController(v1, client)
   controllers.NewOrderproductController(v1, client)
   

   customers := Customers{
    Customer: []Customer{
                Customer{"Dang Dang","Dang@gmail.com"},
                Customer{"AEK Dang","AEK@gmail.com"},
                Customer{"PANG Dang","PANG@gmail.com"},
                Customer{"NW Dang","NW@gmail.com"},
        },
   }

   for _, c := range customer.Customer {
        client.Customer.
            Create().
            SetName(c.Name).
            SetEmail(c.Email).
            Save(context.Background())
   }

   paymentchannels := paymentchannels{
    Paymentchannel: []Paymentchannel{
			Paymentchannel{"KBANK"},
			Paymentchannel{"KTB"},
			Paymentchannel{"TMB"},
			Paymentchannel{"SCB"},
		},
	}

	for _, pay := range paymentchannel.Paymentchannel {
		client.Paymentchannel.
			Create().
			SetBank(pay.bank).
			Save(context.Background())
	}

	products := Products{
		Product: []Product{
			Product{"A", "001", "01-01-2018", "01-01-2024"},
			Product{"B", "002", "01-01-2018", "01-01-2024"},
			Product{"C", "003", "01-01-2018", "01-01-2024"},
			Product{"D", "004", "01-01-2018", "01-01-2024"},
			Product{"E", "005", "01-01-2018", "01-01-2024"},
		},
	}

	for _, pd := range product.Product {
		client.Product.
			Create().
			SetNameProduct(pd.NameProduct).
			SetBarcodeProduct(pd.BarcodeProduct).
			SetMFG(pd.MFG).
			SetEXP(pd.EXP).
			Save(context.Background())
	}

	typeproducts := Typeproducts{
		Typeproduct: []Typeproduct{
			Typeproduct{"KBANK"},
			Typeproduct{"KTB"},
			Typeproduct{"TMB"},
			Typeproduct{"SCB"},
		},
	}

	for _, tp := range typeproduct.Typeproduct {
		client.Typeproduct.
			Create().
			SetTypeProduct(tp.Typeproduct).
			Save(context.Background())
	}

	zoneproducts := Zoneproducts{
		Zoneproduct: []Zoneproduct{
			Zoneproduct{"A"},
			Zoneproduct{"B"},
			Zoneproduct{"C"},
			Zoneproduct{"D"},
		},
	}

	for _, z := range zoneproduct.Zoneproduct {
		client.Zoneproduct.
			Create().
			SetZone(z.Zone).
			Save(context.Background())
    }
    
    Days := Days{
		Day: []Day{
			{"วันอาทิตย์"},
			{"วันจันทร์"},
			{"วันอังคาร"},
			{"วันพุธ"},
			{"วันพฤหัสบดี"},
			{"วันศุกร์"},
			{"วันเสาร์"},
		},
	}

	for _, d := range Days.Day {
		client.Day.
			Create().
			SetDay(d.Day).
			Save(context.Background())
    }
    
    Employees := Employees{
		Employee: []Employee{
			{"Poomin PhimPhimai","Poomin123@gmail.com"},
			{"Pakiafa Kummungkun","Pakiafa456@gmail.com"},
		},
	}

	for _, em := range Employees.Employee {
		client.Employee.
			Create().
            SetName(em.Name).
            SetEmail(em.Email).
			Save(context.Background())
    }

    Roles := Roles{
		Role: []Role{
			{"กวาดพื้น"},
			{"จัดของ"},
			{"เช็ดกระจก"},
			{"CashCheer"},
			{"ส่งของ"},
		},
	}

	for _, r := range Roles.Role {
		client.Role.
			Create().
			SetRole(r.Role).
			Save(context.Background())
    }

    Shifts := Shifts{
		Shift: []Shift{
			{time.Date(0, 0, 0, 6, 30, 0, 0, time.Local), time.Date(0, 0, 0, 14, 0, 0, 0, time.Local)},
            {time.Date(0, 0, 0, 13, 0, 0, 0, time.Local), time.Date(0, 0, 0, 22, 0, 0, 0, time.Local)},
            {time.Date(0, 0, 0, 22, 0, 0, 0, time.Local), time.Date(0, 0, 0, 7, 0, 0, 0, time.Local)},
		},
	}

	for _, sh := range Shifts.Shift {
		client.Shift.
			Create().
            SetTimeStart(sh.TimeStart).
            SetTimeEnd(sh.TimeEnd).
			Save(context.Background())
    }

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
            Save(context.Background())
   }

   zoneproducts := Zoneproducts{
    Zoneproduct: []Zoneproduct{
            Zoneproduct{"A"},
            Zoneproduct{"B"},
            Zoneproduct{"C"},
            Zoneproduct{"D"},
            },
   }

   for _, z := range zoneproduct.Zoneproduct {
        client.Zoneproduct.
            Create().
            SetZone(z.Zone).
            Save(context.Background())
   }

   managers := Managers{
    Manager: []Manager{
         Manager{"Dang Dang","Dang@gmail.com"},
         Manager{"AEK Dang","AEK@gmail.com"},
         Manager{"PANG Dang","PANG@gmail.com"},
         Manager{"NW Dang","NW@gmail.com"},
        },
   }

   for _, m := range manager.Manager {
        client.Manager.
            Create().
            SetName(m.Name).
            SetEmail(m.Email).
            Save(context.Background())
   }

   companys := Companys{
    Company: []Company{
            Company{"Dang Dang"},
            Company{"AEK Dang"},
            Company{"PANG Dang"},
            Company{"Noom Dang"},
        },
   }

   for _, cp := range company.Company {
        client.Company.
            Create().
            SetName(cp.Name).
            Save(context.Background())
   }


   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   router.Run()
}
