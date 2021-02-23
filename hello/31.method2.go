package main

import "fmt"

// 定义一些常量
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

// Color作为byte的别名
type Color byte

// 定义了一个struct:Box，含有三个长宽高字段和一个颜色属性
type Box struct {
	wdith, height, depth float64
	color                Color
}

// 定义了一个slice:BoxList，含有Box
type BoxList []Box // a slice of boxes

// Volume()定义了接收者为Box，返回Box的容量
func (b Box) Volume() float64 {
	return b.wdith * b.height * b.depth
}

// SetColor(c Color)，把Box的颜色改为c
func (b *Box) SetColor(c Color) {
	b.color = c
}

// BiggestColor()定在在BoxList上面，返回list里面容量最大的颜色
func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

// PaintItBlack()把BoxList里面所有Box的颜色全部变成黑色
func (bl BoxList) PaintItBlack() {
	for i := range bl {
		bl[i].SetColor((BLACK))
	}
}

// String()定义在Color上面，返回Color的具体颜色(字符串格式)
func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))                         // We have 6 boxes in our set
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")         // The volume of the first one is 64 cm³
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String()) // The color of the last one is YELLOW
	fmt.Println("The biggest one is", boxes.BiggestColor().String())                // The biggest one is YELLOW

	fmt.Println("Let's paint them all black") // Let's paint them all black
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String()) // The color of the second one is BLACK

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String()) // Obviously, now, the biggest one is BLACK
}
