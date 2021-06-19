sudo pkill core
sleep 5
rm -rf /home/aa/acopy/; mkdir /home/aa/acopy
cp -r /home/aa/.config/bitclout/bitclout/MAINNET/badgerdb /home/aa/acopy
rm home/aa/acopy/badgerdb/*.mem
./clout-cli/clout sync --dir=/home/aa/acopy/badgerdb
