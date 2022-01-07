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
)

type Results struct {
	Key          string
	Value        string
	Description  string
	Value64      float64
	Description2 string
	Description3 string
	Description4 string
	Description5 string
	ListTagih    int
	ListBayar    int
	Time         int
}

func GetDateSql() string {
	var total Results
	err := db.DbCore.Table("custom.fstl").Select("current_date()+1 as Value").Limit(1).Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value
}

func GetData() string {
	var total Results
	err := db.DbCore.Table("partners").Select("sum(otp_timeout_in_seconds) as Value").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value
}

//select count(distinct loan_account) from custom.fstl where transfer_type ='AutoPay'
//and split_part(transaction_date_time,' ', 1) = current_date()

func GetDataNo2() string {
	log.Println("No.2")
	var total Results
	var date = jodaTime.Format("dd-MM-YYYY", time.Now())
	err := db.DbCore.Table("custom.fstl").Select("count(distinct loan_account) as Value").Where("(transfer_type ='AutoPay' and transaction_date_time = to_date('" + date + "','DD-MM-YYYY'))").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value
}

func GetDataNo3() string {
	log.Println("No.3")
	var total Results
	var date = jodaTime.Format("dd-MM-YYYY", time.Now())
	err := db.DbCore.Table("custom.fstl").Select("count(distinct loan_account) as Value").Where("(transfer_type ='AutoPay' and transaction_date_time = to_date('" + date + "','DD-MM-YYYY') and status ='S')").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value
}

func GetDataNo4() string {
	log.Println("No.4")
	var total Results
	var date = jodaTime.Format("dd-MM-YYYY", time.Now())
	err := db.DbCore.Table("custom.fstl").Select("count(distinct loan_account) as Value").Where("(transfer_type ='AutoPay' and transaction_date_time = to_date('" + date + "','DD-MM-YYYY') and status ='F')").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value
}

func GetDataNo5() string {
	log.Println("No.5")
	var total Results
	var date = jodaTime.Format("dd-MM-YYYY", time.Now())
	err := db.DbCore.Table("custom.fstl").Select("count(distinct loan_account) as Value").Where("(transfer_type ='AutoPay' and transaction_date_time = to_date('" + date + "','DD-MM-YYYY') and status ='P')").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value
}

func GetDataNo6() string {
	log.Println("No.6")
	var total Results
	var date = jodaTime.Format("dd-MM-YYYY", time.Now())
	err := db.DbCore.Table("tbaadm.lam").Select("count(1) as Value").Where("(payoff_flg ='Y' and payoff_date = to_date('" + date + "','DD-MM-YYYY'))").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value
}

func GetDataNo7() string {
	log.Println("No.7")
	var total Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	// err := db.DbCore.Table("custom.fstl").Select("count( distinct (loan_account)) as Value").Where("transfer_type ='DisbursementOut' and split_part(rcre_time + interval '11 hour',' ',1) = '" + date + "'").Scan(&total)
	err := db.DbOCH.Table("ececuser.clat").Select("count(*) as Value").Where("application_status = 'LOAN_CREATED' and r_mod_time between '" + date + " 00:00:00' and '" + date + " 23:59:59'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value
}

func GetDataNo8() string {
	log.Println("No.8")
	var total Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	err := db.DbCore.Table("custom.fstl").Select("count(distinct loan_account) as Value").Where("(transfer_type ='DisbursementOut' and split_part(rcre_time + interval '11 hour',' ',1) = '" + date + "' and status ='P')").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value
}

func GetDataHold() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbCore.Table("custom.hstl").Select("status as Key,count(1) as Value").Where("(TRANSFER_TYPE = 'HoldAmt' and TRANSACTION_DATE_TIME ='" + date + "')").Group("status").Rows()
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

	log.Println("asdsd ", aa)

	return aa
}

func GetDataUnhold() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbCore.Table("custom.hstl").Select("status as Key,count(1) as Value").Where("(TRANSFER_TYPE = 'UnHoldAmt' and TRANSACTION_DATE_TIME ='" + date + "')").Group("status").Rows()
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

	log.Println("asdsd ", aa)

	return aa
}

