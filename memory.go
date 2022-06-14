package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cards := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10}

	for i := 0; i < 1; i++ {
		cards_shuffled := shuffle(cards)
		target := 0 // めくられていないカードの最初の位置を表す。
		point_a := 0
		point_b := 0
		point_get := 0
		for {
			// プレイヤーAのターン
			cards_shuffled, target, point_get = simple_draw(cards_shuffled, target)
			fmt.Printf("target: %d\n", target)
			fmt.Printf("A結果: %v\n", cards_shuffled)
			point_a += point_get
			if target == len(cards_shuffled) {
				break
			}

			// プレイヤーBのターン
			cards_shuffled, target, point_get = simple_draw(cards_shuffled, target)
			fmt.Printf("target: %d\n", target)
			fmt.Printf("B結果: %v\n", cards_shuffled)
			point_b += point_get
			if target == len(cards_shuffled) {
				break
			}
		}

		fmt.Printf("A点数: %d\n", point_a)
		fmt.Printf("B点数: %d\n", point_b)
	}
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
