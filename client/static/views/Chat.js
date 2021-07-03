import Parent from "./Parent.js";

export default class Chat extends Parent {
  constructor() {
    super();
  }

  setTitle(title) {
    document.title = title;
  }
  //uuid receiver & send my uuid
  //show all user - except yourself

  async openChat() {
    let chatWindow = document.getElementById("chat");
    let currentUserUuid = super.getUserSession();

    let obj = { receiver: this.value, sender: currentUserUuid };
    let response = await super.fetch("chat/history", obj);
    if (response != null) {
      let res = await response.json();
      console.log.log(res, "mesagegs");
      chatWindow.textContent = res;
    }
  }

  showOnlineUsers(users) {
    let list = document.getElementById("listUser");
    for (let [k, v] of Object.entries(users)) {
      let p = document.createElement("p");
      p.value = k;
      p.onclick = this.openChat;
      p.textContent = v;
      list.append(p);
    }
  }
  //get senderId, receiverId, msg
  async init() {
    //DRY
    //get list users http
    // let wsusers = new WebSocket("ws://localhost:6969/api/chat/users");
    let r = await fetch("http://localhost:6969/api/chat/users");
    if (r != null) {
      let res = await r.json();
      //show after other user add in chat not now is it correct ?
      //show only other user, !own uuid sort
      this.showOnlineUsers(res);
    }
    //add newuser
    let ws = new WebSocket("ws://localhost:6969/api/chat/newuser");
    ws.addEventListener("message", (e) => {
      console.log(JSON.parse(e.data), "get data from back ws");
    });
    //check state -> then send message
    ws.onopen = () => ws.send(JSON.stringify({ type: "newuser" }));

    ws.onclose = function (event) {
      if (event.wasClean) {
        console.log("Соединение закрыто чисто");
      } else {
        console.log("Обрыв соединения"); // например, "убит" процесс сервера
      }
      console.log("Код: " + event.code + " причина: " + event.reason);
    };

    ws.onmessage = function (event) {
      console.log("Получены данные " + event.data, JSON.parse(event.data));
    };

    ws.onerror = function (error) {
      console.log("Ошибка " + error.message);
    };

    // let wsusers = new WebSocket("ws://localhost:6969/api/chat/users");
    // let list = { type: "listusers" };
    // wsusers.onopen = () => wsusers.send(JSON.stringify(list));
    // wsusers.onmessage = (e) => this.showOnlineUsers(JSON.parse(e.data));

    //data list user -> send  by uuid
    // this.showOnlineUsers(JSON.parse(event.data));

    //getLisrUser() & online and offline
    //click -> userId -> getHistoryByChatId()
    //click -> send msg -> webws -> save msg, notify another user
  }

  async getHtml() {
    // /?DRY
    //show online user & sended message - like history users
    //listUser - dynamic, create history window, textarea, and btn -> history dynamic change data
    let body = `
      <div id="listUser" > list users: </div>
      <div id="chat" >message users </div>
      <div id="message_container" >
      <textarea  id="message"> </textarea> 
      <button id="sendMessage" > send </button>
      </div>
    `;
    return super.showHeader() + body;
  }
}
