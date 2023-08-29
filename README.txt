ДЗ 11. Финальный проект                   

Прототип книжного онлайн-магазина (backend)

За основу взята ДЗ 9 плюс кафка и сервис notification из ДЗ 10

Общая схема (system design) и sequesnce диаграммы доступны в папке ./digrams

Ссылка на презентацию:

https://docs.google.com/presentation/d/1FNoHSPuyeBXL_dS1E9yyEeHLESjph_jz4MwtcCegysk/edit#slide=id.p

Инструкция по запуску:

Развернуть кафку:

01-create-kafka-namespace.cmd
02-create-kafka-crd.cmd
03-create-kafka-cluster.cmd

запуск нод кафки может занять несколько минут, для контроля запуска используется команда:

kubectl get pods -n kafka

Развернуть приложение:

05-install-helm-deps.cmd
06-deploy-app.cmd

Выполнить тесты:

07-run-newman-test.cmd


Инструкция по удалению:

08-delete-app.cmd
09-delete-artifacts.cmd
10-delete-kafka.cmd


 
