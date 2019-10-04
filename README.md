## TUGAS 4 NETPRO ##

<p align="center"
  <a><strong>  NAMA KELOMPOK :  </strong></a> 
</p>
<p align="center">
  <a>  I Putu Surya Baratha (1301188566)  </a> 
</p> 

<p align="center">
  <a>  Muhammad Risdham Nur A.P (1301188603)  </a> 
</p>

<p align="center">
  <a>  Sella Tresnasari  (1301188565)  </a> 
</p> 

#### HASIL RUNNING NOMOR 1 ####

Database yang dipakai adalah MongoDB. MongoDB adalah sistem basis data berorentasi dokumen lintas platform. MongoDB diklasifikasikan sebagai basis data yang “NoSQL”, MongoDB menghindari struktur basis data relasional tabel berbasis tradisional yang mendukung JSON seperti dokumen dengan skema dinamis, membuat integrase data dalam beberapa jenis aplikasi lebih mudah dan juga cepat.

[![message-Image-1570012623256.jpg](https://i.postimg.cc/d1JQYZ4v/message-Image-1570012623256.jpg)](https://postimg.cc/McrJ0TCr)

<p align="center">
  <a> Gambar program 1 </a>
</p>

[![Screen-Shot-2019-10-02-at-17-37-50.png](https://i.postimg.cc/T3jZX7SL/Screen-Shot-2019-10-02-at-17-37-50.png)](https://postimg.cc/LJXytTm2)

<p align="center">
  <a> Gambar program 2 </a>
</p>

[![Screen-Shot-2019-10-02-at-17-37-57.png](https://i.postimg.cc/KzsYDvD2/Screen-Shot-2019-10-02-at-17-37-57.png)](https://postimg.cc/JDZm18F6)

<p align="center">
  <a> Gambar program 3 </a>
</p>

[![Screen-Shot-2019-10-02-at-17-38-24.png](https://i.postimg.cc/GmsvB6rt/Screen-Shot-2019-10-02-at-17-38-24.png)](https://postimg.cc/v4b4kq4w)

<p align="center">
  <a> Gambar program 4 </a>
</p>

[![Screen-Shot-2019-10-02-at-17-38-32.png](https://i.postimg.cc/7h8ZWtds/Screen-Shot-2019-10-02-at-17-38-32.png)](https://postimg.cc/Fdpvfp93)

<p align="center">
  <a> Gambar program 5 </a>
</p>

[![Screen-Shot-2019-10-02-at-17-38-41.png](https://i.postimg.cc/JnWwVw7k/Screen-Shot-2019-10-02-at-17-38-41.png)](https://postimg.cc/wtF4LGmq)

<p align="center">
  <a> Gambar program 6 </a>
</p>

[![Screen-Shot-2019-10-02-at-17-38-49.png](https://i.postimg.cc/gjWGDXL1/Screen-Shot-2019-10-02-at-17-38-49.png)](https://postimg.cc/gXg9c0DK)

<p align="center">
  <a> Gambar program 7 </a>
</p>


###### Keterangan Gambar ######

* Gambar 1 : 
* Gambar 2 : Pada gambar kedua ini berfungsi untuk menambahkan data karyawan baru dengan format nama, email, phone number, dan address
* Gambar 3 : Pada gambar ketiga ini adalah hasil dari penambahan data karyawan baru pada gambar sebelumnya
* Gambar 4 : Pada gambar keempat berfungsi untuk mengedit data karyawan. Data-data yang dapat diedit adalah nama, email, phone number dan juga address
* Gambar 5 : Pada gambar kelima adalah hasil edit pada gambar di nomor 4. Hasil editannya adalah namanya berubah dari yang sebelumnya kelompok 4 menjadi kelompok 40
* Gambar 6 : Pada gambar keenam berfungsi untuk menghapus data karyawan, tapi sebelum akan dihapus akan muncul konfirmasi untuk menanyakan kembali apakah benar-benar akan dihapus data karyawannya
* Gambar 7 : Pada gambar terakhir adalah hasil setelah data karyawan kelompok 40 dihapus pada gambar sebelumnya


#### HASIL RUNNING NOMOR 2 ####

[![Screen-Shot-2019-10-02-at-15-33-07.png](https://i.postimg.cc/28wtn6WK/Screen-Shot-2019-10-02-at-15-33-07.png)](https://postimg.cc/HVr3gTBw)

MQTT (Message Quequing Telemetry Transport) merupakan suatu protokol yang biasa digunakan untuk berkomunikasi lewat internet. Beberapa istilah penting dalam MQTT adalah
-	Publish dan Subscribe
-	Topic
-	Messages
-	BrokerPublish dan Subscribe

MQTT menggunakan model Publish Subscribe. Publisher yang dapat melakukan Publish yaitu membagikan suatu “topic” dan Subscriber yang dapat melakukan subscribe yaitu menerima suatu “topic”.  Jadi, Publish dan Topic ini seperti tipe data khusus untuk membagikan atau menerima suatu topic. Jadi dalam penggunaannya pada protokol MQTT, akan dideklarasikan terlebih dahulu suatu variabel misalkan 

funct main() {
		// Config untuk rest3e
		opts := MQTT.NewClientOptions().AddBroker("localhost:1883")
		opts.SetClientID("subscriber")
		opts.SetDefaultPublishHandler(f)
	

dan

func main() {
		// Config
		opts := MQTT.NewClientOptions().AddBroker("localhost:1883")
		opts.SetClientID("publisher")
		opts.SetDefaultPublishHandler(f)
