package main

var input_test = `x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504`

var input = `y=966, x=475..477
y=1193, x=518..533
x=453, y=796..804
x=476, y=157..168
y=1837, x=580..633
x=473, y=202..208
x=528, y=562..573
x=464, y=286..296
x=541, y=638..649
x=469, y=1772..1780
x=479, y=644..650
x=473, y=1570..1585
x=544, y=1683..1694
y=536, x=477..481
x=478, y=176..190
y=368, x=583..585
y=1146, x=591..614
y=1283, x=557..560
x=556, y=202..204
x=469, y=828..842
x=520, y=751..765
x=596, y=743..745
x=514, y=56..58
x=451, y=1366..1376
x=606, y=572..582
x=586, y=1267..1277
x=449, y=1135..1156
x=553, y=331..348
x=452, y=113..124
y=1765, x=624..649
x=595, y=1495..1509
x=455, y=304..314
x=627, y=1462..1478
y=1005, x=641..647
x=605, y=303..306
y=825, x=553..573
x=555, y=798..803
y=1156, x=447..449
x=586, y=799..820
y=1863, x=474..490
x=609, y=700..722
x=457, y=893..919
x=569, y=953..965
x=579, y=1090..1100
y=936, x=610..634
y=528, x=542..550
x=564, y=819..821
y=1297, x=594..601
x=611, y=1691..1701
x=485, y=1702..1706
x=602, y=706..709
y=140, x=491..507
x=624, y=412..414
x=556, y=319..333
x=568, y=1708..1719
y=152, x=596..603
y=47, x=627..652
x=596, y=1094..1097
y=295, x=543..571
x=510, y=160..169
y=501, x=481..497
y=1382, x=471..476
y=719, x=525..543
y=1706, x=477..485
y=498, x=509..532
y=1545, x=557..584
x=472, y=493..520
x=527, y=1454..1466
x=653, y=998..1011
y=1113, x=514..542
y=180, x=599..623
x=566, y=576..580
x=487, y=134..146
x=553, y=816..825
y=477, x=477..495
y=1218, x=457..466
y=1160, x=545..547
x=543, y=852..862
x=573, y=934..944
x=480, y=326..335
y=425, x=470..472
y=1464, x=497..513
x=500, y=155..157
x=600, y=1770..1783
x=535, y=379..383
x=485, y=786..789
y=1384, x=633..649
x=579, y=1492..1504
x=584, y=16..20
y=772, x=459..466
x=652, y=41..47
x=557, y=1650..1656
x=620, y=58..68
y=322, x=593..618
x=479, y=1562..1588
y=1541, x=574..577
y=1589, x=505..531
x=536, y=1455..1466
y=1561, x=553..569
x=548, y=563..573
x=631, y=1677..1679
y=1299, x=499..513
y=47, x=528..531
y=124, x=605..631
x=526, y=1370..1380
y=1777, x=579..581
x=582, y=1693..1696
x=572, y=1392..1413
x=577, y=750..753
x=571, y=321..324
x=493, y=443..456
x=509, y=436..456
y=1258, x=481..500
y=1155, x=545..547
y=1554, x=510..519
y=959, x=488..493
y=1181, x=493..496
y=192, x=592..595
x=559, y=819..821
x=574, y=1538..1541
y=19, x=493..507
y=1677, x=580..599
y=1072, x=474..500
x=562, y=978..992
x=614, y=1375..1398
x=595, y=171..192
x=586, y=624..641
x=502, y=1127..1152
x=562, y=16..20
y=955, x=511..514
y=296, x=464..481
x=627, y=63..65
y=1764, x=469..474
x=645, y=1102..1127
x=515, y=1062..1068
y=1378, x=550..572
x=574, y=910..930
x=547, y=610..624
x=554, y=1111..1123
x=536, y=814..817
x=552, y=1858..1861
x=575, y=281..287
x=526, y=793..803
y=1057, x=582..598
x=607, y=351..353
x=586, y=921..925
x=582, y=433..446
x=636, y=1504..1528
x=557, y=1389..1400
x=588, y=1151..1163
x=571, y=1650..1656
x=511, y=1480..1498
y=1653, x=582..630
y=453, x=477..486
y=1190, x=486..511
x=497, y=849..859
y=605, x=503..510
x=590, y=295..307
y=854, x=592..611
x=652, y=81..103
x=578, y=1706..1715
y=1016, x=586..606
x=466, y=1193..1218
x=474, y=655..682
x=616, y=1505..1528
y=1746, x=499..504
x=475, y=909..916
y=1100, x=556..574
y=953, x=597..603
x=491, y=920..922
x=617, y=1618..1628
x=643, y=508..509
x=543, y=256..265
x=577, y=1450..1462
x=576, y=573..596
x=633, y=887..911
x=601, y=899..902
y=1393, x=549..551
x=537, y=327..352
x=617, y=1409..1424
y=745, x=596..599
x=628, y=620..629
x=561, y=890..901
y=833, x=582..601
x=447, y=307..317
x=593, y=524..533
x=592, y=1044..1053
y=1145, x=545..564
x=556, y=1089..1100
y=63, x=580..600
y=1198, x=572..575
y=1318, x=553..574
y=446, x=582..596
x=603, y=128..152
x=600, y=1191..1194
x=468, y=219..228
y=1444, x=624..642
x=522, y=576..584
y=573, x=528..548
y=582, x=596..606
x=575, y=1265..1272
x=572, y=1489..1499
x=483, y=867..869
x=550, y=420..422
x=488, y=932..959
y=1338, x=549..554
x=546, y=180..188
x=531, y=857..859
y=187, x=490..508
y=456, x=466..493
x=637, y=193..210
x=596, y=799..820
y=492, x=608..624
x=634, y=1541..1567
y=40, x=493..509
x=528, y=392..406
x=566, y=467..472
x=509, y=224..252
y=390, x=465..485
y=146, x=511..535
y=523, x=511..517
y=521, x=590..598
x=449, y=1329..1338
y=1642, x=505..510
x=453, y=304..314
y=1536, x=497..524
y=586, x=449..454
x=504, y=1746..1756
x=580, y=1664..1677
y=1439, x=525..543
y=804, x=453..477
x=610, y=919..936
y=1053, x=590..592
y=1694, x=544..547
y=682, x=474..487
x=537, y=729..741
y=1144, x=617..638
y=817, x=531..536
x=617, y=1096..1098
y=1736, x=559..597
x=536, y=873..884
y=1798, x=615..617
x=604, y=6..12
x=532, y=1152..1163
y=1875, x=571..641
x=498, y=914..926
x=493, y=933..959
x=462, y=1090..1114
x=453, y=477..487
y=1536, x=528..545
x=467, y=1366..1376
y=1272, x=575..577
y=1672, x=590..592
y=1316, x=606..610
x=592, y=1514..1525
x=649, y=1757..1765
x=482, y=1851..1859
x=474, y=26..48
x=569, y=482..483
x=594, y=1284..1297
y=324, x=569..571
x=451, y=936..963
y=756, x=571..588
x=611, y=839..854
x=581, y=397..419
x=496, y=305..315
x=641, y=1717..1731
x=518, y=1420..1422
x=459, y=770..772
x=499, y=1405..1413
y=1199, x=594..613
x=642, y=401..418
y=20, x=562..584
y=36, x=628..637
x=535, y=1079..1080
x=590, y=872..881
x=465, y=1791..1793
x=479, y=1604..1613
x=547, y=1683..1694
x=480, y=913..926
y=325, x=621..647
y=317, x=447..464
x=470, y=417..425
x=554, y=1153..1163
x=545, y=479..502
x=567, y=1068..1079
y=62, x=498..524
x=533, y=937..949
x=529, y=1482..1497
y=940, x=470..482
y=549, x=505..521
x=543, y=1423..1439
x=527, y=1078..1080
x=586, y=241..244
y=103, x=638..652
x=487, y=716..719
x=566, y=953..965
x=613, y=1109..1120
x=490, y=240..252
y=1843, x=514..524
x=596, y=432..446
x=618, y=836..854
y=789, x=478..485
x=592, y=354..372
y=474, x=485..489
x=592, y=1312..1323
y=1822, x=481..485
x=627, y=1678..1679
y=1042, x=504..513
y=752, x=625..649
x=473, y=476..487
x=559, y=576..580
x=636, y=1101..1127
x=625, y=538..563
x=569, y=35..42
y=451, x=477..486
y=688, x=565..590
x=571, y=290..295
y=690, x=536..538
x=611, y=350..353
y=1520, x=579..585
x=496, y=363..372
x=638, y=517..530
y=1307, x=453..469
x=497, y=695..698
y=509, x=643..648
x=487, y=486..497
x=469, y=694..698
y=348, x=550..553
x=469, y=1294..1307
x=562, y=1150..1163
x=560, y=1032..1041
x=595, y=1024..1032
y=268, x=629..632
y=1507, x=448..460
y=1522, x=579..585
y=916, x=475..477
x=472, y=1873..1893
y=1338, x=449..468
x=598, y=509..521
x=582, y=1046..1057
x=501, y=28..35
y=1262, x=581..606
x=636, y=158..169
y=445, x=630..650
x=580, y=318..333
y=369, x=484..487
x=596, y=571..582
x=475, y=1406..1413
y=526, x=489..505
x=484, y=1472..1484
y=1120, x=559..564
y=1898, x=533..546
x=567, y=1268..1277
y=1163, x=532..554
x=590, y=1044..1053
x=470, y=219..228
y=654, x=574..584
x=560, y=381..383
y=1488, x=536..538
y=1708, x=644..650
x=475, y=962..966
x=496, y=725..742
y=1188, x=631..650
x=536, y=1350..1358
x=631, y=1171..1188
x=469, y=1755..1764
y=1397, x=549..551
y=1608, x=527..604
x=545, y=1534..1536
y=1656, x=557..571
y=1399, x=495..520
x=580, y=634..638
x=504, y=1667..1691
x=461, y=665..686
y=188, x=546..556
x=559, y=1726..1736
x=596, y=129..152
x=603, y=1026..1037
y=1715, x=576..578
x=604, y=1168..1179
x=629, y=1270..1279
x=569, y=321..324
x=588, y=899..902
x=556, y=180..188
y=1179, x=576..604
x=559, y=1109..1120
x=592, y=144..157
x=466, y=1033..1036
y=101, x=597..624
x=528, y=45..47
x=574, y=1842..1853
x=538, y=680..690
x=461, y=260..272
x=509, y=620..630
x=527, y=1595..1608
y=1513, x=518..529
y=1567, x=501..527
x=489, y=512..526
x=477, y=797..804
x=455, y=263..275
y=1860, x=637..648
x=525, y=708..719
x=544, y=194..205
y=1613, x=618..644
x=510, y=1743..1760
y=746, x=545..564
x=621, y=795..820
x=453, y=1871..1878
x=560, y=1278..1283
x=618, y=1611..1613
x=636, y=887..911
y=857, x=531..533
y=487, x=453..473
y=42, x=547..569
y=904, x=554..573
x=562, y=1603..1605
x=576, y=1168..1179
x=509, y=31..40
y=1123, x=554..574
x=560, y=914..925
y=1174, x=595..597
y=884, x=536..554
y=1466, x=527..536
y=1756, x=499..504
y=271, x=618..639
x=498, y=50..62
x=521, y=535..549
x=518, y=1507..1513
y=1032, x=595..597
y=1419, x=555..558
x=623, y=113..115
x=493, y=8..19
x=630, y=522..527
y=1567, x=629..634
x=543, y=707..719
y=1480, x=559..577
y=769, x=476..486
y=1807, x=514..521
x=560, y=479..502
y=406, x=483..485
y=1277, x=567..586
x=622, y=1770..1783
x=554, y=1331..1338
y=1793, x=465..491
x=550, y=514..528
y=383, x=513..535
x=604, y=891..894
x=517, y=1523..1525
x=607, y=601..621
x=579, y=1775..1777
y=1478, x=627..643
x=481, y=524..536
x=493, y=30..40
y=99, x=476..525
x=571, y=746..756
y=252, x=472..490
x=515, y=533..544
x=505, y=195..206
y=265, x=528..543
x=643, y=1354..1359
y=1701, x=611..618
x=485, y=464..474
y=1318, x=606..610
x=461, y=111..119
x=554, y=1509..1534
x=617, y=36..52
x=528, y=1535..1536
x=456, y=665..686
x=516, y=1235..1244
y=859, x=531..533
x=575, y=203..204
x=557, y=1532..1545
y=478, x=524..538
x=611, y=1091..1102
y=20, x=627..641
x=524, y=474..478
x=633, y=1826..1837
x=543, y=938..949
x=642, y=612..623
y=1497, x=529..544
x=577, y=1265..1272
x=518, y=1679..1688
x=485, y=402..406
y=1341, x=639..652
y=859, x=497..514
x=570, y=509..514
y=750, x=577..579
y=786, x=489..550
x=592, y=910..930
x=544, y=1048..1070
x=526, y=1352..1362
y=1152, x=480..502
x=525, y=96..99
x=477, y=451..453
y=1679, x=627..631
x=546, y=1048..1070
y=115, x=623..625
x=476, y=578..591
x=599, y=252..266
x=474, y=1056..1072
x=511, y=947..955
x=588, y=1686..1700
x=591, y=1135..1146
x=531, y=1685..1704
x=655, y=1887..1900
x=535, y=1317..1332
y=949, x=533..543
x=490, y=1845..1863
y=1332, x=535..539
y=1760, x=483..510
x=586, y=1003..1016
y=205, x=520..544
x=466, y=770..772
y=1588, x=542..558
x=480, y=345..355
x=611, y=966..980
x=567, y=1412..1422
x=597, y=872..881
x=620, y=1096..1098
y=530, x=621..638
y=1552, x=640..652
x=583, y=846..869
x=622, y=868..883
x=577, y=1430..1441
x=462, y=1611..1636
x=633, y=835..854
x=639, y=1334..1341
x=549, y=1331..1338
y=1859, x=482..484
x=597, y=72..84
x=512, y=302..320
x=618, y=1231..1235
x=503, y=751..765
x=568, y=765..784
x=521, y=1782..1807
x=613, y=95..97
x=454, y=766..777
x=557, y=936..941
x=461, y=177..190
y=65, x=627..634
x=500, y=1480..1498
x=582, y=823..833
x=550, y=713..737
y=1514, x=621..627
x=623, y=159..169
y=321, x=569..571
y=1692, x=482..497
x=624, y=91..101
y=901, x=561..566
x=588, y=524..533
x=586, y=1393..1413
x=472, y=417..425
x=553, y=1559..1561
x=513, y=729..741
x=531, y=1579..1589
x=484, y=1661..1669
x=565, y=684..688
y=1462, x=464..485
x=459, y=844..845
x=475, y=1723..1728
y=1534, x=549..554
x=597, y=1395..1401
x=544, y=1643..1655
x=509, y=889..893
y=353, x=607..611
y=763, x=609..622
x=486, y=765..769
y=837, x=546..548
x=533, y=857..859
x=652, y=1334..1341
x=628, y=26..36
x=569, y=1467..1477
x=482, y=154..157
y=1679, x=516..518
x=587, y=1470..1484
x=553, y=251..270
x=559, y=764..784
y=980, x=611..638
x=574, y=1089..1100
x=586, y=1331..1351
x=649, y=725..752
y=155, x=449..455
x=504, y=1030..1042
x=558, y=998..1004
x=610, y=1316..1318
x=478, y=398..409
y=1528, x=616..636
x=583, y=1384..1388
y=14, x=481..484
x=610, y=1747..1759
x=578, y=1693..1696
x=467, y=1077..1098
x=642, y=665..668
x=643, y=917..937
x=533, y=956..982
x=646, y=59..68
y=1585, x=471..473
y=35, x=592..595
y=232, x=460..478
x=480, y=61..75
x=469, y=201..208
x=458, y=420..431
x=600, y=48..63
y=1220, x=592..609
x=538, y=1857..1861
y=791, x=618..645
x=566, y=1031..1041
y=369, x=624..638
x=471, y=1358..1382
x=595, y=1255..1258
x=489, y=1428..1441
y=1484, x=484..494
x=521, y=1644..1655
y=1818, x=591..609
y=275, x=455..472
y=1893, x=472..500
x=567, y=419..422
y=40, x=584..610
y=1441, x=559..577
x=484, y=8..14
x=464, y=260..272
y=333, x=556..580
x=537, y=69..80
x=635, y=375..390
y=742, x=483..496
x=582, y=1285..1296
x=590, y=510..521
y=146, x=469..487
x=448, y=1871..1878
x=549, y=1746..1759
x=483, y=1742..1760
y=1123, x=603..621
x=517, y=991..1001
y=422, x=550..567
x=500, y=321..342
x=649, y=450..459
y=803, x=555..557
y=902, x=588..601
x=481, y=7..14
x=545, y=1155..1160
x=582, y=1448..1456
x=528, y=161..169
y=1475, x=632..637
y=1102, x=611..626
x=524, y=1667..1691
x=537, y=610..624
x=580, y=1470..1484
x=650, y=1171..1188
x=584, y=144..157
y=210, x=637..653
y=821, x=559..564
x=603, y=1091..1100
x=497, y=1677..1692
x=486, y=451..453
y=681, x=629..633
y=1310, x=531..549
x=519, y=1862..1883
x=638, y=1198..1217
x=472, y=767..777
x=549, y=1411..1422
x=592, y=1001..1012
x=518, y=797..819
x=482, y=1722..1728
y=1376, x=451..467
y=68, x=620..646
x=559, y=1429..1441
y=1565, x=544..594
x=583, y=921..925
x=493, y=599..609
x=527, y=1857..1858
y=922, x=487..491
y=803, x=526..531
x=535, y=1367..1376
y=921, x=583..586
x=591, y=1807..1818
x=598, y=1208..1217
x=590, y=1661..1672
x=460, y=1491..1507
x=601, y=824..833
x=594, y=1554..1565
x=574, y=1112..1123
x=470, y=894..919
x=538, y=473..478
x=594, y=774..784
x=491, y=1319..1334
x=468, y=1492..1515
y=845, x=449..459
x=558, y=956..968
x=455, y=134..155
x=519, y=599..609
x=638, y=368..369
x=550, y=1361..1378
x=643, y=1462..1478
y=609, x=493..519
x=607, y=1191..1194
x=581, y=1775..1777
x=487, y=360..369
x=526, y=956..982
x=595, y=398..419
x=615, y=1788..1798
x=617, y=1788..1798
x=514, y=1319..1334
y=926, x=480..498
x=478, y=1208..1230
x=613, y=868..883
x=457, y=643..650
y=925, x=526..560
x=508, y=180..187
x=584, y=205..229
y=58, x=514..517
x=528, y=620..630
y=169, x=623..636
x=590, y=683..688
y=869, x=466..483
x=572, y=1842..1853
x=455, y=1773..1780
x=559, y=1470..1480
x=618, y=261..271
x=496, y=344..355
x=641, y=1864..1875
x=549, y=1510..1534
y=741, x=513..537
y=1098, x=467..495
y=456, x=509..535
x=572, y=1194..1198
x=576, y=1706..1715
x=614, y=666..668
y=169, x=510..528
x=463, y=325..335
y=1388, x=572..583
x=610, y=1109..1120
x=597, y=1171..1174
x=466, y=61..75
y=1563, x=510..519
y=1398, x=485..487
x=507, y=114..140
x=618, y=1691..1701
y=527, x=630..632
x=545, y=683..696
x=577, y=1469..1480
x=510, y=602..605
y=1456, x=582..585
y=1036, x=462..466
x=614, y=602..621
x=600, y=1208..1217
x=626, y=336..351
y=208, x=485..488
x=504, y=1439..1447
y=1413, x=475..499
y=1380, x=526..544
y=869, x=571..583
x=557, y=1790..1810
x=497, y=488..501
x=573, y=635..638
x=461, y=492..520
x=463, y=111..119
y=1462, x=577..596
y=84, x=597..625
x=472, y=262..275
x=531, y=792..803
y=630, x=509..528
x=632, y=266..268
y=1422, x=509..518
y=1292, x=571..573
x=481, y=286..296
x=597, y=1024..1032
x=484, y=1492..1515
x=449, y=843..845
x=536, y=1759..1772
x=558, y=1208..1221
y=142, x=477..481
x=501, y=1550..1567
x=449, y=135..155
y=623, x=639..642
x=562, y=1845..1856
y=668, x=614..642
x=514, y=946..955
x=621, y=300..325
x=605, y=251..266
x=486, y=1169..1190
y=1243, x=576..592
x=584, y=1027..1037
x=596, y=601..621
y=580, x=559..566
x=563, y=301..307
y=1691, x=504..524
y=1362, x=526..542
x=448, y=1490..1507
x=485, y=1807..1822
x=576, y=58..77
x=574, y=137..165
y=1267, x=614..641
x=637, y=1832..1860
x=558, y=1416..1419
x=609, y=507..511
x=549, y=1393..1397
x=580, y=1844..1856
x=571, y=845..869
x=580, y=1826..1837
y=390, x=631..635
y=596, x=553..576
y=75, x=466..480
y=1271, x=528..541
x=627, y=9..20
x=618, y=296..322
x=631, y=952..959
x=606, y=1316..1318
x=582, y=510..514
y=1484, x=580..587
x=540, y=581..607
x=581, y=1227..1240
x=599, y=1663..1677
x=505, y=1579..1589
y=936, x=557..560
x=592, y=1661..1672
x=604, y=204..229
x=545, y=743..746
x=557, y=137..165
y=97, x=611..613
y=1775, x=579..581
y=168, x=460..476
x=534, y=582..607
x=518, y=1184..1193
y=80, x=483..537
x=520, y=1389..1399
x=604, y=1596..1608
y=1839, x=567..573
x=554, y=873..884
x=629, y=1542..1567
x=587, y=700..722
y=1883, x=517..519
x=619, y=401..418
y=1628, x=547..555
y=1080, x=527..535
x=511, y=503..523
x=498, y=28..35
x=503, y=602..605
x=585, y=364..368
x=449, y=576..586
y=937, x=639..643
y=556, x=556..583
y=883, x=613..622
y=801, x=648..652
y=1221, x=552..558
x=652, y=1540..1552
x=620, y=603..621
x=477, y=909..916
x=552, y=1023..1032
y=925, x=583..586
y=722, x=587..609
y=1114, x=462..464
x=565, y=1284..1296
y=1447, x=504..506
x=560, y=936..941
x=557, y=1279..1283
y=607, x=534..540
x=592, y=839..854
x=542, y=567..570
x=565, y=1310..1312
x=464, y=306..317
x=648, y=789..801
x=528, y=1246..1271
y=35, x=498..501
y=1424, x=596..617
y=820, x=586..596
y=1287, x=612..626
y=372, x=568..592
y=1715, x=496..506
x=597, y=1726..1736
x=593, y=481..483
x=522, y=224..252
x=513, y=533..544
x=528, y=256..265
x=647, y=299..325
x=533, y=1792..1816
x=598, y=456..458
y=563, x=612..625
x=505, y=1635..1642
x=558, y=1886..1900
x=467, y=1661..1662
y=1878, x=448..453
x=526, y=913..925
x=471, y=1570..1585
y=1120, x=610..613
x=621, y=35..52
y=364, x=583..585
x=514, y=849..859
x=639, y=916..937
y=1509, x=589..595
y=533, x=588..593
x=590, y=1068..1079
y=1396, x=506..509
y=881, x=590..597
x=624, y=367..369
x=568, y=1489..1499
x=584, y=649..654
y=820, x=621..632
y=621, x=596..607
x=484, y=1209..1230
y=1376, x=535..537
x=478, y=222..232
y=483, x=569..593
y=1202, x=567..586
y=522, x=630..632
y=1255, x=588..595
y=236, x=641..647
x=513, y=1281..1299
y=314, x=453..455
x=642, y=1442..1444
y=1233, x=448..475
y=1636, x=462..473
y=965, x=566..569
x=575, y=1194..1198
y=270, x=553..579
x=572, y=1385..1388
x=481, y=1251..1258
x=556, y=467..472
y=1238, x=600..628
y=1452, x=553..572
x=487, y=655..682
x=485, y=380..390
x=475, y=1226..1233
y=229, x=584..604
x=632, y=522..527
x=594, y=624..641
x=585, y=457..458
x=526, y=327..352
x=500, y=1057..1072
x=648, y=1616..1628
x=573, y=815..825
y=1244, x=496..516
y=320, x=512..525
x=639, y=262..271
x=468, y=174..186
y=753, x=577..579
y=102, x=469..546
x=593, y=295..322
x=523, y=797..819
x=565, y=1242..1259
y=1783, x=600..622
x=611, y=95..97
y=784, x=594..620
y=671, x=529..555
x=534, y=705..715
x=556, y=539..556
y=877, x=464..469
x=509, y=979..1005
y=52, x=617..621
x=602, y=6..12
x=603, y=938..953
x=487, y=1372..1398
y=1613, x=479..497
x=490, y=978..1005
x=473, y=1258..1264
x=572, y=1361..1378
x=462, y=1033..1036
y=682, x=494..496
x=648, y=507..509
y=719, x=469..487
x=503, y=322..342
x=485, y=1452..1462
x=583, y=364..368
y=157, x=584..592
x=552, y=58..77
x=490, y=179..187
x=608, y=480..492
x=550, y=330..348
x=469, y=875..877
y=1621, x=531..534
x=581, y=1252..1262
x=477, y=524..536
x=544, y=1369..1380
x=494, y=1472..1484
x=630, y=443..445
x=609, y=743..763
x=466, y=866..869
x=482, y=936..940
x=644, y=1610..1613
x=478, y=419..431
x=573, y=1282..1292
x=448, y=1226..1233
x=489, y=775..786
x=627, y=1514..1518
x=491, y=1790..1793
x=629, y=680..681
x=493, y=1066..1069
x=602, y=280..287
x=594, y=1094..1097
y=1669, x=484..492
y=737, x=550..558
x=595, y=26..35
x=653, y=475..484
y=165, x=557..574
x=495, y=1388..1399
x=472, y=174..186
x=641, y=1005..1007
x=493, y=398..409
y=342, x=500..503
x=483, y=1541..1544
x=555, y=658..671
x=477, y=1259..1264
y=621, x=614..620
y=351, x=626..632
y=919, x=457..470
y=459, x=627..649
x=505, y=512..526
x=468, y=1328..1338
y=862, x=525..543
y=1070, x=544..546
x=515, y=556..570
x=627, y=41..47
x=457, y=1563..1588
x=464, y=876..877
x=650, y=443..445
x=482, y=1678..1692
x=447, y=1136..1156
y=119, x=461..463
x=625, y=724..752
y=854, x=618..633
y=352, x=486..489
y=584, x=518..522
y=1477, x=569..571
x=476, y=96..99
x=527, y=682..696
y=1525, x=571..592
y=1851, x=482..484
x=458, y=1165..1171
x=473, y=1612..1636
x=579, y=294..307
x=541, y=1684..1704
y=1486, x=536..538
y=1719, x=568..584
x=517, y=1863..1883
x=599, y=174..180
x=612, y=539..563
x=523, y=412..422
y=1012, x=592..594
x=536, y=680..690
x=570, y=1686..1700
x=649, y=1199..1217
y=1731, x=627..641
y=419, x=581..595
y=698, x=469..497
x=614, y=1134..1146
x=494, y=202..211
x=478, y=362..372
y=992, x=562..586
x=521, y=1760..1772
x=626, y=1280..1287
x=492, y=1662..1669
x=633, y=621..629
x=531, y=1306..1310
x=529, y=1506..1513
x=609, y=1807..1818
y=355, x=480..496
x=484, y=1851..1859
x=551, y=1393..1397
x=478, y=787..789
x=483, y=725..742
y=982, x=526..533
y=1032, x=550..552
y=1728, x=475..482
x=586, y=1227..1240
x=491, y=556..570
y=591, x=476..504
x=548, y=380..383
x=491, y=113..140
x=499, y=1280..1299
y=1499, x=568..572
x=579, y=250..270
x=549, y=1603..1605
x=553, y=572..596
x=468, y=27..48
y=930, x=574..592
x=510, y=1554..1563
x=481, y=1807..1822
x=598, y=1047..1057
x=600, y=1226..1238
x=555, y=1416..1419
y=497, x=487..490
x=513, y=379..383
x=641, y=220..236
x=546, y=90..102
x=592, y=26..35
x=650, y=1698..1708
x=648, y=1832..1860
x=525, y=303..320
x=519, y=1554..1563
x=494, y=679..682
x=621, y=1514..1518
y=567, x=537..542
x=573, y=301..307
x=583, y=539..556
x=476, y=766..769
x=606, y=1004..1016
x=507, y=9..19
x=453, y=1294..1307
x=563, y=795..806
x=567, y=1190..1202
y=911, x=633..636
x=647, y=1005..1007
y=472, x=556..566
x=639, y=612..623
y=715, x=530..534
x=564, y=742..746
y=548, x=530..549
x=498, y=1354..1377
x=525, y=1422..1439
x=637, y=27..36
x=547, y=34..42
y=1334, x=491..514
x=469, y=717..719
y=1186, x=493..496
x=496, y=1711..1715
x=638, y=81..103
x=567, y=1210..1224
x=499, y=1746..1756
x=571, y=1515..1525
y=638, x=573..580
y=963, x=451..460
x=558, y=1310..1312
x=500, y=618..627
y=1523, x=502..517
x=561, y=1328..1341
y=514, x=570..582
x=489, y=464..474
y=1258, x=588..595
x=532, y=1388..1400
x=609, y=1367..1386
x=546, y=1882..1898
x=521, y=637..649
x=629, y=507..511
y=352, x=526..537
x=602, y=1396..1401
x=613, y=1187..1199
y=544, x=513..515
y=1001, x=517..520
x=626, y=1091..1102
x=588, y=1255..1258
x=509, y=1393..1396
x=624, y=1757..1765
x=564, y=1109..1120
y=1414, x=469..471
x=517, y=56..58
y=1341, x=542..561
x=623, y=175..180
y=641, x=586..594
x=569, y=1559..1561
x=514, y=1856..1858
x=614, y=1256..1267
x=647, y=219..236
y=1004, x=545..558
y=458, x=585..598
y=784, x=559..568
x=640, y=1540..1552
x=638, y=965..980
y=1688, x=516..518
x=629, y=266..268
x=638, y=1142..1144
x=606, y=1252..1262
x=490, y=486..497
x=568, y=1791..1810
x=496, y=1203..1218
x=543, y=290..295
y=1217, x=638..649
y=1079, x=567..590
x=511, y=1170..1190
x=465, y=380..390
y=629, x=628..633
y=1171, x=458..479
x=564, y=1134..1145
x=472, y=241..252
x=549, y=536..548
y=409, x=478..493
x=627, y=1716..1731
x=634, y=919..936
y=1068, x=515..524
y=307, x=563..573
y=1240, x=581..586
x=476, y=1357..1382
y=124, x=452..469
x=533, y=1882..1898
x=524, y=1520..1536
x=481, y=1627..1629
x=599, y=742..745
x=470, y=936..940
x=530, y=536..548
x=567, y=1834..1839
y=944, x=547..573
x=531, y=1619..1621
x=469, y=1408..1414
x=575, y=955..968
y=686, x=456..461
x=625, y=113..115
y=1011, x=625..653
y=12, x=602..604
x=447, y=1660..1662
y=841, x=537..554
y=842, x=469..477
y=306, x=601..605
x=540, y=1792..1816
x=497, y=1462..1464
x=597, y=92..101
y=1647, x=611..616
x=586, y=979..992
y=570, x=491..515
y=1230, x=478..484
x=633, y=679..681
x=513, y=1031..1042
x=535, y=435..456
x=594, y=1187..1199
y=1259, x=551..565
y=650, x=457..479
x=586, y=1189..1202
x=596, y=1409..1424
y=211, x=478..494
y=1772, x=521..536
y=570, x=537..542
x=460, y=156..168
x=652, y=788..801
y=1441, x=489..501
x=504, y=578..591
x=536, y=1486..1488
y=1279, x=629..646
x=514, y=1819..1843
x=501, y=1429..1441
x=481, y=489..501
x=616, y=1644..1647
x=518, y=576..584
x=596, y=1451..1462
x=509, y=1421..1422
x=576, y=1230..1243
y=190, x=461..478
x=502, y=1353..1377
y=1218, x=496..504
x=514, y=1782..1807
x=483, y=69..80
y=1704, x=531..541
x=617, y=1142..1144
x=601, y=1284..1297
x=525, y=851..862
x=516, y=1679..1688
x=483, y=402..406
x=646, y=1617..1628
y=1413, x=572..586
x=624, y=1791..1802
x=469, y=91..102
x=605, y=106..124
x=460, y=937..963
x=624, y=481..492
x=577, y=1538..1541
x=562, y=1491..1504
x=531, y=44..47
y=649, x=521..541
x=517, y=504..523
y=1323, x=592..616
y=627, x=465..500
x=545, y=997..1004
x=591, y=1368..1386
x=550, y=1022..1032
x=554, y=893..904
x=542, y=1576..1588
y=1312, x=558..565
y=77, x=552..576
x=530, y=705..715
x=601, y=303..306
y=1816, x=533..540
y=1358, x=534..536
y=1401, x=597..602
x=552, y=1574..1581
x=474, y=1846..1863
x=646, y=1271..1279
x=538, y=1486..1488
x=542, y=1111..1113
x=513, y=1462..1464
y=287, x=575..602
x=634, y=63..65
x=573, y=1834..1839
x=632, y=795..820
y=414, x=624..627
y=1518, x=621..627
x=602, y=890..894
x=574, y=1305..1318
y=1264, x=473..477
y=1416, x=555..558
x=610, y=1332..1351
y=696, x=527..545
y=1069, x=491..493
x=460, y=221..232
x=486, y=342..352
x=547, y=933..944
x=597, y=938..953
x=482, y=890..893
x=632, y=335..351
x=529, y=659..671
y=520, x=461..472
x=592, y=1229..1243
x=611, y=1644..1647
x=606, y=1790..1802
x=542, y=1329..1341
x=495, y=1077..1098
x=500, y=1252..1258
x=653, y=194..210
x=616, y=1313..1323
x=572, y=1447..1452
x=622, y=742..763
x=581, y=242..244
x=494, y=1626..1629
x=550, y=776..786
y=1858, x=514..527
y=941, x=557..560
y=1100, x=579..603
y=186, x=468..472
x=497, y=1604..1613
y=502, x=545..560
x=553, y=1448..1452
x=487, y=920..922
x=547, y=1211..1224
x=625, y=998..1011
x=480, y=1128..1152
x=545, y=1134..1145
y=1853, x=572..574
y=777, x=454..472
x=632, y=1465..1475
x=541, y=1245..1271
y=1194, x=600..607
x=584, y=23..40
x=631, y=375..390
x=477, y=131..142
y=709, x=599..602
y=431, x=458..478
x=620, y=1231..1235
x=508, y=196..206
x=582, y=1640..1653
x=464, y=1091..1114
x=548, y=828..837
y=1504, x=562..579
y=1515, x=468..484
x=534, y=1620..1621
x=489, y=342..352
y=241, x=545..570
x=588, y=746..756
y=1581, x=549..552
y=1628, x=592..617
y=1810, x=557..568
x=457, y=1541..1544
x=549, y=1305..1310
y=95, x=611..613
x=631, y=107..124
x=524, y=1820..1843
x=497, y=1521..1536
y=1465, x=632..637
x=539, y=1318..1332
x=567, y=1769..1780
x=579, y=1520..1522
x=524, y=50..62
x=496, y=679..682
x=544, y=1481..1497
x=610, y=23..40
x=571, y=1467..1477
y=113, x=623..625
x=641, y=10..20
y=1386, x=591..609
x=496, y=1181..1186
x=484, y=360..369
y=418, x=619..642
x=574, y=648..654
x=496, y=1235..1244
x=592, y=1618..1628
x=599, y=706..709
x=592, y=172..192
x=500, y=1873..1893
y=1127, x=636..645
x=568, y=353..372
x=524, y=1062..1068
y=335, x=463..480
y=1629, x=481..494
x=491, y=1066..1069
x=488, y=199..208
x=558, y=1577..1588
x=589, y=1494..1509
x=535, y=144..146
y=48, x=468..474
x=466, y=442..456
y=1525, x=502..517
y=920, x=487..491
x=457, y=1193..1218
y=1098, x=617..620
x=506, y=1711..1715
y=422, x=517..523
y=1662, x=447..467
y=1235, x=618..620
x=533, y=1185..1193
x=612, y=1281..1287
x=554, y=830..841
y=1217, x=598..600
y=272, x=461..464
x=527, y=1551..1567
x=544, y=1554..1565
y=1759, x=549..610
x=520, y=991..1001
x=621, y=1112..1123
y=1224, x=547..567
x=465, y=617..627
x=630, y=1639..1653
x=557, y=798..803
x=504, y=1202..1218
x=624, y=1441..1444
x=637, y=1465..1475
x=613, y=952..959
x=546, y=828..837
y=1097, x=594..596
y=765, x=503..520
x=478, y=201..211
y=1351, x=586..610
y=266, x=599..605
x=539, y=795..806
y=1498, x=500..511
y=228, x=468..470
y=1856, x=562..580
x=649, y=1371..1384
y=1544, x=457..483
x=454, y=575..586
x=594, y=1001..1012
y=1041, x=560..566
x=643, y=474..484
x=505, y=536..549
y=1900, x=558..655
y=1861, x=538..552
x=571, y=1282..1292
y=893, x=482..509
y=383, x=548..560
y=1359, x=616..643
x=553, y=1306..1318
y=1005, x=490..509
x=542, y=513..528
y=894, x=602..604
y=1296, x=565..582
x=517, y=412..422
y=372, x=478..496
x=469, y=114..124
y=1780, x=455..469
y=1007, x=641..647
y=208, x=469..473
y=819, x=559..564
x=534, y=1350..1358
x=509, y=493..498
x=558, y=713..737
x=477, y=467..477
x=514, y=1112..1113
x=537, y=567..570
x=566, y=890..901
y=1802, x=606..624
x=641, y=1257..1267
y=1559, x=553..569
x=479, y=1166..1171
y=1163, x=562..588
x=502, y=1523..1525
x=585, y=1520..1522
x=547, y=1155..1160
x=616, y=1353..1359
x=510, y=1634..1642
x=520, y=193..205
x=506, y=1440..1447
x=532, y=493..498
x=471, y=1408..1414
y=1400, x=532..557
y=624, x=537..547
x=628, y=1225..1238
y=1422, x=549..567
x=645, y=790..791
x=495, y=466..477
x=627, y=412..414
y=511, x=609..629
x=622, y=1375..1398
y=1398, x=614..622
x=573, y=892..904
x=618, y=789..791
x=588, y=1768..1780
x=552, y=1208..1221
x=584, y=1709..1719
x=506, y=1393..1396
y=1628, x=646..648
x=485, y=1372..1398
x=493, y=1181..1186
x=513, y=393..406
x=472, y=306..315
x=551, y=1242..1259
x=592, y=1211..1220
y=1655, x=521..544
x=595, y=1171..1174
x=531, y=813..817
y=1780, x=567..588
x=547, y=1618..1628
y=806, x=539..563
x=585, y=1448..1456
y=1377, x=498..502
y=307, x=579..590
x=633, y=1372..1384
x=485, y=199..208
y=206, x=505..508
y=959, x=613..631
x=464, y=1451..1462
y=1605, x=549..562
x=579, y=750..753
y=1037, x=584..603
y=406, x=513..528
x=627, y=450..459
y=204, x=556..575
x=537, y=831..841
x=477, y=828..842
y=244, x=581..586
x=570, y=236..241
x=571, y=1864..1875
y=968, x=558..575
x=620, y=774..784
y=1696, x=578..582
y=484, x=643..653
x=644, y=1699..1708
x=584, y=1532..1545
x=580, y=47..63
y=1700, x=570..588
y=1331, x=549..554
x=542, y=1353..1362
y=1588, x=457..479
x=477, y=1703..1706
y=252, x=509..522
x=621, y=516..530
x=545, y=235..241
x=549, y=1574..1581
y=315, x=472..496
x=481, y=131..142
x=555, y=1619..1628
x=474, y=1755..1764
y=819, x=518..523
x=477, y=962..966
x=511, y=143..146
x=537, y=1367..1376
x=603, y=1111..1123
x=469, y=133..146
x=609, y=1210..1220
y=157, x=482..500
x=625, y=73..84`
