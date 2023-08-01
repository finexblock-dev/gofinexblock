cd ~
cd go/src

docker rm -f $(docker ps -aq)
docker rmi $(docker images | grep -v 'golang' | awk '{if(NR>1) print $3}')
sudo systemctl start docker

docker build -t finexblock-bitcoin-daemon -f build/bitcoin.daemon.dockerfile .

docker run -d --network=host finexblock-bitcoin-daemon:latest