package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:12345678@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

type Student struct {
	gorm.Model
	ID    int
	Name  string
	Age   int
	Grade string
}

type Account struct {
	gorm.Model
	Name string
}

type Transaction struct {
	gorm.Model
	From_account_id int
	To_account_id   int
	Amount          string
}

type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

type Book struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  int    `db:"price"`
}

type User struct {
	gorm.Model
	Name          string
	Posts         []Post `gorm:"foreignKey:UserID"`
	TotalPostsNum int
}

type Post struct {
	gorm.Model
	Title            string
	Content          string
	Comments         []Comment `gorm:"foreignKey:PostID"`
	UserID           uint
	User             User `gorm:"foreignKey:UserID"`
	TotalCommentsNum int
	Status           string
}

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	userres := User{}
	result := tx.Model(&User{}).Where("id = ?", post.UserID).
		UpdateColumn("total_posts_num", gorm.Expr("total_posts_num + ?", 1)).
		Scan(&userres)
	fmt.Println("TotalCommentNum:", userres.TotalPostsNum)
	return result.Error
}

type Comment struct {
	gorm.Model
	Content string
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
	PostID  uint
	Post    Post `gorm:"foreignKey:PostID"`
}

func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	postres := Post{}
	result := tx.Debug().Model(&Post{}).Where("id = ?", comment.PostID).
		UpdateColumn("total_comments_num", gorm.Expr("total_comments_num - ?", 1)).
		Scan(&postres)
	fmt.Println(postres)
	if postres.TotalCommentsNum == 0 {
		postres.Status = "无评论"
		tx.Save(&postres)
	}
	return result.Error
}

