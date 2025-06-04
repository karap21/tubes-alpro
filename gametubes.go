package main

import (
    "fmt"
    "strings"
    "math/rand"
    "time"
    "encoding/json"
    "os"
)

const MAX = 100
const MAX_HEWAN = 10
const namaFile = "datatubes.json"

type Skor struct {
    TebakAngka   int
    TebakHewan   int
    Perkalian    int
}

type Player struct {
    ID     string
    Nama   string
    Umur   int
    Level  int
    Skor   Skor
    Poin   int
    Item   int
    Trophy int
}

type Hewan struct {
    Nama string
    Ciri string
}

var daftarPlayer [MAX]Player
var nPlayer int

var daftarHewan [MAX_HEWAN]Hewan
var nHewan int

func toLower(s string) string {
    return strings.ToLower(s)
}

func menuBelanja(idx int) {
    for {
        fmt.Println("\n=== MENU BELANJA ===")
        fmt.Println("1. Beli Item (10 poin)")
        fmt.Println("2. Beli Trophy (30 poin)")
        fmt.Println("3. Kembali")
        var pilih int
        fmt.Print("Pilih menu: ")
        fmt.Scan(&pilih)
        if pilih == 1 {
            if daftarPlayer[idx].Poin >= 10 {
                daftarPlayer[idx].Poin -= 10
                daftarPlayer[idx].Item++
                fmt.Println("Berhasil beli 1 Item!")
            } else {
                fmt.Println("Poin tidak cukup!")
            }
        } else if pilih == 2 {
            if daftarPlayer[idx].Poin >= 30 {
                daftarPlayer[idx].Poin -= 30
                daftarPlayer[idx].Trophy++
                fmt.Println("Berhasil beli 1 Trophy!")
            } else {
                fmt.Println("Poin tidak cukup!")
            }
        } else if pilih == 3 {
            break
        }
    }
}

func tambahPlayer(arr *[MAX]Player, n *int, p Player) {
    if *n < MAX {
        arr[*n] = p
        *n++
    } else {
        fmt.Println("Data player penuh!")
    }
}

