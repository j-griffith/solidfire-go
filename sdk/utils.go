package sdk

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func GetInitiatorIqns() ([]string, error) {
	log.Println("Begin utils.GetInitiatorIqns")
	var iqns []string
	out, err := exec.Command("sudo", "cat", "/etc/iscsi/initiatorname.iscsi").CombinedOutput()
	if err != nil {
		log.Printf("Error encountered gathering initiator names: %v\n", err)
		return nil, err
	}
	lines := strings.Split(string(out), "\n")
	for _, l := range lines {
		if strings.Contains(l, "InitiatorName=") {
			iqns = append(iqns, strings.Split(l, "=")[1])
		}
	}
	return iqns, nil
}

func WaitForPathToExist(fileName string, numTries int) bool {
	log.Printf("Being utils.waitForPathToExist: %v\n", fileName)
	for i := 0; i < numTries; i++ {
		_, err := os.Stat(fileName)
		if err == nil {
			log.Printf("path found: %v\n", fileName)
			return true
		}
		if err != nil && !os.IsNotExist(err) {
			return false
		}
		time.Sleep(time.Second)
	}
	return false
}

func GetDeviceFileFromIscsiPath(iscsiPath string) (string, error) {
	log.Printf("Being utils.getDeviceFileFromIscsiPath: %v\n", iscsiPath)
	out, err := exec.Command("sudo", "ls", "-la", iscsiPath).CombinedOutput()
	if err != nil {
		return "", err
	}
	d := strings.Split(string(out), "../../")
	log.Printf("Found d: %v\n", d)
	devFile := "/dev/" + d[1]
	log.Printf("using base of: %v\n", devFile)
	devFile = strings.TrimSpace(devFile)
	return devFile, nil
}

func iscsiSupported() bool {
	_, err := exec.Command("iscsiadm", "-h").CombinedOutput()
	if err != nil {
		log.Printf("iscsiadm tools not found on this host")
		return false
	}
	return true
}

func iscsiDiscovery(portal string) (targets []string, err error) {
	log.Printf("Begin utils.iscsiDiscovery (portal: %s)", portal)
	out, err := exec.Command("sudo", "iscsiadm", "-m", "discovery", "-t", "sendtargets", "-p", portal).CombinedOutput()
	if err != nil {
		log.Printf("Error encountered in sendtargets cmd: %v\n", out)
		return
	}
	targets = strings.Split(string(out), "\n")
	return

}

func iscsiLogin(targetIP, targetIQN string) (err error) {
	log.Printf("Begin utils.iscsiLogin: %s/%s", targetIP, targetIQN)
	_, err = exec.Command("sudo", "iscsiadm", "-m", "node", "-p", targetIP, "-T", targetIQN, "--login").CombinedOutput()
	if err != nil {
		log.Printf("Received error on login attempt: %v", err)
	}
	return err
}

func iscsiDisableDelete(targetIP, targetIQN string) (err error) {
	log.Printf("Begin utils.iscsiDisableDelete: %s/%s\n", targetIP, targetIQN)
	_, err = exec.Command("sudo", "iscsiadm", "-m", "node", "-T", targetIQN, "--portal", targetIP, "-u").CombinedOutput()
	if err != nil {
		log.Printf("Error during iscsi logout: %v\n", err)
		//return
	}
	_, err = exec.Command("sudo", "iscsiadm", "-m", "node", "-o", "delete", "-T", targetIQN).CombinedOutput()
	return
}

func iscsiadmCmd(args []string) error {
	log.Printf("Being utils.iscsiadmCmd: iscsiadm %+v", args)
	resp, err := exec.Command("iscsiadm", args...).CombinedOutput()
	if err != nil {
		log.Printf("Error encountered running iscsiadm, args: %v, response: %v\n", args, resp)
		log.Printf("Error message: %v\n", err)
	}
	return err
}

func LoginWithChap(tiqn, portal, username, password, iface string) error {
	log.Printf("Begin utils.LoginWithChap: iqn: %s, portal: %s, username: %s, password=xxxx, iface: %s", tiqn, portal, username, iface)
	args := []string{"-m", "node", "-T", tiqn, "-p", portal + ":3260"}
	createArgs := append(args, []string{"--interface", iface, "--op", "new"}...)

	if _, err := exec.Command("iscsiadm", createArgs...).CombinedOutput(); err != nil {
		log.Printf("Error running iscsiadm node create: %v\n", err)
		return err
	}

	authMethodArgs := append(args, []string{"--op=update", "--name", "node.session.auth.authmethod", "--value=CHAP"}...)
	if out, err := exec.Command("iscsiadm", authMethodArgs...).CombinedOutput(); err != nil {
		log.Printf("Error running iscsiadm set authmethod: %v\n", err)
		log.Printf("iscsiadm set authmethod output: %v\n", out)
		return err
	}

	authUserArgs := append(args, []string{"--op=update", "--name", "node.session.auth.username", "--value=" + username}...)
	if _, err := exec.Command("iscsiadm", authUserArgs...).CombinedOutput(); err != nil {
		log.Printf("Error running iscsiadm set authuser: %v\n", err)
		return err
	}
	authPasswordArgs := append(args, []string{"--op=update", "--name", "node.session.auth.password", "--value=" + password}...)
	if _, err := exec.Command("iscsiadm", authPasswordArgs...).CombinedOutput(); err != nil {
		log.Printf("Error running iscsiadm set authpassword: %v\n", err)
		return err
	}
	loginArgs := append(args, []string{"--login"}...)
	if _, err := exec.Command("iscsiadm", loginArgs...).CombinedOutput(); err != nil {
		log.Printf("Error running iscsiadm login: %v\n", err)
		return err
	}
	return nil
}
