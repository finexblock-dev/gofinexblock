rm -rf /home/ec2-user/go/src

mkdir -p /home/ec2-user/s3
mkdir -p /home/ec2-user/go/src

aws s3 cp s3://finexblock-bitcoin-daemon/dev/.env /home/ec2-user/go/src/.env