package main

import (
    "fmt"
)
 
const max = 999
type wisata struct {
    nama, lokasi ,jenis string
    fasilitas[max] string
    jp,jumlahfac int
}
 
type listwisata[max] wisata
var jumdat int
 
 
func main () {
    fmt.Println("========================")
    fmt.Println("Pariwisata Bali.Go.id")
    fmt.Println("========================")
    var pilihan ,idx int
    var data listwisata
 
   
   
    pilihan = mainmenu()
    idx = 0
   
    for pilihan !=3 {
        if (pilihan == 1) {
 
            insert(&data, &jumdat, &idx)
        }else {
            view (&data,&jumdat,&idx)
        }
        pilihan = mainmenu()
    }
       
}
 
func mainmenu () int {
    var pilihmenu int
 
   
    fmt.Println("=====Menu======" )
    fmt.Println("1.insert")
    fmt.Println("2.view")
    fmt.Println("3.exit")
    fmt.Println("===============")
    fmt.Print("Pilih [1-3] :")
    fmt.Scanln(&pilihmenu)
 
    for pilihmenu !=1 && pilihmenu !=2 && pilihmenu != 3{
        fmt.Println("Input salah !!!")
        fmt.Print("Pilih [1-3] :")
        fmt.Scanln(&pilihmenu)
 
    }
   
    return pilihmenu
}
 
func insert (data *listwisata, jumdat *int, idx *int){
 
    var idx_lok, isi int
   
    fmt.Println("============Insert Data==================")
    fmt.Print ("Masukkan jumlah data yang ingin di isi : ")
    fmt.Scanln(&isi)
    *jumdat = *jumdat + isi
    for idx_lok < isi {
        fmt.Println("=================================")
        fmt.Println("Wisata ",idx_lok+1)
        fmt.Print ("Nama : ")
        fmt.Scanln(&data[*idx].nama )
        fmt.Print("Lokasi : ")
        fmt.Scanln(&data[*idx].lokasi)
        fmt.Print ("jenis wisata : ")
        fmt.Scanln(&data[*idx].jenis)
        fmt.Println ("Fasilitas ")
        fmt.Print("Masukkan jumlah fasilitas : ")
        fmt.Scanln(&data[*idx].jumlahfac)
        for j := 0 ;j<data[*idx].jumlahfac ; j++ {
            fmt.Print("Fasilitas ",j+1," : ")
            fmt.Scanln(&data[*idx].fasilitas[j])
        }
        fmt.Print ("jumlah pengunjung : ")
        fmt.Scanln(&data[*idx].jp)
        *idx ++
        idx_lok++
	
    }
	fmt.Println("=================================")
    fmt.Println("Data berhasil ditambahkan ")
   exit()
}
 
func view (data *listwisata, jumdat *int,idx *int) {
    var pilih int

    fmt.Println("=======Menu View========")

    fmt.Println("1.edit")
    fmt.Println("2.delete")
    fmt.Println("3.cari")
    fmt.Println("4.sorting")
	fmt.Println("5.Exit")
    fmt.Print("Pilih [1-5] ? : ")
    fmt.Scanln(&pilih)
   
    switch pilih {
        case 1 :
            fmt.Println("==============Edit==============")
            edit(&*data, &*jumdat)
			
        case 2 :
            fmt.Println("==============Delete==============")
            hapus(&*data, &*jumdat,&*idx)
        case 3 :
             fmt.Println("==============Searching==============")
            cari(&*data,&*jumdat)
        case 4 :
             fmt.Println("==============Sorting==============")
            sorting(&*data,&*jumdat)
		case 5 :
			 return
    }
 
}  
 
