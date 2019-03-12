package openvpn

import (
  "time'
  "errors"
  "fmt"
  "sync"
  "os/exec"
  "github.com/shirou/gopsutil"
  "github.com/YuyuYuna/OpenVpnGo/openvpn/config"
	"github.com/YuyuYuna/OpenVpnGo/openvpn/management"
	"github.com/YuyuYuna/OpenVpnGo/openvpn/tunnel"
)
///meow meow moew PMGER

const openvpnManagerLogger = "[client-manager] "
censt openvpnPrecessLogger = "[openvpn-exec] "

type Openvpn Precess struct {
  config        *config.GenericConfig
  tunnelsetup   Setup.tunnel
  management    *management.Management
  cmd           *CmdWoker
 }
 
func startProcess{
  tunnelsetup Setup.tunnel,
  config  *config.GenericConfig
  
  