func GetDataFundTransfer() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbCore.Table("custom.fstl").Select("status as Key,count(1) as Value").Where("(TRANSFER_TYPE = 'AutoPay' and TRANSACTION_DATE_TIME ='" + date + "')").Group("status").Rows()
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

	log.Println("asdsd ", aa)

	return aa
}

func GetDataBreakDownCRApr() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbOCH.Table("ececuser.clat").Select("grade as Key,count(grade) as Value").Where("(application_status !='APP_EXPIRED'and application_status = 'CR_SCORE_APR' and split_part(r_mod_time, ' ', 1) ='" + date + "')").Group("grade").Rows()
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

	log.Println("asdsd ", aa)

	return aa
}

func GetPrevillage() string {
	var listRes []Results
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbOCH.Table("ececuser.capwd a right join ececuser.clat b on a.application_id = b.application_id").Select("case when accttype = 'P' then 'DEBITUR_BIASA' when accttype = 'PP' then 'DEBITUR_PRIVILAGE' end as Key,count(accttype) as Value").Where("(b.application_status = 'LOAN_CREATED' and b.r_mod_time between '" + currDate + " 00:00:00' and '" + currDate + " 23:59:59')").Group("accttype").Rows()
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

	log.Println("asdsd ", aa)

	return aa
}

func GetDataBreakDownCRRej() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbOCH.Table("ececuser.clat").Select("grade as Key,count(grade) as Value").Where("(application_status !='APP_EXPIRED' and application_status = 'CR_SCORE_REJ' and split_part(r_mod_time, ' ', 1) ='" + date + "')").Group("grade").Rows()
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

	log.Println("asdsd ", aa)

	return aa
}

func GetDataDoubleDebet() int {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbCore.Table("custom.fstl").Select("distinct loan_account as Key,count(loan_account) as Value").Where("(transaction_date_time='" + date + "' and status ='S' and transfer_type='AutoPay')").Group("loan_account").Having("count(loan_account)> 1").Rows()
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
	bb := len(listResult)
	// aa := strings.Join(listResult, "")

	log.Println("asdsd ", bb)

	return bb
}

func GetDataNo9() float64 {
	log.Println("No.9")
	var total Results
	var date = jodaTime.Format("dd-MM-YYYY", time.Now())
	err := db.DbCore.Table("(select sum(flow_amt) amt from tbaadm.ltd where flow_id in ('COLL','PDCOL') and entity_cre_flg = 'Y' and del_flg = 'N' and reversal_flg != 'Y' and reversed_flg != 'Y' and tran_date = to_date('" + date + "','DD-MM-YYYY') union select sum(flow_amt) amt from tbaadm.ltdh where flow_id in ('COLL','PDCOL') and entity_cre_flg = 'Y' and del_flg = 'N' and reversal_flg != 'Y' and reversed_flg != 'Y'and tran_date = to_date('" + date + "','DD-MM-YYYY'))").Select("sum(amt) as Value64").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value64
}

func GetDataNo10() float64 {
	log.Println("No.10")
	var total Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	err := db.DbOCH.Table("ececuser.clat").Select("sum(approved_amount) as Value64").Where("application_status = 'LOAN_CREATED' and r_mod_time between '" + date + " 00:00:00' and '" + date + " 23:59:59'").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value64
}

func GetDataTable() string {
	var listRes []Results
	rows, err := db.DbCore.Table("partners").Select("prefix_table as code,count(1) as desc").Group("prefix_table").Order("prefix_table desc").Rows()
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
		a := fmt.Sprintf("%s is %s", "\n"+res.Key, res.Value)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	log.Println("asdsd ", aa)

	return aa
}

func GetDataNo1() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbOCH.Table("ececuser.clat").Select("application_status as key, count(application_status) as value").Where("split_part(r_mod_time,' ', 1) = '" + date + "' and application_status != 'APP_EXPIRED'").Group("application_status").Order("value desc").Rows()

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
		a := fmt.Sprintf("%s is %s", "\n"+res.Key, res.Value)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	log.Println("asdsd ", aa)

	return aa
}

