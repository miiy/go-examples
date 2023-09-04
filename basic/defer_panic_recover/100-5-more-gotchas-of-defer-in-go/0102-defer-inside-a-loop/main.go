package main

// #2 — Defer inside a loop
// #2 — 循环内的defer

// 不要在循环内使用defer，除非你确定自己在做什么。它可能无法按预期工作。

//func() {
//	for {
//		row, err := db.Query("SELECT ...")
//		if err != nil {
//			..
//		}
//		defer row.Close()
//	}
//	// deferred funcs run here
//}

// Solution #1:
// 直接调用它，无需defer
//func() {
//	for {
//		row, err := db.Query("SELECT ...")
//		if err != nil {
//			..
//		}
//		row.Close()
//	}
//}

// Solution #2:
// 将 work 委托给另一个函数并在那里使用 defer，这里，deferred func 将在每次 anonymous func 结束后运行
//func()  {
//	for {
//		func() {
//			row, err := db.Query("SELECT ...")
//			if err != nil {
//				..
//			}
//			defer row.Close()
//			..
//			// deferred func runs here
//		}()
//	}
//}
