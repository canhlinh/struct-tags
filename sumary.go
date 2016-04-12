package main

type WhatIsAStructTag struct {
	Content    string `define:"Là một chuỗi string tùy ý được gắn vào các trường của Struct"`
	Reflection string `key:"value"`
}

type NameSpacing struct {
	Content string `define:"Là một cơ chế cho phép khai báo nhiều tag trên một trường"`
	Example string `json:"example" gorm:"example" validate:"required"`
}

type Serialization struct {
	Content string `purpose:"Quy định tên của trường tương ứng khi thực hiện marshalling and unmarshalling struct"`
	Example string `json:"example_field_name" gorm:"example_mysql_column"`
}

type Validation struct {
	Content string `purpose:"Xác nhận giá trị của trường có hợp lệ hay không"`
	Example string `valid:"required"`
}

type Parameters struct {
	Content   string  `purpose:"Hỗ trợ việc lấy parameters từ request or command line"`
	Example   string  `param:"username"`
	TimeStamp float64 `param:"start_at"`
}
