# 患者登録
class = 01  患者登録
class = 02  患者情報更新
class = 03  患者情報削除
class = 04  保険追加(xml2のみ)
Content-Type : application/xml


http put "http://172.16.205.128:8000/orca12/patientmodv2?class=01" -a ormaster:kadai \
    token@touroku.xml 


    http://172.16.205.128:8000/api01rv2/patientgetv2?id=1&format=json


# 患者１番を取得
http get  "http://172.16.205.128:8000/api01rv2/patientgetv2?id=1&format=json" -a ormaster:kadaiobgy
