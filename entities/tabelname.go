package entities

// TabelUser : Structure tabeluser
type TabelUser struct {
	ID		    string    `json:"ID"`
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	Level       string    `json:"level"`
	Email       string    `json:"email"`
	NoTelpon    string    `json:"notelpon"`
	Token       string    `json:"token"` 
}
// TabelProduct : Structure tabelproduct
type TabelProduct struct {
	ID          string    `json:"ID"`
	ImagePath   string    `json:"imagepath"`
	Title		string    `json:"title"`
	Description string    `json:"description"`
	Price		int       `json:"price,string"`
	IDTemp      string    `json:"IDtemp"`
}
// TabelResume : Structure tabelresume
type TabelResume struct {
	ID          string    `json:"ID"`
	Name		string    `json:"name"`
	Title		string    `json:"title"`
	Photo		string    `json:"photo"`
	Description	string    `json:"description"`
}
// TabelCount : Structure tabelcount
type TabelCount struct {
	ID                string  `json:"ID"`
	CountProduct      int     `json:"countproduct"`
	CountProductTemp  int     `json:"countproducttemp"`
}
// StartEnd : Structure StartId and EndId
type StartEnd struct {
	StartID     int     `json:"startID"`
	EndID		int     `json:"endID"`
}
// DefaultPage : Structure defaultpage
type DefaultPage struct {
	ID          string  `json:"ID"`
	DefaultPage int     `json:"defaultpage"`
}