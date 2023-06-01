echo "*******Installing AWS Inspector Agent********"
curl -o "/tmp/inspector_install" https://inspector-agent.amazonaws.com/linux/latest/installsudo
bash /tmp/inspector_installsudo /etc/init.d/awsagent start
rm -rf /tmp/inspector_install