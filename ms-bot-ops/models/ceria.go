package models

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/bridce/ms-bot-ops/config/db"
	"github.com/leekchan/accounting"
	"github.com/vjeantet/jodaTime"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetReportTbaadmIngest() string {
	var listRes []Results
	// var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	// ac := accounting.Accounting{Symbol: "Rp.", Precision: 2}
	rows, err := db.DbCerSta.Table("tbaadm.ingest_flag").Select("table_name as Key,result as Description2").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Description2)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s ", "\n"+res.Key, res.Description2)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	//log.Println("asdsd ", aa)

	return aa
}

func GetReportCrmuserIngest() string {
	var listRes []Results
	// var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	// ac := accounting.Accounting{Symbol: "Rp.", Precision: 2}
	rows, err := db.DbCerSta.Table("crmuser.ingest_flag").Select("table_name as Key,result as Description2").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Description2)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s ", "\n"+res.Key, res.Description2)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	//log.Println("asdsd ", aa)

	return aa
}

func GetReportCustomIngest() string {
	var listRes []Results
	// var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	// ac := accounting.Accounting{Symbol: "Rp.", Precision: 2}
	rows, err := db.DbCerSta.Table("custom.ingest_flag").Select("table_name as Key,result as Description2").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Description2)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s ", "\n"+res.Key, res.Description2)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	//log.Println("asdsd ", aa)

	return aa
}

func GetReportPublicIngest() string {
	var listRes []Results
	// var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	// ac := accounting.Accounting{Symbol: "Rp.", Precision: 2}
	rows, err := db.DbCerSta.Table("public.ingest_flag").Select("table_name as Key,result as Description2").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Description2)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s |  %s ", "\n"+res.Key, res.Description2)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	//log.Println("asdsd ", aa)

	return aa
}

func GetCashoutBriAll() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	// ac := accounting.Accounting{Symbol: "Rp.", Precision: 2}
	rows, err := db.DbCerCore.Table("custom.fundtrf_sts_tbl").Select("sum(amount) as Key,count(loan_account) as Value").Where("destination_of_fund != '211401000307302' and status='S'and transfer_type='CashOut' and transaction_date_time like '" + date + "%'").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Value)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s", "\n"+res.Key, res.Value)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	//log.Println("asdsd ", aa)

	return aa
}

func getTime() string {

	var listRes []Results

	rows, err := db.DbCerCore.Table("(select split_part(split_part(now()+ '11 hour'::interval,' ',2),':',1) as time)tb ").Select("tb.time as Key").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
		)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s", "\n"+res.Key)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")
	log.Println(aa)

	return aa
}

