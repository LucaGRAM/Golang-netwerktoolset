package main

import (
    "fmt"
    "os/exec"
    "net"
    "time"
    "strings"
    "math/rand"
    "bufio"
    "log"
)

func main() {
    fmt.Printf("Welkom bij Luca's netwerk toolset! \n")
    fmt.Printf("Deze tool set bevat verschillende handige tools voor jou netwerk! \n")
    fmt.Printf("Kies \n")
    fmt.Printf("1: Voor een ping test. \n")
    fmt.Printf("2: Voor een port scan. \n")
    fmt.Printf("3: Voor een wachwoord generator kies \n")
    fmt.Printf("4: Voor een trace route, kies. \n")
    fmt.Printf("5: Voor een internet connectiviteit test. \n")
    fmt.Printf("KIES 0 OM DE APPLICATIE AF TE SLUITEN! \n")
    fmt.Printf("uw keuze: ")
    KeuzeMenu()
}

func KeuzeMenu() {
    var keuze string
    fmt.Scan(&keuze)
    if keuze == ("1") { //Bij keuze 1 wordt PingApplicatie function geactiveert
     PingApp()
   }else if keuze == ("2") { //Bij keuze 2 wordt de Portscan geactiveert
//        fmt.Printf("WAARSCHUWING! Dit kan lang duren, druk op ctrl + C om het proces te stoppen!")
        PortScan()
    }else if keuze == ("3") { //Bij keuze 3 wordt eerst het wachtwoord menu geactiveert
      WachtwoordMenu()
      Terug()
    }else if keuze == ("4") { // bij Keuze 4 wordt de traceroute function geactiveert
      var traceroute string
      fmt.Print("naar welk adres wilt u een trace route? \n")
      fmt.Scan(&traceroute)
      RunTraceroute(traceroute)
    }else if keuze == ("5") { //Bij keuze 5 wordt de Internet test function geactiveert
        InternetTest()
    }else if keuze == ("0") {
       fmt.Printf("De applicatie wordt afgesloten! \n")
      }else{
         fmt.Printf("Sorry dit wordt niet herkent! \n")
         Terug()
    }

}

