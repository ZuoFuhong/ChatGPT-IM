"use strict";
let app = new Vue({
    el: '#container',
    data: {
        activeSession: 0,
        hoverSession: 0,
        activeTab: 1,
        loginUser: {},
        sendText: '',
        users: [],
        sessionUsers: [],
        friends: [],
        groupList: [],
        messageCache: {},
        wsclient: null,
        seq: "0",
        audioDevice: null,
        recording: false
    },
    async created() {
        let userId = getQueryVariable("uid")
        await this.loadLoginUser(userId)
        await this.loadUserFriends(userId)
        await this.initWSConn()
        await this.initAudioDevice()
    },
    updated() {
        this.$nextTick(() => {
            let container = this.$el.querySelector("#message-panel");
            container.scrollTop = container.scrollHeight;
        });
    },
    methods: {
        // 初始化音频驱动
        async initAudioDevice() {
            const that = this
            let audioDevice = new AudioRecorder()
            audioDevice.getAudioRecorderDevice()
            // 注册回调-音频转文本
            audioDevice.addOnStopCallback(function(blob) {
                var data = new FormData()
                data.append('model', 'whisper-1')
                data.append('file', blob, 'translate_tts.mp3')

                const xhr = new XMLHttpRequest()
                // xhr.withCredentials = true
                xhr.addEventListener('readystatechange', function() {
                    if(this.readyState === 4) {
                        const rsp = JSON.parse(this.responseText)
                        // 发送转录的文本
                        that.wsClientSendToUser(rsp.text)
                    }
                })
                xhr.open('POST', config.BASE_URL + '/audio/transcriptions')
                xhr.send(data)
            })
            this.audioDevice = audioDevice;
        },
        // 录制音频
        audioRecord() {
            if (!this.recording) {
                // 开始录制
                console.log('start audio record')
                this.audioDevice.startRecord()
                this.recording = true
            } else {
                // 停止录制
                console.log('stop audio record')
                this.audioDevice.stopRecord()
                this.recording = false
            }
        },
        async loadLoginUser(uid) {
            let res = await AsyncGet(`user/info?uid=${uid}`)
            if (res.error_code !== undefined) {
                alert(res.msg)
            } else {
                this.loginUser = {
                    userId: res.user.user_id,
                    avatar: res.user.avatar_url,
                    username: res.user.nickname,
                    personSignature: res.user.extra,
                    deviceId: res.deviceId,
                    token: res.token
                }
            }
        },
        async loadUserFriends(uid) {
            let res = await AsyncGet(`friend/list?uid=${uid}`)
            if (res.error_code !== undefined) {
                alert(res.msg)
            } else {
                let friends = []
                res.forEach(function(value) {
                    let friend = {
                        userId: value.user_id,
                        avatar: value.avatar_url,
                        username: value.nickname,
                        personSignature: value.extra
                    }
                    friends.push(friend)
                })
                this.friends = friends
            }
        },
        async initWSConn() {
            this.wsclient = new Websocket(config.WS_URL, this.handleWSMessage)
            let that = this;
            setTimeout(function() {
                that.wsClientAuth()
            }, 1500)
        },
        closeDialog: function() {
            this.activeSession = 0;
            this.users = [];
        },
        switchSession: function(userId) {
            this.activeSession = userId
        },
        confirmSendMsg() {
            let content = trim(this.sendText)
            if (content) {
                this.wsClientSendToUser(content)
            }
            this.sendText = '';
        },
        selectStyle(userId) {
            this.hoverSession = userId;
        },
        outStyle(userId) {
            this.hoverSession = 0;
        },
        // 激活会话
        openActiveSession(userId) {
            // 添加用户到会话
            let sessionUsers = this.sessionUsers
            let sessionUserIds = [];
            sessionUsers.forEach(function(value) {
                sessionUserIds.push(value.userId)
            })
            let friends = this.friends
            friends.forEach(function(value) {
                if (userId === value.userId) {
                    if (!sessionUserIds.includes(value.userId)) {
                        sessionUsers.unshift(value)
                    }
                }
            })
            // 添加会话的用户到窗口
            let users = this.users;
            let userIds = [];
            users.forEach(function (value) {
                userIds.push(value.userId)
            })
            sessionUsers.forEach(function (value) {
                if (userId === value.userId) {
                    if (!userIds.includes(userId)) {
                        users.unshift(value)
                    }
                }
            })
            this.activeSession = userId;
        },
        // 关闭会话
        closeSession(userId) {
            let users = this.users;
            let index = 0;
            users.forEach(function (value, k) {
                if (userId === value.userId) {
                    index = k;
                }
            })
            if (users.length > 1) {
                users.splice(index, 1);
                this.activeSession = users[0].userId;
            } else {
                this.activeSession = 0;
                this.users = []
            }
        },
        switchTab(tabId) {
            this.activeTab = tabId
        },
        notSupport() {
            alert("暂不支持，敬请期待")
        },
        handleWSMessage(evt) {
            let packa = JSON.parse(evt.data)
            let that = this;
            switch(packa.Type) {
                case 1:
                    this.wsClientSync();
                    break;
                case 2:
                    packa.Data.Messages.forEach(function(message) {
                        that.syncLocalCacheChatLogs(message)
                    })
                    break;
                case 3:
                    // one heartbeat in 10 seconds
                    break;
                case 4:
                    // TODO: ACK
                    break;
                case 5:
                    that.syncLocalCacheChatLogs(packa.Data)
                    break;
                default:
            }
        },
        // 客户端授权
        wsClientAuth() {
            let data = JSON.stringify({
                appId: "1",
                userId: this.loginUser.userId,
                deviceId: this.loginUser.deviceId,
                token: this.loginUser.token
            })
            this.wsclient.pushToServer(JSON.stringify({
                type: 1,
                requestId: 0,
                data: data
            }))
        },
        // 离线消息同步
        wsClientSync() {
            let data = {
                seq: this.seq
            }
            this.wsclient.pushToServer(JSON.stringify({
                type: 2,
                requestId: 0,
                data: JSON.stringify(data)
            }))
        },
        wsClientSendToUser(content) {
            let data = {
                AppId: "1",
                SenderId: this.loginUser.userId,
                DeviceId: this.loginUser.deviceId,
                ReceiverType: 1,
                ReceiverId: this.activeSession,
                MessageType: 1,
                MessageContent: content,
                ToUserIds: []
            }
            this.wsclient.pushToServer(JSON.stringify({
                type: 5,
                requestId: 0,
                data: JSON.stringify(data)
            }))
        },
        syncLocalCacheChatLogs(message) {
            let userId = this.loginUser.userId;
            let key = `${userId}-${message.SenderId}`
            if (userId === message.SenderId) {
                key = `${userId}-${message.ReceiverId}`
            }
            let messageCache = this.messageCache;
            let messageList = [];
            if (!messageCache.hasOwnProperty(key)) {
                messageCache[key] = messageList;
            } else {
                messageList = messageCache[key];
            }
            let sender = this.extractUserInfo(message.SenderId);
            messageList.push({
                userId: sender.userId,
                avatar: sender.avatar,
                username: sender.username,
                createTime: message.SendTime,
                content: message.Content,
            })
            // trigger computed
            this.messageCache = null;
            this.messageCache = messageCache;
        },
        extractUserInfo(userId) {
            let loginUser = this.loginUser;
            let friends = this.friends;
            let user = {
                userId: userId
            };
            if (userId === this.loginUser.userId) {
                user.avatar = loginUser.avatar;
                user.username = loginUser.username;
            } else {
                friends.forEach(function (value) {
                    if (userId === value.userId) {
                        user.avatar = value.avatar;
                        user.username = value.username;
                    }
                })
            }
            return user;
        },
    },
    computed: {
        curUser: function () {
            let curUser;
            let that = this;
            this.users.forEach(function (value) {
                if (that.activeSession === value.userId) {
                    curUser = value;
                }
            })
            if (!curUser) {
                curUser = {}
            }
            return curUser;
        },
        messageList: function() {
            let userId = this.loginUser.userId;
            let key = `${userId}-${this.activeSession}`
            let messageList = [];
            if (this.messageCache.hasOwnProperty(key)) {
                return this.messageCache[key];
            }
            return messageList;
        }
    }
})