func getAutodebetStatus() Results {
	var total Results
	err := db.DbCerCore.Select("status as Key , continuous_pending_flg as Description").Table("custom.cer_autodb_status").Scan(&total)
	log.Println(total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func getDebitur() Results {
	// var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	var total Results
	err := db.DbCerCore.Table("tbaadm.c_robt").Select("count(*) as Value64").Where("bill_amt != '0' and bill_amt != tot_adj_amt and payment_due_date='2021-01-25'").Scan(&total)

	if err != nil {
		log.Println(err)
	}
	log.Println("crobt", total)

	return total
}

func EodDisbursement() Results {
	var total Results
	err := db.DbCerCore.Table("tbaadm.lam").Select("sum(dis_amt) as Key").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodDailyDisbursement() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.lam").Select("sum(dis_amt) as Key").Where("DIS_SHDL_DATE ='" + yesterday + "%'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalLoanPayment() Results {
	var total Results
	err := db.DbCerCore.Table("(select sum(flow_amt) amt from tbaadm.ltd where flow_id in ('COLL','PDCOL') and entity_cre_flg = 'Y' and del_flg = 'N' and reversal_flg != 'Y' and reversed_flg != 'Y' union select sum(flow_amt) amt from tbaadm.ltdh where flow_id in ('COLL','PDCOL') and entity_cre_flg = 'Y' and del_flg = 'N' and reversal_flg != 'Y' and reversed_flg != 'Y')").Select("sum(amt) as Key").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodDailyLoanPayment() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("(select sum(flow_amt) amt from tbaadm.ltd where flow_id in ('COLL','PDCOL') and entity_cre_flg = 'Y' and del_flg = 'N' and reversal_flg != 'Y' and reversed_flg != 'Y' and tran_date ='" + yesterday + "' union select sum(flow_amt) amt from tbaadm.ltdh where flow_id in ('COLL','PDCOL') and entity_cre_flg = 'Y' and del_flg = 'N' and reversal_flg != 'Y' and reversed_flg != 'Y' and tran_date = '" + yesterday + "')").Select("sum(amt) as Key").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalOSNPL() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.eod_acct_bal_table").Select("nvl(sum(abs(tran_date_bal)),0) as Key").Where("acid in (select acid from tbaadm.general_acct_mast_table where schm_type = 'ODA' and entity_cre_flg = 'Y' and del_flg != 'Y' and acid in (select acid from custom.c_rom where accl_flg = 'Y')) and acid in (select b2k_id from tbaadm.acd where main_classification_user in ('003','004','005') and entity_cre_flg = 'Y' and del_flg != 'Y')and '" + yesterday + "' between eod_date and end_eod_date and tran_date_bal < 0").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalOSPL() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.eab").Select("SUM(abs(TRAN_DATE_BAL))as Key").Where("END_EOD_DATE >= '" + yesterday + "' AND EOD_DATE <= '" + yesterday + "' AND TRAN_DATE_BAL < 0 AND ACID in (SELECT ACID FROM TBAADM.GAM WHERE SCHM_TYPE IN ('LAA'))").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalDPKUser() Results {
	var total Results
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.acd").Select("COUNT(1) as Key").Where("main_classification_user = '002' AND entity_cre_flg = 'Y' AND del_flg != 'Y' and b2k_id in (select acid from tbaadm.gam where acct_cls_flg != 'Y' and schm_type = 'ODA')").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalDPKSystem() Results {
	var total Results
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.acd").Select("COUNT(1) as Key").Where("main_classification_system = '002' AND entity_cre_flg = 'Y' AND del_flg != 'Y' and b2k_id in (select acid from tbaadm.gam where acct_cls_flg != 'Y' and schm_type = 'ODA')").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalNPLUser() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.gam").Select("COUNT(1) as Key").Where("schm_type = 'ODA' and clr_bal_amt < 0 and entity_cre_flg = 'Y' and del_flg != 'Y' and acid in (select acid from custom.c_rom where accl_flg = 'Y' and accl_date <='" + yesterday + "')").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalNPLSystem() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.gam").Select("COUNT(1) as Key").Where("schm_type = 'ODA' and clr_bal_amt < 0 and entity_cre_flg = 'Y' and del_flg != 'Y' and acid in (select acid from custom.c_rom where accl_flg = 'Y' and accl_date <='" + yesterday + "')").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalAmountDPKUser() Results {
	var total Results
	err := db.DbCerCore.Table("tbaadm.gam").Select("nvl(sum(abs(clr_bal_amt)),0) as Key").Where("schm_type = 'LAA' and clr_bal_amt < 0 and entity_cre_flg = 'Y' and del_flg != 'Y' and acid in (SELECT b2k_id FROM TBAADM.ACD WHERE main_classification_user = '002' AND entity_cre_flg = 'Y' AND del_flg != 'Y')").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalAmountDPKSystem() Results {
	var total Results
	err := db.DbCerCore.Table("tbaadm.gam").Select("nvl(sum(abs(clr_bal_amt)),0) as Key").Where("schm_type = 'LAA' and clr_bal_amt < 0 and entity_cre_flg = 'Y' and del_flg != 'Y' and acid in (SELECT b2k_id FROM TBAADM.ACD WHERE main_classification_system = '002' AND entity_cre_flg = 'Y' AND del_flg != 'Y')").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalAmountNPLUser() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.gam").Select("nvl(sum(abs(clr_bal_amt)),0) as Key").Where("schm_type = 'ODA' and clr_bal_amt < 0 and entity_cre_flg = 'Y' and del_flg != 'Y' and acid in (select acid from custom.c_rom where accl_flg = 'Y' and accl_date <= '" + yesterday + "')").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalAmountNPLSystem() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.gam").Select("nvl(sum(abs(clr_bal_amt)),0) as Key").Where("schm_type = 'ODA' and clr_bal_amt < 0 and entity_cre_flg = 'Y' and del_flg != 'Y' and acid in (select acid from custom.c_rom where accl_flg = 'Y' and accl_date <= '" + yesterday + "')").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalLAA() Results {
	var total Results
	err := db.DbCerCore.Table("tbaadm.gam").Select("count(1) as Key").Where("schm_type='LAA'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalODA() Results {
	var total Results
	err := db.DbCerCore.Table("tbaadm.gam").Select("count(1) as Key").Where("schm_type='ODA' and acct_cls_flg = 'N'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodCloseODA() Results {
	var total Results
	err := db.DbCerCore.Table("tbaadm.gam").Select("COUNT(1) as Key").Where("schm_type = 'ODA' and ACCT_CLS_FLG = 'Y'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalLimit() Results {
	var total Results
	err := db.DbCerCore.Table("tbaadm.gam").Select("sum(sanct_lim) as Key").Where("schm_type = 'ODA' and ACCT_CLS_FLG = 'N'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodDailyLimit() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.gam").Select("sum(sanct_lim) as Key").Where("schm_type = 'ODA' and acct_opn_date = '" + yesterday + "' ").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodDailyLAA() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.gam").Select("count(1) as Key").Where("schm_type = 'LAA' and acct_opn_date = '" + yesterday + "' ").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodDailyODA() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.gam").Select("count(1) as Key").Where("schm_type = 'ODA' and acct_opn_date = '" + yesterday + "' ").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodTotalCashout() Results {
	var total Results
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("custom.cmpdt").Select("SUM(PURCHASE_AMT-REVERSAL_AMT) as Key").Where("PURCHASE_AMT - REVERSAL_AMT > 0 AND CASH_OUT_FLG = 'Y'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodDailyAmountCashout() string {
	var listRes []Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	rows, err := db.DbCerCore.Table("custom.cmpdt").Select("CASH_OUT_STATUS as Key, SUM(PURCHASE_AMT) as Value").Where("CASH_OUT_FLG = 'Y' AND POSTING_DATE = '" + yesterday + "'").Group("CASH_OUT_STATUS").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Value)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s", "\n"+res.Key, res.Value)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	return aa
}

func EodDailyDebiturCashout() string {
	var listRes []Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	rows, err := db.DbCerCore.Table("custom.cmpdt").Select("CASH_OUT_STATUS as Key, count(1) as Value").Where("CASH_OUT_FLG = 'Y' AND POSTING_DATE = '" + yesterday + "'").Group("CASH_OUT_STATUS").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Value)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s", "\n"+res.Key, res.Value)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	return aa
}

func EodTotalPurchase() Results {
	var total Results
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("custom.cmpdt C LEFT JOIN (SELECT MERCHANT_ID,MERCHANT_REF_ID,SUM(REFUND_AMT) REFUND_AMT FROM CUSTOM.CRDT GROUP BY MERCHANT_ID,MERCHANT_REF_ID) X ON C.MERCHANT_ID=X.MERCHANT_ID AND C.MERCHANT_REF_ID=X.MERCHANT_REF_ID").Select("SUM(C.PURCHASE_AMT - NVL(X.REFUND_AMT,0)) as Key").Where("C.PURCHASE_AMT - NVL(X.REFUND_AMT,0) > 0 AND CASH_OUT_FLG = 'N'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodDailyAmountPurchaseV1() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("custom.cmpdt").Select("SUM(PURCHASE_AMT) as Key").Where("CASH_OUT_FLG = 'N' AND POSTING_DATE = '" + yesterday + "'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodDailyAmountPurchaseV2() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("custom.cust_mer_pch_dtls_tbl").Select("(sum(purchase_amt)-sum(reversal_amt)) as Key").Where("cash_out_flg ='N' and loan_created = 'Y' and  posting_date = '" + yesterday + "' ").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func EodDailyDebiturPurchase() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("custom.cmpdt").Select("COUNT(1) as Key").Where("CASH_OUT_FLG = 'N' and POSTING_DATE = '" + yesterday + "'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

// daily limit
func GetDailyLimit() Results {
	var total Results
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.gam").Select("sum(sanct_lim) as Value64").Where("schm_type='ODA' and rcre_time between '" + yesterday + " 13:00:00' and '" + currDate + " 12:59:59'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

// hold pending
func GetHoldPending() Results {
	var total Results
	err := db.DbCerCore.Table("custom.hstl").Select("count(loan_account) as Key").Where("status= 'P'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
	// log.Println(total.ListTagih)
}

func GetHoldFailed() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	err := db.DbCerCore.Table("custom.hstl").Select("count(loan_account) as Key").Where("transfer_type = 'HoldAmt' and status= 'F' and rcre_time between '" + yesterday + " 13:00:00' and '" + currDate + " 12:59:59'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
	// log.Println(total.ListTagih)
}

func GetHoldSuccess() Results {
	var total Results
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	err := db.DbCerCore.Table("custom.hstl").Select("count(loan_account) as Key").Where("transfer_type = 'HoldAmt' and status= 'S' and rcre_time between '" + yesterday + " 13:00:00' and '" + currDate + " 12:59:59'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
	// log.Println(total.ListTagih)
}

// briva pending
func GetBrivaPending() Results {
	var total Results
	err := db.DbCerCore.Table("custom.fundtrf_sts_tbl").Select("count(loan_account) as Key").Where("transfer_type ='BRIVAREPAY' and status ='P'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

// oda positif
func getOdaPositive() Results {
	var total Results
	err := db.DbCerCore.Table("tbaadm.gam a join custom.c_rom b on a.acid = b.acid").Select("count(foracid) as Key").Where("clr_bal_amt > 0 and schm_type = 'ODA' and schm_type = 'ODA' and accl_flg != 'Y'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

//oda core
func getOdaCore() Results {
	var total Results
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbCerCore.Table("tbaadm.gam").Select("count(foracid) as Key").Where("rcre_time between '" + yesterday + " 13:00:00' and '" + currDate + " 12:59:59' and schm_type='ODA'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

//No 3 and 5
func GetNumberAndVolCeriaTrx() Results {
	// log.Println("No.3 and 5")
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	var total Results
	err := db.DbCerCore.Table("(select \n\tsum(tot_purchase) - sum(tot_refund) as Key,\n\tcase when sum(refund) is null then sum(purchase) - 0 else sum(purchase) - sum(refund) end as Value64 from custom.cust_mer_profit_shr_table m \nleft join \n\t(select\n\t\ta.merchant_id, \n\t\tcount(*) as tot_purchase, \n\t\tsum(purchase_amt) as purchase \n\t from custom.cmpdt a left join custom.crdt b on a.merchant_ref_id = b.merchant_ref_id where cash_out_flg = 'N' and posting_date = '" + currDate + " 00:00:00'  \n\t group by a.merchant_id) as p\non m.merchant_id=p.merchant_id\nleft join \n\t(select\n\t\tmerchant_id, \n\t\tcount(*) as tot_refund, \n\t\tsum(refund_amt) as refund \n\t from custom.crdt where \n\t merchant_ref_id in (select merchant_ref_id from custom.cmpdt where posting_date = '" + currDate + " 00:00:00') \n\t group by merchant_id) as r\non m.merchant_id=r.merchant_id)").Select("Key,Value64").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total
}

func GetCashoutLinkAja() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	// ac := accounting.Accounting{Symbol: "Rp.", Precision: 2}
	rows, err := db.DbCerCore.Table("custom.fundtrf_sts_tbl").Select("sum(amount) as Key,count(loan_account) as Value").Where("destination_of_fund = '211401000307302' and status='S'and transfer_type='CashOut' and transaction_date_time like '" + date + "%'").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key, &res.Value)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s", "\n"+res.Key, res.Value)
		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	//log.Println("asdsd ", aa)

	return aa
}

func GetStatusCashOut() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())

	rows, err := db.DbCerCore.Table("custom.cmpdt").Select("cash_out_status as Key, count(merchant_ref_id) as Description").Where("posting_date = '" + date + "%'and cash_out_status !='N'").Group("Key").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Description,
		)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s ", "\n"+res.Key, res.Description)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	//log.Println("asdsd ", aa)

	return aa
}

func GetReportAutoDebetCeria() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	ac := accounting.Accounting{Symbol: "Rp.", Precision: 2}
	rows, err := db.DbCerCore.Table("custom.fundtrf_sts_tbl").Select("transfer_type as Key, status as Value, status_code as Description, count(status_code) as Description2, sum(amount) as Value64").Where("transaction_date_time like ('" + date + "%')").Group("Key,Value,Description").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Value,
			&res.Description,
			&res.Description2,
			&res.Value64)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s | %s | %s | %s", "\n"+res.Key, res.Value, res.Description, res.Description2, ac.FormatMoney(res.Value64))

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	//log.Println("asdsd ", aa)

	return aa
}

func GetReportDataPerBlok() string {
	var listRes []Results
	rows, err := db.DbCerCore.Table("(\n\tselect \n\t\ta.cif_id,\n\t\ta.foracid as nomor_oda, \n\t\tc.free_text_1 as nomor_rekening_bri,\n\t\ta.acct_name as nama_nasabah,\n\t\tc.free_code_1 as block,\n\t\tb.main_classification_system as collectibility,\n\t\tc.dpd_cntr as tunggakan_hari,\n\t\ta.schm_type\n\tfrom tbaadm.general_acct_mast_table a\n\t\tjoin tbaadm.asst_class_detail_tbl b on a.acid=b.b2k_id\n\t\tjoin tbaadm.gen_acct_class_table c on b.b2k_id=c.acid\n\twhere\n\t\ta.schm_type='ODA') tb").Select("(case when block is null then 'NULL' \n\twhen block is not null then block\n\tend )as Key,\n\tcount(nomor_oda) as Value").Group("tb.block").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Value)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s ", "\n"+res.Key, res.Value)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	//log.Println("asdsd ", aa)

	return aa
}

func GetReportDataMerchant() string {
	var listRes []Results

	ac := accounting.Accounting{Symbol: "Rp.", Precision: 2}
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	rows, err := db.DbCerCore.Table("(select \n\tmerchant as Key, \n\tcoalesce (tot_purchase, 0) - coalesce (tot_refund, 0) as Value,\n\tcoalesce (purchase, 0) - coalesce (refund, 0) as Value64 \nfrom (select \n\t\tm.merchant_id, \n\t\tm.free_text1 as merchant, \n\t\tpurchase, refund, tot_purchase, tot_refund\n\tfrom custom.cust_mer_profit_shr_table m \nleft join \n\t(select\n\t\ta.merchant_id, \n\t\tcount(*) as tot_purchase, \n\t\tsum(purchase_amt) as purchase \n\t from custom.cmpdt a left join custom.crdt b on a.merchant_ref_id = b.merchant_ref_id where cash_out_flg='N' and posting_date = '" + currDate + " 00:00:00'  group by a.merchant_id) as p\non m.merchant_id=p.merchant_id\nleft join \n\t(select\n\t\tmerchant_id, \n\t\tcount(*) as tot_refund, \n\t\tsum(refund_amt) as refund \n\t from custom.crdt where merchant_ref_id in (select merchant_ref_id from custom.cmpdt where posting_date = '" + currDate + " 00:00:00') group by merchant_id) as r\non m.merchant_id=r.merchant_id)) where KEY != 'CASHOUT BRI'").Select("Key, Value, Value64").Order(" Value64 desc").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Value,
			&res.Value64)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s  | %s ", "\n"+res.Key, res.Value, ac.FormatMoney(res.Value64))

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	// log.Println("nomer 8 ", aa)

	return aa
}

func GetTransactionAnomali() string {
	var listRes []Results

	ac := accounting.Accounting{Symbol: "Rp.", Precision: 2}
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	rows, err := db.DbCerCore.Table("(select a.loc_acct_id as Key,sum(purchase_amt) as Value64,count(a.merchant_id) as Description from custom.cmpdt a left join custom.crdt b on a.merchant_ref_id = b.merchant_ref_id where b.merchant_ref_id is null and reversal_flag = 'N'  and posting_date = '" + currDate + "%' group by a.loc_acct_id ) tb").Select("Key,Value64,Description").Where("tb.Description >= 3").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Value64,
			&res.Description)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s  | %s ", "\n"+res.Key, ac.FormatMoney(res.Value64), res.Description)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	// log.Println("purchase lebih dari 3 kali", aa)

	return aa
}

func GetCashoutAnomali() string {
	var listRes []Results

	// ac := accounting.Accounting{Symbol: "Rp.", Precision: 2}
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	rows, err := db.DbCerCore.Table("(select foracid as Key,b.free_text_1 as Value,c.dest_acct_no as Description from tbaadm.gam a join tbaadm.gac b on a.acid=b.acid join custom.cmpdt c on a.foracid = c.loc_acct_id where posting_date = '" + currDate + "%' and cash_out_flg = 'Y' and cash_out_status in ('S','C','L') and merchant_id ='14017' and b.free_text_1 != c.dest_acct_no order by foracid desc) tb").Select("Key,Value,Description").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		res1 := strings.TrimLeft(res.Description, "0")
		if res.Key == res1 {
			log.Println("sama semua")
		} else {
			log.Println("ada beda")
		}
		rows.Scan(
			&res.Key,
			&res.Value,
			&res1)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s  | %s ", "\n"+res.Key, res.Value, res1)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	log.Println("", aa)

	return aa
}

func GetTransactioBig() string {
	var listRes []Results

	ac := accounting.Accounting{Symbol: "", Precision: 2}
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	rows, err := db.DbCerCore.Table("(select tb.oda as Key ,tb.jumlah as Value,b.sanct_lim as Description,((jumlah/b.sanct_lim)*100) as Value64 from (select a.loc_acct_id as oda,sum(purchase_amt) as jumlah from custom.cmpdt a left join custom.crdt b on a.merchant_ref_id = b.merchant_ref_id where posting_date = '" + currDate + "%' and b.merchant_ref_id is null and cash_out_flg='N' group by a.loc_Acct_id  having sum(purchase_amt) >='2000000') tb join tbaadm.gam b on tb.oda = b.foracid order by foracid desc )").Select("Key,Value,Description,Value64").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Key,
			&res.Value,
			&res.Description,
			&res.Value64)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s  | %s | %s", "\n"+res.Key, res.Value, res.Description, ac.FormatMoney(res.Value64))

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	// log.Println("nomer 8 ", aa)

	return aa
}

func GetTransactioBlokR() Results {

	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	var total Results
	err := db.DbCerCore.Table("(select foracid from tbaadm.gam join tbaadm.gac on gam.acid=gac.acid where free_code_2='R') tb join custom.cmpdt c on tb.foracid = c.loc_acct_id where posting_date = '" + currDate + "%'").Select("COALESCE(count(merchant_ref_id),0) as Key").Scan(&total)

	if err != nil {
		log.Println(err)
	}
	log.Println("Blok R : ", total.Key)

	return total
}

func GetTransactioBlokN() Results {

	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	var total Results
	err := db.DbCerCore.Table("(select foracid from tbaadm.gam join tbaadm.gac on gam.acid=gac.acid where free_code_2='N') tb join custom.cmpdt c on tb.foracid = c.loc_acct_id where posting_date = '" + currDate + "%'").Select("COALESCE(count(merchant_ref_id),0) as Key").Scan(&total)

	if err != nil {
		log.Println(err)
	}
	log.Println("Blok R : ", total.Key)

	return total
}

func GetCashoutMoreThanRule() Results {

	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	var total Results
	err := db.DbCerCore.Table("(select tb.*,sanct_lim,(sanct_lim*0.3)as max_cashout from (select foracid,sum(purchase_amt) as jumlah_cashout,count(merchant_ref_id) from tbaadm.gam a join tbaadm.gac b on a.acid=b.acid join custom.cmpdt c on a.foracid = c.loc_acct_id where posting_date between '2021-05-01%' and '" + currDate + "%' and cash_out_flg = 'Y' and cash_out_status in ('S','C') group by foracid) tb join tbaadm.gam d on tb.foracid = d.foracid) tc").Select("count(tc.foracid) as Key").Where("tc.max_cashout < tc.jumlah_cashout").Scan(&total)

	if err != nil {
		log.Println(err)
	}
	log.Println("Blok R : ", total.Key)

	return total
}

func GetCashoutMoreThanRule2() Results {

	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	// var firstOfMonth = jodaTime.Format("YYYY-MM-dd", time.Now())

	var total Results
	err := db.DbCerCore.Table("(select tb.*,sanct_lim,(sanct_lim*0.3)as max_cashout from (select foracid,sum(purchase_amt) as jumlah_cashout,count(merchant_ref_id) from tbaadm.gam a join tbaadm.gac b on a.acid=b.acid join custom.cmpdt c on a.foracid = c.loc_acct_id where posting_date ='" + currDate + "%' and cash_out_flg = 'Y' and cash_out_status in ('S','C') group by foracid) tb join tbaadm.gam d on tb.foracid = d.foracid) tc").Where("tc.max_cashout < tc.jumlah_cashout").Select("count(tc.foracid) as Key").Scan(&total)

	if err != nil {
		log.Println(err)
	}
	log.Println("Blok R : ", total.Key)

	return total
}

func collectionLos() *mgo.Collection {
	return db.DbMongo.C("loan_origination_system")
}

func collectionUseEkyc() *mgo.Collection {
	return db.DbMongo.C("user_ekyc")
}

func collectionCharges() *mgo.Collection {
	return db.DbMongo.C("charges")
}

func collectionBriva() *mgo.Collection {
	return db.DbMongo.C("briva_payment")
}

func collectionLoanApp() *mgo.Collection {
	return db.DbMongo.C("loan_application")
}
func collectionRegistrations() *mgo.Collection {
	return db.DbMongo.C("registrations")
}

func collectionLogin() *mgo.Collection {
	return db.DbMongo.C("verify_login")
}

func GetWelcomeOnly() string {
	var listResult []string
	SuccessCreateWelcomePackage, _ := GetCountSuccessCreateWelcome()
	listResult = append(listResult, "Success Create Welcome Package  : "+strconv.Itoa(SuccessCreateWelcomePackage))

	aa := strings.Join(listResult, "\n")
	log.Println(aa)

	return aa

}

// func getWhitelistReject() string {
// 	var listResult []string
// 	whitelistReject, _ := getCountRejectWhitelist()
// 	listResult = append(listResult, "Whitelist Reject : "+strconv.Itoa(whitelistReject))

// 	aa := strings.Join(listResult, "\n")
// 	log.Println(aa)

// 	return aa

// }
func GetCekODAOch() string {
	var listResult []string
	SuccessCreateWelcomePackage, _ := GetCountSuccessCreateWelcome()
	listResult = append(listResult, "ODA Och  : "+strconv.Itoa(SuccessCreateWelcomePackage))

	aa := strings.Join(listResult, "\n")
	log.Println(aa)

	return aa
}

func GetBriva() string {
	var listResult []string
	briva, _ := GetBrivaOCH()
	listResult = append(listResult, "Briva OCH  : "+strconv.Itoa(briva))

	aa := strings.Join(listResult, "\n")
	log.Println(aa)

	return aa
}

func GetStatusOchTimeOut() string {
	var listResult []string

	listResult = append(listResult, "Breakdown Status CashOut di OCH :")

	StatusOchTimeOut, _ := GetStatusTimeout()
	listResult = append(listResult, "TIMEOUT : "+strconv.Itoa(StatusOchTimeOut))

	StatusOchInProgress, _ := GetStatusInProgress()
	listResult = append(listResult, "IN_PROGRESS : "+strconv.Itoa(StatusOchInProgress))

	StatusOchSuccess, _ := GetStatusSuccess()
	listResult = append(listResult, "SUCCESS : "+strconv.Itoa(StatusOchSuccess))

	StatusOchFailed, _ := GetStatusFailed()
	listResult = append(listResult, "FAILED : "+strconv.Itoa(StatusOchFailed))

	aa := strings.Join(listResult, "\n")
	log.Println(aa)

	return aa

}

func GetStateCeria() string {
	var listResult []string

	// registered, _ := GetCountRegistered()
	// listResult = append(listResult, "\nRegistered = "+strconv.Itoa(registered))

	// login, _ := GetCountLogin()
	// listResult = append(listResult, "Login = "+strconv.Itoa(login))

	notYet, _ := GetCountNotYet()
	listResult = append(listResult, "NOT_YET = "+strconv.Itoa(notYet))

	inProg, _ := GetCountInProgress()
	listResult = append(listResult, "In Progress = "+strconv.Itoa(inProg))

	verifKtp, _ := GetCountVerifKtp()
	listResult = append(listResult, "KTP verification process = "+strconv.Itoa(verifKtp))

	// notveryktp, _ := GetCountNotVerifKtp()
	// listResult = append(listResult, "KTP not verify process = "+strconv.Itoa(notveryktp))

	listResult = append(listResult, "=============================")
	listResult = append(listResult, "Prosses scoring")
	listResult = append(listResult, "=============================")

	pendCs, _ := GetCountPendingCreditScoring()
	listResult = append(listResult, "Pending Credit Scoring = "+strconv.Itoa(pendCs))

	waitVerEm, _ := GetCountWaitingEmailVerif()
	listResult = append(listResult, "Waiting For email verification = "+strconv.Itoa(waitVerEm))

	appReje, _ := GetCountAppReject()
	listResult = append(listResult, "\nUnique Application Rejected = "+strconv.Itoa(appReje))

	listResult = append(listResult, "=============================")
	listResult = append(listResult, "Break Down Unique Application Reject")
	listResult = append(listResult, "=============================")

	grade8re, _ := GetGradeRej8()
	listResult = append(listResult, "grade 8  = "+strconv.Itoa(grade8re))

	grade9re, _ := GetGradeRej9()
	listResult = append(listResult, "grade 9  = "+strconv.Itoa(grade9re))

	grade10re, _ := GetGradeRej10()
	listResult = append(listResult, "grade 10  = "+strconv.Itoa(grade10re))

	grademin1, _ := GetGrademin1()
	listResult = append(listResult, "grade -1 = "+strconv.Itoa(grademin1))

	listResult = append(listResult, "=============================")

	pendOffer, _ := GetCountPendOffer()
	listResult = append(listResult, "\nPending Offer = "+strconv.Itoa(pendOffer))

	offerRejec, _ := GetCountOfferReje()
	listResult = append(listResult, "Offer Rejected = "+strconv.Itoa(offerRejec))

	offerAccepted, _ := GetCountOfferAccep()
	listResult = append(listResult, "Offer Accepted  = "+strconv.Itoa(offerAccepted))

	selfieNotVerified, _ := GetCountSelfieNot()
	listResult = append(listResult, "Selfie Not Verified  = "+strconv.Itoa(selfieNotVerified))

	selfieVerified, _ := GetCountSelfie()
	listResult = append(listResult, "Selfie Verified  = "+strconv.Itoa(selfieVerified))

	privyRegistered, _ := GetCountPrivyRegistered()
	listResult = append(listResult, "Privy Registered  = "+strconv.Itoa(privyRegistered))

	privySendOtp, _ := GetCountPrivySendOtp()
	listResult = append(listResult, "Privy Send OTP  = "+strconv.Itoa(privySendOtp))

	privySuccessOTP, _ := GetCountPrivySuccessOtp()
	listResult = append(listResult, "Privy Success OTP  = "+strconv.Itoa(privySuccessOTP))

	SuccessFaceCompare, _ := GetCountSuccessFaceCompare()
	listResult = append(listResult, "Success Face Compare  = "+strconv.Itoa(SuccessFaceCompare))

	SuccessSignDocument, _ := GetCountSuccessSignDoc()
	listResult = append(listResult, "Success Sign Document  = "+strconv.Itoa(SuccessSignDocument))

	SuccessCreateWelcomePackage, _ := GetCountSuccessCreateWelcome()
	listResult = append(listResult, "Success Create Welcome Package  = "+strconv.Itoa(SuccessCreateWelcomePackage))

	listResult = append(listResult, "=============================")
	listResult = append(listResult, "Break Down Success Create Welcome Package")
	listResult = append(listResult, "=============================")

	WelcomaPackageGrade1, _ := GetWelcomeGrade1()
	listResult = append(listResult, "grade 1 = "+strconv.Itoa(WelcomaPackageGrade1))

	WelcomaPackageGrade2, _ := GetWelcomeGrade2()
	listResult = append(listResult, "grade 2 = "+strconv.Itoa(WelcomaPackageGrade2))

	WelcomaPackageGrade3, _ := GetWelcomeGrade3()
	listResult = append(listResult, "grade 3 = "+strconv.Itoa(WelcomaPackageGrade3))

	WelcomaPackageGrade4, _ := GetWelcomeGrade4()
	listResult = append(listResult, "grade 4 = "+strconv.Itoa(WelcomaPackageGrade4))

	grade5, _ := GetGrade5()
	listResult = append(listResult, "grade 5 = "+strconv.Itoa(grade5))

	grade6, _ := GetGrade6()
	listResult = append(listResult, "grade 6 = "+strconv.Itoa(grade6))

	grade7, _ := GetGrade7()
	listResult = append(listResult, "grade 7 = "+strconv.Itoa(grade7))

	grade8, _ := GetGrade8()
	listResult = append(listResult, "grade 8 = "+strconv.Itoa(grade8))

	grade9, _ := GetGrade9()
	listResult = append(listResult, "grade 9 = "+strconv.Itoa(grade9))

	listResult = append(listResult, "=============================")

	override, _ := ApproveOverride()
	listResult = append(listResult, "Approve Override	= "+strconv.Itoa(override))

	//gagal, _ := GetCountVerifKtp()
	//listResult = append(listResult, "gagal  = "+strconv.Itoa(gagal))

	aa := strings.Join(listResult, "\n")
	log.Println(aa)

	return aa
}

func GetBrivaOCH() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 14, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	log.Println(yesterday)

	var value = "SUCCESS_PAYMENT_FINACLE"
	var key = "status"
	// var key1 = "is_cashout"
	// var value1 = true
	var tot int
	if tot, err := collectionBriva().Find(bson.M{key: value, "updated_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}

	log.Println(tot)

	return tot, nil
}

func GetStatusTimeout() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 14, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	log.Println(yesterday)

	var value = "TIMEOUT"
	var key = "status"
	var key1 = "is_cashout"
	var value1 = true
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key: value, key1: value1, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetStatusInProgress() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 14, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var value = "IN_PROGRESS"
	var key = "status"
	var key1 = "is_cashout"
	var value1 = true
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key: value, key1: value1, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetStatusFailed() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 14, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var value = "FAILED"
	var key = "status"
	var key1 = "is_cashout"
	var value1 = true
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key: value, key1: value1, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetStatusSuccess() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 13, 30, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 00, 0, 0, utc)
	}

	var value = "SUCCESS"
	var key = "status"
	var key1 = "is_cashout"
	var value1 = true
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key: value, key1: value1, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountRegistered() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("register")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var value = "REGISTERED"
	var key = "status"
	var tot int
	if tot, err := collectionRegistrations().Find(bson.M{key: value, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountLogin() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("login")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "status"
	var value = "OTP_SUCCESS"
	var tot []string
	err := collectionLogin().Find(bson.M{key: value, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Distinct("phoneno", &tot)
	if err != nil {
		log.Fatal(err, " WKWKWK")
	}

	var has = len(tot)
	// fmt.Println(has)
	log.Println(has)

	return has, nil
}

func GetCountNotYet() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("not yet")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var value = "NOT_YET"
	var key = "application_status"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}

	log.Println(tot)
	return tot, nil
}

func GetCountInProgress() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("in progress")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "IN_PROGRESS"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountPendingCreditScoring() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("pending credit scoring")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "PENDING_CREDIT_SCORING"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountWaitingEmailVerif() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("waiting email verif")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}
	var key = "application_status"
	var value = "WAITING_FOR_EMAIL_VERIFICATION"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountAppReject() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("app reject")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "APPLICATION_REJECTED"
	var tot []string
	err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Distinct("detail.account_number", &tot)
	if err != nil {
		log.Fatal(err, " WKWKWK")
	}

	var has = len(tot)
	// fmt.Println(tot)
	log.Println(has)
	return has, nil
}

func GetGradeRej8() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("app reject 8 ")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "APPLICATION_REJECTED"
	var key1 = "credit_scoring_result.grade"
	var value1 = 8
	var tot []string
	err := collectionLos().Find(bson.M{key: value, key1: value1, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Distinct("detail.account_number", &tot)
	if err != nil {
		log.Fatal(err, " WKWKWK")
	}

	var has = len(tot)
	// fmt.Println(has)
	log.Println(has)
	return has, nil
}
func GetGradeRej9() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("app reject 9")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "APPLICATION_REJECTED"
	var key1 = "credit_scoring_result.grade"
	var value1 = 9
	var tot []string
	err := collectionLos().Find(bson.M{key: value, key1: value1, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Distinct("detail.account_number", &tot)
	if err != nil {
		log.Fatal(err, " WKWKWK")
	}

	var has = len(tot)
	// fmt.Println(has)
	log.Println(has)
	return has, nil
}

func GetGradeRej10() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("app reject 10")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "APPLICATION_REJECTED"
	var key1 = "credit_scoring_result.grade"
	var value1 = 10
	var tot []string
	err := collectionLos().Find(bson.M{key: value, key1: value1, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Distinct("detail.account_number", &tot)
	if err != nil {
		log.Fatal(err, " WKWKWK")
	}

	var has = len(tot)
	// fmt.Println(has)
	log.Println(has)
	return has, nil
}

func GetGrademin1() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("app reject -1")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "APPLICATION_REJECTED"
	var key1 = "credit_scoring_result.grade"
	var value1 = -1
	var tot []string
	err := collectionLos().Find(bson.M{key: value, key1: value1, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Distinct("detail.account_number", &tot)
	if err != nil {
		log.Fatal(err, " WKWKWK")
	}

	var has = len(tot)
	// fmt.Println(has)
	log.Println(has)

	return has, nil
}

func GetCountPendOffer() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("pending offer")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "PENDING_OFFER"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountOfferReje() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("offer reject")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "OFFER_REJECTED"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountOfferAccep() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("offer accepted")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "OFFER_ACCEPTED"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountSelfieNot() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("serfie not verif")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "ELEMENT_SELFIE_NOT_VERIFIED"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountSelfie() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("serfie verif")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "ELEMENT_SELFIE_VERIFIED"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountPrivyRegistered() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("privy register")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "PRIVY_REGISTERED"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountPrivySendOtp() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("privy send otp")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "PRIVY_SEND_OTP"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountPrivySuccessOtp() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("privy success otp")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "PRIVY_SUCCESS_VERIFY_OTP"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountSuccessFaceCompare() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("success face compare")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "SUCCESS_FACE_COMPARE"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "last_update_time": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountSuccessSignDoc() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("sign doc")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "status"
	var value = "SUCCESS_SIGN_DOCUMENT"
	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: value, "updated_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountSuccessCreateWelcome() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println("Success Create Welcome")
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "open_date": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetWelcomeGrade1() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}
	var key = "application_status"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var key1 = "credit_scoring_result.grade"
	var value1 = 1
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, key1: value1, "open_date": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetWelcomeGrade2() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}
	var key = "application_status"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var key1 = "credit_scoring_result.grade"
	var value1 = 2
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, key1: value1, "open_date": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetWelcomeGrade3() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var key1 = "credit_scoring_result.grade"
	var value1 = 3
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, key1: value1, "open_date": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetWelcomeGrade4() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var key1 = "credit_scoring_result.grade"
	var value1 = 4
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, key1: value1, "open_date": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetGrade5() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var key1 = "credit_scoring_result.grade"
	var value1 = 5
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, key1: value1, "open_date": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetGrade6() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var key1 = "credit_scoring_result.grade"
	var value1 = 6
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, key1: value1, "open_date": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetGrade7() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var key1 = "credit_scoring_result.grade"
	var value1 = 7
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, key1: value1, "open_date": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetGrade8() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var key1 = "credit_scoring_result.grade"
	var value1 = 8
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, key1: value1, "open_date": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetGrade9() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var key1 = "credit_scoring_result.grade"
	var value1 = 9
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, key1: value1, "open_date": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ApproveOverride() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "credit_scoring_result.desc"
	var value = "APPROVE OVERRIDE"
	// var key1 = "credit_scoring_result.grade"
	// var value1 = 8
	var tot int
	if tot, err := collectionLos().Find(bson.M{key: value, "open_date": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountGagalBelanjaCeria() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "status"
	//var value = "FAILED,REVERSAL"
	var totRev int
	var totFailed int
	if totRev, err := collectionCharges().Find(bson.M{key: "REVERSAL", "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return totRev, err
	}

	if totFailed, err := collectionCharges().Find(bson.M{key: "FAILED", "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return totFailed, err
	}

	tot := totFailed + totRev

	log.Println(tot)

	return tot, nil
}

func GetCountVerifKtp() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "legal_checklist.ektp"
	var value = "Y"
	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: value, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountRefundTokpedFailed() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	// var key = "status"
	// var value = "SUCCESS"
	var key1 = "refunds.status"
	var value1 = "FAILED"
	var key2 = "merchant_name"
	var value2 = "Tokopedia"
	var tot int

	if tot, err := collectionCharges().Find(bson.M{key1: value1, key2: value2, "refunds.created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountRefundTokpedSuccess() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	// var key = "status"
	// var value = "SUCCESS"
	var key1 = "refunds.status"
	var value1 = "SUCCESS"
	var key2 = "merchant_name"
	var value2 = "Tokopedia"
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key1: value1, key2: value2, "refunds.created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountRefundLinkAjaSuccess() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}
	// var key = "status"
	// var value = "SUCCESS"
	var key1 = "refunds.status"
	var value1 = "SUCCESS"
	var key2 = "merchant_name"
	var value2 = "LinkAja"
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key1: value1, key2: value2, "refunds.created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountRefundLinkAjaFailed() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	// var key = "status"
	// var value = "SUCCESS"
	var key1 = "refunds.status"
	var value1 = "FAILED"
	var key2 = "merchant_name"
	var value2 = "LinkAja"
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key1: value1, key2: value2, "refunds.created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountRefundDMFailed() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	// var key = "status"
	// var value = "SUCCESS"
	var key1 = "refunds.status"
	var value1 = "FAILED"
	var key2 = "merchant_name"
	var value2 = "Dinomarket"
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key1: value1, key2: value2, "refunds.created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountRefundDMSuccess() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	// var key = "status"
	// var value = "SUCCESS"
	var key1 = "refunds.status"
	var value1 = "SUCCESS"
	var key2 = "merchant_name"
	var value2 = "Dinomarket"
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key1: value1, key2: value2, "refunds.created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountTotalHitWebmerchant() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "is_cashout"
	var value = false
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key: value, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountTotalHitWebmerchantPending() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}
	var key = "is_cashout"
	var value = false
	var key1 = "status"
	var value1 = "PENDING"
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key: value, key1: value1, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountTotalHitWebmerchantSuccess() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}
	var key = "is_cashout"
	var value = false
	var key1 = "status"
	var value1 = "SUCCESS"
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key: value, key1: value1, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountTotalHitWebmerchantFailed() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}
	var key = "is_cashout"
	var value = false
	var key1 = "status"
	var value1 = "FAILED"
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key: value, key1: value1, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func getCountRejectWhitelist() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var whitelist = "whitelist.error"
	var notFound = "no data found"
	var loc = "location.place"
	var exist = "$exists"
	var bool = false
	var tot []string
	err := collectionLoanApp().Find(bson.M{loc: bson.M{exist: bool}, whitelist: notFound, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Distinct("cams.AccNo", &tot)
	if err != nil {
		log.Fatal(err, " WKWKWK")
	}
	log.Println(tot)

	var has = len(tot)
	// fmt.Println(tot)

	return has, nil
}

func GetRejectFromPalembang() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "location.place"
	var regex = "$regex"
	var sumsel = ".*sumatera selatan.*"
	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{regex: sumsel}, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetLoginFromPalembang() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	// log.Println()
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "login_location.place"
	var regex = "$regex"
	var sumsel = ".*Sumatera Selatan.*"
	var tot int
	if tot, err := collectionRegistrations().Find(bson.M{key: bson.M{regex: sumsel}, "login_location.created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetPengajuanUkerPalembang() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	log.Println()
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "'unit_kerja.kode_uker"
	var in = "$in"
	var uker = "['00008','00020','00030','00040','00059','00063','00128','00129','00131','00138','00160','00164','00179','00184','00200','00275','00315','00324','00342','00602','00604','00606','00607','00670','00697','01001','01100','01102','01103','01104','01304','01463','01471','01487','01500','01669','01670','01676','01679','01707','01821','01863','01869','01887','02015','02076','02088','02089','02090','02091','02093','02145','02174','02193','02194','02205','02206','02207','02239','02263','02272','02277','02278','02279','02292','03286','03373','03388','03389','03393','03395','03396','03397','03550','03552','03553','03554','03555','03556','03557','03558','03559','03560','03696','05602','05603','05604','05605','05606','05607','05608','05609','05610','05611','05612','05613','05614','05628','05629','05630','05631','05632','05633','05634','05635','05636','05637','05638','05639','05640','05641','05642','05643','05644','05645','05646','05647','05648','05666','05667','05668','05669','05670','05671','05672','05673','05674','05675','05676','05677','05678','05679','05680','05681','05682','05683','05713','05714','05715','05716','05717','05718','05719','05720','05721','05722','05723','05724','05725','05726','05727','05728','05729','05730','05731','05732','05733','05734','05735','05736','05737','05738','05739','05740','05741','05742','05743','05744','05745','05746','05747','05748','05749','05750','05751','05752','05753','05754','05755','05756','05757','05758','05759','05760','05761','05762','05763','05764','05765','05766','05767','05768','05769','05770','05782','05783','05784','05785','05786','05787','05788','05789','05790','05791','05792','05793','05803','05804','05805','05806','07017','07041','07042','07043','07044','07045','07046','07047','07048','07049','07050','07051','07052','07053','07054','07080','07101','07102','07103','07104','07105','07106','07107','07108','07139','07140','07141','07142','07177','07178','07186','07196','07217','07219','07246','07288','07289','07338','07339','07360','07412','07432','07454','07461','07462','07498','07499','07520','07527','07606','07607','07608','07663','07671','07690','07691','07738','07739','07751','07752','07767','07770','07771','07785','07786','07850','07889','07902','07915','07916','07975','08052','08056','08057','08058','08059','08060','08061','08062','08099','08100','08144','08149','08150','08151','08152','08153','08154','08162','09838']"
	var tot int
	log.Println(key)
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{in: uker}, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountTotalHitWebmerchantReversal() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}
	var key = "is_cashout"
	var value = false
	var key1 = "status"
	var value1 = "REVERSAL"
	var tot int
	if tot, err := collectionCharges().Find(bson.M{key: value, key1: value1, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetWebMerchant() string {
	var listResult []string
	total_hit, _ := GetCountTotalHitWebmerchant()
	listResult = append(listResult, "\n11.Total Hit Web Merchant Payment  : "+strconv.Itoa(total_hit))

	aa := strings.Join(listResult, "\n")
	log.Println(aa)

	return aa
}
func ToInProgressToday() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "application_status_history"
	var element = "$elemMatch"
	var value = "IN_PROGRESS"
	var key1 = "to"
	// var value1 = 9
	var tot []string
	err := collectionLos().Find(bson.M{key: bson.M{element: bson.M{key1: value, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Distinct("detail.account_number", &tot)
	if err != nil {
		log.Fatal(err, " WKWKWK")
	}

	var has = len(tot)
	log.Println(has)

	return has, nil
}

func ToInProgress() (int, error) {
	log.Println("start report webview")
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	log.Println(yesterday)

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "IN_PROGRESS"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromInProgress() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "IN_PROGRESS"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"
	// var value1 = 9
	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToPendingCreditScoring() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "PENDING_CREDIT_SCORING"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"
	// var value1 = 9
	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromPendingCreditScoringToday() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "PENDING_CREDIT_SCORING"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"
	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromPendingCreditScoring() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "PENDING_CREDIT_SCORING"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToWaitingEmail() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "WAITING_FOR_EMAIL_VERIFICATION"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromWaitingEmail() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "WAITING_FOR_EMAIL_VERIFICATION"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToAppReject() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "APPLICATION_REJECTED"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToPendingOffer() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "PENDING_OFFER"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromPendingOffer() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "PENDING_OFFER"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToOfferReject() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "OFFER_REJECTED"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToOfferAccepted() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "OFFER_ACCEPTED"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromOfferAccepted() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "application_status_history"
	var element = "$elemMatch"
	var value = "OFFER_ACCEPTED"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToSelfieNotVerif() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "ELEMENT_SELFIE_NOT_VERIFIED"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromSelfieNotVerif() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "ELEMENT_SELFIE_NOT_VERIFIED"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToSelfieVerif() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "ELEMENT_SELFIE_VERIFIED"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromSelfieVerif() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "ELEMENT_SELFIE_VERIFIED"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToPrivyRegister() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "PRIVY_REGISTERED"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromPrivyRegister() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "PRIVY_REGISTERED"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToPrivySendOTP() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "PRIVY_SEND_OTP"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromPrivySendOTP() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "PRIVY_SEND_OTP"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToPrivySuccessVerifOTP() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "PRIVY_SUCCESS_VERIFY_OTP"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromPrivySuccessVerifOTP() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "PRIVY_SUCCESS_VERIFY_OTP"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToSuccessFaceCompare() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "SUCCESS_FACE_COMPARE"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromSuccessFaceCompare() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "SUCCESS_FACE_COMPARE"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToSuccessSignDocument() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "SUCCESS_SIGN_DOCUMENT"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func FromSuccessSignDocument() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "SUCCESS_SIGN_DOCUMENT"
	var key1 = "from"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}
func ToWelcomePackage() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var key1 = "to"
	var key2 = "source"
	var value1 = "1401714017"

	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{key: bson.M{element: bson.M{key1: value, key2: value1, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToWelcomePackageApprove() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var key1 = "to"
	var desc = "credit_scoring_result.desc"
	var value1 = "APPROVE"
	var key2 = "source"
	var value2 = "1401714017"
	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{desc: value1, key: bson.M{element: bson.M{key1: value, key2: value2, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func ToWelcomePackageApproveOverride() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 16, 59, 59, 59, utc)
	}

	var key = "status_changes"
	var element = "$elemMatch"
	var value = "SUCCESS_CREATE_WELCOME_PACKAGE"
	var key1 = "to"
	var desc = "credit_scoring_result.desc"
	var value1 = "APPROVE OVERRIDE"
	var key2 = "source"
	var value2 = "1401714017"
	var tot int
	if tot, err := collectionUseEkyc().Find(bson.M{desc: value1, key: bson.M{element: bson.M{key1: value, key2: value2, "at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func Waterfall() string {
	var listResult []string
	ToInProgress, _ := ToInProgress()
	listResult = append(listResult, "\n - TO IN_PROGRESS  : "+strconv.Itoa(ToInProgress))
	FromInProgress, _ := FromInProgress()
	listResult = append(listResult, "FROM IN_PROGRESS  : "+strconv.Itoa(FromInProgress))
	ToPendingCreditScoring, _ := ToPendingCreditScoring()
	listResult = append(listResult, "TO PENDING_CREDIT_SCORING  : "+strconv.Itoa(ToPendingCreditScoring))
	FromPendingCreditScoring, _ := FromPendingCreditScoring()
	listResult = append(listResult, "FROM  PENDING_CREDIT_SCORING  : "+strconv.Itoa(FromPendingCreditScoring))
	ToWaitingEmail, _ := ToWaitingEmail()
	listResult = append(listResult, "TO WAITING_EMAIL_VERIFICATION  : "+strconv.Itoa(ToWaitingEmail))
	FromWaitingEmail, _ := FromWaitingEmail()
	listResult = append(listResult, "FROM WAITING_EMAIL_VERIFICATION  : "+strconv.Itoa(FromWaitingEmail))
	ToAppReject, _ := ToAppReject()
	listResult = append(listResult, "TO APPLICATIN_REJECT  : "+strconv.Itoa(ToAppReject))
	ToPendingOffer, _ := ToPendingOffer()
	listResult = append(listResult, "TO PENDING_OFFER  : "+strconv.Itoa(ToPendingOffer))
	FromPendingOffer, _ := FromPendingOffer()
	listResult = append(listResult, "FROM PENDING_OFFER  : "+strconv.Itoa(FromPendingOffer))
	ToOfferReject, _ := ToOfferReject()
	listResult = append(listResult, "TO OFFER_REJECTED  : "+strconv.Itoa(ToOfferReject))
	ToOfferAccepted, _ := ToOfferAccepted()
	listResult = append(listResult, "FROM OFFER_ACCEPTED  : "+strconv.Itoa(ToOfferAccepted))
	// FromOfferAccepted, _ := FromOfferAccepted()
	// listResult = append(listResult, "-OFFER_ACCEPTED  : "+strconv.Itoa(FromOfferAccepted))
	ToSelfieNotVerif, _ := ToSelfieNotVerif()
	listResult = append(listResult, "TO ELEMENT_SELFIE_NOT_VERIFIED  : "+strconv.Itoa(ToSelfieNotVerif))
	FromSelfieNotVerif, _ := FromSelfieNotVerif()
	listResult = append(listResult, "FROM ELEMENT_SELFIE_NOT_VERIFIED  : "+strconv.Itoa(FromSelfieNotVerif))
	ToSelfieVerif, _ := ToSelfieVerif()
	listResult = append(listResult, "TO ELEMENT_SELFIE_VERIFIED  : "+strconv.Itoa(ToSelfieVerif))
	FromSelfieVerif, _ := FromSelfieVerif()
	listResult = append(listResult, "FROM ELEMENT_SELFIE_VERIFIED  : "+strconv.Itoa(FromSelfieVerif))
	ToPrivyRegister, _ := ToPrivyRegister()
	listResult = append(listResult, "TO PRIVY_REGISTER  : "+strconv.Itoa(ToPrivyRegister))
	FromPrivyRegister, _ := FromPrivyRegister()
	listResult = append(listResult, "FROM PRIVY_REGISTER  : "+strconv.Itoa(FromPrivyRegister))
	ToPrivySendOTP, _ := ToPrivySendOTP()
	listResult = append(listResult, "TO PRIVY_SEND_OTP  : "+strconv.Itoa(ToPrivySendOTP))
	FromPrivySendOTP, _ := FromPrivySendOTP()
	listResult = append(listResult, "FROM PRIVY_SEND_OTP  : "+strconv.Itoa(FromPrivySendOTP))
	ToPrivySuccessVerifOTP, _ := ToPrivySuccessVerifOTP()
	listResult = append(listResult, "TO PRIVY_SUCCESS_VERIFY_OTP  : "+strconv.Itoa(ToPrivySuccessVerifOTP))
	FromPrivySuccessVerifOTP, _ := FromPrivySuccessVerifOTP()
	listResult = append(listResult, "FROM PRIVY_SUCCESS_VERIFY_OTP  : "+strconv.Itoa(FromPrivySuccessVerifOTP))
	ToSuccessFaceCompare, _ := ToSuccessFaceCompare()
	listResult = append(listResult, "TO SUCCESS_FACE_COMPARE  : "+strconv.Itoa(ToSuccessFaceCompare))
	FromSuccessFaceCompare, _ := FromSuccessFaceCompare()
	listResult = append(listResult, "FROM SUCCESS_FACE_COMPARE  : "+strconv.Itoa(FromSuccessFaceCompare))
	ToSuccessSignDocument, _ := ToSuccessSignDocument()
	listResult = append(listResult, "TO SUCCESS_SIGN_DOCUMENT  : "+strconv.Itoa(ToSuccessSignDocument))
	FromSuccessSignDocument, _ := FromSuccessSignDocument()
	listResult = append(listResult, "FROM SUCCESS_SIGN_DOCUMENT  : "+strconv.Itoa(FromSuccessSignDocument))
	ToWelcomePackage, _ := ToWelcomePackage()
	listResult = append(listResult, "TO SUCCESS_CREATE_WELCOME_PACKAGE  : "+strconv.Itoa(ToWelcomePackage))
	// ToWelcomePackageApprove, _ := ToWelcomePackageApprove()
	// listResult = append(listResult, "APPROVE  : "+strconv.Itoa(ToWelcomePackageApprove))
	// ToWelcomePackageApproveOverride, _ := ToWelcomePackageApproveOverride()
	// listResult = append(listResult, "APPROVE OVERRIDE  : "+strconv.Itoa(ToWelcomePackageApproveOverride))

	aa := strings.Join(listResult, "\n - ")
	log.Println(aa)
	return aa

}

func GetPalembang() string {
	var listResult []string
	// palembang, _ := GetRejectFromPalembang()
	// listResult = append(listResult, "\n=============================")
	// listResult = append(listResult, "\nPengajuan dari Palembang  : "+strconv.Itoa(palembang))

	// loginLocation, _ := GetLoginFromPalembang()
	// listResult = append(listResult, "Login dari Palembang  : "+strconv.Itoa(loginLocation))
	// ukerPalembang, _ := GetPengajuanUkerPalembang()
	// listResult = append(listResult, "Unit Kerja dari Palembang  : "+strconv.Itoa(ukerPalembang))
	camsFreeze, _ := CamsFreeze()
	listResult = append(listResult, "User Salah Input Data Cams > 3x  : "+strconv.Itoa(camsFreeze))

	aa := strings.Join(listResult, "\n")
	log.Println(aa)
	return aa
}

func CamsFreeze() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "cams_attempt"
	var value = 3
	var tot int
	if tot, err := collectionLoanApp().Find(bson.M{key: value, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetRegistarsiWebView() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()
	// log.Println()
	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "last_device_ide"
	var regex = "$regex"
	var webview = ".*webview.*"
	var tot int
	if tot, err := collectionRegistrations().Find(bson.M{key: bson.M{regex: webview}, "updated_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Count(); err == nil {
		return tot, err
	}
	log.Println(tot)

	return tot, nil
}

func GetCountLoginWeb() (int, error) {
	now := time.Now()
	utc, _ := time.LoadLocation("")
	year, month, day := now.Date()

	var yesterday time.Time

	if day == 1 {
		yesterday = time.Date(year, month, 0, 17, 0, 0, 0, utc)
	} else {
		yesterday = time.Date(time.Now().Year(), time.Now().Month(), time.Now().AddDate(0, 0, -1).Day(), 17, 0, 0, 0, utc)
	}

	var key = "status"
	var value = "AUTHENTICATED_WEB"
	var tot []string
	err := collectionLogin().Find(bson.M{key: value, "created_at": bson.M{"$gte": yesterday, "$lte": time.Now().UTC()}}).Distinct("phoneno", &tot)
	if err != nil {
		log.Fatal(err, " WKWKWK")
	}

	var has = len(tot)
	// fmt.Println(has)

	return has, nil
}

func GetRefund() string {
	var listResult []string
	Refund_Tokped_Failed, _ := GetCountRefundTokpedFailed()
	listResult = append(listResult, "\nRefund Tokopedia Gagal  : "+strconv.Itoa(Refund_Tokped_Failed))

	Refund_Tokped_Success, _ := GetCountRefundTokpedSuccess()
	listResult = append(listResult, "Refund Tokopedia Success  : "+strconv.Itoa(Refund_Tokped_Success))

	Refund_Linkaja_failed, _ := GetCountRefundLinkAjaFailed()
	listResult = append(listResult, "Refund LinkAja gagal : "+strconv.Itoa(Refund_Linkaja_failed))

	Refund_Linkaja_Success, _ := GetCountRefundLinkAjaSuccess()
	listResult = append(listResult, "Refund LinkAja Success : "+strconv.Itoa(Refund_Linkaja_Success))

	Refund_DM_Failed, _ := GetCountRefundDMFailed()
	listResult = append(listResult, "Refund Dinomarket gagal : "+strconv.Itoa(Refund_DM_Failed))

	Refund_DM_Success, _ := GetCountRefundDMSuccess()
	listResult = append(listResult, "Refund Dinomarket Success : "+strconv.Itoa(Refund_DM_Success))

	aa := strings.Join(listResult, "\n")
	log.Println(aa)

	return aa

}

func GetBodyMesageCeria() string {

	var bodys = []string{}
	ac := accounting.Accounting{Symbol: "Rp.", Precision: 2}
	//bodys = append(bodys, "\nDear "+"partnerName"+",\n\nWe want to inform you that the transaction for refund cannot be processed\nbecause the balance on your checking account is insufficient:")
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\nREPORT CERIA : "+jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now()))
	bodys = append(bodys, "\n==========================")

	var res = GetNumberAndVolCeriaTrx()
	// var gagal, _ = GetCountGagalBelanjaCeria()
	var oda = getOdaPositive()
	var cekOda = getOdaCore()
	var briva = GetBrivaPending()
	var hold = GetHoldPending()
	var limit = GetDailyLimit()
	var holdSuccess = GetHoldSuccess()
	var holdFailed = GetHoldFailed()
	var autodebet = getAutodebetStatus()

	bodys = append(bodys, "\n1. daily download\t\t\t:\t")
	bodys = append(bodys, "\n2. user di msg2 step di aplikasi ceria sampai limitnya diapprove\t\t\t:\t"+GetStateCeria())
	bodys = append(bodys, "\n3. Jml user yg sukses belanja menggunakan ceria\t\t\t:\t"+res.Key)
	// bodys = append(bodys, "\n4. Jml user yg belanja menggunakan ceria tp gagal\t\t\t:\t"+strconv.Itoa(gagal))
	bodys = append(bodys, "\n4. Besar volume transaksi ceria based on data no 3\t\t\t:\t"+ac.FormatMoney(res.Value64))
	bodys = append(bodys, "\n5. AutoDebet Ceria \t\t\t:\t"+GetReportAutoDebetCeria())
	bodys = append(bodys, "\n6. Data Per Blok Ceria \t\t\t:\t"+GetReportDataPerBlok())
	bodys = append(bodys, "\n7. Merchant \nMERCHANT | JUMLAH | AMOUNT \t\t\t:\t"+GetReportDataMerchant())
	bodys = append(bodys, "\n8. Transaksi cashout ceria \t\t\t:\t")
	bodys = append(bodys, "\n Cashout Bri\t\t\t: \nAMOUNT | JUMLAH\t"+GetCashoutBriAll())
	bodys = append(bodys, "\n Cashout LinkAja\t\t\t:\nAMOUNT | JUMLAH\t"+GetCashoutLinkAja())
	bodys = append(bodys, "\n9. Detail Transaksi Refund \t\t\t:\t"+GetRefund())
	bodys = append(bodys, ""+GetWebMerchant())
	bodys = append(bodys, "\n10. ODA Positive \t\t\t:\t"+oda.Key)
	bodys = append(bodys, "\n11. Briva Pending \t\t\t:\t"+briva.Key)
	bodys = append(bodys, "\n12. Hold Pending  \t\t\t:\t"+hold.Key)
	bodys = append(bodys, "\n13. Hold Failed  \t\t\t:\t"+holdFailed.Key)
	bodys = append(bodys, "\n14. Hold Success \t\t\t:\t"+holdSuccess.Key)
	bodys = append(bodys, "\n15. Break Down Status Cashout \t\t\t:\t"+GetStatusCashOut())
	// bodys = append(bodys, "\n18. "+GetStatusOchTimeOut())
	// bodys = append(bodys, "\n19. "+getWhitelistReject())
	bodys = append(bodys, "\n16. Autodebet Status \t\t\t:\t"+autodebet.Key)
	// bodys = append(bodys, "\n21. "+GetBriva())
	// bodys = append(bodys, ""+number21and22())
	bodys = append(bodys, "\n\nFraud detection \t\t\t\t"+GetPalembang())

	bodys = append(bodys, "\n\nWe would be very happy to assist you.\nThank you.\n\nBest Regards,\n\nCeria Ops Team")

	bodys = append(bodys, "\nOur Summary Report")

	bodys = append(bodys, "\na. "+GetWelcomeOnly())
	bodys = append(bodys, "\nb. Transaksi \t\t\t:\t"+ac.FormatMoney(res.Value64))
	bodys = append(bodys, "\nc. Login Ceria \t\t\t:\tNormal / Tidak Normal")
	bodys = append(bodys, "\nd. Status CashOut \t\t\t:\tNormal / Tidak Normal")
	bodys = append(bodys, "\ne. Status Pengajuan \t\t\t:\tNormal / Tidak Normal")
	bodys = append(bodys, "\nf. Cek ODA : ")
	bodys = append(bodys, "\nODA Core \t:\t"+cekOda.Key)
	bodys = append(bodys, "\n"+GetCekODAOch())
	bodys = append(bodys, "\ng. Double Debet Check Report \t\t\t:\t")
	bodys = append(bodys, "\nh. Daily Limit\t\t\t:\t"+ac.FormatMoney(limit.Value64))

	return strings.Join(bodys, "")
}

func GetCeriaFraud() string {
	var bodys = []string{}
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\nREPORT FRAUD CERIA : "+jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now()))
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n1. Transaksi dan cashout lebih dari 3 kali dalam waktu berdekatan \nODA \t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t | Amount \t\t\t\t\t\t\t\t| Count \t\t\t"+GetCashoutAnomali())
	bodys = append(bodys, "\n==========================")
	// bodys = append(bodys, "\nOK:")
	// bodys = append(bodys, "\nNOT OK:")
	// bodys = append(bodys, "\n==========================")
	// bodys = append(bodys, "\n2. Cashout beda rekening \nODA \t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t | Rekening Autodebet \t\t\t\t\t\t\t\t| Rekening Tujuan  \t\t\t"+GetCashoutAnomali())
	// bodys = append(bodys, "\n==========================")
	return strings.Join(bodys, "")
}

func GetCeriaFraudExtend() string {
	var bodys = []string{}

	var BlokR = GetTransactioBlokR()
	var BlokN = GetTransactioBlokN()
	var Akumulasi = GetCashoutMoreThanRule()
	var cashoutRule = GetCashoutMoreThanRule2()

	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\nREPORT FRAUD CERIA : "+jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now()))
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n3. Transaksi diatas 2 Juta \nODA \t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t | Amount Transaksi | Limit | Percentage \t\t\t"+GetTransactioBig())
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\nOK:")
	bodys = append(bodys, "\nNOT OK:")
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n4. Transaksi di Lakukan Blok R \t\t\t:\t"+BlokR.Key)
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n5. Transaksi di Lakukan Blok N \t\t\t:\t"+BlokN.Key)
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n6. Akumulasi cashout dengan amount lebih dari 30% \t\t\t:\t"+Akumulasi.Key)
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n7. Cashout dengan amount lebih dari 30% dari limit \t\t\t:\t"+cashoutRule.Key)
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n8. Cashout lebih dari 3 kali dalam 1 bulun \t\t\t:\t"+BlokN.Key)
	bodys = append(bodys, "\n==========================")
	return strings.Join(bodys, "")
}

func GetCeriaWaterfall() string {
	var bodys = []string{}
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\nREPORT WATERFALL WEBVIEW: ")
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, ""+Waterfall())
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n\nWe would be very happy to assist you.\nThank you.\n\nBest Regards,\n\nCeria Ops Team")
	return strings.Join(bodys, "")
}

func GetReportIngestFlag() string {
	var bodys = []string{}
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\nREPORT FRAUD CERIA : "+jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now()))
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n1. tbaadm status ingest\t\t\t:\t"+GetReportTbaadmIngest())
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n2. crmuser status ingest\t\t\t:\t"+GetReportCrmuserIngest())
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n3. custom status ingest\t\t\t:\t"+GetReportCustomIngest())
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n4. public status ingest\t\t\t:\t"+GetReportPublicIngest())
	return strings.Join(bodys, "")
}

func GetUpdateGDS() string {
	var bodys = []string{}

	var Disbursement = EodDisbursement()
	var DailyDisbursement = EodDailyDisbursement()
	var LoanPayment = EodTotalLoanPayment()
	var DailyLoanPayment = EodDailyLoanPayment()
	var TotalOSNPL = EodTotalOSNPL()
	var TotalOSPL = EodTotalOSPL()
	var TotaDPKUser = EodTotalDPKUser()
	var TotalDPKSystem = EodTotalDPKSystem()
	var TotalNPLUser = EodTotalNPLUser()
	var TotalNPLSystem = EodTotalNPLSystem()
	var TotalAmountDPKUser = EodTotalAmountDPKUser()
	var TotalAmountDPKSystem = EodTotalAmountDPKSystem()
	var TotalAmountNPLUser = EodTotalAmountNPLUser()
	var TotalAmountNPLSystem = EodTotalAmountNPLSystem()
	var TotalLAA = EodTotalLAA()
	var TotalODA = EodTotalODA()
	var TotalCloseODA = EodCloseODA()
	var TotalLimit = EodTotalLimit()
	var DailyLimit = EodDailyLimit()
	var DailyLAA = EodDailyLAA()
	var DailyODA = EodDailyODA()
	var TotalCashout = EodTotalCashout()
	var TotalPurchase = EodTotalPurchase()
	var DailyAmountPurchaseV1 = EodDailyAmountPurchaseV1()
	var DailyAmountPurchaseV2 = EodDailyAmountPurchaseV2()
	var DailyDebiturPurchse = EodDailyDebiturPurchase()

	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\nEOD REPORT : "+jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now()))
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n1. Total disbursement \t\t:\t"+Disbursement.Key)
	bodys = append(bodys, "\n2. Total loan payment \t\t:\t"+LoanPayment.Key)
	bodys = append(bodys, "\n3. Total outstanding PL \t\t:\t"+TotalOSPL.Key)
	bodys = append(bodys, "\n4. Total outstanding NPL \t\t:\t"+TotalOSNPL.Key)
	bodys = append(bodys, "\n5. - Total debitur of DPK (User Modified Collectability) \t\t:\t"+TotaDPKUser.Key)
	bodys = append(bodys, "\n   - Total debitur of DPK (System Calculated Collectability) \t\t:\t"+TotalDPKSystem.Key)
	bodys = append(bodys, "\n6. - Total debitur of NPL (User Modified Collectability) \t\t:\t"+TotalNPLUser.Key)
	bodys = append(bodys, "\n   - Total debitur of NPL (System Calculated Collectability) \t\t:\t"+TotalNPLSystem.Key)
	bodys = append(bodys, "\n7. Total amount of DPK \na. Total Amount Based on User Modified Collectability \t\t:\t"+TotalAmountDPKUser.Key)
	bodys = append(bodys, "\nb. Total Amount Based on System Calculated Collectability \t\t:\t"+TotalAmountDPKSystem.Key)
	bodys = append(bodys, "\n8. Total amount of NPL \na. Total Amount Based on User Modified Collectability \t\t:\t"+TotalAmountNPLUser.Key)
	bodys = append(bodys, "\nb. Total Amount Based on System Calculated Collectability \t\t:\t"+TotalAmountNPLSystem.Key)
	bodys = append(bodys, "\n10. Total number of loan accounts \n a. LAA \t\t:\t"+TotalLAA.Key)
	bodys = append(bodys, "\nb. ODA \t\t:\t"+TotalODA.Key)
	bodys = append(bodys, "\n11. Total number of closed ODA accounts \t\t:\t"+TotalCloseODA.Key)
	bodys = append(bodys, "\n12. Daily Disbursement \t\t:\t"+DailyDisbursement.Key)
	bodys = append(bodys, "\n13. Daily Loan Payment \t\t:\t"+DailyLoanPayment.Key)
	bodys = append(bodys, "\n14. Daily LAA \t\t:\t"+DailyLAA.Key)
	bodys = append(bodys, "\n15. Daily ODA \t\t:\t"+DailyODA.Key)
	bodys = append(bodys, "\n16. Total Limit \t\t:\t"+TotalLimit.Key)
	bodys = append(bodys, "\n17. Daily Limit \t\t:\t"+DailyLimit.Key)
	bodys = append(bodys, "\n18. Total Cashout\t\t:\t"+TotalCashout.Key)
	bodys = append(bodys, "\n19. Daily amount of cashout (success, failed, pending)\t\t:\t"+EodDailyAmountCashout())
	bodys = append(bodys, "\n20. Cashout debitur per day (success, failed, pending)\t\t:\t"+EodDailyDebiturCashout())
	bodys = append(bodys, "\n21. Total Purchase\t\t:\t"+TotalPurchase.Key)
	bodys = append(bodys, "\n22. Daily amount of purchase\t\t:\n V1\t:\t"+DailyAmountPurchaseV1.Key)
	bodys = append(bodys, "\nV2\t:\t"+DailyAmountPurchaseV2.Key)
	bodys = append(bodys, "\n23. purchase debitur per day\t\t:\t"+DailyDebiturPurchse.Key)

	bodys = append(bodys, "\n==========================")
	return strings.Join(bodys, "")
}

// func number21and22() string {
// 	var listResult []string
// 	FromPendingCreditScoring, _ := FromPendingCreditScoringToday()
// 	ToInProgressToday, _ := ToInProgressToday()
// 	RejectWhitelist, _ := getCountRejectWhitelist()
// 	hasil := ToInProgressToday + RejectWhitelist
// 	listResult = append(listResult, "\n21. Total Pengajuan: "+strconv.Itoa(hasil))
// 	listResult = append(listResult, "22. Total Scoring: "+strconv.Itoa(FromPendingCreditScoring))

// 	aa := strings.Join(listResult, "\n")
// 	log.Println(aa)
// 	return aa

// }

// type KKD struct {
// 	Key         string
// 	Value       string
// 	Description string
// 	Ket         string
// 	Status      string
// }

// func GetKKD(foracid string) string {
// 	var result []KKD
// 	query := fmt.Sprintf("select "+
// 		"gam.acct_name as Key, "+
// 		"acc.uniqueid as Value, "+
// 		"acd.main_classification_system as Description, "+
// 		"case when acd.main_classification_system = '001' then 'LANCAR'"+
// 		"when acd.main_classification_system = '002' then 'DPK'"+
// 		"when acd.main_classification_system = '003' then 'NPL'"+
// 		"when acd.main_classification_system = '004' then 'NPL'"+
// 		"else 'NPL' end as Ket, "+
// 		"case when gam.acct_cls_flg = 'N' then 'OPEN ACC' "+
// 		"else 'CLOSE ACC' end as Status "+
// 		"from crmuser.accounts acc "+
// 		"join tbaadm.general_acct_mast_table gam on gam.cif_id=acc.orgkey "+
// 		"join tbaadm.asst_class_detail_tbl acd on acd.b2k_id=gam.acid "+
// 		"where acc.uniqueidtype = 'KTP'  and gam.schm_type = 'ODA' "+
// 		"and gam.foracid= '%s';", foracid)

// 	rows, err := db.DbCerCore.Raw(query).Scan(&result).Rows()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer rows.Close()
// 	var listResult []string

// 	for rows.Next() {
// 		var res KKD
// 		rows.Scan(
// 			&res.Key,
// 			&res.Value,
// 			&res.Description,
// 			&res.Ket,
// 			&res.Status,
// 		)
// 		result = append(result, res)
// 		a := fmt.Sprintf("Name = %s \nKTP = %s\nKolek = %s\nKeterangan = %s\nStatus = %s", res.Key, res.Value, res.Description, res.Ket, res.Status)

// 		listResult = append(listResult, a)
// 	}

// 	aa := strings.Join(listResult, "")

// 	//log.Println("asdsd ", aa)

// 	return aa
// }
