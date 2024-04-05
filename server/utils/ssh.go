package utils

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
)

func Exec(str string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", str)
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	if len(errBuffer.String()) != 0 {
		log.Println(errBuffer.String())
		return "", errors.New(errBuffer.String())
	}
	return outBuffer.String(), err
}

//func SSH(cmd string) []byte {
//	// SSH连接配置
//	config := &ssh.ClientConfig{
//		User: "root",
//		Auth: []ssh.AuthMethod{
//			ssh.Password("Pam#6bre"),
//		},
//		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//	}
//
//	// 连接SSH服务器
//	client, err := ssh.Dial("tcp", "43.128.79.149:22", config)
//	if err != nil {
//		global.GVA_LOG.Error("无法连接到主机!", zap.Error(err))
//		return []byte{}
//	}
//
//	// 创建一个新的会话
//	session, err := client.NewSession()
//	if err != nil {
//		global.GVA_LOG.Error("无法开启会话!", zap.Error(err))
//		return []byte{}
//	}
//	defer session.Close()
//
//	// 执行多个命令
//	output, err := session.CombinedOutput(cmd)
//	if err != nil {
//		global.GVA_LOG.Error("执行失败!", zap.Error(err))
//		return []byte{}
//	}
//
//	return output
//}
//
//func EXEC(host, user, pwd, cmd string) ([]byte, error) {
//	// SSH连接配置
//	config := &ssh.ClientConfig{
//		User: user,
//		Auth: []ssh.AuthMethod{
//			ssh.Password(pwd),
//		},
//		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//	}
//
//	// 连接SSH服务器
//	client, err := ssh.Dial("tcp", host+":22", config)
//	if err != nil {
//		global.GVA_LOG.Error("无法连接到主机!", zap.Error(err))
//		return []byte{}, err
//	}
//
//	// 创建一个新的会话
//	session, err := client.NewSession()
//	if err != nil {
//		global.GVA_LOG.Error("无法开启会话!", zap.Error(err))
//		return []byte{}, err
//	}
//	defer session.Close()
//
//	fmt.Println(cmd)
//	output, err := session.Output(cmd)
//	if err != nil {
//		global.GVA_LOG.Error("执行失败!", zap.Error(err))
//		global.GVA_LOG.Error(string(output))
//		return []byte{}, err
//	}
//
//	return output, nil
//}
