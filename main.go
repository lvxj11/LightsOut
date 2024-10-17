package main

import "fmt"

// 判断传入灯泡状态是否全开，返回布尔值
func all_on(state [9]bool) bool {
	for _, v := range state {
		if !v {
			return false
		}
	}
	return true
}

// 根据传入状态和灯泡编号翻转指定灯泡及横向纵向相邻灯泡，返回翻转结果
func flip(state [9]bool, num int) [9]bool {
	// 翻转指定灯泡
	state[num] = !state[num]
	row := num / 3
	col := num % 3
	// 翻转横向相邻灯泡
	if col == 0 {
		state[num+1] = !state[num+1]
	} else if col == 1 {
		state[num+1] = !state[num+1]
		state[num-1] = !state[num-1]
	} else {
		state[num-1] = !state[num-1]
	}
	// 翻转纵向相邻灯泡
	if row == 0 {
		state[num+3] = !state[num+3]
	} else if row == 1 {
		state[num+3] = !state[num+3]
		state[num-3] = !state[num-3]
	} else {
		state[num-3] = !state[num-3]
	}
	// 等待任意键
	// fmt.Scanln()
	return state
}

// 循环遍历指定范围灯泡并检测是否全开
func check(state [9]bool, steps []int, start int) (bool, []int) {
	// 循环遍历
	for i := start; i <= 8; i++ {
		// 判断翻转后状态是否全开，全开则返回步骤
		if all_on(flip(state, i)) {
			// 添加当前步骤到步骤列表
			steps = append(steps, i)
			return true, steps
		}
	}
	return false, steps
}

// 解出3*3灯泡谜题，返回解决步骤，为空为未解决
func solve(state [9]bool, steps []int, level int) []int {
	for i := level; i <= 8; i++ {
		// 尝试单次遍历灯泡，如果可以解决则返回步骤
		if all, steps := check(state, steps, level); all {
			return steps
		}
		// 未找到解法，翻转后继续递归
		tmp_steps := append(steps, i)
		state := flip(state, i)
		re_steps := solve(state, tmp_steps, i+1)
		// 对比返回值长度，如果有变化返回返回值，没有变化则舍弃翻转返回原步骤
		if len(re_steps) != len(tmp_steps) {
			return re_steps
		}
	}
	return steps
}
func main() {
	// 设置灯泡初始状态
	var state = [9]bool{true, false, false, false, false, false, false, false, false}
	// 检查初始状态是否为全开
	if all_on(state) {
		fmt.Println("已经是全开状态！")
		return
	}
	// 运行解决函数传入初始状态
	steps := solve(state, []int{}, 0)
	if len(steps) == 0 {
		fmt.Println("无法解决！")
	} else {
		fmt.Println("解决步骤：")
		for _, step := range steps {
			// 输出格式为：第几行第几列
			fmt.Printf("第%d行第%d列\n", step/3+1, step%3+1)
		}
	}
}
