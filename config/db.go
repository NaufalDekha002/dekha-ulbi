package config

import (
	"os"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/whatsauth"
)

//var IteungIPAddress string = os.Getenv("ITEUNGBEV1")

var MongoString string = os.Getenv("MONGOSTRING")

var MariaStringAkademik string = os.Getenv("MARIASTRINGAKADEMIK")

var DBUlbimariainfo = atdb.DBInfo{
	DBString: MariaStringAkademik,
	DBName:   "us7p83oqsjfxqbn7",
}

var DBUlbimongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "tugas5_db",
}

var Ulbimariaconn = atdb.MariaConnect(DBUlbimariainfo)

var Ulbimongoconn = atdb.MongoConnect(DBUlbimongoinfo)

var Usertables = [4]whatsauth.LoginInfo{mhs, dosen, user, user1}

var mhs = whatsauth.LoginInfo{
	Userid:   "MhswID",
	Password: "Password",
	Phone:    "Telepon",
	Username: "Login",
	Uuid:     "simak_mst_mahasiswa",
	Login:    "2md5",
}

var dosen = whatsauth.LoginInfo{
	Userid:   "NIDN",
	Password: "Password",
	Phone:    "Handphone",
	Username: "Login",
	Uuid:     "simak_mst_dosen",
	Login:    "2md5",
}

var user = whatsauth.LoginInfo{
	Userid:   "user_id",
	Password: "user_password",
	Phone:    "phone",
	Username: "user_name",
	Uuid:     "simak_besan_users",
	Login:    "2md5",
}

var user1 = whatsauth.LoginInfo{
	Userid:   "user_id",
	Password: "user_password",
	Phone:    "user_phone",
	Username: "user_name",
	Uuid:     "besan_users",
	Login:    "2md5",
}
