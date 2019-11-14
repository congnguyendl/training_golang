package main

import (
	"fmt"
	"github.com/twinj/uuid"
)

func main() {

	// In ra chuỗi dữ liệu
	fmt.Println("Hello","dasdas",1)

	// khai báo biến pas1, pas2 nhân dữ liệu trả về từ hàm LayPassWifi()
	pas1, pas2 := LayPassWifi()
	fmt.Println("Mat Khau 1 :",pas1)
	fmt.Println("Mat Khau 2 :",pas2)

	DinhNghiaBien()
}

// Khai báo hàm wifi có kiểu dữ liệu trả về nhiều đối số (trả về 2 biến string)
func LayPassWifi() (string, string) {
	matKhau1 := uuid.NewV4()
	matKhau2 := uuid.NewV4()
	return matKhau1.String(), matKhau2.String()
}

func DinhNghiaBien() {
	//Khai báo không khởi tạo mặc định giá trị biến la 0
	var a int
	fmt.Println(a)

	// Khai báo có khởi tạo giá trị
	var so1 float64 = 9
	fmt.Println(so1)

	// Khai báo 2 biến i, j có kiểu dữ liệu là int và khởi tạo luôn giá trị cho chúng
	var i, j int = 1, 2

	fmt.Println(i, j)

	// Khai báo 3 biến c, python, java đều có kiểu dữ liệu là bool
	var c, python, java bool
	fmt.Println(c, python, java)

	// Khai báo ngắn gọn biến k và khởi tạo giá trị luôn cho nó.
	// Không dùng từ khóa var mà dùng dấu hai chấm,
	// lúc này kiểu dữ liệu sẽ được ngầm định tùy theo giá trị của biến.
	k := 3
	fmt.Println(k)

	/* Khi khai báo biến mà không khởi tạo giá trị ban đầu cho nó thì biến đó sẽ có giá trị zero tùy thuộc vào kiểu dữ liệu:

	- 0 cho kiểu số.
	-	false cho kiểu boolean.
	- "" cho kiểu chuỗi.

	*/
}
