package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	_ "./config/db"
	"./models"
	"github.com/joho/godotenv"
	cr "github.com/robfig/cron"
	"github.com/vjeantet/jodaTime"
	"github.com/yanzay/tbot"
)

type application struct {
	client *tbot.Client
}

var (
	app   application
	bot   *tbot.Server
	token string
)

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	token = os.Getenv("TELEGRAM_TOKEN")
	log.Println("token ", token)
}

func main() {
	// log.Println("Time Server Sql", models.GetDateSql())

	// log.Println("dads ", models.GetStateCeria())
	// m := tbot.Message{}
	bot = tbot.New(token)
	app.client = bot.Client()
	bot.HandleMessage("/startDaily", app.startHandler)
	bot.HandleMessage("/getData", app.getDataHandler)
	bot.HandleMessage("/Waterfall", app.getWaterfall)
	bot.HandleMessage("/getDataDaily", app.getDataDaily)
	bot.HandleMessage("/getAnomaly", app.getDataAnomaly)
	bot.HandleMessage("/getCeriaDaily", app.getDailyCeria)
	bot.HandleMessage("/getCeria", app.getDataCeria)
	bot.HandleMessage("/getDetailTransfer", app.getDetailTransfer)
	bot.HandleMessage("/getTestBriapi", app.getTestBriapi)
	bot.HandleMessage("/getFraud1", app.getCeriaFraud)
	bot.HandleMessage("/getFraud2", app.getCeriaFraud2)
	bot.HandleMessage("/Ceria", app.getCeriaWaterfall)
	bot.HandleMessage("/GdsUpdate", app.getUpdateGDS)
	bot.HandleMessage("/Report", app.getReportIngest)
	// bot.HandleMessage("/Transaction", app.getTransaction)
	// bot.HandleMessage("/DailyTransaction", app.DailtTransaction)
	// bot.HandleMessage(m.Text, app.getKKDReport)
	// log.Println(m.Text)
	// log.Println(app.getKKDReport)
	log.Println("asuu : ", bot.Start())
}

// func (app *application) DailtTransaction(m *tbot.Message) {
// 	msg := "Initiate Scheduler Execute Daily Transaction Ceria, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
// 	app.client.SendMessage(m.Chat.ID, msg)

// 	c := cr.New()
// 	c.AddFunc("@every 1h0m0s", func() { app.client.SendMessage(m.Chat.ID, models.GetTransaction()) })
// 	c.Start()

// }

// func (app *application) DailtTransaction(m *tbot.Message) {
// 	msg := "Initiate Scheduler Execute Daily Ceria, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
// 	app.client.SendMessage(m.Chat.ID, msg)

// 	c := cr.New()
// 	c.AddFunc("@every 0h60m0s", func() { app.client.SendMessage(m.Chat.ID, models.GetTransaction()) })
// 	c.Start()
// }

func (app *application) startHandler(m *tbot.Message) {
	msg := "Initiate Scheduler Execute Daily Pinang, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)

	c := cr.New()
	c.AddFunc("@every 1h0m0s", func() { app.client.SendMessage(m.Chat.ID, models.GetBodyMesage()) })
	c.Start()
}

// func (app *application) getKKDReport(m *tbot.Message) {
// 	getMessageSplit := strings.Split(m.Text, "/")
// 	log.Println(getMessageSplit)
// 	app.client.SendMessage(m.Chat.ID, models.GetKKD(getMessageSplit[2]))
// }

// func (app *application) getTransaction(m *tbot.Message) {
// 	msg := "Get Data Transaction, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
// 	app.client.SendMessage(m.Chat.ID, msg)
// 	app.client.SendMessage(m.Chat.ID, models.GetTransaction())
// }

func (app *application) getReportIngest(m *tbot.Message) {
	msg := "Get Report Ingest, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)
	app.client.SendMessage(m.Chat.ID, models.GetReportIngestFlag())
}

func (app *application) getUpdateGDS(m *tbot.Message) {
	msg := "Get Data EOD, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)
	app.client.SendMessage(m.Chat.ID, models.GetUpdateGDS())
}

func (app *application) getCeriaWaterfall(m *tbot.Message) {
	msg := "Get Data Waterfall, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)
	app.client.SendMessage(m.Chat.ID, models.GetCeriaWaterfall())
}

func (app *application) getCeriaFraud2(m *tbot.Message) {
	msg := "Get Data Fraud Ceria, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)
	app.client.SendMessage(m.Chat.ID, models.GetCeriaFraudExtend())
}

func (app *application) getCeriaFraud(m *tbot.Message) {
	msg := "Get Data Fraud Ceria, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)
	app.client.SendMessage(m.Chat.ID, models.GetCeriaFraud())
}

func (app *application) getTestBriapi(m *tbot.Message) {
	msg := "Get Data Test BRIAPI, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)
	app.client.SendMessage(m.Chat.ID, models.GetBodyMessageBriapi())
}

func (app *application) getDataHandler(m *tbot.Message) {
	msg := "Get Data Daily Pinang, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)
	app.client.SendMessage(m.Chat.ID, models.GetBodyMesage())
}

func (app *application) getWaterfall(m *tbot.Message) {
	msg := "Get Data Daily Pinang, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)
	app.client.SendMessage(m.Chat.ID, models.GetWaterfall())
}

func (app *application) getDataAnomaly(m *tbot.Message) {
	msg := "Get Data Anomaly Pinang, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)
	app.client.SendMessage(m.Chat.ID, models.GetBodyMesageAnomaly())
}

func (app *application) getDailyCeria(m *tbot.Message) {
	msg := "Initiate Scheduler Execute Daily Ceria, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)

	c := cr.New()
	c.AddFunc("@every 1h0m0s", func() { app.client.SendMessage(m.Chat.ID, models.GetBodyMesageCeria()) })
	c.Start()
}

func (app *application) getDataDaily(m *tbot.Message) {
	msg := "Initiate Scheduler Execute Daily Ceria, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)

	c := cr.New()
	c.AddFunc("@every 1h0m0s", func() { app.client.SendMessage(m.Chat.ID, models.GetBodyMesage()) })
	c.Start()
}

func (app *application) getDataCeria(m *tbot.Message) {
	msg := "Get Data Ceria, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)
	app.client.SendMessage(m.Chat.ID, models.GetBodyMesageCeria())
}

func (app *application) getDetailTransfer(m *tbot.Message) {
	msg := "Get Detail Transfer Type Pinang, " + jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now())
	app.client.SendMessage(m.Chat.ID, msg)
	app.client.SendMessage(m.Chat.ID, models.GetBodyMesageDetilTransferType())
}

func randomId(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func GenerateRandomId() string {
	myrand := randomId(100000000000, 999999999999)

	seqNum := strconv.Itoa(myrand)
	return seqNum
}
