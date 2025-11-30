package main 
import (
	"fmt"
	"io"
	"log"
	"os"
)

//go run で実行すると引数ではなくコンパイル対象として解釈されるため、正常に機能しない
//go buildを実行して./Unixcat ch5/fanc.goを実行しないと正常に動かない

func main (){
	fmt.Println("os.Args:", os.Args)  // 引数を確認
	fmt.Println("len(os.Args):", len(os.Args)) //引数の数を確認
	if len (os.Args) < 2 {
		log.Fatal("ファイルが指定されていません")
	}
	f,closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer closer()
	data := make([]byte, 2248)
	//2248バイト分0で初期化している
	fmt.Println(data)
	for {
		//Read(data)は読み込んだファイルでdataに上書きする
		count,err := f.Read(data)
		//ファイルから読み取れたバイト列そのもの末尾の無数の０がバッファの部分
		fmt.Println(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

//ファイルを開いて、クロージャーを返すヘルパー関数
func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, nil
}