<template>
  <div id="app">
    <div class="container">
      <div class="contacts">
          <div class="contact-item" v-for="user in users" :key="user.userId">
              <div class="human" :class="{'human-active': user.userId === activeSession}"
                    v-on:click="switchSession(user.userId)" @mouseover="selectStyle(user.userId)" @mouseout="outStyle(user.userId)">
                  <div class="avatar">
                      <img :src="user.avatar" alt="">
                  </div>
                  <div class="username">
                      {{user.username}}
                  </div>
                  <div class="remove">
                      <button class="removeBtn" @click.stop="closeSession(user.userId)" v-show="hoverSession === user.userId">x</button>
                  </div>
              </div>
          </div>
      </div>
      <div class="session-windows">
          <div class="session-info">
            <div class="avatar">
              <img :src="curUser.avatar" alt="">
            </div>
            <div class="username">
              {{ curUser.username }}
            </div>
          </div>
          <div class="message-panel" id="message-panel">
              <div class="message-item" v-for="message in messageList" :key="message.Seq">
                  <div class="user-info" v-show="message.userId !== loginUser.userId">
                      <div class="avatar">
                          <img :src="message.avatar" alt="">
                      </div>
                      <div class="username">
                          {{ message.username }}
                      </div>
                      <div class="createTime">
                          {{ message.createTime }}
                      </div>
                  </div>
                  <div class="message-info" style="justify-content: flex-start; margin-left: 30px;" v-show="message.userId !== loginUser.userId">
                      <div class="message-icon">
                          <img src="./assets/icon/left-msg.svg" alt="">
                      </div>
                      <div class="message-content">
                          {{ message.content }}
                      </div>
                  </div>
                  <!-- me -->
                  <div class="user-info right" v-show="message.userId === loginUser.userId">
                      <div class="createTime">
                          {{ message.createTime }}
                      </div>
                      <div class="username">
                          {{ message.username }}
                      </div>
                      <div class="avatar">
                          <img :src="message.avatar" alt="">
                      </div>
                  </div>
                  <div class="message-info" style="justify-content: flex-end; margin-right: 28px;" v-show="message.userId === loginUser.userId">
                      <div class="message-content right">
                          {{ message.content }}
                      </div>
                      <div class="message-icon">
                          <img src="./assets/icon/right-msg.svg" alt="">
                      </div>
                  </div>
              </div>
          </div>
          <div class="message-input">
            <div class="plugins">
                <div class="plugin-item" v-on:click="notSupport">
                    <img src="./assets/icon/smile.svg" alt="">
                </div>
                <div class="plugin-item" v-on:click="notSupport">
                    <img src="./assets/icon/image.svg" alt="">
                </div>
                <div class="plugin-item" v-on:click="notSupport">
                    <img src="./assets/icon/file.svg" alt="">
                </div>
                <div class="plugin-item" v-on:click="audioRecord">
                    <img v-if="recording" src="./assets/icon/audio-record.svg" alt="">
                    <img v-if="!recording" src="./assets/icon/audio.svg" alt="">
                </div>
                <div class="plugin-item" v-on:click="notSupport">
                    <img src="./assets/icon/video.svg" alt="">
                </div>
                <div class="plugin-item history" v-on:click="notSupport">
                    <img src="./assets/icon/history.svg" alt="">
                    <span class="history-txt">聊天记录</span>
                </div>
            </div>
            <div class="input-area">
              <textarea rows="10" maxlength="140" @keyup.enter="confirmSendMsg" v-model="sendText"></textarea>
            </div>
            <div class="input-button-groups">
              <div class="button-groups">
                <button v-on:click="closeSession(activeSession)">关闭</button>
                <button v-on:click="confirmSendMsg">发送</button>
              </div>
            </div>
          </div>
      </div>
    </div>
  </div>
</template>

<script>
import Axios from './utils/axios'
import Websocket from './utils/websocket'
import AudioRecorder from './utils/audio'
import {getUserSettings} from './settings'
import {trim} from './utils/strings'