func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&Student{})

	//task03_1_1_1
	// student01 := Student{
	// 	Name:  "张三",
	// 	Age:   20,
	// 	Grade: "三年级",
	// }
	// db.Create(&student01)

	//task03_1_1_2
	// students := []Student{}
	// db.Debug().Where("age > 18").Find(&students)
	// fmt.Println(students)

	//task03_1_1_3
	// db.Model(&Student{}).Debug().Where("name", "张三").Update("grade", "四年级")

	//task03_1_1_4
	// db.Debug().Where("age < 15").Delete(&Student{})
	// db.Debug().Unscoped().Where("age < 15").Delete(&Student{})

	//task03_1_2_1
	// db.AutoMigrate(&Account{})
	// db.AutoMigrate(&Transaction{})

	// account_a := Account{
	// 	Name:    "A",
	// 	Balance: big.NewInt(10000).String(),
	// }

	// account_b := Account{
	// 	Name:    "B",
	// 	Balance: big.NewInt(10000).String(),
	// }

	// db.Create(&account_a)
	// db.Create(&account_b)

	// err_t := db.Transaction(func(tx *gorm.DB) error {

	// 	account_a := Account{}
	// 	result_a := tx.Model(&Account{}).Where("ID = ?", "1").First(&account_a)
	// 	if result_a.Error != nil {
	// 		return result_a.Error // 查询报错，回滚
	// 	}
	// 	if result_a.RowsAffected == 0 {
	// 		return fmt.Errorf("account not found") // 自定义 error，也会回滚
	// 	}

	// 	account_b := Account{}
	// 	result_b := tx.Model(&Account{}).Where("ID = ?", "2").First(&account_b)
	// 	if result_b.Error != nil {
	// 		return result_b.Error // 查询报错，回滚
	// 	}
	// 	if result_b.RowsAffected == 0 {
	// 		return fmt.Errorf("account not found") // 自定义 error，也会回滚
	// 	}

	// 	account_a_Balance_int := new(big.Int)
	// 	if _, ok := account_a_Balance_int.SetString(account_a.Balance, 10); !ok {
	// 		return errors.New("转换失败")
	// 	}
	// 	account_a.Balance = new(big.Int).Sub(account_a_Balance_int, big.NewInt(100)).String()
	// 	account_b_Balance_int := new(big.Int)
	// 	if _, ok := account_b_Balance_int.SetString(account_b.Balance, 10); !ok {
	// 		return errors.New("转换失败")
	// 	}
	// 	account_b.Balance = new(big.Int).Add(account_b_Balance_int, big.NewInt(100)).String()

	// 	result_save_a := tx.Save(&account_a)
	// 	if result_save_a.Error != nil {
	// 		return result_save_a.Error // 查询报错，回滚
	// 	}
	// 	if result_save_a.RowsAffected == 0 {
	// 		return fmt.Errorf("account not found") // 自定义 error，也会回滚
	// 	}
	// 	result_save_b := tx.Save(&account_b)
	// 	if result_save_b.Error != nil {
	// 		return result_save_b.Error // 查询报错，回滚
	// 	}
	// 	if result_save_b.RowsAffected == 0 {
	// 		return fmt.Errorf("account not found") // 自定义 error，也会回滚
	// 	}
	// 	transaction_01 := Transaction{
	// 		From_account_id: int(account_a.ID),
	// 		To_account_id:   int(account_b.ID),
	// 		Amount:          big.NewInt(100).String(),
	// 	}

	// 	result_t := tx.Model(&Transaction{}).Create(&transaction_01)
	// 	if result_t.Error != nil || result_t.RowsAffected == 0 {
	// 		return result_t.Error
	// 	}

	// 	return nil
	// })
	// fmt.Println(err_t)

	// account_a := Account{}
	// account_b := Account{}

	// tx := db.Begin()

	// result_a := tx.Model(&Account{}).Where("ID = ?", "1").First(&account_a)
	// if result_a.Error != nil {
	// 	tx.Rollback()
	// 	fmt.Println(result_a.Error)
	// 	return
	// }
	// if result_a.RowsAffected == 0 {
	// 	tx.Rollback()
	// 	fmt.Println("account not found")
	// 	return
	// }

	// result_b := tx.Model(&Account{}).Where("ID = ?", "2").First(&account_b)
	// if result_b.Error != nil {
	// 	tx.Rollback()
	// 	fmt.Println(result_b.Error)
	// 	return
	// }
	// if result_b.RowsAffected == 0 {
	// 	tx.Rollback()
	// 	fmt.Println("account not found")
	// 	return
	// }

	// account_a_Balance_int := new(big.Int)
	// if _, ok := account_a_Balance_int.SetString(account_a.Balance, 10); !ok {
	// 	tx.Rollback()
	// 	fmt.Println("转换失败")
	// 	return
	// }
	// account_a.Balance = new(big.Int).Sub(account_a_Balance_int, big.NewInt(100)).String()
	// account_b_Balance_int := new(big.Int)
	// if _, ok := account_b_Balance_int.SetString(account_b.Balance, 10); !ok {
	// 	tx.Rollback()
	// 	fmt.Println("转换失败")
	// 	return
	// }
	// account_b.Balance = new(big.Int).Add(account_b_Balance_int, big.NewInt(100)).String()

	// result_save_a := tx.Save(&account_a)
	// if result_save_a.Error != nil {
	// 	tx.Rollback()
	// 	fmt.Println(result_save_a.Error)
	// 	return
	// }
	// if result_save_a.RowsAffected == 0 {
	// 	tx.Rollback()
	// 	fmt.Println("account not found")
	// 	return
	// }
	// // result_save_b := tx.Model(&Account{}).Save(&account_b)
	// result_save_b := tx.Save(&account_b)
	// if result_save_b.Error != nil {
	// 	tx.Rollback()
	// 	fmt.Println(result_save_b.Error)
	// 	return
	// }
	// if result_save_b.RowsAffected == 0 {
	// 	tx.Rollback()
	// 	fmt.Println("account not found")
	// 	return
	// }
	// transaction_01 := Transaction{
	// 	From_account_id: int(account_a.ID),
	// 	To_account_id:   int(account_b.ID),
	// 	Amount:          big.NewInt(100).String(),
	// }

	// result_t := tx.Model(&Transaction{}).Create(&transaction_01)
	// if result_t.Error != nil || result_t.RowsAffected == 0 {
	// 	tx.Rollback()
	// 	fmt.Println(result_t.Error)
	// }

	// tx.Commit()
	sqlDB, err := db.DB()

	// sqlxDB := sqlx.NewDb(sqlDB, "mysql")
	/*
		CREATE TABLE employees (
		    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		    name VARCHAR(255),
		    department VARCHAR(255),
		    salary INT
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

		CREATE TABLE books (
		    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		    title VARCHAR(255),
		    author VARCHAR(255),
		    price INT
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	*/
	// employees_org := []Employee{
	// 	{
	// 		Name:       "张三",
	// 		Department: "技术部",
	// 		Salary:     300000,
	// 	},
	// 	{
	// 		Name:       "里斯",
	// 		Department: "销售部",
	// 		Salary:     100000,
	// 	},
	// 	{
	// 		Name:       "王五",
	// 		Department: "后勤部",
	// 		Salary:     50000,
	// 	},
	// }
	// db.Create(employees_org)
	// //task03_2_1_1
	// employees := []Employee{}
	// sqlxDB.Select(&employees, "select * from employees where department = ?", "技术部")
	// fmt.Println(employees)
	// //task03_2_1_2
	// employees_max := []Employee{}
	// sqlxDB.Select(&employees_max, "select * from employees where salary = (select max(salary) from employees)")
	// fmt.Println(employees_max)

	// //task03_2_2_1
	// books := []Book{}
	// sqlxDB.Select(&books, "select * from books where price > 50")
	// fmt.Println(books)

	//task03_3_1_1&2
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Post{})
	// db.AutoMigrate(&Comment{})

	//task03_3_2_1

	// comment1 := Comment{
	// 	Content: "comment_1",
	// 	UserID:  1,
	// 	PostID:  1,
	// }
	// comment2 := Comment{
	// 	Content: "comment_2",
	// 	UserID:  1,
	// 	PostID:  1,
	// }
	// comment3 := Comment{
	// 	Content: "comment_3",
	// 	UserID:  2,
	// 	PostID:  2,
	// }
	// post1 := Post{
	// 	Title:            "title_1",
	// 	Content:          "post_content_1",
	// 	UserID:           1,
	// 	TotalCommentsNum: 2,
	// 	Status:           "",
	// }
	// post2 := Post{
	// 	Title:            "title_2",
	// 	Content:          "post_content_2",
	// 	UserID:           2,
	// 	TotalCommentsNum: 1,
	// 	Status:           "",
	// }
	// user1 := User{
	// 	Name:          "张三",
	// 	TotalPostsNum: 0,
	// }
	// user2 := User{
	// 	Name:          "里斯",
	// 	TotalPostsNum: 0,
	// }
	// db.Create(&user1)
	// db.Create(&user2)
	// db.Create(&post1)
	// db.Create(&post2)
	// db.Create(&comment1)
	// db.Create(&comment2)
	// db.Create(&comment3)

	// user_query_1 := User{}
	// db.Model(&User{}).Preload("Posts").Preload("Posts.Comments").Where("id = ?", "2").First(&user_query_1)
	// fmt.Println(user_query_1)
	// for _, post := range user_query_1.Posts {
	// 	for _, comment := range post.Comments {
	// 		fmt.Println(comment.Content)
	// 	}
	// }
	//task03_3_2_2
	// type Result struct {
	// 	PostID uint
	// 	Count  int
	// }

	// var res Result
	// db.Model(&Comment{}).
	// 	Select("post_id, COUNT(*) as count").
	// 	Group("post_id").
	// 	Order("count DESC").
	// 	Limit(1).
	// 	Scan(&res)

	// fmt.Println(res.PostID, res.Count)

	// comment := Comment{}
	// db.Where("id = ?", 8).Preload("User").Preload("Post").First(&comment)
	// db.Delete(&comment)

	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

}