func binarySearchID(arr [MAX]Player, n int, id string) int {
    left, right := 0, n-1
    for left <= right {
        mid := (left + right) / 2
        if arr[mid].ID == id {
            return mid
        } else if arr[mid].ID < id {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}


func hapusPlayer(arr *[MAX]Player, n *int, idx int) {
    for i := idx; i < *n-1; i++ {
        arr[i] = arr[i+1]
    }
    *n--
}

func selectionSortUmur(arr *[MAX]Player, n int, ascending bool) {
    for i := 0; i < n-1; i++ {
        idx := i
        for j := i+1; j < n; j++ {
            if ascending {
                if arr[j].Umur < arr[idx].Umur {
                    idx = j
                }
            } else {
                if arr[j].Umur > arr[idx].Umur {
                    idx = j
                }
            }
        }
        arr[i], arr[idx] = arr[idx], arr[i]
    }
}

func insertionSortNama(arr *[MAX]Player, n int, ascending bool) {
    for i := 1; i < n; i++ {
        temp := arr[i]
        j := i - 1
        if ascending {
            for j >= 0 && arr[j].Nama > temp.Nama {
                arr[j+1] = arr[j]
                j--
            }
        } else {
            for j >= 0 && arr[j].Nama < temp.Nama {
                arr[j+1] = arr[j]
                j--
            }
        }
        arr[j+1] = temp
    }
}

func tampilInfoPlayer(arr [MAX]Player, n int) {
    fmt.Println("Daftar Player:")
    for i := 0; i < n; i++ {
        fmt.Printf("%d. ID: %s | Nama: %s | Umur: %d | Level: %d | Skor Angka: %d | Skor Hewan: %d | Skor Perkalian: %d | Poin: %d | Item: %d | Trophy: %d\n",
            i+1, arr[i].ID, arr[i].Nama, arr[i].Umur, arr[i].Level, arr[i].Skor.TebakAngka, arr[i].Skor.TebakHewan, arr[i].Skor.Perkalian, arr[i].Poin, arr[i].Item, arr[i].Trophy)
    }
}

func gameTebakAngka(idx int) {
    rand.Seed(time.Now().UnixNano())
    angka := rand.Intn(10) + 1
    var tebakan int
    fmt.Print("Tebak angka 1-10: ")
    fmt.Scan(&tebakan)
    if tebakan == angka {
        fmt.Println("Benar!")
        daftarPlayer[idx].Skor.TebakAngka++
        daftarPlayer[idx].Poin += 5
    } else {
        fmt.Printf("Salah! Angka: %d\n", angka)
    }
}

func gameTebakHewan(idx int) {
    if nHewan == 0 {
        fmt.Println("Belum ada data hewan.")
        return
    }
    idxH := rand.Intn(nHewan)
    fmt.Println("Ciri-ciri hewan:", daftarHewan[idxH].Ciri)
    var jawaban string
    fmt.Print("Jawaban: ")
    fmt.Scan(&jawaban)
    if toLower(jawaban) == toLower(daftarHewan[idxH].Nama) {
        fmt.Println("Benar!")
        daftarPlayer[idx].Skor.TebakHewan++
        daftarPlayer[idx].Poin += 5
    } else {
        fmt.Printf("Salah! Jawaban: %s\n", daftarHewan[idxH].Nama)
    }
}

func gamePerkalian(idx int) {
    rand.Seed(time.Now().UnixNano())
    skor := 0
    for i := 1; i <= 3; i++ {
        a := rand.Intn(10) + 1
        b := rand.Intn(10) + 1
        var jawab int
        fmt.Printf("Soal %d: %d x %d = ", i, a, b)
        fmt.Scan(&jawab)
        if jawab == a*b {
            fmt.Println("Benar!")
            skor++
            daftarPlayer[idx].Poin += 3
        } else {
            fmt.Printf("Salah! Jawaban: %d\n", a*b)
        }
    }
    daftarPlayer[idx].Skor.Perkalian += skor
}

func tambahHewan(arr *[MAX_HEWAN]Hewan, n *int, nama, ciri string) {
    if *n < MAX_HEWAN {
        arr[*n] = Hewan{Nama: nama, Ciri: ciri}
        *n++
    }
}

func skorTotal(p Player) int {
    return p.Skor.TebakAngka + p.Skor.TebakHewan + p.Skor.Perkalian
}

func selectionSortSkor(arr *[MAX]Player, n int) {
    for i := 0; i < n-1; i++ {
        idx := i
        for j := i+1; j < n; j++ {
            if skorTotal(arr[j]) > skorTotal(arr[idx]) {
                idx = j
            }
        }
        arr[i], arr[idx] = arr[idx], arr[i]
    }
}

func insertionSortID(arr *[MAX]Player, n int) {
    for i := 1; i < n; i++ {
        temp := arr[i]
        j := i - 1
        for j >= 0 && arr[j].ID > temp.ID {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = temp
    }
}

func tampilRangking(arr [MAX]Player, n int) {
    fmt.Println("\n=== RANGKING PLAYER BERDASARKAN SKOR TOTAL ===")
    for i := 0; i < n; i++ {
        fmt.Printf("%d. %s (ID: %s) | Skor Total: %d | Angka: %d | Hewan: %d | Perkalian: %d\n",
            i+1, arr[i].Nama, arr[i].ID, skorTotal(arr[i]), arr[i].Skor.TebakAngka, arr[i].Skor.TebakHewan, arr[i].Skor.Perkalian)
    }
}

func cariPlayerByID(arr *[MAX]Player, n int) {
    insertionSortID(arr, n)
    var id string
    fmt.Print("Masukkan ID player yang dicari: ")
    fmt.Scan(&id)
    idx := binarySearchID(*arr, n, id)
    if idx != -1 {
        fmt.Printf("Ditemukan: %s | Umur: %d | Level: %d | Skor Angka: %d | Skor Hewan: %d | Skor Perkalian: %d | Poin: %d | Item: %d | Trophy: %d\n",
            arr[idx].Nama, arr[idx].Umur, arr[idx].Level, arr[idx].Skor.TebakAngka, arr[idx].Skor.TebakHewan, arr[idx].Skor.Perkalian, arr[idx].Poin, arr[idx].Item, arr[idx].Trophy)
    } else {
        fmt.Println("Player tidak ditemukan.")
    }
}

func simpanPlayer(arr [MAX]Player, n int) {
    temp := make([]Player, n)
    for i := 0; i < n; i++ {
        temp[i] = arr[i]
    }
    file, err := os.Create(namaFile)
    if err != nil {
        fmt.Println("Gagal menyimpan file:", err)
        return
    }
    defer file.Close()
    encoder := json.NewEncoder(file)
    err = encoder.Encode(temp)
    if err != nil {
        fmt.Println("Gagal encode data pemain:", err)
    }
}

func muatPlayer(arr *[MAX]Player, n *int) {
    file, err := os.Open(namaFile)
    if err != nil {
        if os.IsNotExist(err) {
            *n = 0
            return
        }
        fmt.Println("Gagal membuka file:", err)
        os.Exit(1)
    }
    defer file.Close()
    var temp []Player
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&temp)
    if err != nil {
        if err.Error() == "EOF" {
            *n = 0
            return
        }
        fmt.Println("Gagal decode data pemain:", err)
        os.Exit(1)
    }
    *n = 0
    for i := 0; i < len(temp) && i < MAX; i++ {
        arr[i] = temp[i]
        *n++
    }
}

func main() {
   muatPlayer(&daftarPlayer, &nPlayer)
    rand.Seed(time.Now().UnixNano()) 
    tambahHewan(&daftarHewan, &nHewan, "Ayam", "Berkaki dua, unggas, bersayap, bertelur, sering dipelihara")
    tambahHewan(&daftarHewan, &nHewan, "Kucing", "Berkaki empat, berbulu, sering dipelihara, suka mengeong")
    tambahHewan(&daftarHewan, &nHewan, "Bebek", "Berkaki dua, unggas, bersayap, suka berenang, bertelur")
    tambahHewan(&daftarHewan, &nHewan, "Gajah", "Berkaki empat, bertubuh besar, belalai panjang, bertelinga lebar")
    tambahHewan(&daftarHewan, &nHewan, "Singa", "Berkaki empat, raja hutan, berbulu coklat, jantan berjanggut tebal")
    tambahHewan(&daftarHewan, &nHewan, "Lumba-lumba", "Hidup di laut, mamalia, cerdas, suka melompat di air")
    tambahHewan(&daftarHewan, &nHewan, "Nyamuk", "Bertubuh kecil, bersayap, suka menghisap darah, sering bersuara mendengung")
    tambahHewan(&daftarHewan, &nHewan, "Ular", "Bertubuh panjang, tidak berkaki, melata, ada yang berbisa")
 
    var menu int
    for {
        fmt.Println("\n=== MENU UTAMA ===")
        fmt.Println("1. Tambah Player")
        fmt.Println("2. Info Player")
        fmt.Println("3. Hapus Player")
        fmt.Println("4. Main Game")
        fmt.Println("5. Urutkan Player (Umur/Nama)")
        fmt.Println("6. Rangking Player (Skor Total Descending)")
        fmt.Println("7. Cari Player (Binary Search ID)")
        fmt.Println("8. Belanja Item/Trophy")
        fmt.Println("9. Akhiri Program")
        fmt.Print("Pilih menu: ")
        fmt.Scan(&menu)

        if menu == 1 {
            var p Player
            fmt.Print("ID: "); fmt.Scan(&p.ID)
            fmt.Print("Nama: "); fmt.Scan(&p.Nama)
            fmt.Print("Umur: "); fmt.Scan(&p.Umur)
            p.Level = 1
            p.Skor = Skor{}
            p.Poin = 0
            p.Item = 0
            p.Trophy = 0
            tambahPlayer(&daftarPlayer, &nPlayer, p)
            simpanPlayer(daftarPlayer, nPlayer)
        } else if menu == 2 {
            tampilInfoPlayer(daftarPlayer, nPlayer)
        } else if menu == 3 {
            var id string
            fmt.Print("ID player yang ingin dihapus: "); fmt.Scan(&id)
            idx := binarySearchID(daftarPlayer, nPlayer, id)
            if idx != -1 {
                hapusPlayer(&daftarPlayer, &nPlayer, idx)
                fmt.Println("Player dihapus.")
                simpanPlayer(daftarPlayer, nPlayer)
            } else {
                fmt.Println("Tidak ditemukan.")
            }
        } else if menu == 4 {
            var id string
            fmt.Print("ID player: "); fmt.Scan(&id)
            idx := binarySearchID(daftarPlayer, nPlayer, id)
            if idx != -1 {
                fmt.Println("Pilih Game: 1.Tebak Angka 2.Tebak Hewan 3.Perkalian")
                var g int
                fmt.Scan(&g)
                if g == 1 {
                    gameTebakAngka(idx)
                } else if g == 2 {
                    gameTebakHewan(idx)
                } else if g == 3 {
                    gamePerkalian(idx)
                }
                simpanPlayer(daftarPlayer, nPlayer)
            } else {
                fmt.Println("Player tidak ditemukan.")
            }
        } else if menu == 5 {
            fmt.Println("1. Selection Sort Umur")
            fmt.Println("2. Insertion Sort Nama")
            var sortMenu, asc int
            fmt.Print("Pilih: "); fmt.Scan(&sortMenu)
            fmt.Print("Ascending (1) / Descending (0): "); fmt.Scan(&asc)
            if sortMenu == 1 {
                selectionSortUmur(&daftarPlayer, nPlayer, asc == 1)
            } else if sortMenu == 2 {
                insertionSortNama(&daftarPlayer, nPlayer, asc == 1)
            }
            fmt.Println("Data diurutkan.")
        } else if menu == 6 {
            selectionSortSkor(&daftarPlayer, nPlayer)
            tampilRangking(daftarPlayer, nPlayer)
        } else if menu == 7 {
            cariPlayerByID(&daftarPlayer, nPlayer)
        } else if menu == 9 {
            simpanPlayer(daftarPlayer, nPlayer)
            fmt.Println("Terima kasih telah bermain!")
            return
        } else if menu == 8 {
            var id string
            fmt.Print("ID player: "); fmt.Scan(&id)
            idx := binarySearchID(daftarPlayer, nPlayer, id)
            if idx != -1 {
                menuBelanja(idx)
                simpanPlayer(daftarPlayer, nPlayer)
            } else {
                fmt.Println("Player tidak ditemukan.")
            }
        }
    }
}

//kuda makan tahu 2 kali sehari