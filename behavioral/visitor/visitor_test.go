package visitor

import "testing"

func Test1(t *testing.T) {

	/* 本测试代码只做使用演示：
	JsonFile和YamlFile是两个基本的类。原始的需求是，对json和yaml文件读取后，需要进行压缩和加密。本来这两个功能（压缩和加密）是可以
	分别放在JsonFile和YamlFile两个类中自己实现的。但是为了后续的功能扩展，比如对文件中的某些字段过滤、合并等。新增功能的时候，需要对JsonFile
	和YamlFile进行改动，并且调用这些新的功能时候，需要额外的工作。使用访问者模式，就可以解决功能上扩展的问题，只要原始类JsonFile中实现accept
	方法即可。把功能的实现放在Visitor中实现（准确说是，实现了Visitor接口的类中实现）。因此，后面新增功能，只需要新增一个实现了Visitor接口的
	类即可。
	*/

	// 模拟JsonFile和YamlFile已经有数据读取到了
	files := []FileInterface{&JsonFile{msg: "Json File Content"}, &YamlFile{msg: "Yaml File Content"}}

	// 实现压缩的功能
	compressor := &Compressor{}
	for _, f := range files {
		f.accept(compressor)
	}

	// 实现加密的功能
	encryptor := &Encryptor{}
	for _, f := range files {
		f.accept(encryptor)
	}

}