func edit (data *listwisata, jumdat *int) {
    var index,pilih ,datedit ,newjp , jmark,editfasi,fscbaru,pilnew int
	var new string
    //var new string
    if (*jumdat==0){
        fmt.Println("data belum diisi")
		exit()
    }
 
    for index < *jumdat {
        fmt.Println (index+1,".",data[index].nama)
        index++
    }
 
    fmt.Print("Edit data no : ")
    fmt.Scanln(&datedit)
	for datedit > *jumdat{
		fmt.Println("Data tidak ditemukan")
		fmt.Print("Edit data no : ")
		fmt.Scanln(&datedit)
	}
	
    fmt.Println("========================")

    fmt.Println("Nama:", data[datedit-1].nama)
    fmt.Println("lokasi:", data[datedit-1].lokasi)
    fmt.Println("jenis wIsata:", data[datedit-1].jenis)
    fmt.Println("fasilitas:")
    for j:=0 ; j< data[datedit-1].jumlahfac ; j++ {
        fmt.Println("fasilitas",j+1, ":", data[datedit-1].fasilitas[j])//menampilkan fasilitas sebelum di edit
    }
    fmt.Println("jumlah pengunjung:", data[datedit-1].jp)
	fmt.Println(" ")
    fmt.Println("Edit ? 1.nama   2.lokasi   3.jenis wisata   4.fasilitas   5.jumlah pengunjung")
    fmt.Print("pilih [1-5] :")
    fmt.Scanln(&pilih)
    for pilih != 1 && pilih != 2 && pilih!=3 && pilih!=4 && pilih !=5 {
		fmt.Println("Inputan anda salah")
		fmt.Print("pilih [1-5] :")
		fmt.Scanln(&pilih)
	}
 
    if (pilih == 1) {
        fmt.Print("Nama baru :")
		fmt.Scanln(&new)
        data[datedit-1].nama = new
		fmt.Println("Nama sudah berhasil di edit")
		exit()
    }else if (pilih == 2) {
        fmt.Print("lokasi :")
        fmt.Scanln(&new)
        data[datedit-1].lokasi = new
		fmt.Println("Lokasi sudah berhasil di edit")
		exit()
    }else if (pilih == 3) {
        fmt.Print("jenis wisata:")
        fmt.Scanln(&new)
        data[datedit-1].jenis = new
		fmt.Println("Jenis wisata sudah berhasil di edit")
	exit()
    }else if (pilih == 4) {
		for j:=0 ; j< data[datedit-1].jumlahfac ; j++ {
			jmark = j
        fmt.Println("fasilitas",j+1, ":", data[datedit-1].fasilitas[j])//menampilkan fasilitas sebelum di edit
		}
		
		fmt.Print("Tambah atau ganti? ")
		fmt.Print("1.Tambah ")
		fmt.Print("2.Ganti : ")
		fmt.Scanln(&editfasi)
		for editfasi !=1 && editfasi!=2 {
			fmt.Print("Tambah atau ganti? ")
			fmt.Print("1.Tambah ")
			fmt.Print("2.Ganti : ")
			fmt.Scanln(&editfasi)
		}
		if editfasi == 1 {
			fmt.Print("Masukkan jumlah fasilitas baru:")
			fmt.Scanln(&fscbaru)
			if (data[datedit-1].jumlahfac==0){
				jmark = -1
			}			
			data[datedit-1].jumlahfac = data[datedit-1].jumlahfac + fscbaru 
			for j:= jmark + 1 ;j< data[datedit-1].jumlahfac ; j++{
				fmt.Print("fasilitas ",j+1,": ")
				fmt.Scanln(&new)
				data[datedit-1].fasilitas[j] = new
			}
			fmt.Println("Fasilitas sudah berhasil di tambah")
			exit()
		}else if editfasi == 2 {
			if (data[datedit-1].jumlahfac==0){
				fmt.Println("Tempat ini belum ada fasilitas")
				exit()
			}
		
			for j:=0 ; j< data[datedit-1].jumlahfac ; j++ {
				fmt.Println("fasilitas",j+1, ":", data[datedit-1].fasilitas[j])//menampilkan fasilitas sebelum di edit
			}
			fmt.Print("Pilih nomor fasilitas yang ingin diubah: ")
			fmt.Scanln(&pilnew)
			fmt.Print("Fasilitasi baru : ")
			fmt.Scanln(&new)
			data[datedit-1].fasilitas[pilnew-1]= new
			fmt.Println("Fasilitas sudah berhasil di ganti")
		exit()
		}
		
		
    } else if pilih == 5{
        fmt.Print("jumlah pengunjung :")
        fmt.Scanln(&newjp)
        data[datedit-1].jp = newjp
		fmt.Println("Jumlah Pengunjung sudah berhasil di ganti")
		exit()
    }
 
}

 
func cari (data *listwisata, jumdat *int) {
    var index ,pilih int
    var found bool
	var lok, jenis string
	var top,btm,mid int
	var sorttmp listwisata
	var temp wisata
	var pass int
    //var lok ,jen string

    found = false
 
    fmt.Print("Cari berdasarkan : \n1.Lokasi wisata\n2.jenis wisata\npilih[1-2] : ")
    fmt.Scanln(&pilih)

    switch pilih {
        case 1 :
            fmt.Print("Lokasi Wisata : ")
            fmt.Scanln(&lok)
		
            for index < *jumdat && !found {
                if(lok == data[index].lokasi){
                    found = true
                }else{
                    index++
                }
            }
 
            if (!found) {
                fmt.Println("Data tidak ditemukan")
            }else {
                for index := 0 ; index < *jumdat ;index ++ {
                    if (data[index].lokasi== lok){
                        fmt.Println("========================")
                        fmt.Println("Nama :",data[index].nama)
                        fmt.Println("Jenis Wisata :",data[index].jenis)
                        fmt.Println("Lokasi :",data[index].lokasi,"\nfasilitas :")
                        for j:= 0; j <data[index].jumlahfac ; j ++ {
                            fmt.Println("fasilitas",j+1,":",data[index].fasilitas[j])
                        }
                        fmt.Println("jumlah pengunjung :",data[index].jp)
						fmt.Println(" ")
                    }
                }
            }
			exit()
			
        case 2 :
		    fmt.Print("Jenis wisata  : ")
            fmt.Scanln(&jenis)
			for i:= 0 ; i<*jumdat ; i++ {
				sorttmp[i]=data[i]       
			}
			pass = 0
			for pass < *jumdat-1 {
				index = pass + 1
				temp = sorttmp[index]
				for index > 0 && sorttmp[index].jenis < sorttmp[index-1].jenis{
					sorttmp[index] = sorttmp[index-1]
					index--
				}
				sorttmp[index] = temp
				pass++
			}
			top = *jumdat
			btm = 0 
			found = false
            for top >=btm && !found {
				mid = (top+btm)/2
				if jenis < sorttmp[mid].jenis{
					top = mid -1
				}else if jenis > sorttmp[mid].jenis{
					btm = mid +1
				}else {
					found= true
				}
			}
			
            if (!found) {
                fmt.Println("Data tidak ditemukan")
            }else {
                for index := 0 ; index < *jumdat ;index ++ {
                    if (sorttmp[index].jenis == jenis){
                        fmt.Println("========================")
                        fmt.Println("Nama :",sorttmp[index].nama)
                        fmt.Println("Jenis Wisata :",sorttmp[index].jenis)
                        fmt.Println("Lokasi :",sorttmp[index].lokasi,"\nfasilitas :")
                        for j:= 0; j <sorttmp[index].jumlahfac ; j ++ {
                            fmt.Println("fasilitas",j+1,":",sorttmp[index].fasilitas[j])
                        }
                        fmt.Println("jumlah pengunjung :",sorttmp[index].jp)
                    }
                }
			}
		exit()
    }
}
 
