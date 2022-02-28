package structs

// Posts is a representation of a post
type User struct {
	Userid			   int		 `json:"userid" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name         	   string    `json:"name"`
	Age      		   int   	 `json:"age"`
	Password      	   string    `json:"password"`
}
type Risk_profile struct {
	Id			 int	 	 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	MM           float32     `json:"mm"`
	BOND      	 float32   	 `json:"bond"`
	STOCK      	 float32   	 `json:"stock"`
	Userid     	 int		 `json:"userid"`
	User     	 User		 `gorm:"foreignKey:Userid"`


}

// Result is an array of post
type Result struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
}