func GetDataNo0() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbOCH.Table("(select *, case when split_part(last_unsuccessful_login_time,' ', 1) = '" + date + "' then 'F' else 'S' end as status from ececuser.cusr where split_part(login_date,' ', 1) = '" + date + "')").Select("status as key, count(status) as value").Where("split_part(login_date,' ', 1) = '" + date + "'").Group("status").Rows()
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
		a := fmt.Sprintf("%s is %s", "\n"+res.Key, res.Value)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	log.Println("asdsd ", aa)

	return aa
}

func GetDataCLrBalAmt0AndAccClsFlgN() string {
	log.Println("No.1")
	var total Results
	err := db.DbCore.Table("(select a.foracid from tbaadm.gam a join tbaadm.eit b on a.acid = b.entity_id where clr_bal_amt = 0 and acct_cls_flg = 'N')").Select("count(1) as Value").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value
}

func GetDataCLrBalAmtGreaterThanAndAccClsFlgN() string {
	log.Println("No.2")
	var total Results
	err := db.DbCore.Table("( select a.foracid from tbaadm.gam a join tbaadm.eit b on a.acid = b.entity_id join tbaadm.acd c on a.acid = c.b2k_id and  b.entity_id = c.b2k_id where clr_bal_amt > 0 and acct_cls_flg = 'N')").Select("count(1) as Value").Scan(&total)
	if err != nil {
		log.Println(err)
	}
	return total.Value
}

func GetDataExtractAppExpired() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbOCH.Table("( select a.application_status, case when b.status_desc is null then '' else b.status_desc end as status_desc, rank() over(partition by a.application_id order by a.r_mod_time desc) from ececuser.clht a left join ececuser.clat b on a.application_id = b.application_id where a.application_id in ( select application_id from (select application_id, status_desc, rank() over(partition by user_id order by r_mod_time desc) from ececuser.clat where application_status = 'APP_EXPIRED' and split_part(r_mod_time, ' ', 1) = '" + date + "') where rank = 1))").Select("application_status as key, status_desc as desc, count(status_desc) as value").Where("rank = 2").Group("application_status,status_desc").Order("application_status").Rows()
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
			&res.Value)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s is %s is %s", "\n"+res.Key, res.Value, res.Description)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	log.Println("asdsd ", aa)

	return aa
}

func GetDataDetailAutoPay() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbCore.Table("(select distinct UPPER(NVL(SUBSTR(status,0,1),'S')) as Key,status_code as Value,count(1) as Description from custom.fstl \nwhere transfer_type = 'AutoPay' and (transaction_date_time) = '" + date + "' \ngroup by UPPER(NVL(SUBSTR(status,0,1),'S')),status_code)").Select("key,value,description").Rows()
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
			&res.Description)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s | %s", "\n"+res.Key, res.Value, res.Description)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	return aa
}

func GetDataDetailDisbursementOut() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	rows, err := db.DbCore.Table("(select distinct UPPER(NVL(SUBSTR(status,0,1),'S')) as Key,status_code as Value,count(1) as Description from custom.fstl \nwhere transfer_type = 'DisbursementOut' and transaction_date_time = '" + date + " 00:00:00' \ngroup by UPPER(NVL(SUBSTR(status,0,1),'S')),status_code)").Select("key,value,description").Rows()
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
			&res.Description)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s | %s", "\n"+res.Key, res.Value, res.Description)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	return aa
}

func GetDataDetailFullPayoffOut() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbCore.Table("(select distinct UPPER(NVL(SUBSTR(status,0,1),'S')) as Key,status_code as Value,count(1) as Description from custom.fstl \nwhere transfer_type = 'FullPayoffOut' and (transaction_date_time) = '" + date + "' \ngroup by UPPER(NVL(SUBSTR(status,0,1),'S')),status_code)").Select("key,value,description").Rows()
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
			&res.Description)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s | %s | %s", "\n"+res.Key, res.Value, res.Description)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	return aa
}

