./script/apistop.sh
nohup ./gorm_generate -t api >> log/api.log 2>&1 &
echo "restart ok"