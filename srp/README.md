```
# echo -ne '\xff\xff\xff\xff\x21\x4c\x5f\xa0\x05\x00\x00\x00\x08\xd2\x09\x10\x00'|nc -u 90.127.28.132 27036 |strings
Phantom0
90-61-AE-B3-42-79z
00-15-5D-44-80-A7z
30-9C-23-48-BF-1Az
90-61-AE-B3-42-78z
00-E0-4C-01-0C-58z
90-61-AE-B3-42-7C
192.168.4.49
192.168.1.10
90.127.28.132
````

```
                 Server Hello(13)                     (9)      (2)  (2)       (V)                    (6)        (6)       L     NAME                                 (9)           (4)                        (5)       (4)    (3)     (3)     (4)  D  L  MACs
                ffffffff214c5fa01600000008 fc89cee79992a7c67e 1001 18f8 d1a387e3dfdffa4cd1      000000080810 06189cd30122 07 5068616e746f6d                   300238 1040014a0e09 704ce804                    0100100110 b4b5d121 580160 bddac8 ef057000 7a 11 39302d36312d41452d42332d34322d3739 7a 11   30302d31352d35442d34342d38302d4137 7a 11 33302d39432d32332d34382d42462d3141 7a 11 39302d36312d41452d42332d34322d3738 7a 11 30302d45302d34432d30312d30432d3538 7a 11 39302d36312d41452d42332d34322d3743 a2 01 0c 3139322e3136382e342e3439 a2 01 0c 3139322e3136382e312e3130 aa 01 0d 39302e3132372e32382e313332
119.195.193.239,ffffffff214c5fa01800000008 bbea8eb5c98bf9ec8a 0110 0118 cfb7a9d588c0b5eb9e0170  000000080810 06189cd30122 0f 4445534b544f502d47383346494935   300238 1040014a0e09 1706381c                    0100100110 88e7c82e 580160 86e6cb ef057000 7a 11 42302d36452d42462d43452d31392d3334 a2 01 0f 3131392e3139352e3139332e323339 aa 01 0f 3131392e31393652e3139332e323339
119.195.193.239,ffffffff214c5fa01800000008 bbea8eb5c98bf9ec8a 0110 0118 cfb7a9d588c0b5eb9e0170  000000080810 06189cd30122 0f 4445534b544f502d47383346494935   300238 1040014a0e09 1706381c                    0100100110 88e7c82e 580160 86e6cb ef057000 7a 11 42302d36452d42462d43452d31392d3334 a2 01 0f 3131392e3139352e3139332e323339 aa 01 0f 3131392e3139352e3139332e323339
176.72.184.35,  ffffffff214c5fa01600000008 a3d5f6e785c790fb44 1001 1894 f7b6e3cf8de6dd30ac      000000080810 06189cd30122 03 4d5349                           300238 1040014a0e09 331fd309                    0100100110 a8dc8721 580160 86e6cb ef057000 7a 11 30382d44342d30432d38302d33452d3531 7a 11   30382d44342d30432d38302d33452d3530 7a 11 44382d43422d38412d46312d31462d3931 7a 11 30382d44342d30432d38302d33452d3534 7a 11 30412d44342d30432d38302d33452d3530 a2 01 0d 3137362e37322e3138342e3335 aa 01 0d 3137362e37322e3138342e3335
107.3.142.211,  ffffffff214c5fa01800000008 cfc9a5feebe1c5b3ba 0110 0118 caefaacfefaedcb1bf01a2  000000080810 06189cd30122 0f 4445534b544f502d504548334c4c43   300238 1040014a0e09 e5631203                    0100100110 ddcecd28 580160 6e6cb  ef057000 7a 11 41342d32422d42302d44412d45322d3446 7a 11   33302d39432d32332d36412d45392d3431 7a 11 31362d32422d42302d44412d45322d3446 7a 11 32362d32422d42302d44412d45322d3446 a2 01 0a 31302e302e302e313838aa010d3130372e332e3134322e323131
68.80.145.186,  ffffffff214c5fa01600000008 abe096f2e4f4aa923b 1001 1890 a7c6fcbef9c2c43c3f      010000080810 06189cd30122 0f 4445534b544f502d443639304e4233   300238 1040014a0e09 3b7dfb02                    0100100110 d7c0da30 580160 85e6cb ef057000 7a 11 30302d30352d39412d33432d37412d3030 7a 11   35302d45302d38352d38362d34352d4439 7a 11 42342d32452d39392d33462d33372d3342 7a 11 30302d31352d35442d30302d30412d3030 7a 11 30322d30302d34432d34462d34462d3530 7a 11 35302d45302d38352d38362d34352d4441 7a 11 39362d31352d45342d31442d45452d3036 7a 11 42342d32452d39392d33462d33372d3339 7a 11 35322d45302d38352d38362d34352d4439 a2010d 3139322e3136382e32342e3230 a2 01 0a 31302e302e302e323033 a2 01 09 31302e302e37352e31 a2 01 0f 3136392e3235342e3234372e323135 a2 01 0d 3137322e32382e31332e313435 aa 01 0d 36382e38302e3134352e313836
                ffffffff214c5fa01600000008 a387b1e68895e1e93a 1001 18c9 d0e7e383899b957a76      000000080810 06189cd30122 10 706576657576652d6c696e75782d7063 300238 c6feffffffff ffffff01 40014a0e09d76ce000 0100100110 daba961f 580160 b9e2d4 ef057001 7a 11 42433a41453a43353a37433a36303a4246 a2 01 0c 3139322e3136382e312e3131 aa 01 0e 39302e3132372e3131372e323234
```

```
00000000: ffff ffff 214c 5fa0 1600 0000 08fc 89ce  ....!L_.........
00000010: e799 92a7 c67e 1001 18f8 d1a3 87e3 dfdf  .....~..........
00000020: fa4c d100 0000 0808 1006 189c d301 2207  .L............".
00000030: 5068 616e 746f 6d30 0238 1040 014a 0e09  Phantom0.8.@.J..
00000040: 704c e804 0100 1001 10b4 b5d1 2158 0160  pL..........!X.`
00000050: 89da c9ef 0570 007a 1139 302d 3631 2d41  .....p.z.90-61-A
00000060: 452d 4233 2d34 322d 3739 7a11 3030 2d31  E-B3-42-79z.00-1
00000070: 352d 3544 2d34 342d 3830 2d41 377a 1133  5-5D-44-80-A7z.3
00000080: 302d 3943 2d32 332d 3438 2d42 462d 3141  0-9C-23-48-BF-1A
00000090: 7a11 3930 2d36 312d 4145 2d42 332d 3432  z.90-61-AE-B3-42
000000a0: 2d37 387a 1130 302d 4530 2d34 432d 3031  -78z.00-E0-4C-01
000000b0: 2d30 432d 3538 7a11 3930 2d36 312d 4145  -0C-58z.90-61-AE
000000c0: 2d42 332d 3432 2d37 43a2 010c 3139 322e  -B3-42-7C...192.
000000d0: 3136 382e 342e 3439 a201 0c31 3932 2e31  168.4.49...192.1
000000e0: 3638 2e31 2e31 30aa 010d 3930 2e31 3237  68.1.10...90.127
```