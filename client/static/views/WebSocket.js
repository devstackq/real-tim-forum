import Parent from "./Parent.js";

export default class WebSocket extends Parent {
  constructor() {
    super();
    this.wsType = "";
    this.wsMessage = {};
    this.ws = null;
  }

  getWsMessages() {
    return this.wsMessage;
  }
  getWSType() {
    return this.wsType;
  }
  getWebsocket() {
    return this.ws;
  }
  openNewWs() {
    this.ws = new WebSocket("ws://localhost:6969/api/chat");
  }
  addNewUser() {
    let newuser = {
      sender: super.getUserSession(),
      type: "newuser",
    };
    this.ws.onopen = () => {
      this.ws.send(JSON.stringify(newuser));
    };
    console.log('add user', newuser)
  }

  handleWsMessages() {
    // if (super.getAuthState() == "true") {
    this.ws.onmessage = (e) => {
      let msg = JSON.parse(e.data);
      this.wsType = msg.type;
      this.wsMessage = JSON.parse(e.data);
      console.log(this.wsType, 'msg type');
      this.ws.onclose = function (event) {
        console.log("Обрыв соединения"); // например, "убит" процесс сервера
        console.log("Код: " + event.code + " причина: " + event.reason);
      };
      this.ws.onerror = function (error) {
        console.log("Ошибка " + error.message);
      };
    };
  }
}
