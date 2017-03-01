package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出 {{end}} \n"))
	tEmpty.Execute(os.Stdout, nil)

	//template.Must是检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写
	tWithValue := template.New("test")
	tWithValue = template.Must(tWithValue.Parse("不为空 : {{if `sss`}} 这里会输出 {{end}} \n"))
	tWithValue.Execute(os.Stdout, nil)

	tIf := template.New("test if else")
	tIf = template.Must(tIf.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIf.Execute(os.Stdout, nil)

	a := false
	tWithStuct := template.New("test if")
	//注意 这里‘{{}}’中就不需要再套一个‘{{}}’了  比如{{if {{.}} }}会报错
	//另外：if里面只能是bool值  Mail=="astaxie@gmail.com"这种是不行的！
	tWithStuct = template.Must(tWithStuct.Parse("if-else demo: {{if . }} if部分 {{else}} else部分.{{end}}\n"))
	tWithStuct.Execute(os.Stdout, a)

	fmt.Println("\n")

	//模板变量的声明方式
	tValT := template.New("test value")
	//这里输出 output
	tValT, _ = tValT.Parse(`{{with $x := "output"}} {{printf "%q" $x}} {{end}}`)
	tValT, _ = tValT.Parse(`{{with $x := "output" | printf "%q"}} {{$x}} {{end}}`)
	tValT, _ = tValT.Parse(`{{with $x := "123" }} {{$x | printf "%q"}} {{end}}`)
	tValT.Execute(os.Stdout, nil)

}
