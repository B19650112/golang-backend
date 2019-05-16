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
	ID          int       `json:"ID,string"`
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
