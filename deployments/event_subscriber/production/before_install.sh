rm -rf /home/ec2-user/go/src
rm -rf /home/ec2-user/s3

ls
mkdir -p /home/ec2-user/s3
mkdir -p /home/ec2-user/go/src

aws s3 cp s3://finexblock-event-subscriber/prod/.env /home/ec2-user/go/src/.env
