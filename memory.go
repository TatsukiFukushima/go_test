package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	win_a := 0
	win_b := 0
	draw := 0

	for i := 0; i < 10000; i++ {
		cards := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10}
		cards_shuffled := shuffle(cards)
		target := 0 // めくられていないカードの最初の位置を表す。
		point_a := 0
		point_b := 0
		point_get := 0
		for {
			// プレイヤーAのターン
			// cards_shuffled, target, point_get = simple_draw(cards_shuffled, target)
			// cards_shuffled, target, point_get = smart_draw(cards_shuffled, target)
			cards_shuffled, target, point_get = super_smart_draw(cards_shuffled, target, 2.3)
			point_a += point_get
			if target == len(cards_shuffled) {
				break
			}

			// プレイヤーBのターン
			cards_shuffled, target, point_get = simple_draw(cards_shuffled, target)
			// cards_shuffled, target, point_get = smart_draw(cards_shuffled, target)
			// cards_shuffled, target, point_get = super_smart_draw(cards_shuffled, target, 1.1)
			point_b += point_get
			if target == len(cards_shuffled) {
				break
			}
		}

		if point_a > point_b {
			win_a++
		} else if point_a < point_b {
			win_b++
		} else {
			draw++
		}
	}

	fmt.Printf("A勝利: %d\n", win_a)
	fmt.Printf("B勝利: %d\n", win_b)
	fmt.Printf("引き分け: %d\n", draw)
}

// shuffle 与えられた配列をシャッフルして返す
func shuffle(array []int) []int {
	array_len := len(array)
	array_rand := []int{}
	for i := 0; i < array_len; i++ {
		num := rand.Intn(len(array))
		array_rand = append(array_rand, array[num])
		array = append(array[:num], array[num+1:]...)
	}
	return array_rand
}

// simple_draw 従来通りのドロー方法。めくられてない場所を順番にめくる。
// 戻り値: 処理済みカード配列、めくられていないカード位置、ゲットしたカード枚数
func simple_draw(cards []int, target int) ([]int, int, int) {
	point := 0
	if target == 0 {
		// まだ1枚もめくられていない場合
		target += 2
		if cards[0] == cards[1] {
			cards[0] = 0
			cards[1] = 0
			point += 2
		} else {
			return cards, target, point
		}
	}

	// すでにめくられているカードでペアができているか確認
	before_target := target - 1
	before_target_num := cards[before_target]
	if before_target_num != 0 {
		for i := 0; i < before_target; i++ {
			if cards[i] == before_target_num {
				cards[i] = 0
				cards[before_target] = 0
				point += 2
				break
			}
		}
	}

	// 未知のカードをめくる場合（通常）
	for {
		target_num := cards[target]
		already_exist := false
		for i := 0; i < target; i++ {
			if cards[i] == target_num {
				cards[i] = 0
				cards[target] = 0
				point += 2
				target += 1
				already_exist = true
				break
			}
		}

		// すでにめくられているカードの中に、今めくったカードとのペアがなかった場合
		if !already_exist {
			if cards[target] == cards[target+1] {
				cards[target] = 0
				cards[target+1] = 0
				point += 2
				target += 2
			} else {
				target += 2
				break
			}
		}

		if target == len(cards) {
			break
		}
	}

	return cards, target, point
}

// smart_draw 賢いドロー方法。ヒントをなるべく増やさない。
// 戻り値: 処理済みカード配列、めくられていないカード位置、ゲットしたカード枚数
func smart_draw(cards []int, target int) ([]int, int, int) {
	point := 0
	if target == 0 {
		// まだ1枚もめくられていない場合
		target += 2
		if cards[0] == cards[1] {
			cards[0] = 0
			cards[1] = 0
			point += 2
		} else {
			return cards, target, point
		}
	}

	// すでにめくられているカードでペアができているか確認
	before_target := target - 1
	before_target_num := cards[before_target]
	if before_target_num != 0 {
		for i := 0; i < before_target; i++ {
			if cards[i] == before_target_num {
				cards[i] = 0
				cards[before_target] = 0
				point += 2
				break
			}
		}
	}

	// 未知のカードをめくる場合（通常）
	for {
		target_num := cards[target]
		already_exist := false
		exist_cards_count := 0
		for i := 0; i < target; i++ {
			if cards[i] == target_num {
				cards[i] = 0
				cards[target] = 0
				point += 2
				target += 1
				already_exist = true
				break
			}
			if cards[i] != 0 {
				exist_cards_count++
			}
		}

		// すでにめくられているカードの中に、今めくったカードとのペアがなかった場合
		if !already_exist {
			// すでにめくられたカードがない、または全てのペアが分かってしまっている場合、新規カードをめくる
			if exist_cards_count == 0 || len(cards)-target == exist_cards_count+2 {
				if cards[target] == cards[target+1] {
					cards[target] = 0
					cards[target+1] = 0
					point += 2
					target += 2
				} else {
					target += 2
					break
				}
			} else {
				target++
				break
			}
		}

		if target == len(cards) {
			break
		}
	}

	return cards, target, point
}

// super_smart_draw 超賢いドロー方法。めくられた枚数、めくられてない枚数の比率で戦略を変える。
// 戻り値: 処理済みカード配列、めくられていないカード位置、ゲットしたカード枚数
func super_smart_draw(cards []int, target int, param float64) ([]int, int, int) {
	point := 0
	if target == 0 {
		// まだ1枚もめくられていない場合
		target += 2
		if cards[0] == cards[1] {
			cards[0] = 0
			cards[1] = 0
			point += 2
		} else {
			return cards, target, point
		}
	}

	// すでにめくられているカードでペアができているか確認
	before_target := target - 1
	before_target_num := cards[before_target]
	if before_target_num != 0 {
		for i := 0; i < before_target; i++ {
			if cards[i] == before_target_num {
				cards[i] = 0
				cards[before_target] = 0
				point += 2
				break
			}
		}
	}

	// 未知のカードをめくる場合（通常）
	for {
		target_num := cards[target]
		already_exist := false
		exist_cards_count := 0
		for i := 0; i < target; i++ {
			if cards[i] == target_num {
				cards[i] = 0
				cards[target] = 0
				point += 2
				target += 1
				already_exist = true
				break
			}
			if cards[i] != 0 {
				exist_cards_count++
			}
		}

		// すでにめくられているカードの中に、今めくったカードとのペアがなかった場合
		if !already_exist {
			// すでにめくられたカードがない、またはめくられたカードが多い場合、新規カードをめくる
			if exist_cards_count == 0 || float64(len(cards)-target-1)/float64(exist_cards_count+1) < param {
				if cards[target] == cards[target+1] {
					cards[target] = 0
					cards[target+1] = 0
					point += 2
					target += 2
				} else {
					target += 2
					break
				}
			} else {
				target++
				break
			}
		}

		if target == len(cards) {
			break
		}
	}

	return cards, target, point
}
