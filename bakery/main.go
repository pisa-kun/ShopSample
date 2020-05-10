package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

)

// Json化するときはフィールドのサフィックスはキャメルケースにする
type chashier struct {
	Product_list []bread `json:"bread"`
	Shop_name    string  `json:"sname"`
	Name         string  `json:"name"`
}

type bread struct {
	Name   string `json:"name"`   // 商品名
	Value  int    `json:"value"`  // 価格
	Number int    `json:"number"` // 個数
}

var list []bread

var version = "1.0.0.0"

func (c *chashier) productShow() (b bool) {
	if c.Product_list == nil {
		return false
	}
	fmt.Printf("---[%v shop] show now item list ! ---\n chashier : %v\n", c.Shop_name, c.Name)
	for _, item := range c.Product_list {
		fmt.Printf("[item name : %v] [item value : %v] [item last number : %v]\n",
			item.Name, item.Value, item.Number)
	}
	return true
}

func (c *chashier) productReCalculate(n string, num int) (b bool) {
	if c.Product_list == nil {
		return false
	}

	for count, item := range c.Product_list {
		if n == item.Name {
			if item.Number < num {
				fmt.Println("入力された数より在庫数が少ないです")
				fmt.Println(item.Name)
				fmt.Printf("在庫 : %v, 入力数 : %v\n", item.Number, num)
				return false
			}
			c.Product_list[count].Number = item.Number - num
			return true
		}
	}
	return false
}

func (c *chashier) saveJson() (b bool) {
	output, err := json.MarshalIndent(&c, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to Json:", err)
		return false
	}
	dstname := c.Shop_name + ".json"
	if err = ioutil.WriteFile(dstname, output, 0644); err != nil {
		fmt.Println("Error writing Json to File:", err)
		return false
	}
	return true
}

func main() {

	var versionshow bool
	var help bool
	// -v, -versionが指定されたときに真になるよう定義
	// -help, -hが指定されたときにApplicationの説明
	flag.BoolVar(&versionshow, "v", false, "show version")
	flag.BoolVar(&versionshow, "version", false, "show version ")
	flag.BoolVar(&help, "help", false, "show use app")
	flag.BoolVar(&help, "h", false, "show use app")
	flag.Parse() // 引数からオプションをパース

	if versionshow {
		// バージョン情報を表示して終了
		fmt.Println("version", version)
		return
	}

	if help {
		fmt.Println("this application is bread sample")
		fmt.Println("show bread list and number")

		return
	}

	// Hello World in Command Line Tool
	printWelcome()

	var pantazia chashier
	jsonData, e := ReadandOpenJson()

	// jsonがなかった場合などは初期化
	if e != true {
		InitializeList()
		pantazia = chashier{
			Product_list: list,
			Shop_name:    "pantazia",
			Name:         "yamamoto",
		}
	} else {
		// jsonから読み込んだデータを使う
		json.Unmarshal(jsonData, &pantazia)
	}

	if b := pantazia.productShow(); b != true {
		fmt.Println("product is nil")
	}

	var name string
	fmt.Println("欲しい商品は？")
	fmt.Scan(&name)

	var nameBool bool = false
	var target bread
	for _, l := range pantazia.Product_list {
		if name == l.Name {
			target = l
			nameBool = true
			break
		}
	}

	if !nameBool {
		fmt.Println("欲しい商品名がなかったのでやり直し！")
		return
	}

	var nameN string
	fmt.Println("何個ご購入されますか？")
	fmt.Scan(&nameN)

	i, err := strconv.Atoi(nameN)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := pantazia.productReCalculate(target.Name, i); err != true {
		fmt.Println("return false")
	}
	pantazia.productShow()

	if err := pantazia.saveJson(); err != true {
		fmt.Println("save json error")
	}

}

func ReadandOpenJson() (by []byte, b bool) {

	// jsonのオープン、読み込み
	jsonFile, err := os.Open("pantazia.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return by, false
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return by, false
	}
	return jsonData, true
}

func InitializeList() {
	var sandwitch bread = bread{Name: "sandwitch", Value: 120, Number: 3}
	kurowasan := bread{
		Name:   "kurowassan",
		Value:  80,
		Number: 5,
	}
	syoppan := bread{
		Name:   "syoppan",
		Value:  250,
		Number: 10,
	}
	donuts := bread{
		Name:   "donuts",
		Value:  130,
		Number: 20,
	}
	agepan := bread{
		Name:   "agepan",
		Value:  130,
		Number: 10,
	}
	denish := bread{
		Name:   "denish",
		Value:  160,
		Number: 2,
	}
	list = []bread{sandwitch, kurowasan, syoppan, donuts, agepan, denish}
}

// ようこそっぽい文字表示
func printWelcome() {
	fmt.Println("-------------------------------")
	fmt.Println("-------                 -------")
	fmt.Println("-------  w e l c o m e  -------")
	fmt.Println(time.Now().UTC())
	fmt.Println("-------------------------------")
	fmt.Println("-------------------------------")
}
