Sending folder 

1. first send root folder
   root folder name is justo 
        4a 75 73 74 6f  justo in hex

        (number of files in hex) the extra  7 zeros + (size of folder in 8 bytes either little endian or small endian) + foldername + zero + 8 ff's

        0b(number of files in hex)   +   00 00 00 00 00 00 00   +    a5 fc 08 00 00 00 00 00   +    4a 75 73 74 6f + 00   +    ff ff ff ff ff ff ff ff

0040   0b  00 00 00 00 00 00 00 a5 fc 08 00 00 00
0050   00 00 4a 75 73 74 6f 00 ff ff ff ff ff ff ff ff


start Zellij themes 14 (20 in decimal) length 13
start Justo 0b (14 in decimal) length 5
start testDir1 04 (4 in decimal) length 8




2. Sending the files
 - get the path to the filename relative to root folder
   path to filename + 1 byte + size of file + then stream the contents of the file