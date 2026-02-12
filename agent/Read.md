验证白名单IP 
curl -i http://127.0.0.1:9999/ws/activity-script
验证 token 有效性
wscat -c ws://127.0.0.1:9999/ws/activity-script -H "X-Agent-Token:AGENT_14ee06a05db8de56bb5ff5ff0b81793c46af9657e4c06d1a39a3e377d8523225"
