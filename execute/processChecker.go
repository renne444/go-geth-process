package execute

import (
	"fmt"
	"os/exec"
	"time"
	"util/Network"
	"util/localFileIO"
)

/*
 * this struct check the process geth by checking if the RPC port and GETH port has been occupied
 */
type ProcessChecker struct{
	config *Config
}

func NewProcessChecker(config *Config) *ProcessChecker {
	if config != nil {
		return &ProcessChecker{config:config}
	}else{
		return nil
	}
}

func (pc *ProcessChecker)composeInitCmd() string{
	cmd := fmt.Sprintf("%s --datadir %s init %s",
					pc.config.GethPath,
					pc.config.DataDir,
					pc.config.GenesisFile)

	return cmd
}

func (pc *ProcessChecker)InitProcess() bool{
	if !localFileIO.IsNonEmptyDir(pc.config.DataDir) {
		//util.CreateFolder(pc.config.DataDir)
		cmd := exec.Command("/bin/sh","-c",pc.composeInitCmd())
		fmt.Printf("[INFO]init cmd `%s` will be exec\n",pc.composeInitCmd())

		err := cmd.Run()
		if err != nil{
			fmt.Println("[ERROR]error occured while initializing")
			return false
		}
		fmt.Println("[INFO]Initialize finished")
	}
	return true
}

//check geth process by checking the occupation of the port and the RPC port
func (pc *ProcessChecker) checkProcessExist() bool{
	return Network.PortInUse(pc.config.Port) && Network.PortInUse(pc.config.RpcPort)
}

func (pc *ProcessChecker) composeRunCmd() string{
	cmd := fmt.Sprintf(`%s --datadir %s --networkid %d  --etherbase "%s" --port %d`,
								pc.config.GethPath,
								pc.config.DataDir,
								pc.config.NetworkId,
								pc.config.CoinBase,
								pc.config.Port,
								)
	if pc.config.Rpc {
		cmd += fmt.Sprintf(` --rpc --rpcaddr "%s" --rpcport "%d" --rpcapi "db,eth,net,web3,personal,admin"`,
														pc.config.RpcAddr,
														pc.config.RpcPort,)
	}

	if pc.config.Mine {
		cmd += fmt.Sprintf(` --mine --minerthreads %d`,pc.config.MineThread)
	}
	//cmd += fmt.Sprintf(` >> %s`,pc.config.GethLogPath)
	return cmd
}

// cmd | tee -a $LOG
func (pc *ProcessChecker) RunProcess() bool{
	if pc.checkProcessExist() == false{
		fmt.Printf("[INFO]run cmd `%s` will be exec\n",pc.composeRunCmd())
		runCmd := exec.Command("/bin/sh","-c",pc.composeRunCmd())
		outputCmd := exec.Command("/usr/bin/tee","-a",pc.config.GethLogPath)

//		runCmd.Stdout, _ = outputCmd.StdinPipe()
		runCmd.Stderr, _ = outputCmd.StdinPipe() //tmd居然在错误输出..

		err := outputCmd.Start()
		if err!= nil{
			fmt.Println(`[ERROR]error occured while starting output process`)
			return false
		}

		err = runCmd.Run()
		if err!= nil{
			fmt.Println(`[ERROR]error occured while running process`)
			return false
		}

//		err = outputCmd.Wait()
//		if err!= nil{
//			fmt.Println(`[ERROR]error occured while waiting output process`)
//			return false
//		}
	}
	return true
}

func (pc *ProcessChecker)Execute(){
	for{
		pc.InitProcess()
		pc.RunProcess()
		fmt.Println("Process has been broken , will be restart at 1 minute")
		time.Sleep(time.Minute)
	}
}