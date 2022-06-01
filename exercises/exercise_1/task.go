package exercise_1

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

func main() {
	ntpTime, err := ntp.Time("time.nist.gov")
	if err == nil {
		log.Fatalf("Error: %v", err)
	}
	ntpTimeFormatted := ntpTime.Format(time.UnixDate)

	fmt.Printf("NTP Time: %v\n", ntpTime)
	fmt.Printf("NTP Unix Date Time: %v\n", ntpTimeFormatted)
	fmt.Println()
	fmt.Printf("TIme: %v\n", time.Now())
	fmt.Printf("Unix Date Time: %v\n", time.Now().Format(time.UnixDate))
}