func PingApp() {
     var keuzeIP string
     var keuzeIP1 string
     fmt.Print("   \n")
     fmt.Print("================================================================================ \n")
     fmt.Print("   \n")
     fmt.Print("welkom bij de ping applicatie! \n")
     fmt.Print("U kunt kiezen uit verschillende opties: \n")
     fmt.Print("Voor een ping test naar een zelf uitgekozen IP adres, kies: 1 \n")
     fmt.Print("Voor het snel pingen naar 1 van de hosts in het netwerk, kies: 2 \n")
     fmt.Print("Voor het pingen naar elke host in het netwerk, kies: 3 \n")
     fmt.Print("Uw keuze: ")
     fmt.Scan(&keuzeIP)
     if keuzeIP == ("1") {
        fmt.Printf("Welk IP-adres wilt u pingen? \n")
        fmt.Print("Uw keuze: ")
        fmt.Scan(&keuzeIP1)
        Command := fmt.Sprintf("ping -c 1 %v > /dev/null && echo Het pingen naar het gekozen adres is gelukt. || echo 1 Het pingen naar het gekozen a>
        output, err := exec.Command("/bin/sh", "-c", Command).Output()
        fmt.Print(string(output))
        fmt.Print(err)
        Terug()
     }
     if keuzeIP == ("2") {
          fmt.Print("Naar welke host wilt u pingen? \n")
          fmt.Print("Voor Router 1, kies: 1 \n")
     }
}

func PortScan() {
 var portrange int
 fmt.Print("kies t/m welke port u wilt scannen:  \n")
 fmt.Scan(&portrange)
 fmt.Printf("WAARSCHUWING! Dit kan lang duren, druk op ctrl + C om het proces te stoppen! \n")
 for i := 1; i <= portrange; i++ {
             address := fmt.Sprintf("scanme.nmap.org:%d", i)

             conn, err := net.Dial("tcp", address)
             if err != nil {
                 continue
             }
             conn.Close()
             fmt.Printf("Open poort: %d\n", i)
         }
 Terug()
}


func WachtwoordMenu() {
     rand.Seed(time.Now().Unix())
      var passwordLength int
      var minUpperCase int
      var minNum int
      var minSpecialChar int
      fmt.Print("welkom bij de wachtwoord generator\n")
      fmt.Print("gemaakt door Luca Gram\n")
      fmt.Print("hoelang moet het wachtwoord minimaal zijn?\n")
      fmt.Scan(&passwordLength)
      fmt.Print("hoeveel speciale karakters moeten er in het wachtwoord?\n")
      fmt.Scan(&minSpecialChar)
      fmt.Print("Hoeveel nummers moeten er in het wachtwoord?\n")
      fmt.Scan(&minNum)
      fmt.Print("hoeveel hoofdletters moeten er in het wachtwoord?\n")
      fmt.Scan(&minUpperCase)
      password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
      fmt.Println(password)
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
    var password strings.Builder
     var lowerCharSet   = "abcdedfghijklmnopqrst"
     var upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
     var specialCharSet = "!@#$%&*"
     var numberSet      = "0123456789"
     var allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet


    for i := 0; i < minSpecialChar; i++ {
        random := rand.Intn(len(specialCharSet))
        password.WriteString(string(specialCharSet[random]))
    }

    for i := 0; i < minNum; i++ {
        random := rand.Intn(len(numberSet))
        password.WriteString(string(numberSet[random]))
    }

    for i := 0; i < minUpperCase; i++ {
        random := rand.Intn(len(upperCharSet))
        password.WriteString(string(upperCharSet[random]))
    }

    remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
    for i := 0; i < remainingLength; i++ {
        random := rand.Intn(len(allCharSet))
        password.WriteString(string(allCharSet[random]))
    }
    inRune := []rune(password.String())
    rand.Shuffle(len(inRune), func(i, j int) {
        inRune[i], inRune[j] = inRune[j], inRune[i]
    })
    return string(inRune)
}


func RunTraceroute(host string) {
    errch := make(chan error, 1)
    cmd := exec.Command("traceroute", host)

    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatal(err)
    }

    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }

    go func() {
        errch <- cmd.Wait()
    }()

    go func() {
        for _, char := range "|/-\\" {
            fmt.Printf("\r%s...%c", "Running traceroute", char)
            time.Sleep(100 * time.Millisecond)
        }
        scanner := bufio.NewScanner(stdout)
        fmt.Println("")
        for scanner.Scan() {
            line := scanner.Text()
            log.Println(line)
        }
    }()

    select {
    case <-time.After(time.Second * 15):
        log.Println("Timeout hit..")
        return
    case err := <-errch:
        if err != nil {
            log.Println("traceroute failed:", err)
        }
    }
}


func InternetTest() {
      fmt.Print("Internet connectiviteit test \n")
      fmt.Print("Er wordt eerst een ping test gedaan naar Router 1.\n")
      fmt.Print("daarna worden er verschillende pingtests gedaan naar servers op het internet! \n")
      Command := fmt.Sprintf("ping -c 1 172.16.0.1 > /dev/null && echo Het pingen naar Router 1 is gelukt! || echo Het pingen naar Router 1 is misluk>
        output, err := exec.Command("/bin/sh", "-c", Command).Output()
        fmt.Print(string(output))
        fmt.Print(err)
      Command2 := fmt.Sprintf("ping -c 1 8.8.8.8 > /dev/null && echo Google is bereikbaar! || echo Google was niet te bereiken")
        output2, err := exec.Command("/bin/sh", "-c", Command2).Output()
        fmt.Print(string(output2))
        fmt.Print(err)
        Terug()
}

func Terug() {
      var terug string
      fmt.Print("Wilt u terug naar het hoofd menu? \n")
      fmt.Print("Kies ja of nee:  ")
      fmt.Scan(&terug)
      if terug == ("ja") {
        main()
      }else if terug == ("nee") {
         fmt.Print("De applicatie wordt afgesloten! \n")
      }else{
         fmt.Print("Dit wordt niet herkent! \n")
         Terug()
      }

}