func sorting(data *listwisata, jumdat *int) {
    var index ,pilih ,pass,idxmax,idxmin int
    var ad  string
    var t wisata //variabel untuk melakukan penukaran tipe data
    var sorttmp listwisata //variabel lokal agar setelah keluar menu sorting data tetap dan tidak berubah

	if (*jumdat==0){
        fmt.Println("data belum diisi")
		exit()
    }
    fmt.Print("Urutkan berdasarkan \n1.Nama\n2.jumlah pengunjung\npilih[1-2] : ")
    fmt.Scanln(&pilih)
	
    for i:= 0 ; i<*jumdat ; i++ {
        sorttmp[i]=data[i]       
    }
 
    if (pilih == 1) {
        fmt.Print("membesar/mengecil ?: ")
        fmt.Scanln(&ad)
        if (ad == "membesar"){
            fmt.Println(" ")
            pass = 0
            for pass < *jumdat - 1 {
                index = pass + 1
                t = sorttmp[index]
                for index > 0 && sorttmp[index].nama < sorttmp[index-1].nama {
                    sorttmp[index]= sorttmp[index - 1]
                    index --
                }
                sorttmp[index] = t
                pass++
            }
 
            for i:= 0 ;i< *jumdat ; i ++ {
                fmt.Println("========================")
                fmt.Println("Nama :" ,sorttmp[i].nama)
                fmt.Println("Lokasi :" ,sorttmp[i].lokasi)
                fmt.Println("jenis wisata:",sorttmp[i].jenis)
                fmt.Println("fasilitas :")
                for j:= 0 ; j < sorttmp[i].jumlahfac ; j++ {
                    fmt.Println("fasilitas ",j+1,":",sorttmp[i].fasilitas[j])
                }
                fmt.Println("jumlah pengunjung :",sorttmp[i].jp)
				fmt.Println(" ")
				
            }
			exit()
		
        }else if (ad =="mengecil"){
            fmt.Println(" ")
            pass = 0
            for pass < *jumdat - 1 {
                index = pass + 1
                t = sorttmp[index]
                for index > 0 && sorttmp[index].nama > sorttmp[index-1].nama {
                    sorttmp[index]= sorttmp[index - 1]
                    index --
                }
                sorttmp[index] = t
                pass++
            }
 
            for i:= 0 ;i < *jumdat ; i ++ {
                fmt.Println("========================")
                fmt.Println("Nama :" ,sorttmp[i].nama)
                fmt.Println("Lokasi :" ,sorttmp[i].lokasi)
                fmt.Println("jenis wisata:",sorttmp[i].jenis)
                fmt.Println("fasilitas :")
                for j:= 0 ; j < sorttmp[i].jumlahfac ; j++ {
                    fmt.Println("fasilitas ",j+1,":",sorttmp[i].fasilitas[j])
                }
                fmt.Println("jumlah pengunjung :",sorttmp[i].jp)
				fmt.Println(" ")
				
            }
		exit()
        }
    }else if (pilih == 2){
        fmt.Print("membesar/mengecil ?: ")
        fmt.Scanln(&ad)
        if (ad == "mengecil"){
            fmt.Println(" ")
            pass = 0
            for pass < *jumdat - 1 {
                idxmax = pass
                index = pass + 1
                for index < *jumdat {
                    if sorttmp[idxmax].jp<sorttmp[index].jp{
                        idxmax = index
                    }
                    index ++
                }
                t = sorttmp[idxmax]
                sorttmp[idxmax]=sorttmp[pass]
                sorttmp[pass]=t
                pass ++
            }
            for i:= 0 ;i< *jumdat ; i ++ {
                fmt.Println("========================")
                fmt.Println("Nama :" ,sorttmp[i].nama)
                fmt.Println("Lokasi :" ,sorttmp[i].lokasi)
                fmt.Println("jenis wisata :",sorttmp[i].jenis)
                fmt.Println("fasilitas :")
                for j:= 0 ; j < sorttmp[i].jumlahfac ; j++ {
                    fmt.Println("fasilitas ",j+1,":",sorttmp[i].fasilitas[j])
                }
                fmt.Println("jumlah pengunjung :",sorttmp[i].jp)
				fmt.Println(" ")
				
            }
			exit()
        }else if(ad == "membesar"){
            fmt.Println(" ")
            pass = 0
            for pass < *jumdat - 1 {
                idxmin = pass
                index = pass + 1
                for index < *jumdat {
                    if sorttmp[idxmin].jp>sorttmp[index].jp{
                        idxmin = index
                    }
                    index ++
                }
                t = sorttmp[idxmin]
                sorttmp[idxmin]=sorttmp[pass]
                sorttmp[pass]=t
                pass ++
            }
            for i:= 0 ;i< *jumdat ; i ++ {
                fmt.Println("========================")
                fmt.Println("Nama :" ,sorttmp[i].nama)
                fmt.Println("Lokasi :" ,sorttmp[i].lokasi)
                fmt.Println("jenis wisata :",sorttmp[i].jenis)
                fmt.Println("fasilitas :")
                for j:= 0 ; j < sorttmp[i].jumlahfac ; j++ {
                    fmt.Println("fasilitas ",j+1,":",sorttmp[i].fasilitas[j])
                }
                fmt.Println("jumlah pengunjung :",sorttmp[i].jp)
				fmt.Println(" ")
				
            }
			exit()
			
        }
 
 
    }
 
}
 
