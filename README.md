# GAMESTORE

● As User (Employee GAMESTORE) I Can Store Data Customer GAMESTORE<br />
● As User (Owner GAMESTORE) I Can See Detail Transaction (Revenue every Visit, PurchaseDate on Visit, Unique ID Customer)<br />
● As User (Owner GAMESTORE) I Can see List Of Disc Coupon By Customer (Unique ID)<br />
● As User (Employee GAMESTORE) Can settle the transaction with input (UniqueID (Hsay123Sas))<br />
● As User (Employee GAMESTORE) Can Settle The Transaction UniqueID with coupon disc code and the normalize price<br />
● Can Give / Generate Discount to Stimulate Next Revenue *Coupon one time use<br />

-Pelanggan yang mempunyai revenuenya
more than 25.000.000 give 30% discount (Accessories & New Game)<br />
-Pelanggan yang mempunyai revenuenya melebihi
13.000.000 get disc 15% (Service Console)<br />
-Pelanggan yang purchase more than 6.000.000 get disc 5% (Second Game)<br />

Generate ketika transaksi sudah mengikuti threshold

Contoh Coupon Disc :  <br />
•kupon 30% dengan prefix ULTI: ULTI-RND7821387123456(16 digit)<br />
•kupon 15% dengan prefix PREMI: PREMI-RND1209318092312(16 digit)<br />
•kupon 5% dengan prefix BASIC: BASIC-RND1923808132345(16 digit)<br />

When Use Coupon :<br />
•Cek Expired DATE Coupon<br />
•Cek Unique ID == Unique ID yang terdaftar pada Coupon<br />
•Cek revenue pelanggan dengan Unique ID<br />

Can use by criteria typenya,<br />
Criteria type :<br />
•Service Console<br />
•Buy New Console<br />
•Buy New Game<br />
•Buy Second Game<br />
•Buy Accessories Console<br />

![alt text](https://github.com/rubutar/pc-shop-final-project/blob/main/Picture1.png?raw=true)

