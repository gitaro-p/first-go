package main
 
import "fmt"
 
func main() {
 
    // ダブルクォートで囲うリテラル記述例
    var s string
    s = "ダブルクォートで囲う文字列リテラル"
    fmt.Println(s)
 
    // エスケープシーケンスの使用例
    escape := "エスケープシーケンスの使用例\n改行と\tタブを使用"
    fmt.Println(escape)
 
}