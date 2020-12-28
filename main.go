package main
 
import (
    "fmt"
)
 
// 関数外で宣言した変数と定数は、同じファイル内のどこからでもアクセスできる
var (
    global1 = 1
    global2 = 10
)
const GlobalConst = "1 + 100 ="
 
func main() {
    // 関数内で宣言した変数や定数はその関数内でしか使用できない
    local1, local2 := 100, 1000
    const LocalConst = "10 * 1000 ="
 
    fmt.Println(GlobalConst, global1+local1)
    fmt.Println(LocalConst, global2*local2)
}