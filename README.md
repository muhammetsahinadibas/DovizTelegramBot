# DovizTelegramBot

Her saat başında bir (ör: 13:00, 18:00, 00:00) Telegram kanalında güncel döviz kurunu paylaşan Telegram botu.

### 🚀 Aktif olarak çalışan hali: https://t.me/kurfiyatlari

## ✔️ Örnek Çıktı
➔ Dolar: X,X ₺ <br>
➔ Euro: X,X ₺ <br>
➔ İngiliz Sterlini: X,X ₺ <br>
➔ Gram Altın: X,X ₺ <br>
📍 @kurfiyatlari

## 🔗 Gereksinimler
```sh
go get github.com/PuerkitoBio/goquery
```

## 🛠️ Kurulum

1. https://t.me/BotFather üzerinden bir Telegram botu oluşturunuz.
2. Botunuzun paylaşım yapmasını istediğiniz Telegram kanalınızı oluşturunuz.
3. Botunuzu Telegram kanalınıza yönetici olarak ve tam yetki vererek ekleyiniz.
4. Telegram kanalınızın ID değerini öğreniniz. (https://t.me/username_to_id_bot ile öğrenebilirsiniz.)
5. @BotFather üzerinden aldığınız Telegram botunuzun token değerini ve Telegram kanalınızın ID değerini <b>main.go</b> dosyası içerisinde bulunan <b>BotToken</b> ve <b>ChannelID</b> değişkenlerinin içerisine yazınız.

    ```sh
    BotToken := "1234567890:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"
    ```
    ```sh
    ChannelID := "-1234567890"
    ```
6. Kodu terminalinizde çalıştırınız.
    ```sh
    go run main.go
    ```
    Kodu arka planda çalıştırmak için:
    ```sh
    nohup go run main.go &
    ```
## 📫 İletişim
[Twitter](https://twitter.com/muhammetadibas) <br>
[Linkedln](https://linkedin.com/in/muhammetadibas) </br>
[Blog](https://muhammetsahinadibas.com.tr)