func hapus(data *listwisata, jumdat *int, idx *int){
    var index,pilih int
    var yakin string
    var t wisata
	

    if (*jumdat==0){
        fmt.Println("data belum diisi")
		exit()
    }
 
    for index < *jumdat {
        fmt.Println(index+1,".",data[index].nama)
        index ++
    }
 
    fmt.Print("pilih data yang akan di hapus : ")
    fmt.Scanln(&pilih)

    fmt.Println("========================")
    fmt.Println("Nama:", data[pilih-1].nama)
    fmt.Println("lokasi:", data[pilih-1].lokasi)
    fmt.Println("jenis wisata:", data[pilih-1].jenis)
    fmt.Println("fasilitas:")
    for j:=0 ; j< data[pilih-1].jumlahfac ; j++ {
        fmt.Println("fasilitas",j+1, ":", data[pilih-1].fasilitas[j])
    }
    fmt.Println("jumlah pengunjung:", data[pilih-1].jp)
    fmt.Print("Apakah anda yakin ntuk menghapus data ini ('y/n')? : ")
    fmt.Scanln(&yakin)
 
    if (yakin == "y") {
        *idx = pilih - 1
        for *idx < *jumdat {
            t = data[*idx]
            data[*idx] = data[*idx+1]
            data[*idx + 1] = t
            *idx++
        }
        *jumdat = *jumdat - 1
        *idx = *idx - 1
        fmt.Println("Data berhasil dihapus")
       exit()
    }else if (yakin == "n"){
       exit()
    }
 
 
}

func exit(){
	var exit int
	fmt.Print("Tekan 7 untuk kembali: ")
	fmt.Scanln(&exit)
	for exit != 7 {
		fmt.Print("Tekan 7 untuk kembali: ")
		fmt.Scanln(&exit)
	}
	return
}