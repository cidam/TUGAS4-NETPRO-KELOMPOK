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
