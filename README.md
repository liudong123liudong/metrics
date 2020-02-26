安装位置
    Linux：/usr/local/cloudmonitor
    Windows： C:\Program Files\futong\cloudmonitor

进程信息
    插件安装后，会在您的服务器上运行以下进程：
    Linux 32位：futongAgent.linux-386
    Linux 64位：futongAgent.linux-amd64
    Windows 32位：futongAgent.windows-386.exe
    Windows 64位：futongAgent.windows-amd64.exe

插件日志
    监控数据日志位于安装位置logs目录下

资源占用情况
    插件进程占用10~20M左右内存和1~2%的单核CPU。
    安装包大小在10~15M。
    日志最多占用40M空间，超过40M会进行清除。
    每15秒发送一次监控数据

Linux插件下载说明
    xxx

Linux插件安装说明
    # 注册为系统服务
    /usr/local/cloudmonitor/futongAgent.linux-${ARCH} install
    # 从系统服务中移除
    /usr/local/cloudmonitor/futongAgent.linux-${ARCH} uninstall
    # 启动
    /usr/local/cloudmonitor/futongAgent.linux-${ARCH} start
    # 停止
    /usr/local/cloudmonitor/futongAgent.linux-${ARCH} stop
    # 重启
    /usr/local/cloudmonitor/futongAgent.linux-${ARCH} restart
    # 卸载
    /usr/local/cloudmonitor/futongAgent.linux-${ARCH} stop && \
    /usr/local/cloudmonitor/futongAgent.linux-${ARCH} uninstall && \
    rm -rf /usr/local/cloudmonitor