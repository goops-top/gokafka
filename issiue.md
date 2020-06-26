## gokafka相关梳理

`注意:` 由于内部kafka集群为多个版本，1.0.1和2.5.0版本，因此在使用过程中需要注意选择合适的api版本进行覆盖。

- [X] list 子命令对于低版本的kafka不兼容 `使用V1_0_0_0`版本的admin api可以解决该问题

```
$ ./build/gokafka.mac --cluster dev-kafka list  topic
panic: EOF

goroutine 1 [running]:
gokafka/controller.ListTopic(0xc000066a80, 0x3, 0x4, 0xc0001d7cc8, 0x0, 0x0)
	/Users/xuebiaoxu/Desktop/gokafka/controller/list.go:27 +0x6a1
gokafka/cmd.glob..func5(0x1b496e0, 0xc00000eda0, 0x0, 0x2)
	/Users/xuebiaoxu/Desktop/gokafka/cmd/list.go:53 +0x1f9
github.com/spf13/cobra.(*Command).execute(0x1b496e0, 0xc00000ed80, 0x2, 0x2, 0x1b496e0, 0xc00000ed80)
	/Users/xuebiaoxu/go/pkg/mod/github.com/spf13/cobra@v1.0.0/command.go:846 +0x29d
github.com/spf13/cobra.(*Command).ExecuteC(0x1b49980, 0x0, 0x0, 0x10078cf)
	/Users/xuebiaoxu/go/pkg/mod/github.com/spf13/cobra@v1.0.0/command.go:950 +0x349
github.com/spf13/cobra.(*Command).Execute(...)
	/Users/xuebiaoxu/go/pkg/mod/github.com/spf13/cobra@v1.0.0/command.go:887
gokafka/cmd.Execute(...)
	/Users/xuebiaoxu/Desktop/gokafka/cmd/root.go:30
main.main()
	/Users/xuebiaoxu/Desktop/gokafka/main.go:28 +0x98

```
