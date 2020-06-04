let app = new Vue({
    el: '#container',
    data: {
        activeSession: 0,
        hoverSession: 0,
        activeTab: 3,
        loginUser: {
            userId: 2,
            avatar: "./assets/avatar2.jpg",
            username: "熊二",
            personSignature: "俺要吃蜂蜜"
        },
        sendText: '',
        users: [],
        sessionUsers: [],
        friends: [
            {
                userId: 1,
                avatar: "./assets/avatar1.jpg",
                username: "熊大",
                personSignature: "做熊，就要有个熊样"
            },
            {
                userId: 2,
                avatar: "./assets/avatar2.jpg",
                username: "熊二",
                personSignature: "俺要吃蜂蜜"
            },
            {
                userId: 3,
                avatar: "./assets/avatar3.jpg",
                username: "光头强",
                personSignature: "臭狗熊，我饶不了你们！"
            }
        ],
        groupList: [
            {
                groupId: 1,
                avatar: "./assets/group1.png",
                groupName: "伐木经验分享群"
            },
            {
                groupId: 2,
                avatar: "./assets/group2.png",
                groupName: "保卫森林交流群"
            },
        ],
        messageList: [
            {
                userId: 1,
                avatar: "./assets/avatar1.jpg",
                username: "熊大",
                createTime: "2020-06-04 23:10:10",
                content: "hello 熊二"
            },
            {
                userId: 2,
                avatar: "./assets/avatar2.jpg",
                username: "熊二",
                createTime: "2020-06-04 23:10:10",
                content: "好呀好呀"
            }
        ]
    },
    methods: {
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
                let loginUser = this.loginUser
                this.messageList.push({
                    userId: loginUser.userId,
                    avatar: loginUser.avatar,
                    username: loginUser.username,
                    createTime: dateFormat("YYYY-mm-dd HH:MM:SS", new Date()),
                    content: content
                })
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
        }
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
        }
    }
})