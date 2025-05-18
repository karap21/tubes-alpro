package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "os"
    "time"
)

type Player struct {
    Name  string
    Age   int
    Score int
}

var players []Player 
const dataFile = "datatubes.json"

func loadPlayers() {
    file, err := os.Open(dataFile)
    if err != nil {
        if os.IsNotExist(err) {
            players = []Player{}
            return
        }
        fmt.Println("Gagal membuka file:", err)
        os.Exit(1)
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&players)
    if err != nil {
        fmt.Println("Gagal decode data pemain:", err)
        os.Exit(1)
    }
}

func savePlayers() {
    file, err := os.Create(dataFile)
    if err != nil {
        fmt.Println("Gagal menyimpan file:", err)
        return
    }
    defer file.Close()
    encoder := json.NewEncoder(file)
    err = encoder.Encode(players)
    if err != nil {
        fmt.Println("Gagal encode data pemain:", err)
    }
}

func bubbleSortByScore(players []Player) {
    n := len(players)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if players[j].Score < players[j+1].Score {
                players[j], players[j+1] = players[j+1], players[j]
            }
        }
    }
}

func isSorted(arr []int) bool {
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] > arr[i+1] {
            return false
        }
    }
    return true
}

func shuffle(arr []int) {
    for i := range arr {
        j := rand.Intn(len(arr))
        arr[i], arr[j] = arr[j], arr[i]
    }
}

func bogoSort(arr []int) float64 {
    start := time.Now()
    for !isSorted(arr) {
        shuffle(arr)
    }
    return time.Since(start).Seconds()
}

func addPlayer() {
    var name string
    var age int
    fmt.Print("Masukkan Nama Pemain: ")
    fmt.Scan(&name)
    fmt.Print("Masukkan Umur Pemain: ")
    fmt.Scan(&age)
    players = append(players, Player{Name: name, Age: age, Score: 0})
    savePlayers()
}

func removePlayer() {
    if len(players) == 0 {
        fmt.Println("Belum ada data pemain.")
        return
    }
    fmt.Println("Pilih nomor pemain yang ingin dihapus:")
    for i, p := range players {
        fmt.Printf("%d. %s - %d tahun (Skor: %d)\n", i+1, p.Name, p.Age, p.Score)
    }
    var idx int
    fmt.Print("Nomor pemain: ")
    fmt.Scan(&idx)
    if idx < 1 || idx > len(players) {
        fmt.Println("Nomor tidak valid.")
        return
    }
    players = append(players[:idx-1], players[idx:]...)
    fmt.Println("Pemain berhasil dihapus.")
    savePlayers()
}

func showPlayers() {
    if len(players) == 0 {
        fmt.Println("Belum ada data pemain.")
        return
    }
    bubbleSortByScore(players)
    fmt.Println("\nDaftar Pemain (Skor Tertinggi ke Terendah):")
    for _, p := range players {
        fmt.Printf("%s - %d tahun (Skor: %d)\n", p.Name, p.Age, p.Score)
    }
}

func playGame() {
    var numArray int
    fmt.Print("Masukkan jumlah angka yang ingin diurutkan dengan Bogo Sort: ")
    fmt.Scan(&numArray)

    var arr []int
    fmt.Println("Masukkan angka-angka:")
    for i := 0; i < numArray; i++ {
        var num int
        fmt.Scan(&num)
        arr = append(arr, num)
    }

    var guesses []float64
    fmt.Println("\nSetiap pemain harus menebak waktu sorting:")
    for _, p := range players {
        var guess float64
        fmt.Printf("%s, tebak berapa detik Bogo Sort akan selesai: ", p.Name)
        fmt.Scan(&guess)
        guesses = append(guesses, guess)
    }

    elapsed := bogoSort(arr)

    fmt.Printf("\nBogo Sort selesai dalam %.3f detik.\n", elapsed)

    winners := []int{}
    for i, guess := range guesses {
        if guess >= elapsed-0.5 && guess <= elapsed+0.5 {
            winners = append(winners, i)
        }
    }

    if len(winners) > 0 {
        fmt.Print("Pemenang: ")
        for _, idx := range winners {
            fmt.Print(players[idx].Name, " ")
            players[idx].Score++ 
        }
        fmt.Println()
        savePlayers()
    } else {
        fmt.Println("Tidak ada pemenang kali ini! Coba lagi.")
    }

    fmt.Println("Array setelah diurutkan:", arr)
}

func main() {
    rand.Seed(time.Now().UnixNano())
    loadPlayers()
    for {
        fmt.Println("\n--- MENU UTAMA ---")
        fmt.Println("1. Masukkan Data Pemain")
        fmt.Println("2. Lihat Semua Data Pemain")
        fmt.Println("3. Hapus Pemain")
        fmt.Println("4. Mulai Permainan Tebakan Bogo Sort")
        fmt.Println("5. Keluar Program")
        fmt.Print("Pilih opsi: ")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            addPlayer()
        case 2:
            showPlayers()
        case 3:
            removePlayer()
        case 4:
            if len(players) == 0 {
                fmt.Println("Tambahkan pemain terlebih dahulu sebelum bermain!")
            } else {
                playGame()
            }
        case 5:
            savePlayers()
            fmt.Println("Terima kasih telah bermain! ")
            return
        default:
            fmt.Println("Opsi tidak valid. Coba lagi.")
        }
    }
}