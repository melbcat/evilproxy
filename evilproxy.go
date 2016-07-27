/* EvilProxy*/

package main

import (
	"github.com/op/go-logging"
	/*"evilproxy/config"*/
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"runtime"
	"flag"
	"encoding/json"
	"encoding/binary"
)

var log = logging.MustGetLogger("evilproxy")

func main() {
	fmt.Println("Only a new GO EVILPROXY for cSploit")
}
