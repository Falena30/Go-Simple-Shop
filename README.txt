TODO

1. Checkout :
    a. HTML (done)
    b. System (done)
    c. Database (done)
2. Login :
    a. Database (done)
    b. HTML (not yet start)
    c. System (not yet start)
3. Super Admin
    a. Database (done)
    b. HTML (not yet start)
    c. system (not yet start)
	
SQL Example Join 2 table
SELECT Nama_Pemebeli, Keranjang.ID_Barang,Nama_Barang FROM Keranjang // untuk menampilkan hasil query nanti PS : untuk relasi tidak perlu menambhkan nama tablenya langsung panggil saja nama row dari tabelnya 
INNER JOIN Daftar_Barang // nama table yang akan di join
ON Keranjang.ID_Barang = Daftar_Barang.ID_Barang // mencaari secondary keynya
WHERE Nama_Pemebeli = "dimas"

SQL Example join 3 table
SELECT U_Username, Nama_Pemebeli,Nama_Barang FROM Keranjang
INNER JOIN Daftar_Barang ON Keranjang.ID_Barang = Daftar_Barang.ID_Barang
INNER JOIN User ON User.ID_User = Keranjang.ID_User
WHERE Nama_Pemebeli = "dimas"