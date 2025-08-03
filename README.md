# TLSPass

TLSPass是一个简单的重定向工具，监听443端口，并将访问转发到特定url

certPath: /etc/TLSPass/tlspass.pem,
keyPath: /etc/TLSPass/tlspass.key,
config: /etc/TLSPass/config

~~~config
config:
/path1>http://localhost:port
/path2>http://localhost:port/targetpath
~~~

为了方便使用工具，可以在执行`go build -o TLSPass`后，将执行文件放在`/usr/local/bin/`目录下。执行`TLSPass help`查看指令

监听服务在终端退出会自动退出，可以用systemdctl托管

`sudo nano /etc/systemd/system/TLSPass.service`

~~~ini
[Unit]
Description=TLSPass Server

[Service]
ExecStart=/usr/local/bin/TLSPass start

[Install]
WantedBy=multi-user.target
~~~

~~~bash
sudo systemctl daemon-reload
sudo systemctl enable TLSPass
sudo systemctl start TLSPass
~~~
