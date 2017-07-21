package signalMgr

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Jordanzuo/goutil/debugUtil"
	"github.com/Jordanzuo/goutil/logUtil"
)

//Start 启动信号管理器
// reloadFunc:收到重启信号时调用的方法
// exitFunc:收到退出信号时调用的方法
func Start(reloadFunc func() []error, exitFunc func() error) {
	go func() {
		// 处理内部未处理的异常，以免导致主线程退出，从而导致系统崩溃
		defer func() {
			if r := recover(); r != nil {
				logUtil.LogUnknownError(r)
			}
		}()

		sigs := make(chan os.Signal)
		signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

		for {
			// 准备接收信息
			sig := <-sigs

			// 输出信号
			debugUtil.Println("sig:", sig)

			if sig == syscall.SIGHUP {
				logUtil.Log("收到重启的信号，准备重新加载配置", logUtil.Info, true)

				// 重新加载
				if reloadFunc != nil {
					errList := reloadFunc()
					for _, err := range errList {
						logUtil.Log(fmt.Sprintf("重启失败，错误信息为:%s", err), logUtil.Error, true)
					}
				}

				logUtil.Log("收到重启的信号，重新加载配置完成", logUtil.Info, true)
			} else {
				logUtil.Log("收到退出程序的信号，开始退出……", logUtil.Info, true)

				// 做一些收尾的工作
				if exitFunc != nil {
					if err := exitFunc(); err != nil {
						logUtil.Log(fmt.Sprintf("执行exitFunc失败，错误信息为:%s", err), logUtil.Error, true)
					}
				}

				logUtil.Log("收到退出程序的信号，退出完成……", logUtil.Info, true)

				// 一旦收到信号，则表明管理员希望退出程序，则先保存信息，然后退出
				os.Exit(0)
			}
		}
	}()
}
