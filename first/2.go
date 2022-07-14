package main

import (
	"fmt"
	"log"
)

const spaceCh = "."
const barrierCh = "*"
const cowCh = "C"
const peopleCh = "F"

type node struct {
	direct int //1上 2右 3下 4左
	x      int
	y      int
}

func main() {
	table := [10][10]string{}
	scan := ""
	step := 0
	maxSteps := 1000
	c := new(node)
	p := new(node)
	for i := 0; i < len(table); i++ {
		_, err := fmt.Scan(&scan)
		if err != nil {
			log.Fatal("输入错误: ", err)
		}
		for k := 0; k < len(scan); k++ {
			//初始化牛和人的坐标
			if string(scan[k]) == cowCh {
				c.x = i
				c.y = k
				c.direct = 1
			} else if string(scan[k]) == peopleCh {
				p.x = i
				p.y = k
				p.direct = 1
			}
			table[i][k] = string(scan[k])
		}
	}

	fmt.Println("矩阵：")
	for _, v := range table {
		for _, v1 := range v {
			fmt.Print(v1)
		}
		fmt.Println()
	}

	//最大步数不超过1000，超过就陷入死循环，或者记录某个状态，如果再走一次，说明也进入死循环了
	for step <= maxSteps {
		//牛往上走
		if c.direct == 1 {
			if !isBlocked(c.x, c.y, c.direct, table) {
				//下一步是人
				if table[c.x-1][c.y] == peopleCh {
					//人是否阻塞，阻塞才能捕捉
					if isBlocked(p.x, p.y, p.direct, table) {
						//移动
						table[c.x-1][c.y] = cowCh
						table[c.x][c.y] = spaceCh
						//更新坐标
						c.x = c.x - 1
						fmt.Println(step + 1)
						break
					}
				} else if table[c.x-1][c.y] == spaceCh { //下一步是空格
					//移动
					table[c.x-1][c.y] = cowCh
					table[c.x][c.y] = spaceCh
					//更新坐标
					c.x = c.x - 1
				}
			} else { //遇到墙或者障碍物，不能移动，顺时针改变方向
				c.direct = 2
			}
		} else if c.direct == 2 { //牛往右走
			if !isBlocked(c.x, c.y, c.direct, table) {
				//下一步是人，捕获
				if table[c.x][c.y+1] == peopleCh {
					if isBlocked(p.x, p.y, p.direct, table) {
						//移动
						table[c.x][c.y+1] = cowCh
						table[c.x][c.y] = spaceCh
						//更新坐标
						c.y = c.y + 1
						fmt.Println(step + 1)
						break
					}
				} else if table[c.x][c.y+1] == spaceCh { //下一步是空格

					//移动
					table[c.x][c.y+1] = cowCh
					table[c.x][c.y] = spaceCh
					//更新坐标
					c.y = c.y + 1

				}
			} else { //遇到墙或者障碍物，不能移动,顺时针改变方向
				c.direct = 3
			}
		} else if c.direct == 3 { //牛往下走
			if !isBlocked(c.x, c.y, c.direct, table) {
				//下一步是人，捕获
				if table[c.x+1][c.y] == peopleCh {
					if isBlocked(p.x, p.y, p.direct, table) {
						//移动
						table[c.x+1][c.y] = cowCh
						table[c.x][c.y] = spaceCh
						//更新坐标
						c.x = c.x + 1
						fmt.Println(step + 1)
						break
					}
				} else if table[c.x+1][c.y] == spaceCh { //下一步是空格
					//移动
					table[c.x+1][c.y] = cowCh
					table[c.x][c.y] = spaceCh
					//更新坐标
					c.x = c.x + 1
				}
			} else { //遇到墙或者障碍物，不能移动,顺时针改变方向
				c.direct = 4
			}
		} else if c.direct == 4 { //牛往左走
			if !isBlocked(c.x, c.y, c.direct, table) {
				//下一步是人，捕获
				if table[c.x][c.y-1] == peopleCh {
					if isBlocked(p.x, p.y, p.direct, table) {
						//移动
						table[c.x][c.y-1] = cowCh
						table[c.x][c.y] = spaceCh
						//更新坐标
						c.y = c.y - 1
						fmt.Println(step + 1)
						break
					}
				} else if table[c.x][c.y-1] == spaceCh { //下一步是空格
					//移动
					table[c.x][c.y-1] = cowCh
					table[c.x][c.y] = spaceCh
					//更新坐标
					c.y = c.y - 1
				}
			} else { //遇到墙或者障碍物，不能移动,顺时针改变方向
				c.direct = 1
			}
		}

		//人往上走
		if p.direct == 1 {
			if !isBlocked(p.x, p.y, p.direct, table) {
				//下一步是牛
				if table[p.x-1][p.y] == cowCh {
					if isBlocked(c.x, c.y, c.direct, table) {
						//移动
						table[p.x-1][p.y] = peopleCh
						table[p.x][p.y] = spaceCh
						//更新坐标
						p.x = p.x - 1
						fmt.Println(step + 1)
						break
					}
				} else if table[p.x-1][p.y] == spaceCh { //下一步是空格
					//移动
					table[p.x-1][p.y] = peopleCh
					table[p.x][p.y] = spaceCh
					//更新坐标
					p.x = p.x - 1
				}
			} else { //遇到墙或者障碍物，不能移动,顺时针改变方向
				p.direct = 2
			}
		} else if p.direct == 2 { //人往右走
			if !isBlocked(p.x, p.y, p.direct, table) {
				//下一步是人，捕获
				if table[p.x][p.y+1] == cowCh {
					if isBlocked(c.x, c.y, c.direct, table) {
						//移动
						table[p.x][p.y+1] = peopleCh
						table[p.x][p.y] = spaceCh
						//更新坐标
						p.y = p.y + 1
						fmt.Println(step + 1)
						break
					}
				} else if table[p.x][p.y+1] == spaceCh { //下一步是空格
					//移动
					table[p.x][p.y+1] = peopleCh
					table[p.x][p.y] = spaceCh
					//更新坐标
					p.y = p.y + 1
				}
			} else { //遇到墙或者障碍物，不能移动,顺时针改变方向
				p.direct = 3
			}
		} else if p.direct == 3 { //人往下走
			if !isBlocked(p.x, p.y, p.direct, table) {
				//下一步是人，捕获
				if table[p.x+1][p.y] == cowCh {
					if isBlocked(c.x, c.y, c.direct, table) {
						//移动
						table[p.x+1][p.y] = peopleCh
						table[p.x][p.y] = spaceCh
						//更新坐标
						p.x = p.x + 1
						fmt.Println(step + 1)
						break
					}
				} else if table[p.x+1][p.y] == spaceCh { //下一步是空格
					//移动
					table[p.x+1][p.y] = peopleCh
					table[p.x][p.y] = spaceCh
					//更新坐标
					p.x = p.x + 1
				}
			} else { //遇到墙或者障碍物，不能移动,顺时针改变方向
				p.direct = 4
			}
		} else if p.direct == 4 {
			if !isBlocked(p.x, p.y, p.direct, table) {
				//下一步是人，捕获
				if table[p.x][p.y-1] == cowCh {
					if isBlocked(c.x, c.y, c.direct, table) {
						//移动
						table[p.x][p.y-1] = peopleCh
						table[p.x][p.y] = spaceCh
						//更新坐标
						p.y = p.y - 1
						fmt.Println(step + 1)
						break
					}
				} else if table[p.x][p.y-1] == spaceCh { //下一步是空格
					//移动
					table[p.x][p.y-1] = peopleCh
					table[p.x][p.y] = spaceCh
					//更新坐标
					p.y = p.y - 1
				}
			} else { //遇到墙或者障碍物，不能移动,顺时针改变方向
				p.direct = 1
			}
		}
		step++
	}
	fmt.Println(step)
	fmt.Println("矩阵：")
	for _, v := range table {
		for _, v1 := range v {
			fmt.Print(v1)
		}
		fmt.Println()
	}
}

// x,y处的是否在direct方向上阻塞
func isBlocked(x, y, direct int, table [10][10]string) bool {
	switch direct {
	case 1:
		return x-1 == -1 || table[x-1][y] == barrierCh
	case 2:
		return y+1 == 10 || table[x][y+1] == barrierCh
	case 3:
		return x+1 == 10 || table[x+1][y] == barrierCh
	case 4:
		return y-1 == -1 || table[x][y-1] == barrierCh
	default:
		fmt.Println("输入错误")
	}
	return false
}
