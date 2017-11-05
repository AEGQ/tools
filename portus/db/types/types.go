package types

type Registry struct {
	ID       int
	Name     string
	Hostname string
	//Created_at string
	//Updated_at string
	Use_ssl bool
}

type Team struct {
	ID   int
	Name string
	//Created_at  string
	//Updated_at  string
	Hidden int
	//Description string
}

type NameSpace struct {
	ID   int
	Name string
	//Created_at  string
	//Updated_at  string
	Team_id     int
	Registry_id int
	Global      bool
	//Description string
	Visibility int
}

type Repository struct {
	ID           int
	Name         string
	Namespace_id int
	//Created_at   string
	//Updated_at   string
	Marked int
}

type User struct {
	ID       int
	Username string
	//Email                  string
	//Encrypted_password     string
	//Reset_password_token   string
	//Reset_password_sent_at string
	//Remember_created_at    string
	//Sign_in_count          int
	//Current_sign_in_at     string
	//Last_sign_in_at        string
	//Current_sign_in_ip     string
	//Last_sign_in_ip        string
	//Created_at             string
	//Updated_at             string
	Admin     int
	Enabled   int
	Ldap_name string
	//Failed_attempts        int
	//Locked_at              string
	Namespace_id int
	//Display_name           string
}

type Tag struct {
	ID            int
	Name          string
	Repository_id int
	//Created_at    string
	//Updated_at    string
	User_id int
	//Digest        string
	//Image_id      string
	Marked int
}

type Star struct {
	ID            int
	User_id       int
	Repository_id int
	Created_at    string
	Updated_at    string
}

type TeamUser struct {
	ID      int
	User_id int
	Team_id int
	//Created_at string
	//Updated_at string
	Role int
}
