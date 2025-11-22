package cli

import (
    "bufio"
    "fmt"
    "os"
    "parkee/internal/service"
    "strconv"
    "strings"
)

type CLI struct {
    service *service.ParkingService
}

func NewCLI(s *service.ParkingService) *CLI {
    return &CLI{service: s}
}

func (c *CLI) RunFromFile(filename string) {
    file, _ := os.Open(filename)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        c.processCommand(line)
    }
}

func (c *CLI) processCommand(line string) {
    parts := strings.Split(line, " ")

    switch parts[0] {
    case "create_parking_lot":
        n, _ := strconv.Atoi(parts[1])
        c.service.CreateParkingLot(n)
        fmt.Println("Created parking lot with", n, "slots")

    case "park":
        carNumber := parts[1]
        lot, err := c.service.Park(carNumber)
        if err != nil {
            fmt.Println("Sorry, parking lot is full")
            return
        }
        fmt.Println("Allocated slot number:", lot.SlotNumber)

    case "leave":
        carNumber := parts[1]
        hours, _ := strconv.Atoi(parts[2])

        lot, car, charge, err := c.service.Leave(carNumber, hours)
        if err != nil {
            fmt.Println("Registration number", carNumber, "not found")
            return
        }

        fmt.Printf(
            "Registration number %s with Slot Number %d is free with Charge $%d\n",
            car.CarNumber,
            lot.SlotNumber,
            charge,
        )

    case "status":
        lots := c.service.Status()

        fmt.Println("Slot No.   Registration No.")
        for _, lot := range lots {
            if !lot.IsEmpty() {
                fmt.Printf("%d          %s\n",
                    lot.SlotNumber,
                    lot.Car.CarNumber,
                )
            }
        }
    }
}