func GetRegister() string {
	var listRes []Results
	var date = jodaTime.Format("YYYY-MM-dd", time.Now())
	rows, err := db.DbOCH.Table("ececuser.cusr").Select("count(user_id) as Value").Where("r_cre_time like '" + date + "%'").Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var listResult []string
	for rows.Next() {
		var res Results
		rows.Scan(
			&res.Value)
		listRes = append(listRes, res)
		a := fmt.Sprintf("%s", "\n"+res.Value)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	log.Println("asdsd ", aa)

	return aa
}

func WaterFall() string {
	var listRes []Results
	// var date = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	rows, err := db.DbOCH.Table("( select application_status,split_part(r_mod_time, ' ', 1) as date, case application_status when 'APP_CREATED'  then 1 when 'PAYROLL_REJ'  then 2 when 'PAYROLL_APP'  then 3 when 'OCR_VER_FAIL|1'  then 4 when 'OCR_VER_FAIL|2'  then 5 when 'KTP_SAVED'  then 6 when 'PER_SAVED'  then 7 when 'PAY_SAVED'  then 8 when 'CON_SAVED'  then 9 when 'CR_SCORE_SUB'  then 10 when 'CR_SCORE_REJ'  then 11 when 'CR_SCORE_APR'  then 12 when 'KTP_VERIFIED'  then 13 when 'USR_REJECT'  then 14 when 'DISB_ACC_CONF'  then 15 when 'EKYC_COM'  then 16 when 'DIG_SIGN_COM'  then 17 when 'DOCUMENT_SIGNED'  then 18 when 'LOAN_CREATED'  then 19 when 'LOAN_PAID'  then 20 when 'APP_EXPIRED'  then 21 end as id from ececuser.clht where split_part(r_mod_time, ' ', 1) = current_date())").Select("application_status as Key,count(application_status) as Description").Group("id,application_status").Order("id").Rows()

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
		a := fmt.Sprintf("%s is %s", "\n"+res.Key, res.Value)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	log.Println("waterfall ", aa)

	return aa
}

func nomer3() string {
	var listRes []Results
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	rows, err := db.DbOCH.Table("( select \n\tapplication_status,\n\tsplit_part(r_mod_time, ' ', 1) as date, \n\t\tcase application_status when 'APP_CREATED'  then 1 \n\t\twhen 'PAYROLL_REJ'  then 2 \n\t\twhen 'PAYROLL_APP'  then 3 \n\t\twhen 'OCR_VER_FAIL|1'  then 4 \n\t\twhen 'OCR_VER_FAIL|2'  then 5 \n\t\twhen 'KTP_SAVED'  then 6 \n\t\twhen 'PER_SAVED'  then 7 \n\t\twhen 'PAY_SAVED'  then 8 \n\t\twhen 'CON_SAVED'  then 9 \n\t\twhen 'CR_SCORE_SUB'  then 10 \n\t\twhen 'CR_SCORE_REJ'  then 11 \n\t\twhen 'CR_SCORE_APR'  then 12 \n\t\twhen 'KTP_VERIFIED'  then 13 \n\t\twhen 'USR_REJECT'  then 14 \n\t\twhen 'DISB_ACC_CONF'  then 15 \n\t\twhen 'EKYC_COM'  then 16 \n\t\twhen 'DIG_SIGN_COM'  then 17 \n\t\twhen 'DOCUMENT_SIGNED'  then 18 \n\t\twhen 'LOAN_CREATED'  then 19 \n\t\twhen 'LOAN_PAID'  then 20 \n\t\twhen 'APP_EXPIRED'  then 21 end as id from ececuser.clat \n\t\twhere r_mod_time between '" + currDate + " 00:00:00' and '" + currDate + " 23:59:59' )").Select("application_status as Key,count(application_status) as Description").Group("id,application_status").Order("id").Rows()

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
		a := fmt.Sprintf("%s is %s", "\n"+res.Key, res.Value)

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	log.Println("waterfall ", aa)

	return aa
}

func disbursementPinang() string {
	var listRes []Results
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	ac := accounting.Accounting{Symbol: "Rp.", Precision: 0}
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	rows, err := db.DbOCH.Table("ececuser.capwd a right join ececuser.clat b on a.application_id = b.application_id").Select("case when accttype = 'P' then 'DEBITUR_BIASA' when accttype = 'PP' then 'DEBITUR_PRIVILAGE' end as Key,  scheme_code as Value, count(scheme_code) as Description, sum(approved_amount) as Value64").Where("b.application_status = 'LOAN_CREATED' and scheme_code ='PNANG' and b.r_mod_time between '" + currDate + " 00:00:00' and '" + currDate + " 23:59:59'").Group("accttype,scheme_code").Rows()

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
		a := fmt.Sprintf("%s | %s | %s | %s", "\n"+res.Key, res.Value, res.Description, ac.FormatMoney(res.Value64))

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	log.Println("waterfall ", aa)

	return aa
}

func disbursementPyltr() string {
	var listRes []Results
	var currDate = jodaTime.Format("YYYY-MM-dd", time.Now())
	ac := accounting.Accounting{Symbol: "Rp.", Precision: 0}
	// var yesterday = jodaTime.Format("YYYY-MM-dd", time.Now().AddDate(0, 0, -1))
	rows, err := db.DbOCH.Table("ececuser.capwd a right join ececuser.clat b on a.application_id = b.application_id").Select("case when accttype = 'P' then 'DEBITUR_BIASA' when accttype = 'PP' then 'DEBITUR_PRIVILAGE' end as Key,  scheme_code as Value, count(scheme_code) as Description, sum(approved_amount) as Value64").Where("b.application_status = 'LOAN_CREATED' and scheme_code ='PYLTR' and b.r_mod_time between '" + currDate + " 00:00:00' and '" + currDate + " 23:59:59'").Group("accttype,scheme_code").Rows()

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
		a := fmt.Sprintf("%s | %s | %s | %s", "\n"+res.Key, res.Value, res.Description, ac.FormatMoney(res.Value64))

		listResult = append(listResult, a)
	}

	aa := strings.Join(listResult, "")

	log.Println("waterfall ", aa)

	return aa
}

