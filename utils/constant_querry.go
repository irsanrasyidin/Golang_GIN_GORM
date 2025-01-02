package utils

const (

	//customer
	INSERT_CUST        = "INSERT INTO ms_user(id,username, password, role, is_active ) VALUES($1, $2, $3, $4, $5) RETURNING id"
	INSERT_CUST_USR    = "INSERT INTO ms_customer (id, id_user, full_name, NIK, noPhone, email, address, created_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"
	GET_CUST_ID        = "SELECT id,id_user,full_name,NIK,noPhone,email,address,created_at,updated_at,created_by,updated_by FROM ms_customer WHERE id=$1"
	GET_CUST_ID_MEMBER = "SELECT msm.type FROM ms_customer AS msc JOIN ms_member AS msm ON msc.id = msm.id_customer WHERE msc.id = $1 AND msm.expire > CURRENT_DATE"
	GET_CUST_USRID     = "SELECT id,id_user,full_name,NIK,noPhone,email,address,created_at,updated_at,created_by,updated_by FROM ms_customer WHERE id_user = $1"
	GET_ALL_CUSTOMER   = "SELECT id,id_user,full_name,NIK,noPhone,email,address,created_at,updated_at,created_by,updated_by FROM ms_customer"
	GET_CUST_NAME      = "SELECT id,id_user,full_name,NIK,noPhone,email,address,created_at,updated_at,created_by,updated_by FROM ms_customer WHERE full_name = $1"
	EDIT_CUST_ID       = "UPDATE ms_customer SET full_name=$1,NIK=$2,noPhone=$3,email=$4,address=$5,updated_at=$6,updated_by=$7 WHERE id = $8"

	// Scheduler
	GET_ALL_STRIKE        = "SELECT tic.id, tic.id_vehicle, tic.id_credit, tc.interest FROM tx_installment_credit as tic join tx_credit tc on tic.id_credit = tc.id where status = true and  date_payment <= (CURRENT_DATE - INTERVAL '3 months');"
	UPDATE_STATUS_SUSPEND = "update tx_installment_credit set status=$1, suspend=$2 where id = $3"
	UPDATE_INTEREST       = "update tx_credit set interest=$1 where id=$2"
)
