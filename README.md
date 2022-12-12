# GAMESTORE

●As User (Employee GAMESTORE) I Can Store Data Customer GAMESTORE
●As User (Owner GAMESTORE) I Can See Detail Transaction (Revenue every Visit, PurchaseDate on Visit, Unique ID Customer)
●As User (Owner GAMESTORE) I Can see List Of Disc Coupon By Customer (Unique ID)
●As User (Employee GAMESTORE) Can settle the transaction with input (UniqueID (Hsay123Sas))
●As User (Employee GAMESTORE) Can Settle The Transaction UniqueID with coupon disc code and the normalize price
●Can Give / Generate Discount to Stimulate Next Revenue *Coupon one time use

-Pelanggan yang mempunyai revenuenya
more than 25.000.000 give 30% discount (Accessories & New Game)
-Pelanggan yang mempunyai revenuenya melebihi
13.000.000 get disc 15% (Service Console)
-Pelanggan yang purchase more than 6.000.000 get disc 5% (Second Game)

Generate ketika transaksi sudah mengikuti threshold

Contoh Coupon Disc :  
•kupon 30% dengan prefix ULTI: ULTI-RND7821387123456(16 digit)
•kupon 15% dengan prefix PREMI: PREMI-RND1209318092312(16 digit)
•kupon 5% dengan prefix BASIC: BASIC-RND1923808132345(16 digit)

When Use Coupon :
•Cek Expired DATE Coupon
•Cek Unique ID == Unique ID yang terdaftar pada Coupon
•Cek revenue pelanggan dengan Unique ID

Can use by criteria typenya,
Criteria type :
•Service Console
•Buy New Console
•Buy New Game
•Buy Second Game
•Buy Accessories Console

![alt text](https://github.com/rubutar/pc-shop-final-project/blob/main/Picture1.png?raw=true)