func GetBodyMesage() string {

	ac := accounting.Accounting{Symbol: "Rp.", Precision: 2}
	var bodys = []string{}
	//bodys = append(bodys, "\nDear "+"partnerName"+",\n\nWe want to inform you that the transaction for refund cannot be processed\nbecause the balance on your checking account is insufficient:")
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\nREPORT PINANG : "+jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now()))
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\n1. Jumlah Nasabah Login, sukses dan gagal\t\t\t:\t"+GetDataNo0())
	bodys = append(bodys, "\n2. Jumlah Nasabah Register\t\t\t:\t"+GetRegister())
	bodys = append(bodys, "\n3. Jumlah Pengguna Untuk Setiap Langkah Pengajuan\t\t\t:\t"+nomer3())
	// bodys = append(bodys, "\nBreakdown APP_EXPIRED\t\t\t:\t"+GetDataExtractAppExpired())
	bodys = append(bodys, "\n4. Jumlah Nasabah Yang Melakukan Pembayaran\t\t\t:\t"+GetDataNo2())
	bodys = append(bodys, "\n5. Jumlah Nasabah Yang Melakukan Pembayaran Sukses\t\t\t:\t"+GetDataNo3())
	bodys = append(bodys, "\n6. Jumlah Nasabah Yang Melakukan Pembayaran Gagal\t\t\t:\t"+GetDataNo4())
	bodys = append(bodys, "\n7. Jumlah Nasabah Yang Melakukan Pembayaran Pending\t\t\t:\t"+GetDataNo5())
	bodys = append(bodys, "\n8. Jumlah Nasabah Yang Melakukan Pelunasan\t\t\t:\t"+GetDataNo6())
	bodys = append(bodys, "\n9. Jumlah Pencairan Sukses\t\t\t:\t"+GetDataNo7())
	bodys = append(bodys, "\n10. Jumlah Pencairan Gagal\t\t\t:\t"+GetDataNo8())
	bodys = append(bodys, "\n11. Jumlah Pendebetan\t\t\t:\t"+ac.FormatMoneyFloat64(GetDataNo9()))
	bodys = append(bodys, "\n12. Jumlah Pencairan\t\t\t:\t"+ac.FormatMoneyFloat64(GetDataNo10()))
	bodys = append(bodys, "\n\nWe would be very happy to assist you.\nThank you.\n\nBest Regards,\n\nPinang Ops Team")

	return strings.Join(bodys, "")
}
func GetWaterfall() string {

	var bodys = []string{}
	//bodys = append(bodys, "\nDear "+"partnerName"+",\n\nWe want to inform you that the transaction for refund cannot be processed\nbecause the balance on your checking account is insufficient:")
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\nREPORT WATERFALL : "+jodaTime.Format("dd-MM-YYYY", time.Now().AddDate(0, 0, -1)))
	bodys = append(bodys, "\n==========================")

	bodys = append(bodys, "\nWaterfall report\t\t\t:\t"+WaterFall())

	bodys = append(bodys, "\n\nWe would be very happy to assist you.\nThank you.\n\nBest Regards,\n\nPinang Ops Team")

	return strings.Join(bodys, "")
}