export default {
  name: 'App',
  data() {
    return {
      settings: {},
      activeSession: 0,
      hoverSession: 0,
      activeTab: 1,
      loginUser: {},
      sendText: '',
      users: [],
      sessionUsers: [],
      friends: [],
      messageCache: {},
      wsclient: null,
      seq: "0",
      audioDevice: null,
      recording: false,
      axiosIns: null
    }
  },
  async created() {
    const settings = await getUserSettings()
    this.axiosIns = new Axios(settings.http_url)
    await this.loadLoginUser(settings.user_id)
    await this.loadUserFriends(settings.user_id)
    await this.initWSConn(settings.ws_url)
    await this.initAudioDevice(settings.http_url)
    // 激活会话
    await this.autoActiveUserSession()
  },
  updated() {
    this.$nextTick(() => {
      let container = this.$el.querySelector("#message-panel");
      container.scrollTop = container.scrollHeight;
    });
  },
  methods: {
    async loadLoginUser(uid) {
      let res = await this.axiosIns.get(`user/info?uid=${uid}`)
      if (res.error_code !== undefined) {
          console.log('loadLoginUser', res.msg)
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
      let res = await this.axiosIns.get(`friend/list?uid=${uid}`)
      if (res.error_code !== undefined) {
        console.log('loadUserFriends', res.msg)
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
    // 自动激活会话
    async autoActiveUserSession() {
      for (var i = 0; i < this.friends.length; i++) {
        this.openActiveSession(this.friends[i].userId)
      }
    },
    // 初始化音频驱动
    async initAudioDevice(address) {
      const that = this
      let audioDevice = new AudioRecorder()
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
        xhr.open('POST', address + '/audio/transcriptions')
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
    async initWSConn(addr) {
      this.wsclient = new Websocket(addr, this.handleWSMessage)
      let that = this;
      setTimeout(function() {
          that.wsClientAuth()
      }, 1500)
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
    outStyle() {
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
      console.log("功能暂不支持")
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
      if (!messageCache[key]) {
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
    curUser () {
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
    messageList() {
      let userId = this.loginUser.userId;
      let key = `${userId}-${this.activeSession}`
      let messageList = [];
      if (this.messageCache[key] !== undefined) {
          return this.messageCache[key];
      }
      return messageList;
    }
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
* {
  margin: 0;
  padding: 0;
}
html {
  height: 100%;
}
body {
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0);
  overflow: hidden;
}
button:hover {
    cursor: pointer;
}
button:focus {
    outline: none;
    border-color: transparent;
    box-shadow:none;
}
#app {
  height: 100%;
}
.container {
  width: 100%;
  height: 100%;
  display: flex;
}
.contacts {
  width: 25%;
  height: 100%;
  background-color: #d9d9d9;
}
.contact-item {
  padding: 5px 2% 0;
}
.human {
  display: flex;
  padding: 5px;
  border-radius: 4px;
}
.human:hover {
  cursor: pointer;
  background-color: #e2e2e2;
}
.human-active {
  background-color: white;
}
.human > .avatar {
  height: 40px;
}
.avatar img {
  width: 40px;
  height: 40px;
  border-radius: 100%;
}
.human > .username {
  height: 40px;
  line-height: 40px;
  margin-left: 10px;
  font-size: 16px;
  width: 70%;
  color: #474747;
}
.human > .remove {
  display: flex;
  align-items: center;
}
.removeBtn {
  border: none;
  background-color: #555555;
  width: 18px;
  height: 18px;
  border-radius: 9px;
  color: white;
  text-align: center;
  line-height: 18px;
  margin-left: 10px;
}
.removeBtn:hover {
    cursor: pointer;
}
.session-windows {
  width: 75%;
}
.session-info {
  height: 60px;
  background-color: #f7f7f7;
  display: flex;
  align-items: center;
  position: relative;
}
.session-info > .avatar {
  margin-left: 15px;
  display: flex;
  align-items: center;
}
.session-info > .avatar > img{
    width: 50px;
    height: 50px;
    border-radius: 100%;
}
.session-info > .username {
    margin-left: 10px;
    font-size: 18px;
    color: #474747;
}
.tools {
    position: absolute;
    top: 0;
    right: 0;
}
.tools > .closeBtn {
    border: none;
    width: 50px;
    color: #6c6c74;
    font-size: 24px;
    font-weight: lighter;
    line-height: 40px;
    background-color: #f7f7f7;
}
.session-info > .tools > .closeBtn:hover {
    cursor: pointer;
    color: #adadaf;
}
.message-panel {
    height: 60%;
    background-color: white;
    overflow: auto;
    padding-bottom: 20px;
}
.message-panel .message-item {
    padding: 20px 20px 0;
}
.message-item > .user-info {
    display: flex;
}
.message-item > .right {
    justify-content: flex-end;
}
.user-info > .username {
    font-size: 14px;
    padding-top: 5px;
    margin-left: 15px;
    color: #999999;
}
.user-info > .createTime {
    font-size: 14px;
    padding-top: 5px;
    margin-left: 15px;
    color: #999999;
}
.message-info {
    display: flex;
    justify-content: center;
}
.message-info > .message-icon > img {
    height: 16px;
    width: 16px;
}
.message-info > .message-content {
  max-width: 440px;
  background-color: #e2e2e2;
  padding: 8px 12px;
  border-radius: 4px;
  font-size: 14px;
}
.message-info .right {
  background-color: #75b57e;
  color: #ffffff;
}
.message-input {
  border-top: 1px solid #e3e2e2;
  height: 30%;
}
.plugins {
  display: flex;
}
.plugins > .plugin-item {
  padding: 10px;
  margin-left: 10px;
}
.plugins > .plugin-item > img {
  height: 25px;
  width: 25px;
}
.plugins > .history {
  display: flex;
}
.plugins > .plugin-item:hover {
  cursor: pointer;
}
.history > .history-txt {
  margin-left: 10px;
  font-size: 14px;
  line-height: 25px;
}
.input-area {
  padding: 0 15px;
}
.input-area > textarea {
  border: none;
  outline: none;
  resize: none;
  width: 100%;
  font-size: 14px;
}
.input-button-groups {
  position: fixed;
  bottom: 20px;
  right: 10px;
}
.button-groups {
  height: 33px;
  background-color: white;
  display: flex;
  align-items: center;
  justify-content: flex-end;
}
.button-groups > button {
  border: none;
  background-color: #75b47e;
  color: white;
  font-size: 14px;
  padding: 8px 20px;
  margin-right: 10px;
  border-radius: 2px;
}
.button-groups > button:active {
  background-color: #89c492;
}
</style>
