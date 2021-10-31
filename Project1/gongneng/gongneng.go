package gongneng

import (
	"fmt"
)

//單條信息
type Mx struct {
	Zhanghu float64
	Money   float64
	Beizhu  string
}

//總信息
type Book struct {
	total float64
	m     []Mx
}

//接受参数为1时
func (b Book) Mingxi() {
	fmt.Println("----------当前收支明细----------")
	fmt.Println("收支    账户金额   收支金额   说明")
	for _, v := range b.m {
		if v.Money >= 0 {
			fmt.Printf("收入     %v      %v    %v\n", v.Zhanghu, v.Money, v.Beizhu)

		} else {
			fmt.Printf("支出    %v      %v    %v\n", v.Zhanghu, v.Money, v.Beizhu)
		}
	}
	fmt.Println()
}

//接收参数为2时 增加收入

func (b *Book) Shouru() {
	var f float64
	var bz string
	for {
		fmt.Println("收入:")
		_, err := fmt.Scanln(&f)
		if err != nil {
			fmt.Println("输入错误请重新输入！(数值类型)")
			continue
		}
		break
	}
	for {
		fmt.Println("说明:")
		_, er := fmt.Scanln(&bz)
		if er != nil {
			fmt.Println("输入错误请重新输入！(文本类型)")
			continue
		}
		break
	}
	b.total += f
	b.m = append(b.m, Mx{Zhanghu: b.total, Money: f, Beizhu: bz})
	fmt.Println("登记成功")
}

//接受参数为3时 登记支出

func (b *Book) Zhichu() {
	var f float64
	var bz string
	for {
		fmt.Println("支出:")
		_, err := fmt.Scanln(&f)
		if err != nil {
			fmt.Println("输入错误请重新输入！(数值类型)")
			continue
		}
		break
	}
	for {
		fmt.Println("说明:")
		_, er := fmt.Scanln(&bz)
		if er != nil {
			fmt.Println("输入错误请重新输入！(文本类型)")
			continue
		}
		break
	}
	b.total -= f
	b.m = append(b.m, Mx{Zhanghu: b.total, Money: -f, Beizhu: bz})
	fmt.Println("登记成功")
}
