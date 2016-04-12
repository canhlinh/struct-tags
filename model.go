package main

type User struct {
	ID        uint64 `sql:"primary_key" json:"-"`
	FirstName string `sql:"type:varchar(20)" json:"first_name" valid:"alpha,required,length(2|45)" param:"first_name"`
	LastName  string `sql:"type:varchar(20)" json:"last_name" valid:"alpha,required,length(0|10)" param:"last_name"`
}

type ErrorResponse struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}

type JsonObject struct {
	Object ObjectType `json:"object"`
}

type ObjectType struct {
	BufferSize int             `json:"buffer_size"`
	Databases  []DatabasesType `json:"databases"`
}

type DatabasesType struct {
	Host   string       `json:"host"`
	User   string       `json:"user"`
	Pass   string       `json:"password"`
	Type   string       `json:"type"`
	Name   string       `json:"name"`
	Tables []TablesType `json:"tables"`
}

type TablesType struct {
	Name     string      `json:"name"`
	Statment string      `json:"statment"`
	Regex    string      `json:"regex"`
	Types    []TypesType `json:"types"`
}

type TypesType struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type Post struct {
	Title    string `valid:"alphanum,required"`
	Message  string `valid:"duck,ascii"`
	AuthorIP string `valid:"ipv4"`
	Date     string `valid:"-"`
}

type SignupForm struct {
	Name  string `param:"name"`
	Email string `param:"email_address"`
	// We use a struct tag with "-" to ignore a value.
	Password string `param:"-"`
}
