package thinkLog

import "testing"

func TestSetLogFileTask(t *testing.T) {
	var i int
	SetLogFileTask()
	for true {
		i++
		DebugLog.Println()
		DebugLog.Println()
		DebugLog.Println()
		DebugLog.Println()
		DebugLog.Println()
		DebugLog.Print(`const (
    // OK表示没有出错
    OK  ErrorCode = iota
    // 当上下文环境有歧义时导致ErrAmbigContext：
    // 举例：
    //   <a href="{{if .C}}/path/{{else}}/search?q={{end}}{{.X}}"&rt;
    // 说明：
    //   {{.X}}的URL上下文环境有歧义，因为根据{{.C}}的值，
    //   它可以是URL的后缀，或者是查询的参数。
    //   将{{.X}}移动到如下情况可以消除歧义：
    //   <a href="{{if .C}}/path/{{.X}}{{else}}/search?q={{.X}}{{end}}"&rt;
    ErrAmbigContext
    // 期望空白、属性名、标签结束标志而没有时，标签名或无引号标签值包含非法字符时，
    // 会导致ErrBadHTML；举例：
    //   <a href = /search?q=foo&rt;
    //   <href=foo&rt;
    //   <form na<e=...&rt;
    //   <option selected<
    // 讨论：
    //   一般是因为HTML元素输入了错误的标签名、属性名或者未用引号的属性值，导致解析失败
    //   将所有的属性都用引号括起来是最好的策略
    ErrBadHTML
    // {{if}}等分支不在相同上下文开始和结束时，导致ErrBranchEnd
    // 示例：
    //   {{if .C}}<a href="{{end}}{{.X}}
    // 讨论：
    //   html/template包会静态的检验{{if}}、{{range}}或{{with}}的每一个分支，
    //   以对后续的pipeline进行转义。该例出现了歧义，{{.X}}可能是HTML文本节点，
    //   或者是HTML属性值的URL的前缀，{{.X}}的上下文环境可以确定如何转义，但该
    //   上下文环境却是由运行时{{.C}}的值决定的，不能在编译期获知。
    //   这种问题一般是因为缺少引号或者角括号引起的，另一些则可以通过重构将两个上下文
    //   放进if、range、with的不同分支里来避免，如果问题出现在参数长度一定非0的
    //   {{range}}的分支里，可以通过添加无效{{else}}分支解决。
    ErrBranchEnd
    // 如果以非文本上下文结束，则导致ErrEndContext
    // 示例：
    //   <div
    //   <div title="no close quote&rt;
    //   <script>f()
    // 讨论：
    //   执行模板必须生成HTML的一个文档片段，以未闭合标签结束的模板都会引发本错误。
    //   不用在HTML上下文或者生成不完整片段的模板不应直接执行。
    //   {{define "main"}} <script&rt;{{template "helper"}}</script> {{end}}
    //   {{define "helper"}} document.write(' <div title=" ') {{end}}
    //   模板"helper"不能生成合法的文档片段，所以不直接执行，用js生成。
    ErrEndContext
    // 调用不存在的模板时导致ErrNoSuchTemplate
    // 示例：
    //   {{define "main"}}<div {{template "attrs"}}&rt;{{end}}
    //   {{define "attrs"}}href="{{.URL}}"{{end}}
    // 讨论：
    //   html/template包略过模板调用计算上下文环境。
    //   此例中，当被"main"模板调用时，"attrs"模板的{{.URL}}必须视为一个URL；
    //   但如果解析"main"时，"attrs"还未被定义，就会导致本错误
    ErrNoSuchTemplate
    // 不能计算输出位置的上下文环境时，导致ErrOutputContext
    // 示例：
    //   {{define "t"}}{{if .T}}{{template "t" .T}}{{end}}{{.H}}",{{end}}
    // 讨论：
    //   一个递归的模板，其起始和结束的上下文环境不同时；
    //   不能计算出可信的输出位置上下文环境时，就可能导致本错误。
    //   检查各个命名模板是否有错误；
    //   如果模板不应在命名的起始上下文环境调用，检查在不期望上下文环境中对该模板的调用；
    //   或者将递归模板重构为非递归模板；
    ErrOutputContext
    // 尚未支持JS正则表达式插入字符集
    // 示例：
    //     <script>var pattern = /foo[{{.Chars}}]/</script&rt;
    // 讨论：
    //   html/template不支持向JS正则表达式里插入字面值字符集
    ErrPartialCharset
    // 部分转义序列尚未支持
    // 示例：
    //   <script>alert("\{{.X}}")</script&rt;
    // 讨论：
    //   html/template包不支持紧跟在反斜杠后面的action
    //   这一般是错误的，有更好的解决方法，例如：
    //     <script>alert("{{.X}}")</script&rt;
    //   可以工作，如果{{.X}}是部分转义序列，如"xA0"，
    //   可以将整个序列标记为安全文本：JSStr()
    ErrPartialEscape
    // range循环的重入口出错，导致ErrRangeLoopReentry
    // 示例：
    //   <script>var x = [{{range .}}'{{.}},{{end}}]</script&rt;
    // 讨论：
    //   如果range的迭代部分导致其结束于上一次循环的另一上下文，将不会有唯一的上下文环境
    //   此例中，缺少一个引号，因此无法确定{{.}}是存在于一个JS字符串里，还是一个JS值文本里。
    //   第二次迭代生成类似下面的输出：
    //     <script>var x = ['firstValue,'secondValue]</script&rt;
    ErrRangeLoopReentry
    // 斜杠可以开始一个除法或者正则表达式
    // 示例：
    //   <script&rt;
    //     {{if .C}}var x = 1{{end}}
    //     /-{{.N}}/i.test(x) ? doThis : doThat();
    //   </script&rt;
    // 讨论：`)
	}
}