func GetBodyMesageAnomaly() string {

	var bodys = []string{}
	//bodys = append(bodys, "\nDear "+"partnerName"+",\n\nWe want to inform you that the transaction for refund cannot be processed\nbecause the balance on your checking account is insufficient:")
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\nREPORT ANOMALY PINANG : "+jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now()))
	bodys = append(bodys, "\n==========================")

	bodys = append(bodys, "\n1. number of GetDataCLrBalAmt0AndAccClsFlgN\t\t\t:\t"+GetDataCLrBalAmt0AndAccClsFlgN())
	bodys = append(bodys, "\n2. Number of GetDataCLrBalAmtGreaterThanAndAccClsFlgN\t\t\t:\t"+GetDataCLrBalAmtGreaterThanAndAccClsFlgN())

	bodys = append(bodys, "\n\nWe would be very happy to assist you.\nThank you.\n\nBest Regards,\n\nPinang Ops Team")

	return strings.Join(bodys, "")
}

func GetBodyMesageDetilTransferType() string {

	var bodys = []string{}
	var test = GetDataDoubleDebet()
	//bodys = append(bodys, "\nDear "+"partnerName"+",\n\nWe want to inform you that the transaction for refund cannot be processed\nbecause the balance on your checking account is insufficient:")
	bodys = append(bodys, "\n==========================")
	bodys = append(bodys, "\nREPORT DETAIL TRANSFER TYPE : "+jodaTime.Format("dd-MM-YYYY HH:mm:ss", time.Now()))
	bodys = append(bodys, "\n==========================")

	bodys = append(bodys, "\n1. Autodebet\t:\t"+GetDataDetailAutoPay())
	bodys = append(bodys, "\n2. Pencairan\t:\t"+GetDataDetailDisbursementOut())
	// bodys = append(bodys, "\n3. Previllage Atau Non Previllage\t:\t"+GetPrevillage())
	bodys = append(bodys, "\n3. Breakdown Pencairan\t:\t")
	bodys = append(bodys, "\na. Detail Pencairan Pinang \t:\t"+disbursementPinang())
	bodys = append(bodys, "\nb. Detail Pencairan Paylater\t:\t"+disbursementPyltr())
	bodys = append(bodys, "\n4. Pelunasan\t:\t"+GetDataDetailFullPayoffOut())
	bodys = append(bodys, "\n============================")
	bodys = append(bodys, "\n5. Breakdown Hold\t\t\t:\t"+GetDataHold())
	bodys = append(bodys, "\n6. Breakdown UnHold\t\t\t:\t"+GetDataUnhold())
	// bodys = append(bodys, "\n8. Breakdown Fund Transfer\t\t\t:\t"+GetDataFundTransfer())
	bodys = append(bodys, "\n============================")
	bodys = append(bodys, "\n7. Breakdown Credit Scoring Approve\t\t\t:\t"+GetDataBreakDownCRApr())
	bodys = append(bodys, "\n8. Breakdown Credit Scoring Reject\t\t\t:\t"+GetDataBreakDownCRRej())
	bodys = append(bodys, "\n============================")
	bodys = append(bodys, "\n9. Pendebetan > 1\t\t\t:\t"+strconv.Itoa(test))

	bodys = append(bodys, "\n\nWe would be very happy to assist you.\nThank you.\n\nBest Regards,\n\nPinang Ops Team")

	return strings.Join(bodys, "")
}
