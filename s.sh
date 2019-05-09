USERNAME="root"
#数据库密码
PASSWORD="surpass@2018"
#数据库端口
PORT="3306"
#数据SERVER
HOSTNAME="218.17.71.139"
#数据库
DBNAME="test_baisiyi"

select_sql="SHOW TABLES;"
datas=`/usr/local/mysql/bin/mysql -h${HOSTNAME} -P${PORT} -u${USERNAME} -p${PASSWORD} ${DBNAME} --show-warnings=false -e "${select_sql}" | awk 'NR>1'`

if [[ ! "$datas" ]]; then
	echo -e "[$(date "+%Y-%m-%d %H:%M:%S")]\t数据表查询失败"
	exit 1
fi
for name in ${datas}
do
	structname=`echo $name | sed -r -e 's/_[a-zA-Z]/\U&/g'`
	structname=`echo $structname|awk '{print toupper(substr($0,0,1))substr($0,2,length($0))}'`
	structname=`echo ${structname//\_/}`
	/home/go/src/github.com/Shelnutt2/db2struct/db2struct/db2struct -H ${HOSTNAME} -d ${DBNAME} --package=models --json --gorm -u ${USERNAME} -p ${PASSWORD} -t ${name} --struct=${structname} &>>"/home/go/src/baisiyi.net/models/${structname}.go"
done

echo "执行完毕"
exit 0
