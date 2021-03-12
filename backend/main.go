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

type Customers struct {
	Customer []Customer
}

type Customer struct {
	Name     string
	Email    string
	Password string
	age      int
}

type Paymentchannels struct {
	Paymentchannel []Paymentchannel
}

type Paymentchannel struct {
	bank string
}

type Products struct {
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

type Companys struct {
	Company []Company
}

type Company struct {
	Name string
}

type Managers struct {
	Manager []Manager
}

type Manager struct {
	Name     string
	Email    string
	Password string
}

type Zoneproducts struct {
	Zoneproduct []Zoneproduct
}

type Zoneproduct struct {
	Zone string
}

type Days struct {
	Day []Day
}

type Day struct {
	Day string
}

type Employees struct {
	Employee []Employee
}

type Employee struct {
	Name     string
	Email    string
	Password string
	Age      int
}

type Roles struct {
	Role []Role
}

type Role struct {
	Role string
}

type BeginWorks struct {
	BeginWork []BeginWork
}

type BeginWork struct {
	BeginWork time.Time
}

type GetOffWorks struct {
	GetOffWork []GetOffWork
}

type GetOffWork struct {
	GetOffWork time.Time
}

type Assessments struct {
	Assessment []Assessment
}

type Assessment struct {
	AssessmentName string
}

type Positions struct {
	Position []Position
}

type Position struct {
	PositionName string
}

type Discounts struct {
	Discount []Discount
}

type Discount struct {
	Sale int
}

type Giveaways struct {
	Giveaway []Giveaway
}

type Giveaway struct {
	GiveawayName string
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

	client, err := ent.Open("sqlite3", "file:Database.db?&cache=shared&_fk=1")
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

	controllers.NewDayController(v1, client)
	controllers.NewEmployeeController(v1, client)
	controllers.NewEmployeeWorkingHoursController(v1, client)
	controllers.NewRoleController(v1, client)
	controllers.NewBeginWorkController(v1, client)
	controllers.NewGetOffWorkController(v1, client)

	controllers.NewAssessmentController(v1, client)
	controllers.NewPositionController(v1, client)
	controllers.NewSalaryController(v1, client)

	controllers.NewDiscountController(v1, client)
	controllers.NewGiveawayController(v1, client)
	controllers.NewPromotionController(v1, client)

	customers := Customers{
		Customer: []Customer{
			Customer{"AEK", "AEK@gmail.com", "123", 19},
			Customer{"KEN", "KEN@gmail.com", "1234", 20},
			Customer{"NAME", "NAME@gmail.com", "1235", 21},
			Customer{"TJ", "TJ@gmail.com", "1236", 22},
		},
	}

	for _, c := range customers.Customer {
		client.Customer.
			Create().
			SetName(c.Name).
			SetPassword(c.Password).
			SetAge(c.age).
			SetEmail(c.Email).
			Save(context.Background())
	}

	paymentchannels := Paymentchannels{
		Paymentchannel: []Paymentchannel{
			Paymentchannel{"KBANK"},
			Paymentchannel{"KTB"},
			Paymentchannel{"TMB"},
			Paymentchannel{"SCB"},
		},
	}

	for _, pay := range paymentchannels.Paymentchannel {
		client.Paymentchannel.
			Create().
			SetBank(pay.bank).
			Save(context.Background())
	}

	products := Products{
		Product: []Product{
			Product{"สบู่", "001", "01-01-2018", "01-01-2024"},
			Product{"แชมพู", "002", "02-01-2018", "02-01-2024"},
			Product{"อาหารหมา", "003", "03-01-2018", "03-01-2024"},
			Product{"อาหารแมว", "004", "04-01-2018", "04-01-2024"},
			Product{"เนื้อหมัก", "005", "05-01-2018", "05-01-2024"},
			Product{"ผักกาดขาว", "006", "06-01-2018", "06-01-2024"},
			Product{"สายชาร์จแบต", "007", "07-01-2018", "07-01-2024"},
			Product{"ปลั๊กไฟ", "008", "08-01-2018", "08-01-2024"},
			Product{"ขนมเลย์", "009", "09-01-2018", "09-01-2024"},
			Product{"ป๊อกกี้", "010", "10-01-2018", "10-01-2024"},
		},
	}

	for _, pd := range products.Product {
		client.Product.
			Create().
			SetNameProduct(pd.NameProduct).
			SetBarcodeProduct(pd.BarcodeProduct).
			SetMFG(pd.MFG).
			SetEXP(pd.Example).
			Save(context.Background())
	}

	typeproducts := Typeproducts{
		Typeproduct: []Typeproduct{
			Typeproduct{"ของใช้ทั่วไป"},
			Typeproduct{"อาหารสัตว"},
			Typeproduct{"ของสด"},
			Typeproduct{"เครื่องใช้ไฟฟ้า"},
			Typeproduct{"ขนมขบเคี้ยว"},

		},
	}

	for _, tp := range typeproducts.Typeproduct {
		client.Typeproduct.
			Create().
			SetTypeproduct(tp.Typeproduct).
			Save(context.Background())
	}

	zoneproducts := Zoneproducts{
		Zoneproduct: []Zoneproduct{
			Zoneproduct{"A"},
			Zoneproduct{"B"},
			Zoneproduct{"C"},
			Zoneproduct{"D"},
			Zoneproduct{"E"},
			Zoneproduct{"F"},
			Zoneproduct{"G"},
			Zoneproduct{"H"},
			Zoneproduct{"I"},
			Zoneproduct{"J"},
		},
	}

	for _, z := range zoneproducts.Zoneproduct {
		client.Zoneproduct.
			Create().
			SetZone(z.Zone).
			Save(context.Background())
	}

	managers := Managers{
		Manager: []Manager{
			Manager{"Dang Dang", "Dang@gmail.com", "456"},
			Manager{"AEK Dang", "AEK@gmail.com", "4567"},
			Manager{"PANG Dang", "PANG@gmail.com", "4568"},
			Manager{"NW Dang", "NW@gmail.com", "4569"},

			Manager{"Poommin Phinphimai", "poommin2543@gmail.com", "1234"},
			Manager{"Nontakorn Payusuk", "frashrock@gmail.com", "0862234089"},

		},
	}

	for _, m := range managers.Manager {
		client.Manager.
			Create().
			SetName(m.Name).
			SetEmail(m.Email).
			SetPassword(m.Password).
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

	for _, cp := range companys.Company {
		client.Company.
			Create().
			SetName(cp.Name).
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
			{"Poomin PhimPhimai", "Poomin123@gmail.com", "1123", 19},
			{"Pakiafa Kummungkun", "Pakiafa456@gmail.com", "11234", 20},
		},
	}

	for _, em := range Employees.Employee {
		client.Employee.
			Create().
			SetName(em.Name).
			SetEmail(em.Email).
			SetPassword(em.Password).
			SetAge(em.Age).
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

	BeginWorks := BeginWorks{
		BeginWork: []BeginWork{
			{time.Date(0, 0, 0, 6, 30, 0, 0, time.Local)},
			{time.Date(0, 0, 0, 13, 0, 0, 0, time.Local)},
			{time.Date(0, 0, 0, 22, 0, 0, 0, time.Local)},
		},
	}

	for _, bw := range BeginWorks.BeginWork {
		client.BeginWork.
			Create().
			SetBeginWork(bw.BeginWork).
			Save(context.Background())
	}

	GetOffWorks := GetOffWorks{
		GetOffWork: []GetOffWork{
			{time.Date(0, 0, 0, 14, 0, 0, 0, time.Local)},
			{time.Date(0, 0, 0, 22, 0, 0, 0, time.Local)},
			{time.Date(0, 0, 0, 7, 0, 0, 0, time.Local)},
		},
	}

	for _, bw := range GetOffWorks.GetOffWork {
		client.GetOffWork.
			Create().
			SetGetOffWork(bw.GetOffWork).
			Save(context.Background())
	}

	assessments := Assessments{
		Assessment: []Assessment{
			Assessment{"ยอดเยี่ยม"},
			Assessment{"ดีมาก"},
			Assessment{"ดี"},
			Assessment{"ปานกลาง"},
			Assessment{"ต่ำ"},
			Assessment{"ต่ำมาก"},
			Assessment{"ไม่ผ่านเกณฑ์"},
		},
	}

	for _, ass := range assessments.Assessment {
		client.Assessment.
			Create().
			SetAssessment(ass.AssessmentName).
			Save(context.Background())
	}

	positions := Positions{
		Position: []Position{
			Position{"พนักงานทั่วไป"},
			Position{"ผู้ช่วยผู้จัดการ"},
			Position{"ผู้จัดการ"},
		},
	}

	for _, po := range positions.Position {
		client.Position.
			Create().
			SetPosition(po.PositionName).
			Save(context.Background())
	}

	Discounts := Discounts{
		Discount: []Discount{
			{10},
			{50},
			{80},
		},
	}

	for _, d := range Discounts.Discount {
		client.Discount.
			Create().
			SetSale(d.Sale).
			Save(context.Background())
	}

	Giveaways := Giveaways{
		Giveaway: []Giveaway{
			{"ตุ๊กตาหมี"},
			{"ร่มแบบพกพา"},
			{"เก้าอี้"},
		},
	}

	for _, gc := range Giveaways.Giveaway {
		client.Giveaway.
			Create().
			SetGiveawayName(gc.GiveawayName).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
