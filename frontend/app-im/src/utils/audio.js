class AudioRecorder {

  //麦克风
  mDevice = null;

  //从麦克风获取的音频流
  mMediaStream = null;

  mAudioContext = null;

  mAudioFromMicrophone = null;

  mMediaRecorder = null;

  mStatus = "stop";

  mChunks = [];

  onStopCallBack = null;

  constructor() {
    this.getAudioRecorderDevice()
  }

  /**
  * 获取录音机对象设备
  * @method getAudioRecorderDevice
  * @for AudioRecorder
  * @returns {Promise} 返回一个promise对象
  */
  getAudioRecorderDevice(){
    //仅用来进行录音
    var constraints = { audio: true};
    // 老的浏览器可能根本没有实现 mediaDevices，所以我们可以先设置一个空的对象
    if(navigator.mediaDevices === undefined) {
        navigator.mediaDevices = {};
    }
    // 一些浏览器部分支持 mediaDevices。我们不能直接给对象设置 getUserMedia
    // 因为这样可能会覆盖已有的属性。这里我们只会在没有getUserMedia属性的时候添加它。
    if(navigator.mediaDevices.getUserMedia === undefined) {
        navigator.mediaDevices.getUserMedia = function(constraints) {

          // 首先，如果有getUserMedia的话，就获得它
          var getUserMedia = navigator.webkitGetUserMedia || navigator.mozGetUserMedia;

          // 一些浏览器根本没实现它 - 那么就返回一个error到promise的reject来保持一个统一的接口
          if(!getUserMedia) {
              return Promise.reject(new Error('getUserMedia is not implemented in this browser'));
          }

          // 否则，为老的navigator.getUserMedia方法包裹一个Promise
          this.mDevice = new Promise(function(resolve, reject) {
              getUserMedia.call(navigator, constraints, resolve, reject);
          });
        }
    }
    else
    {
      this.mDevice = navigator.mediaDevices.getUserMedia(constraints);
    }

    if(this.mDevice != null)
    {
      this.mDevice.then((mediaStream) => { this.openDeviceSuccess.call(this,mediaStream) },this.openDeviceFailure);
    }
  }

  addOnStopCallback(onStop){
    this.onStopCallBack = onStop;
  }

  openDeviceSuccess(mediaStream){
    this.mMediaStream = mediaStream;
  }

  openDeviceFailure(reason){
    let errorMessage;
    switch(reason.name) {
      // 用户拒绝
      case 'NotAllowedError':
      case 'PermissionDeniedError':
        errorMessage = '用户已禁止网页调用录音设备';
        break;
      // 没接入录音设备
      case 'NotFoundError':
      case 'DevicesNotFoundError':
        errorMessage = '录音设备未找到';
        break;
      // 其它错误
      case 'NotSupportedError':
        errorMessage = '不支持录音功能';
        break;
      default:
        errorMessage = '录音调用错误';
        window.console.log(errorMessage);
    }
    console.log(errorMessage);
  }

  /**
  * 开始录音
  * @method startRecord
  * @for AudioRecorder
  * @return {Boolean}
  */
  startRecord(){
    let retValue = false;
    if(this.mStatus == "stop")
    {
        this.mChunks = [];
        if(this.mMediaRecorder == null)
        {
            const AudioContext = window.AudioContext || window.webkitAudioContext;
            this.mAudioContext = new AudioContext();
            //创建音频源
            this.mAudioFromMicrophone = this.mAudioContext.createMediaStreamSource(this.mMediaStream);
            //创建目的节点
            var destination = this.mAudioContext.createMediaStreamDestination();
            this.mMediaRecorder = new MediaRecorder(destination.stream);
            this.mAudioFromMicrophone.connect(destination);
            this.mMediaRecorder.ondataavailable = (audioData) => { this.onProcessData.call(this,audioData)};
            this.mMediaRecorder.onstop = (event) => { this.onStop.call(this,event)};
        }
        this.mMediaRecorder.start();
        this.mStatus = "record";
        retValue = true;
    }
    return retValue;
  }

  onProcessData(audioData)
  {
    this.mChunks.push(audioData.data);
  }

  onStop(){
    var blob = new Blob(this.mChunks, { 'type' : 'audio/mp3' });
    if(this.onStopCallBack != null)
    {
      this.onStopCallBack(blob);
    }
  }

  /**
  * 结束录音
  * @method stopRecord
  * @for AudioRecorder
  */
  stopRecord(){
    if(this.mStatus == "record")
    {
      this.mMediaRecorder.requestData();
      this.mMediaRecorder.stop();
      this.mStatus = "stop";
    }
  }
}

export default AudioRecorder